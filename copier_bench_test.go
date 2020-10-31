package copy

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unsafe"

	"github.com/jinzhu/copier"
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

var repeats = []int{10, 100, 1000}

func name(i int) string {
	name := strconv.Itoa(i)
	return strings.Repeat(" ", 6-len(name)) + name
}

func BenchmarkCopiers_____(b *testing.B) {
	c := New()
	c.Prepare(&dst, &src)
	for _, repeat := range repeats {
		b.Run(name(repeat), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for i := 0; i < repeat; i++ {
					c.Copy(&dst, &src)
				}
			}
		})
	}
}

func BenchmarkCopier______(b *testing.B) {
	c := New()
	copier := c.Get(&dst, &src)
	for _, repeat := range repeats {
		b.Run(name(repeat), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for i := 0; i < repeat; i++ {
					copier.Copy(&dst, &src)
				}
			}
		})
	}
}

func copysruct(dst, src *testStruct) {
	*dst = *src
}

func BenchmarkDirectCopy__(b *testing.B) {
	for _, repeat := range repeats {
		b.Run(name(repeat), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for i := 0; i < repeat; i++ {
					copysruct(&dst, &src)
				}
			}
		})
	}
}

func BenchmarkJinzhuCopier(b *testing.B) {
	for _, repeat := range repeats {
		b.Run(name(repeat), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for i := 0; i < repeat; i++ {
					copier.Copy(&dst, &src)
				}
			}
		})
	}
}

func BenchmarkManualCopy__(b *testing.B) {
	for _, repeat := range repeats {
		b.Run(name(repeat), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for i := 0; i < repeat; i++ {
					dst = testStruct{
						S:  src.S,
						I:  src.I,
						BB: src.BB,
						V:  src.V,
					}
				}
			}
		})
	}
}

func BenchmarkFiledIndex(b *testing.B) {
	src := struct {
		I int
	}{
		I: 10,
	}
	dst := src
	dst.I = 0
	vSrc := reflect.ValueOf(&src).Elem()
	vDest := reflect.ValueOf(&dst).Elem()
	for i := 0; i < b.N; i++ {
		vDest.Field(0).Set(vSrc.Field(0))
	}
}

func BenchmarkFiledPointer(b *testing.B) {
	src := struct {
		I int
	}{
		I: 10,
	}
	dst := src
	dst.I = 0

	vSrc := reflect.ValueOf(&src)
	srcOffset := vSrc.Elem().Type().Field(0).Offset
	// srcFieldType := vSrc.Elem().Type().Field(0).Type
	srcPtr := unsafe.Pointer(vSrc.Pointer())
	size := vSrc.Elem().Type().Field(0).Type.Size()

	vDest := reflect.ValueOf(&dst)
	dstOffset := vDest.Elem().Type().Field(0).Offset
	// dstFieldType := vDest.Elem().Type().Field(0).Type
	dstPtr := unsafe.Pointer(vDest.Pointer())
	for i := 0; i < b.N; i++ {
		memcopy(unsafe.Pointer(uintptr(dstPtr)+dstOffset), unsafe.Pointer(uintptr(srcPtr)+srcOffset), size)
		// src := reflect.NewAt(srcFieldType, unsafe.Pointer(uintptr(srcPtr)+srcOffset)).Elem()
		// dst := reflect.NewAt(dstFieldType, unsafe.Pointer(uintptr(dstPtr)+dstOffset)).Elem()
		// dst.Set(src)
	}
}

// func BenchmarkFuncClosure(b *testing.B) {
// 	src := struct {
// 		I int
// 	}{
// 		I: 10,
// 	}
// 	dst := src
// 	dst.I = 0

// 	vSrc := reflect.ValueOf(&src)
// 	srcOffset := vSrc.Elem().Type().Field(0).Offset
// 	// srcFieldType := vSrc.Elem().Type().Field(0).Type
// 	srcPtr := unsafe.Pointer(vSrc.Pointer())
// 	size := vSrc.Elem().Type().Field(0).Type.Size()

// 	vDest := reflect.ValueOf(&dst)
// 	dstOffset := vDest.Elem().Type().Field(0).Offset
// 	// dstFieldType := vDest.Elem().Type().Field(0).Type
// 	dstPtr := unsafe.Pointer(vDest.Pointer())

// 	f := func(dst, src unsafe.Pointer) {
// 		memcopy(unsafe.Pointer(uintptr(dst)+dstOffset), unsafe.Pointer(uintptr(src)+srcOffset), size)
// 	}
// 	for i := 0; i < b.N; i++ {
// 		f(dstPtr, srcPtr)
// 	}
// }

// type funcState struct {
// 	SrcOffset uintptr
// 	DstOffset uintptr
// 	Size      uintptr
// }

// func (s funcState) do(dst, src unsafe.Pointer) {
// 	memcopy(unsafe.Pointer(uintptr(dst)+s.DstOffset), unsafe.Pointer(uintptr(src)+s.SrcOffset), s.Size)
// }

// func BenchmarkFuncObject(b *testing.B) {
// 	src := struct {
// 		I int
// 	}{
// 		I: 10,
// 	}
// 	dst := src
// 	dst.I = 0

// 	vSrc := reflect.ValueOf(&src)
// 	srcOffset := vSrc.Elem().Type().Field(0).Offset
// 	// srcFieldType := vSrc.Elem().Type().Field(0).Type
// 	srcPtr := unsafe.Pointer(vSrc.Pointer())
// 	size := vSrc.Elem().Type().Field(0).Type.Size()

// 	vDest := reflect.ValueOf(&dst)
// 	dstOffset := vDest.Elem().Type().Field(0).Offset
// 	// dstFieldType := vDest.Elem().Type().Field(0).Type
// 	dstPtr := unsafe.Pointer(vDest.Pointer())

// 	f := funcState{
// 		DstOffset: dstOffset,
// 		SrcOffset: srcOffset,
// 		Size:      size,
// 	}
// 	for i := 0; i < b.N; i++ {
// 		f.do(dstPtr, srcPtr)
// 	}
// }
