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

func equal(t *testing.T, actual, expected interface{}) {
	actualData, err := json.Marshal(actual)
	if err != nil {
		t.Fatalf("equal: %s", err)
	}

	t.Logf("actual: %s", string(actualData))

	expectedData, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("equal: %s", err)
	}

	t.Logf("expected: %s", string(expectedData))

	diff(t, "actual and expected is not equal", actualData, expectedData)
}

func TestCopiers_Copy(t *testing.T) {
	type internal1 struct {
		I int
	}

	type internal2 struct {
		I int
	}

	type Embedded struct {
		E string
	}

	type testStruct1 struct {
		Embedded
		S  string
		I  int
		BB []bool
		V  internal1
	}

	type testStruct2 struct {
		Embedded
		S  string
		I  int
		BB []bool
		V  internal2
	}

	src := testStruct1{
		Embedded: Embedded{
			E: "embedded",
		},
		S:  "string",
		I:  10,
		BB: []bool{true, false},
		V:  internal1{I: 5},
	}
	dst := testStruct2{}

	Prepare(&dst, &src)
	Copy(&dst, &src)
	equal(t, dst, src)
}

func TestCopier_Copy(t *testing.T) {
	type internal1 struct {
		I int
	}

	type internal2 struct {
		I int
	}

	type Embedded struct {
		E string
	}

	type testStruct1 struct {
		Embedded
		S  string
		I  int
		BB []bool
		A  [500]byte
		V  internal1
	}

	type testStruct2 struct {
		Embedded
		S  string
		I  int
		BB []bool
		A  [500]byte
		V  internal2
	}

	src := testStruct1{
		Embedded: Embedded{
			E: "embedded",
		},
		S:  "string",
		I:  10,
		BB: []bool{true, false},
		A:  [500]byte{5},
		V:  internal1{I: 5},
	}
	dst := testStruct2{}

	New().Get(&dst, &src).Copy(&dst, &src)
	equal(t, dst, src)
}

func TestCopier_Struct(t *testing.T) {
	v := struct{}{}
	func() {
		defer func() {
			if recover() == nil {
				t.Error("must panic on non struct source")
			}
		}()
		Copy(&v, new(int))
	}()
	func() {
		defer func() {
			if recover() == nil {
				t.Error("must panic on non struct destination")
			}
		}()
		Copy(new(int), &v)
	}()

	func() {
		defer func() {
			if recover() == nil {
				t.Error("must panic on non struct source")
			}
		}()
		Get(&v, new(int))
	}()
	func() {
		defer func() {
			if recover() == nil {
				t.Error("must panic on non struct destination")
			}
		}()
		Get(new(int), &v)
	}()
}

func TestCopier_Pointer(t *testing.T) {
	v := struct{}{}
	func() {
		defer func() {
			if recover() == nil {
				t.Error("must panic on non pointer source")
			}
		}()
		Copy(&v, v)
	}()
	func() {
		defer func() {
			if recover() == nil {
				t.Error("must panic on non pointer destination")
			}
		}()
		Copy(v, &v)
	}()

	c := New().Get(&v, &v)
	func() {
		defer func() {
			if recover() == nil {
				t.Error("must panic on when source pointer is nil")
			}
		}()
		c.Copy(&v, nil)
	}()
	func() {
		defer func() {
			if recover() == nil {
				t.Error("must panic on when destination pointer is nil")
			}
		}()
		Copy(nil, &v)
	}()
}

func TestCopier_Copy_CheckParams(t *testing.T) {
	v := struct{ i int }{}
	vWrong := struct{}{}

	c := New().Get(&v, &v)
	// Check src parameter
	func() {
		defer func() {
			if recover() == nil {
				t.Error("must panic on when parameters types does not match the copier types")
			}
		}()
		c.Copy(&vWrong, &v)
	}()
	// Check dst parameter
	func() {
		defer func() {
			if recover() == nil {
				t.Error("must panic on when parameters types does not match the copier types")
			}
		}()
		c.Copy(&v, &vWrong)
	}()
}

