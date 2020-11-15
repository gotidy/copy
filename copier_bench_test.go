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
func BenchmarkCopiers(b *testing.B) {
	c := New()
	c.Prepare(&dst, &src)

	for i := 0; i < b.N; i++ {
		c.Copy(&dst, &src)
	}
}
