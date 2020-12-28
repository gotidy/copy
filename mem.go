package copy

import (
	"reflect"
	"unsafe"
)

func memcopy(dst, src unsafe.Pointer, size int) {
	var srcSlice []byte
	srcSH := (*reflect.SliceHeader)((unsafe.Pointer(&srcSlice)))
	srcSH.Data = uintptr(src)
	srcSH.Cap = int(size)
	srcSH.Len = int(size)

	var dstSlice []byte
	dstSH := (*reflect.SliceHeader)((unsafe.Pointer(&dstSlice)))
	dstSH.Data = uintptr(dst)
	dstSH.Cap = int(size)
	dstSH.Len = int(size)

	copy(dstSlice, srcSlice)
}

func alloc(size int) unsafe.Pointer {
	if size == 0 {
		return nil
	}
	size = (size + 7) / 8 // size in int64
	return unsafe.Pointer(&(make([]int64, size)[0]))
}

type slice struct {
	data unsafe.Pointer
	size uintptr

	Len int
}

func sliceAt(ptr unsafe.Pointer, size uintptr) slice {
	s := (*reflect.SliceHeader)(ptr)
	return slice{data: unsafe.Pointer(s.Data), size: size, Len: s.Len}
}

func makeSliceAt(ptr unsafe.Pointer, size uintptr, len int) slice {
	s := (*reflect.SliceHeader)(ptr)
	if s.Cap < len || s.Cap > len*2 {
		s.Data = uintptr(alloc(int(size) * len))
		s.Cap = len
	}
	s.Len = len
	return slice{data: unsafe.Pointer(s.Data), size: size, Len: s.Len}
}

func (s slice) Index(i int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(s.data) + s.size*uintptr(i))
}
