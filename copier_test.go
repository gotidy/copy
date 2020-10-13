package copy

import (
	"reflect"
	"testing"
)

func TestCopier_Copy(t *testing.T) {
	type internal struct {
		I int
	}
	type testStruct struct {
		S  string
		I  int
		BB []bool
		V  internal
	}

	src := testStruct{
		S:  "string",
		I:  10,
		BB: []bool{true, false},
		V:  internal{I: 5},
	}
	dest := testStruct{}

	c := New("") // New("json")
	c.Prefetch(&src, &dest)
	c.Copy(&src, &dest)
	if !reflect.DeepEqual(dest, src) {
		t.Errorf("getFieldCopier() = %v, want %v", dest, src)
	}
}
