// +build !safe

package copy

import "unsafe"

// More safe and independent from internal structs
// func ifaceData(i interface{}) unsafe.Pointer {
// 	v := reflect.ValueOf(i)
// 	if v.Kind() != reflect.Ptr {
// 		panic("source must be pointer to struct")
// 	}
// 	return unsafe.Pointer(v.Pointer())
// }

type Type = unsafe.Pointer

func DataOf(i interface{}) (typ Type, data unsafe.Pointer) {
	if i == nil {
		return nil, nil
	}
	iface := *(*iface)(unsafe.Pointer(&i))
	return iface.Type, iface.Data
}

func TypeOf(i interface{}) (typ Type) {
	if i == nil {
		return nil
	}
	return (*iface)(unsafe.Pointer(&i)).Type
}

type iface struct {
	Type, Data unsafe.Pointer
}

func PtrOf(i interface{}) unsafe.Pointer {
	if i == nil {
		return nil
	}
	// eface := *(*iface)(unsafe.Pointer(&i))
	return (*iface)(unsafe.Pointer(&i)).Data
}
