package copy

import (
	"fmt"
	"reflect"
	"unsafe"
)

type SliceCopier struct {
	BaseCopier

	copier  func(dst, src unsafe.Pointer)
	dstSize uintptr // Size of the destination element
	srcSize uintptr // Size of the source element
}

func NewSliceCopier(c *Copiers) *SliceCopier {
	copier := &SliceCopier{BaseCopier: NewBaseCopier(c)}
	return copier
}

func (c *SliceCopier) init(dst, src reflect.Type) {
	c.BaseCopier.init(dst, src)

	c.copier = c.getCopierFunc(dst.Elem(), src.Elem(), 0, 0)
	if c.copier == nil && !c.options.Skip {
		panic(fmt.Errorf(`slice element of type «%s» is not assignable to slice element of type «%s»`, src.String(), dst.String()))
	}
	c.dstSize = dst.Elem().Size()
	c.srcSize = src.Elem().Size()
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
	if c.copier == nil {
		return
	}
	srcSlice := sliceAt(src, c.srcSize)
	dstSlice := makeSliceAt(dst, c.dstSize, srcSlice.Len)

	for i := 0; i < srcSlice.Len; i++ {
		c.copier(dstSlice.Index(i), srcSlice.Index(i))
	}
}
