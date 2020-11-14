//go:generate go run gen.go
//go:generate gofmt -s -w funcs.gen.go
// Package funcs provides copy functions for specified types.
package funcs

import (
	"reflect"
	"sync"
	"unsafe"
)

type funcKey struct {
	Src reflect.Type
	Dst reflect.Type
	// Zero bool // Initiate pointer to zero value
}

func typeOf(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}

func typeOfPointer(v interface{}) reflect.Type {
	return reflect.PtrTo(reflect.TypeOf(v))
}

// CopyFuncs is the storage of functions intended for copying data.
type CopyFuncs struct {
	mu    sync.RWMutex
	funcs map[funcKey]func(dst, src unsafe.Pointer)
	sizes []func(dst, src unsafe.Pointer)
}

// Get the copy function for the pair of types, if it is not found then nil is returned.
func (t *CopyFuncs) Get(dst, src reflect.Type) func(dst, src unsafe.Pointer) {
	t.mu.RLock()
	f := t.funcs[funcKey{Src: src, Dst: dst}]
	t.mu.RUnlock()
	if f != nil {
		return f
	}

	if dst.Kind() != src.Kind() {
		return nil
	}

	if dst.Kind() == reflect.String {
		// TODO
		return nil
	}

	same := dst == src

	switch dst.Kind() {
	case reflect.Array, reflect.Chan, reflect.Ptr, reflect.Slice:
		same = same || dst.Elem() == src.Elem()
	case reflect.Map:
		same = same || (dst.Elem() == src.Elem() && dst.Key() == src.Key())
	}

	if same && dst.Size() == src.Size() && src.Size() > 0 && src.Size() <= uintptr(len(t.sizes)) {
		return t.sizes[src.Size()-1]
	}

	return nil
}

// Set the copy function for the pair of types.
func (t *CopyFuncs) Set(dst, src reflect.Type, f func(dst, src unsafe.Pointer)) {
	t.mu.Lock()
	t.funcs[funcKey{Src: src, Dst: dst}] = f
	t.mu.Unlock()
}

// Get the copy function for the pair of types, if it is not found then nil is returned.
func Get(dst, src reflect.Type) func(dst, src unsafe.Pointer) {
	return funcs.Get(dst, src)
}

// Set the copy function for the pair of types.
func Set(dst, src reflect.Type, f func(dst, src unsafe.Pointer)) {
	funcs.Set(dst, src, f)
}

var funcs = &CopyFuncs{
	funcs: map[funcKey]func(dst, src unsafe.Pointer){},
	sizes: []func(dst, src unsafe.Pointer){},
}
