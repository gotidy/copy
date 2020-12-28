// Package copy provides a copy library to make copying of structs to/from others structs a bit easier.
package copy

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/gotidy/copy/funcs"
	"github.com/gotidy/copy/internal/cache"
)

type copierFunc = func(dst, src unsafe.Pointer)

// StructCopier fills a destination from source.
type StructCopier struct {
	BaseCopier

	copiers []copierFunc
}

func NewStructCopier(c *Copiers) *StructCopier {
	copier := &StructCopier{BaseCopier: NewBaseCopier(c)}
	return copier
}

func (c *StructCopier) init(dst, src reflect.Type) {
	c.BaseCopier.init(dst, src)

	srcStruct := c.cache.GetByType(src)
	dstStruct := c.cache.GetByType(dst)

	for i := 0; i < srcStruct.NumField(); i++ {
		srcField := srcStruct.Field(i)
		if dstField, ok := dstStruct.FieldByName(srcField.Name); ok {
			if f := c.fieldCopier(dstField, srcField); f != nil {
				c.copiers = append(c.copiers, f)
			}
		}
	}
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

func (c *StructCopier) fieldCopier(dst, src cache.Field) copierFunc {
	// TODO: Unify with the slice copier, may be do it inside Copiers.get
	dstOffset := dst.Offset
	srcOffset := src.Offset
	copierFunc := funcs.Get(dst.Type, src.Type)
	if copierFunc != nil {
		return func(dstPtr, srcPtr unsafe.Pointer) {
			copierFunc(unsafe.Pointer(uintptr(dstPtr)+dstOffset), unsafe.Pointer(uintptr(srcPtr)+srcOffset))
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

	copier, err := c.get(dst.Type, src.Type)
	if err == nil {
		return func(dstPtr, srcPtr unsafe.Pointer) {
			copier.copy(unsafe.Pointer(uintptr(dstPtr)+dst.Offset), unsafe.Pointer(uintptr(srcPtr)+src.Offset))
		}
	}

	if !c.options.Skip {
		panic(fmt.Errorf(`field «%s» of type «%s» is not assignable to field «%s» of type «%s»`, src.Name, src.Type.String(), dst.Name, dst.Type.String()))
	}

	return nil
}
