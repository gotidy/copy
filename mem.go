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
	size = (size + 7) / 8 // size in int64
	return unsafe.Pointer(&(make([]int64, size)[0]))
}
