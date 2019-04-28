package main

import (
	"fmt"
	"testing"
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
		{"1000000", 1000000, false},
		{"1000000", 1000000, true},
	}

	z := Zzz{
		Name:      "Vasia",
		Count:     11,
		Hide:      true,
		ID:        112,
		List:      []string{"book", "bug"},
		Internal:  InternalZzz{Left: "L", Right: "R"},
		Internal2: &InternalZzz{Left: "L", Right: "R"},
	}
	opt := traverseOpt{
		Flat:      true, //false,
		Depth:     0,
		SkipSlice: false,
		TagName:   "json",
	}

	for _, test := range tests {
		name := fmt.Sprintf("Count: %10v, Flat: %-6v", test.count, test.flat)
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
