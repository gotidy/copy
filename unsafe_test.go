package copy

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestPtrOf(t *testing.T) {
	var i int
	p := &i
	if PtrOf(p) != unsafe.Pointer(reflect.ValueOf(p).Pointer()) {
		t.Errorf("PtrOf() = %v, want %v", PtrOf(p), unsafe.Pointer(reflect.ValueOf(p).Pointer()))
	}
	if PtrOf(nil) != nil {
		t.Errorf("PtrOf() = %v, want %v", PtrOf(p), nil)
	}
}
