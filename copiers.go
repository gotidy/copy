// Package copy provides a copy library to make copying of structs to/from others structs a bit easier.
package copy

import (
	"fmt"
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

type indirectCopierKey struct {
	Src  Type
	Dest Type
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
}

type internalCopier interface {
	Copier
	copy(dst, src unsafe.Pointer)
	init(dst, src reflect.Type)
}

// Copiers is a structs copier.
type Copiers struct {
	cache   *cache.Cache
	options Options

	mu              sync.RWMutex
	copiers         map[copierKey]internalCopier
	indirectCopiers map[indirectCopierKey]Copier
}

// New create new internalCopier.
func New(options ...Option) *Copiers {
	var opts Options

	for _, option := range options {
		option(&opts)
	}

	return &Copiers{
		cache:           cache.New(opts.Tag),
		options:         opts,
		copiers:         make(map[copierKey]internalCopier),
		indirectCopiers: make(map[indirectCopierKey]Copier),
	}
}

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
	c.Get(dst, src).Copy(dst, src)
}

func (c *Copiers) get(dst, src reflect.Type) internalCopier {
	copier, ok := c.copiers[copierKey{Src: src, Dest: dst}]
	if ok {
		return copier
	}

	copier = getCopier(c, dst, src)
	if copier == nil {
		panic(fmt.Sprintf("the combination of destination(%s) and source(%s) types is not supported", dst, src))
	}

	c.copiers[copierKey{Src: src, Dest: dst}] = copier

	copier.init(dst, src)

	return copier
}

// Get Copier for a specific destination and source.
func (c *Copiers) Get(dst, src interface{}) Copier {
	c.mu.RLock()
	copier, ok := c.indirectCopiers[indirectCopierKey{Dest: TypeOf(dst), Src: TypeOf(src)}]
	c.mu.RUnlock()
	if ok {
		return copier
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	srcType := reflect.TypeOf(src)
	if srcType.Kind() != reflect.Ptr {
		panic("source must be pointer")
	}
	srcType = srcType.Elem()

	dstType := reflect.TypeOf(dst)
	if dstType.Kind() != reflect.Ptr {
		panic("destination must be pointer")
	}
	dstType = dstType.Elem()

	copier = c.get(dstType, srcType)

	c.indirectCopiers[indirectCopierKey{Dest: TypeOf(dst), Src: TypeOf(src)}] = copier

	return copier
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
