package copy

import (
	"encoding/json"
	"testing"
)

func trim(b []rune) []rune {
	if len(b) > 20 {
		return b[0:20]
	}
	return b
}

func diff(t *testing.T, prefix string, a, b []byte) {
	ar := []rune(string(a))
	br := []rune(string(b))
	l := len(ar)
	if l < len(br) {
		l = len(br)
	}
	for i := 0; i < l; i++ {
		if i >= len(ar) || i >= len(br) || ar[i] != br[i] {
			j := i - 10
			if j < 0 {
				j = 0
			}
			t.Errorf(prefix+": diverge at %d: «%s» vs «%s»", i, string(trim(ar[j:])), string(trim(br[j:])))
			return
		}
	}
}

func equal(t *testing.T, dst, src interface{}) {
	dstData, err := json.Marshal(dst)
	if err != nil {
		t.Fatalf("equal: %s", err)
	}
	t.Logf("dst: %s", string(dstData))
	srcData, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("equal: %s", err)
	}
	t.Logf("src: %s", string(srcData))
	// if string(srcData) != string(dstData) {
	// 	t.Error("dst and src is not equal")
	// }
	diff(t, "dst and src is not equal", dstData, srcData)
}

func TestCopiers_Copy(t *testing.T) {
	type internal1 struct {
		I int
	}
	type internal2 struct {
		I int
	}
	type Embeded struct {
		E string
	}
	type testStruct1 struct {
		Embeded
		S  string
		I  int
		BB []bool
		V  internal1
	}

	type testStruct2 struct {
		Embeded
		S  string
		I  int
		BB []bool
		V  internal2
	}

	src := testStruct1{
		Embeded: Embeded{
			E: "embeded",
		},
		S:  "string",
		I:  10,
		BB: []bool{true, false},
		V:  internal1{I: 5},
	}
	dst := testStruct2{}

	c := New()
	c.Prepare(&dst, &src)
	c.Copy(&dst, &src)
	equal(t, dst, src)
}

func TestCopier_Copy(t *testing.T) {
	type internal1 struct {
		I int
	}
	type internal2 struct {
		I int
	}
	type Embeded struct {
		E string
	}
	type testStruct1 struct {
		Embeded
		S  string
		I  int
		BB []bool
		V  internal1
	}

	type testStruct2 struct {
		Embeded
		S  string
		I  int
		BB []bool
		V  internal2
	}

	src := testStruct1{
		Embeded: Embeded{
			E: "embeded",
		},
		S:  "string",
		I:  10,
		BB: []bool{true, false},
		V:  internal1{I: 5},
	}
	dst := testStruct2{}

	New().Get(&dst, &src).Copy(&dst, &src)
	equal(t, dst, src)
}
