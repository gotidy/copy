package copy

import (
	"reflect"
	"unsafe"
)

type SliceCopier struct {
	BaseCopier
}

func NewSliceCopier(c *Copiers) *SliceCopier {
	copier := &SliceCopier{BaseCopier: NewBaseCopier(c)}
	return copier
}

func (c *SliceCopier) init(dst, src reflect.Type) {
	c.BaseCopier.init(dst, src)

	// TODO: Init
}

// Copy copies the contents of src into dst. Dst and src each must be a pointer to struct.
func (c *SliceCopier) Copy(dst, src interface{}) {
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

func (c *SliceCopier) copy(dst, src unsafe.Pointer) {
	// for _, c := range c.copiers {
	// 	c(dst, src)
	// }
}

type slice struct {
	data unsafe.Pointer
	size int
	len  int
}

func sliceAt(ptr unsafe.Pointer, size int) slice {
	s := (*reflect.SliceHeader)(ptr)
	return slice{data: unsafe.Pointer(s.Data), size: size, len: s.Len}
}

// TODO: Init
func makeSliceAt(ptr unsafe.Pointer, size int, len int) slice {
	s := (*reflect.SliceHeader)(ptr)
	return slice{data: unsafe.Pointer(s.Data), size: size, len: s.Len}
}

func (s slice) Index(i int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(s.data) + uintptr(s.size*i))
}
