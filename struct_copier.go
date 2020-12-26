// Package copy provides a copy library to make copying of structs to/from others structs a bit easier.
package copy

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/gotidy/copy/funcs"
	"github.com/gotidy/copy/internal/cache"
)

type fieldCopier = func(dst, src unsafe.Pointer)

// StructCopier fills a destination from source.
type StructCopier struct {
	BaseCopier

	copiers []fieldCopier
}

func NewStructCopier(c *Copiers, dst, src reflect.Type) *StructCopier {
	copier := &StructCopier{BaseCopier: NewBaseCopier(c, dst, src)}

	srcStruct := c.cache.GetByType(src)
	dstStruct := c.cache.GetByType(dst)

	for i := 0; i < srcStruct.NumField(); i++ {
		srcField := srcStruct.Field(i)
		if dstField, ok := dstStruct.FieldByName(srcField.Name); ok {
			if f := copier.fieldCopier(dstField, srcField); f != nil {
				copier.copiers = append(copier.copiers, f)
			}
		}
	}

	return copier
}

// Copy copies the contents of src into dst. Dst and src each must be a pointer to struct.
func (c *StructCopier) Copy(dst, src interface{}) {
	dstType, dstPtr := DataOf(dst)
	srcType, srcPtr := DataOf(src)

	if c.src.Check(srcType) {
		panic("source expected type " + c.src.Name + ", but has " + reflect.TypeOf(src).String())
	}
	if c.dst.Check(dstType) {
		panic("destination expected type " + c.dst.Name + ", but has " + reflect.TypeOf(dst).String())
	}

	c.copy(dstPtr, srcPtr)
}

func (c *StructCopier) copy(dst, src unsafe.Pointer) {
	for _, c := range c.copiers {
		c(dst, src)
	}
}

func (c *StructCopier) fieldCopier(dst, src cache.Field) fieldCopier {
	dstOffset := dst.Offset
	srcOffset := src.Offset
	copier := funcs.Get(dst.Type, src.Type)
	if copier != nil {
		return func(dstPtr, srcPtr unsafe.Pointer) {
			copier(unsafe.Pointer(uintptr(dstPtr)+dstOffset), unsafe.Pointer(uintptr(srcPtr)+srcOffset))
		}
	}

	// same type -> same type
	if src.Type == dst.Type {
		size := int(src.Type.Size())

		return func(dstPtr, srcPtr unsafe.Pointer) {
			// More safe and independent from internal structs
			// src := reflect.NewAt(src.Type, unsafe.Pointer(uintptr(srcPtr)+src.Offset)).Elem()
			// dst := reflect.NewAt(dst.Type, unsafe.Pointer(uintptr(dstPtr)+dst.Offset)).Elem()
			// dst.Set(src)
			memcopy(unsafe.Pointer(uintptr(dstPtr)+dst.Offset), unsafe.Pointer(uintptr(srcPtr)+src.Offset), size)
		}
	}

	// struct -> struct
	if src.Type.Kind() == reflect.Struct && dst.Type.Kind() == reflect.Struct {
		copier := c.get(dst.Type, src.Type)

		return func(dstPtr, srcPtr unsafe.Pointer) {
			copier.copy(unsafe.Pointer(uintptr(dstPtr)+dst.Offset), unsafe.Pointer(uintptr(srcPtr)+src.Offset))
		}
	}

	// *struct -> struct
	if src.Type.Kind() == reflect.Ptr && src.Type.Elem().Kind() == reflect.Struct && dst.Type.Kind() == reflect.Struct {
		copier := c.get(dst.Type, src.Type.Elem())

		return func(dstPtr, srcPtr unsafe.Pointer) {
			srcFieldPtr := (**struct{})(unsafe.Pointer(uintptr(srcPtr) + src.Offset))
			if *srcFieldPtr == nil {
				return
			}
			copier.copy(unsafe.Pointer(uintptr(dstPtr)+dst.Offset), unsafe.Pointer(*srcFieldPtr))
		}
	}

	// struct -> *struct
	if src.Type.Kind() == reflect.Struct && dst.Type.Kind() == reflect.Ptr && dst.Type.Elem().Kind() == reflect.Struct {
		copier := c.get(dst.Type.Elem(), src.Type)

		dstSize := int(dst.Type.Elem().Size())

		return func(dstPtr, srcPtr unsafe.Pointer) {
			dstFieldPtr := (**struct{})(unsafe.Pointer(uintptr(dstPtr) + dst.Offset))
			if *dstFieldPtr == nil {
				*dstFieldPtr = (*struct{})(alloc(dstSize))
			}

			copier.copy(unsafe.Pointer(*dstFieldPtr), unsafe.Pointer(uintptr(srcPtr)+src.Offset))
		}
	}

	// *struct -> *struct
	if src.Type.Kind() == reflect.Ptr && src.Type.Elem().Kind() == reflect.Struct &&
		dst.Type.Kind() == reflect.Ptr && dst.Type.Elem().Kind() == reflect.Struct {
		copier := c.get(dst.Type.Elem(), src.Type.Elem())

		dstSize := int(dst.Type.Elem().Size())

		return func(dstPtr, srcPtr unsafe.Pointer) {
			srcFieldPtr := (**struct{})(unsafe.Pointer(uintptr(srcPtr) + src.Offset))
			if *srcFieldPtr == nil {
				return
			}
			dstFieldPtr := (**struct{})(unsafe.Pointer(uintptr(dstPtr) + dst.Offset))
			if *dstFieldPtr == nil {
				*dstFieldPtr = (*struct{})(alloc(dstSize))
			}

			copier.copy(unsafe.Pointer(*dstFieldPtr), unsafe.Pointer(*srcFieldPtr))
		}
	}

	if !c.options.Skip {
		panic(fmt.Errorf(`field «%s» of type «%s» is not assignable to field «%s» of type «%s»`, src.Name, src.Type.String(), dst.Name, dst.Type.String()))
	}

	return nil
}

type StructToPStructCopier struct {
	BaseCopier

	structCopier func(dst, src unsafe.Pointer)
	size         int
}

func NewStructToPStructCopier(c *Copiers, dst, src reflect.Type) *StructToPStructCopier {
	copier := &StructToPStructCopier{BaseCopier: NewBaseCopier(c, dst, src)}
	dst = dst.Elem() // *SomeStruct -> SomeStruct
	copier.structCopier = copier.get(dst, src).copy
	copier.size = int(dst.Size())
	return copier
}

func (c *StructToPStructCopier) Copy(dst, src interface{}) {
	dstType, dstPtr := DataOf(dst)
	srcType, srcPtr := DataOf(src)

	if c.src.Check(srcType) {
		panic("source expected type " + c.src.Name + ", but has " + reflect.TypeOf(src).String())
	}
	if c.dst.Check(dstType) {
		panic("destination expected type " + c.dst.Name + ", but has " + reflect.TypeOf(dst).String())
	}

	c.copy(dstPtr, srcPtr)
}

func (c *StructToPStructCopier) copy(dst, src unsafe.Pointer) {
	dstFieldPtr := (**struct{})(dst)
	if *dstFieldPtr == nil {
		*dstFieldPtr = (*struct{})(alloc(c.size))
	}

	c.structCopier(unsafe.Pointer(*dstFieldPtr), src)
}
