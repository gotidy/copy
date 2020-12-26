// +build safe

package copy

import (
	"reflect"
	"unsafe"
)

type Type = reflect.Type

func DataOf(i interface{}) (typ Type, data unsafe.Pointer) {
	if i == nil {
		return nil, nil
	}
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		return nil, nil
	}
	return reflect.TypeOf(i), unsafe.Pointer(v.Pointer())
}

func TypeOf(i interface{}) (typ Type) {
	if i == nil {
		return nil
	}
	return reflect.TypeOf(i)
}

type iface struct {
	Type, Data unsafe.Pointer
}

func PtrOf(i interface{}) unsafe.Pointer {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		return nil
	}
	return unsafe.Pointer(v.Pointer())
}
