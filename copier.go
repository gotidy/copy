// Package copy provides a copy library to make copying of structs to/from others structs a bit easier.
package copy

import (
	"reflect"
	"unsafe"
)

// Copier fills a destination from source.
type Copier struct {
	srcType     Type
	srcTypeName string
	dstType     Type
	dstTypeName string
	copiers     []fieldCopier
}

// func NewCopier(cache   *cache.Cache, opt Options) Copier {
// 	copier := Copier{}.Init(dst, src interface{})
// }

// Copy copies the contents of src into dst. Dst and src each must be a pointer to struct.
func (c *Copier) Copy(dst, src interface{}) {
	dstType, dstPtr := DataOf(dst)
	srcType, srcPtr := DataOf(src)

	if c.srcType != srcType {
		panic("source expected type " + c.srcTypeName + ", but has " + reflect.TypeOf(src).String())
	}
	if c.dstType != dstType {
		panic("destination expected type " + c.dstTypeName + ", but has " + reflect.TypeOf(dst).String())
	}

	c.copy(dstPtr, srcPtr)
}

func (c *Copier) copy(dst, src unsafe.Pointer) {
	for _, c := range c.copiers {
		c(dst, src)
	}
}
