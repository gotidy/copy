package copy

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/gotidy/copy/cache"
	"github.com/gotidy/copy/funcs"
)

type fieldCopier = func(src, dest unsafe.Pointer)

func getFieldCopier(src, dest cache.Field) fieldCopier {
	copier := funcs.Get(src.Type, dest.Type)
	if copier != nil {
		return func(srcPtr, destPtr unsafe.Pointer) {
			copier(unsafe.Pointer(uintptr(srcPtr)+src.Offset), unsafe.Pointer(uintptr(destPtr)+dest.Offset))
		}
	}
	if !(src.Type.AssignableTo(dest.Type)) {
		panic(fmt.Errorf(`field "%s" of type %s is not assignable to field %s of type %s`, src.Name, src.Type.String(), dest.Name, dest.Type.String()))
	}
	return func(srcPtr, destPtr unsafe.Pointer) {
		src := reflect.NewAt(src.Type, unsafe.Pointer(uintptr(srcPtr)+src.Offset)).Elem()
		dest := reflect.NewAt(dest.Type, unsafe.Pointer(uintptr(destPtr)+dest.Offset)).Elem()
		dest.Set(src)
	}
}
