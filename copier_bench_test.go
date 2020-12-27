package copy

import (
	"testing"
)

type internal struct {
	I int
}

type testStruct struct {
	S  string
	I  int
	B  bool
	F  float64
	BB []bool
	V  internal
	PV *internal
}

var src = testStruct{
	S:  "string",
	I:  10,
	B:  true,
	F:  4.9,
	BB: []bool{true, false},
	V:  internal{I: 5},
	PV: &internal{I: 15},
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
			V:  internal{I: src.V.I},
			PV: &internal{I: src.V.I},
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

func BenchmarkCopiers(b *testing.B) {
	c := New()
	c.Prepare(&dst, &src)

	for i := 0; i < b.N; i++ {
		c.Copy(&dst, &src)
	}
}
