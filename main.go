package main

import (
	"fmt"
	"log"

	"github.com/kr/pretty"
)

type InternalZzz struct {
	Left  string `json:"left,omitempty"`
	Right string `json:"right,omitempty"`
}

type Zzz struct {
	Name      string       `json:"name,omitempty"`
	Count     int          `json:"count,omitempty"`
	Hide      bool         `json:"hide,omitempty"`
	ID        int64        `json:"id,omitempty"`
	List      []string     `json:"list,omitempty"`
	Internal  InternalZzz  `json:"internal,omitempty"`
	Internal2 *InternalZzz `json:"internal2,omitempty"`
}

func main() {
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
	m, err := Traverse(z, opt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%# v", pretty.Formatter(m))
}
