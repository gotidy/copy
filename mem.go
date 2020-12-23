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

// More safe and independent from internal structs
// func ifaceData(i interface{}) unsafe.Pointer {
// 	v := reflect.ValueOf(i)
// 	if v.Kind() != reflect.Ptr {
// 		panic("source must be pointer to struct")
// 	}
// 	return unsafe.Pointer(v.Pointer())
// }

func iface(i interface{}) (typ, data unsafe.Pointer) {
	if i == nil {
		panic("input parameter is nil")
	}

	type iface struct {
		Type, Data unsafe.Pointer
	}

	return (*iface)(unsafe.Pointer(&i)).Type, (*iface)(unsafe.Pointer(&i)).Data
}

func ifaceType(i interface{}) (typ unsafe.Pointer) {
	// if i == nil {
	// 	panic("input parameter is nil")
	// }

	type iface struct {
		Type, Data unsafe.Pointer
	}

	eface := *(*iface)(unsafe.Pointer(&i))
	return eface.Type
}

func ifaceData(i interface{}) unsafe.Pointer {
	if i == nil {
		panic("input parameter is nil")
	}

	type iface struct {
		Type, Data unsafe.Pointer
	}

	return (*iface)(unsafe.Pointer(&i)).Data
}
