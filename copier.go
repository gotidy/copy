// Package copy provides a copy library to make copying of structs to/from others structs a bit easier.
package copy

import (
	"unsafe"
)

// Copier fills a destination from source.
type Copier struct {
	// srcType reflect.Type
	// dstType reflect.Type
	// srtType unsafe.Pointer
	// dstType unsafe.Pointer
	copiers []fieldCopier
}

// func NewCopier(cache   *cache.Cache, opt Options) Copier {
// 	copier := Copier{}.Init(dst, src interface{})
// }

// Copy copies the contents of src into dst. Dst and src each must be a pointer to struct.
func (c *Copier) Copy(dst, src interface{}) {
	// fmt.Println("Copy.dst", reflect.TypeOf(dst))
	// fmt.Println("Copy.src", reflect.TypeOf(src))

	// fmt.Println("copier.dst", reflect.PtrTo(c.dstType))
	// fmt.Println("copier.src", reflect.PtrTo(c.srcType))

	// if srcType := reflect.TypeOf(src); c.srcType != srcType {
	// 	panic("source expected type «" + c.srcType.String() + "», but has " + srcType.String())
	// }
	// if dstType := reflect.TypeOf(dst); c.dstType != dstType {
	// 	panic("destination expected type «" + c.dstType.String() + "», but has " + dstType.String())
	// }

	dstPtr := ifaceData(dst)
	srcPtr := ifaceData(src)

	c.copy(dstPtr, srcPtr)
}

func (c *Copier) copy(dst, src unsafe.Pointer) {
	for _, c := range c.copiers {
		c(dst, src)
	}
}
