// Package copy provides a copy library to make copying of structs to/from others structs a bit easier.
package copy

import (
	"reflect"
	"sync"
	"unsafe"

	"github.com/gotidy/copy/internal/cache"
)

const defaultTagName = "copy"

type copierKey struct {
	Src  reflect.Type
	Dest reflect.Type
}

// Options is Copiers parameters.
type Options struct {
	Tag  string
	Skip bool
}

// Option changes default Copiers parameters.
type Option func(c *Options)

// Tag set tag name.
func Tag(tag string) Option {
	return func(o *Options) {
		o.Tag = tag
	}
}

// Skip nonassignable types else cause panic.
func Skip() Option {
	return func(o *Options) {
		o.Skip = true
	}
}

// StructCopier fills a destination from source.
type Copier interface {
	Copy(dst interface{}, src interface{})
	copy(dst, src unsafe.Pointer)
}

// Copiers is a structs copier.
type Copiers struct {
	cache   *cache.Cache
	options Options

	mu      sync.RWMutex
	copiers map[copierKey]*StructCopier
}

// New create new Copier.
func New(options ...Option) *Copiers {
	var opts Options

	for _, option := range options {
		option(&opts)
	}

	return &Copiers{cache: cache.New(opts.Tag), options: opts, copiers: make(map[copierKey]*StructCopier)}
}

type fieldCopier = func(dst, src unsafe.Pointer)

// Prepare caches structures of src and dst. Dst and src each must be a pointer to struct.
// contents is not copied. It can be used for checking ability of copying.
//
//   c := copy.New()
//   c.Prepare(&dst, &src)
func (c *Copiers) Prepare(dst, src interface{}) {
	_ = c.Get(dst, src)
}

// Copy copies the contents of src into dst. Dst and src each must be a pointer to struct.
func (c *Copiers) Copy(dst, src interface{}) {
	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() != reflect.Ptr {
		panic("source must be pointer to struct")
	}
	srcPtr := unsafe.Pointer(srcValue.Pointer())
	srcValue = srcValue.Elem()
	if srcValue.Kind() != reflect.Struct {
		panic("source must be pointer to struct")
	}

	dstValue := reflect.ValueOf(dst)
	if dstValue.Kind() != reflect.Ptr {
		panic("destination must be pointer to struct")
	}
	dstPtr := unsafe.Pointer(dstValue.Pointer())
	dstValue = dstValue.Elem()
	if dstValue.Kind() != reflect.Struct {
		panic("destination must be pointer to struct")
	}

	copier := c.get(dstValue.Type(), srcValue.Type())

	copier.copy(dstPtr, srcPtr)
}

func (c *Copiers) get(dst, src reflect.Type) Copier {
	c.mu.RLock()
	copier, ok := c.copiers[copierKey{Src: src, Dest: dst}]
	c.mu.RUnlock()
	if ok {
		return copier
	}

	copier = NewStructCopier(c, dst, src)

	c.mu.Lock()
	c.copiers[copierKey{Src: src, Dest: dst}] = copier
	c.mu.Unlock()

	return copier
}

// Get Copier for a specific destination and source.
func (c *Copiers) Get(dst, src interface{}) Copier {
	srcValue := reflect.Indirect(reflect.ValueOf(src))
	if srcValue.Kind() != reflect.Struct {
		panic("source must be struct")
	}

	dstValue := reflect.Indirect(reflect.ValueOf(dst))
	if dstValue.Kind() != reflect.Struct {
		panic("destination must be struct")
	}

	return c.get(dstValue.Type(), srcValue.Type())
}

// defaultCopier uses Copier with a "copy" tag.
var defaultCopier = New(Tag(defaultTagName))

// Prepare caches structures of src and dst.  Dst and src each must be a pointer to struct.
// contents is not copied. It can be used for checking ability of copying.
//
//   copy.Prepare(&dst, &src)
func Prepare(dst, src interface{}) {
	defaultCopier.Prepare(dst, src)
}

// Copy copies the contents of src into dst. Dst and src each must be a pointer to a struct.
func Copy(dst, src interface{}) {
	defaultCopier.Copy(dst, src)
}

// Get Copier for a specific destination and source.
func Get(dst, src interface{}) Copier {
	return defaultCopier.Get(dst, src)
}
