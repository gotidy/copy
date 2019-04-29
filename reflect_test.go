package main

import (
	"fmt"
	"reflect"
	"testing"

	null "gopkg.in/guregu/null.v3"
)

func BenchmarkToMap(b *testing.B) {
	tests := []struct {
		name  string
		count int
		flat  bool
	}{
		{"10000", 10000, false},
		{"10000", 10000, true},
		{"100000", 100000, false},
		{"100000", 100000, true},
		// {"1000000", 1000000, false},
		// {"1000000", 1000000, true},
	}

	z := Zzz{
		Name:       "Vasia",
		Count:      11,
		Hide:       true,
		ID:         112,
		List:       []string{"book", "bug"},
		Internal:   InternalZzz{Left: "L", Right: "R"},
		Internal2:  &InternalZzz{Left: "L", Right: "R"},
		NullString: null.StringFrom("Sss"),
		NullBool:   null.BoolFrom(true),
		NullInt:    null.IntFrom(100),
		NullFloat:  null.FloatFrom(1.23),
	}
	opt := traverseOpt{
		Flat:      false,
		Depth:     0,
		SkipSlice: false,
		TagName:   "json",
	}

	for _, test := range tests {
		name := fmt.Sprintf("Count: %10v, Flat: %-6v", test.count, test.flat)
		opt.Flat = test.flat
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for i := 0; i < test.count; i++ {
					_, err := Traverse(z, opt)
					if err != nil {
						b.Fatal(err)
					}
				}
			}
		})
	}
}

var result interface{}

func BenchmarkValue(b *testing.B) {

	type Ttt struct {
		String string
		Null   bool
	}
	t := Ttt{String: "Zzzzzzzzzzzzzzzzzzzzzz", Null: false}
	v := reflect.ValueOf(t)
	//list := []reflect.Value{v, v, v, v, v, v, v, v}

	var r bool

	b.Run("Inteface", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for i := 0; i < 10000; i++ {
				val, _ := (v.Interface()).(Ttt)
				ii := interface{}(val.String)
				r = r != (ii == nil)
			}
		}
	})

	result = r

	b.Run("Value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for i := 0; i < 10000; i++ {
				val := v.Field(0)
				ii := val.Interface()
				r = r != (ii == nil)
			}
		}
	})

	result = r

}
