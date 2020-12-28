package copy

import (
	"reflect"
	"unsafe"

	"github.com/gotidy/copy/funcs"
)

type TypeInfo struct {
	Type Type
	Name string
}

func NewTypeInfo(typ reflect.Type) TypeInfo {
	p := reflect.New(typ)
	return TypeInfo{Type: TypeOf(p.Interface()), Name: p.String()}
}

func (t TypeInfo) Check(typ Type) bool {
	return t.Type != typ
}

type BaseCopier struct {
	*Copiers

	dst TypeInfo
	src TypeInfo
}

func NewBaseCopier(c *Copiers) BaseCopier {
	return BaseCopier{Copiers: c}
}

func (b *BaseCopier) init(dst, src reflect.Type) {
	b.dst = NewTypeInfo(dst)
	b.src = NewTypeInfo(src)
}

func (b *BaseCopier) getCopierFunc(dst, src reflect.Type, dstOffset, srcOffset uintptr) copierFunc {
	copierFunc := funcs.Get(dst, src)
	if copierFunc != nil {
		return func(dstPtr, srcPtr unsafe.Pointer) {
			copierFunc(unsafe.Pointer(uintptr(dstPtr)+dstOffset), unsafe.Pointer(uintptr(srcPtr)+srcOffset))
		}
	}

	// same type -> same type
	if src == dst {
		size := int(src.Size())

		return func(dstPtr, srcPtr unsafe.Pointer) {
			// More safe and independent from internal structs
			// src := reflect.NewAt(src, unsafe.Pointer(uintptr(srcPtr)+src.Offset)).Elem()
			// dst := reflect.NewAt(dst, unsafe.Pointer(uintptr(dstPtr)+dst.Offset)).Elem()
			// dst.Set(src)
			memcopy(unsafe.Pointer(uintptr(dstPtr)+dstOffset), unsafe.Pointer(uintptr(srcPtr)+srcOffset), size)
		}
	}

	copier, err := b.get(dst, src)
	if err == nil {
		return func(dstPtr, srcPtr unsafe.Pointer) {
			copier.copy(unsafe.Pointer(uintptr(dstPtr)+dstOffset), unsafe.Pointer(uintptr(srcPtr)+srcOffset))
		}
	}

	return nil
}

type ValueToPValueCopier struct {
	BaseCopier

	structCopier func(dst, src unsafe.Pointer)
	size         int
}

func NewValueToPValueCopier(c *Copiers) *ValueToPValueCopier {
	copier := &ValueToPValueCopier{BaseCopier: NewBaseCopier(c)}
	return copier
}

func (c *ValueToPValueCopier) init(dst, src reflect.Type) {
	c.BaseCopier.init(dst, src)
	dst = dst.Elem()                                // *struct -> struct
	c.structCopier = checkGet(c.get(dst, src)).copy // Get struct copier for struct -> struct
	c.size = int(dst.Size())
}

func (c *ValueToPValueCopier) Copy(dst, src interface{}) {
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

func (c *ValueToPValueCopier) copy(dst, src unsafe.Pointer) {
	dstFieldPtr := (**struct{})(dst)
	if *dstFieldPtr == nil {
		*dstFieldPtr = (*struct{})(alloc(c.size))
	}

	c.structCopier(unsafe.Pointer(*dstFieldPtr), src)
}

type PValueToValueCopier struct {
	BaseCopier

	structCopier func(dst, src unsafe.Pointer)
}

func NewPValueToValueCopier(c *Copiers) *PValueToValueCopier {
	copier := &PValueToValueCopier{BaseCopier: NewBaseCopier(c)}
	return copier
}

func (c *PValueToValueCopier) init(dst, src reflect.Type) {
	c.BaseCopier.init(dst, src)
	src = src.Elem()                                // *struct -> struct
	c.structCopier = checkGet(c.get(dst, src)).copy // Get struct copier for struct -> struct
}

func (c *PValueToValueCopier) Copy(dst, src interface{}) {
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

func (c *PValueToValueCopier) copy(dst, src unsafe.Pointer) {
	srcFieldPtr := (**struct{})(src)
	if *srcFieldPtr == nil {
		return
	}

	c.structCopier(dst, unsafe.Pointer(*srcFieldPtr))
}

type PValueToPValueCopier struct {
	BaseCopier

	structCopier func(dst, src unsafe.Pointer)
	size         int
}

func NewPValueToPValueCopier(c *Copiers) *PValueToPValueCopier {
	copier := &PValueToPValueCopier{BaseCopier: NewBaseCopier(c)}
	return copier
}

func (c *PValueToPValueCopier) init(dst, src reflect.Type) {
	c.BaseCopier.init(dst, src)
	dst = dst.Elem()                                // *struct -> struct
	src = src.Elem()                                // *struct -> struct
	c.structCopier = checkGet(c.get(dst, src)).copy // Get struct copier for struct -> struct
	c.size = int(dst.Size())
}

func (c *PValueToPValueCopier) Copy(dst, src interface{}) {
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

func (c *PValueToPValueCopier) copy(dst, src unsafe.Pointer) {
	srcFieldPtr := (**struct{})(src)
	if *srcFieldPtr == nil {
		return
	}

	dstFieldPtr := (**struct{})(dst)
	if *dstFieldPtr == nil {
		*dstFieldPtr = (*struct{})(alloc(c.size))
	}

	c.structCopier(unsafe.Pointer(*dstFieldPtr), unsafe.Pointer(*srcFieldPtr))
}
