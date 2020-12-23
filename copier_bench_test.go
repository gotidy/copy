package copy

import (
	"reflect"
	"testing"
	"unsafe"
)

type internal struct {
	I int
}

type testStruct struct {
	S  string
	I  int
	BB []bool
	V  internal
}

var src = testStruct{
	S:  "string",
	I:  10,
	BB: []bool{true, false},
	V:  internal{I: 5},
}

var dst = testStruct{}

func copyStruct(dst, src *testStruct) {
	*dst = *src
}

func BenchmarkDirectCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copyStruct(&dst, &src)
	}
}

func BenchmarkManualCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dst = testStruct{
			S:  src.S,
			I:  src.I,
			BB: src.BB,
			V:  src.V,
		}
	}
}

func BenchmarkCopier(b *testing.B) {
	c := New()
	copier := c.Get(&dst, &src)

	for i := 0; i < b.N; i++ {
		copier.Copy(&dst, &src)
	}
}
func BenchmarkCopierIf(b *testing.B) {
	type Copier interface {
		Copy(dst interface{}, src interface{})
	}

	c := New()
	var copier Copier = c.Get(&dst, &src)

	for i := 0; i < b.N; i++ {
		copier.Copy(&dst, &src)
	}
}
func BenchmarkCopiers(b *testing.B) {
	c := New()
	c.Prepare(&dst, &src)

	for i := 0; i < b.N; i++ {
		c.Copy(&dst, &src)
	}
}

var resultTypeOf reflect.Type

func BenchmarkTypeTypeOf(b *testing.B) {
	var v int
	for i := 0; i < b.N; i++ {
		resultTypeOf = reflect.TypeOf(&v)
		resultTypeOf = reflect.TypeOf(&v)
		// resultTypeOf = reflect.TypeOf(&v)
		// resultTypeOf = reflect.TypeOf(&v)
		// resultTypeOf = reflect.TypeOf(&v)
		// resultTypeOf = reflect.TypeOf(&v)
		// resultTypeOf = reflect.TypeOf(&v)
		// resultTypeOf = reflect.TypeOf(&v)
		// resultTypeOf = reflect.TypeOf(&v)
		// resultTypeOf = reflect.TypeOf(&v)
	}
}

var resultIface unsafe.Pointer

func BenchmarkTypeIface(b *testing.B) {
	var v int
	for i := 0; i < b.N; i++ {
		resultIface = ifaceType(&v)
		resultIface = ifaceType(&v)
		// resultIface = ifaceType(&v)
		// resultIface = ifaceType(&v)
		// resultIface = ifaceType(&v)
		// resultIface = ifaceType(&v)
		// resultIface = ifaceType(&v)
		// resultIface = ifaceType(&v)
		// resultIface = ifaceType(&v)
		// resultIface = ifaceType(&v)
	}
}
