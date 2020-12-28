package copy

import (
	"testing"
	"unsafe"
)

func Test_sliceAt(t *testing.T) {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := sliceAt(unsafe.Pointer(&ii), unsafe.Sizeof(ii[0]))
	for i := 0; i < s.Len; i++ {
		if *(*int)(s.Index(i)) != ii[i] {
			t.Errorf("*(*int)(s.Index(i)) = %v, want %v", *(*int)(s.Index(i)), ii[i])
		}
	}
}

func Test_makeSliceAt(t *testing.T) {
	const size = 10
	var ii []int
	s := makeSliceAt(unsafe.Pointer(&ii), unsafe.Sizeof(ii[0]), size)
	if len(ii) != size {
		t.Errorf("actual slice size is %v, expected %v", len(ii), size)
	}
	for i := 0; i < size; i++ {
		*(*int)(s.Index(i)) = i + 1
		if ii[i] != i+1 {
			t.Errorf("actual %v, expected %v", ii[i], i+1)
		}
	}
}