func TestCopier_Skip(t *testing.T) {
	src := struct{ S string }{S: "string"}
	dst := struct{ S int }{}

	func() {
		defer func() {
			if recover() == nil {
				t.Error("must panic when fields with the same name are of different types")
			}
		}()
		New().Get(&dst, &src)
	}()

	func() {
		defer func() {
			if err := recover(); err != nil {
				t.Errorf("do not must panic with Skip option when fields with the same name are of different types: %s", err)
			}
		}()
		New(Skip()).Get(&dst, &src)
	}()
}

func TestCopier_Expand(t *testing.T) {
	type Expanded struct {
		E string
	}

	type testStruct1 struct {
		Exp Expanded `copy:"+"`
	}

	type testStruct2 struct {
		E string
	}

	src := testStruct1{
		Exp: Expanded{
			E: "Expanded",
		},
	}
	dst := testStruct2{}

	Get(&dst, &src).Copy(&dst, &src)
	if src.Exp.E != dst.E {
		t.Errorf("want «%s» got «%s»", src.Exp.E, dst.E)
	}
}

func TestCopier_StructPtr(t *testing.T) {
	type internal0 struct {
		I int
	}

	type testStruct0 struct {
		V internal0
	}

	type internal1 struct {
		I int
	}

	type testStruct1 struct {
		V internal1
	}

	type internal2 struct {
		I int
	}

	type testStruct2 struct {
		V *internal2
	}

	type internal3 struct {
		I int
	}

	type testStruct3 struct {
		V *internal3
	}

	type test struct {
		src      interface{}
		dst      interface{}
		expected interface{}
	}
	tests := []test{
		// struct -> struct
		{
			src:      &testStruct0{V: internal0{I: 5}},
			dst:      &testStruct1{V: internal1{I: 0}},
			expected: &testStruct1{V: internal1{I: 5}},
		},
		// struct -> *struct
		{
			src:      &testStruct1{V: internal1{I: 5}},
			dst:      &testStruct2{V: nil},
			expected: &testStruct2{V: &internal2{I: 5}},
		},
		{
			src:      &testStruct1{V: internal1{I: 5}},
			dst:      &testStruct2{V: &internal2{I: 0}},
			expected: &testStruct2{V: &internal2{I: 5}},
		},
		// *struct -> struct
		{
			src:      &testStruct2{V: &internal2{I: 5}},
			dst:      &testStruct1{V: internal1{I: 0}},
			expected: &testStruct1{V: internal1{I: 5}},
		},
		{
			src:      &testStruct2{V: nil},
			dst:      &testStruct1{V: internal1{I: 5}},
			expected: &testStruct1{V: internal1{I: 5}},
		},
		// *struct -> *struct
		{
			src:      &testStruct2{V: nil},
			dst:      &testStruct3{V: nil},
			expected: &testStruct3{V: nil},
		},
		{
			src:      &testStruct2{V: nil},
			dst:      &testStruct3{V: &internal3{I: 5}},
			expected: &testStruct3{V: &internal3{I: 5}},
		},
		{
			src:      &testStruct2{V: &internal2{I: 5}},
			dst:      &testStruct3{V: &internal3{I: 0}},
			expected: &testStruct3{V: &internal3{I: 5}},
		},
		{
			src:      &testStruct2{V: &internal2{I: 5}},
			dst:      &testStruct3{V: nil},
			expected: &testStruct3{V: &internal3{I: 5}},
		},
	}

	for _, test := range tests {
		New().Get(test.dst, test.src).Copy(test.dst, test.src)
		equal(t, test.dst, test.expected)
		// if !reflect.DeepEqual(&test.dst, &test.expected) {
		// 	t.Errorf("Expected «%» actual ")
		// }
	}
}
