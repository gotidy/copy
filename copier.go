// Package copy provides a copy library to make copying of structs to/from others structs a bit easier.
package copy

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"

	"github.com/gotidy/copy/funcs"
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

// Copiers is a structs copier.
type Copiers struct {
	cache   *cache.Cache
	options Options

	mu      sync.RWMutex
	copiers map[copierKey]Copier
}

// New create new Copier.
func New(options ...Option) *Copiers {
	var opts Options

	for _, option := range options {
		option(&opts)
	}

	return &Copiers{cache: cache.New(opts.Tag), options: opts, copiers: make(map[copierKey]Copier)}
}

type fieldCopier = func(dst, src unsafe.Pointer)

func (c *Copiers) fieldCopier(dst, src cache.Field) fieldCopier {
	dstOffset := dst.Offset
	srcOffset := src.Offset
	copier := funcs.Get(dst.Type, src.Type)
	if copier != nil {
		return func(dstPtr, srcPtr unsafe.Pointer) {
			copier(unsafe.Pointer(uintptr(dstPtr)+dstOffset), unsafe.Pointer(uintptr(srcPtr)+srcOffset))
		}
	}

	// same type -> same type
	if src.Type == dst.Type {
		size := int(src.Type.Size())

		return func(dstPtr, srcPtr unsafe.Pointer) {
			// More safe and independent from internal structs
			// src := reflect.NewAt(src.Type, unsafe.Pointer(uintptr(srcPtr)+src.Offset)).Elem()
			// dst := reflect.NewAt(dst.Type, unsafe.Pointer(uintptr(dstPtr)+dst.Offset)).Elem()
			// dst.Set(src)
			memcopy(unsafe.Pointer(uintptr(dstPtr)+dst.Offset), unsafe.Pointer(uintptr(srcPtr)+src.Offset), size)
		}
	}

	// struct -> struct
	if src.Type.Kind() == reflect.Struct && dst.Type.Kind() == reflect.Struct {
		copier := c.get(dst.Type, src.Type)

		return func(dstPtr, srcPtr unsafe.Pointer) {
			copier.copy(unsafe.Pointer(uintptr(dstPtr)+dst.Offset), unsafe.Pointer(uintptr(srcPtr)+src.Offset))
		}
	}

	// *struct -> struct
	if src.Type.Kind() == reflect.Ptr && src.Type.Elem().Kind() == reflect.Struct && dst.Type.Kind() == reflect.Struct {
		copier := c.get(dst.Type, src.Type.Elem())

		return func(dstPtr, srcPtr unsafe.Pointer) {
			srcFieldPtr := (**struct{})(unsafe.Pointer(uintptr(srcPtr) + src.Offset))
			if *srcFieldPtr == nil {
				return
			}
			copier.copy(unsafe.Pointer(uintptr(dstPtr)+dst.Offset), unsafe.Pointer(*srcFieldPtr))
		}
	}

	// struct -> *struct
	if src.Type.Kind() == reflect.Struct && dst.Type.Kind() == reflect.Ptr && dst.Type.Elem().Kind() == reflect.Struct {
		copier := c.get(dst.Type.Elem(), src.Type)

		dstSize := int(src.Type.Elem().Size())

		return func(dstPtr, srcPtr unsafe.Pointer) {
			dstFieldPtr := (**struct{})(unsafe.Pointer(uintptr(dstPtr) + dst.Offset))
			if *dstFieldPtr == nil {
				*dstFieldPtr = (*struct{})(alloc(dstSize))
			}

			copier.copy(unsafe.Pointer(*dstFieldPtr), unsafe.Pointer(uintptr(srcPtr)+src.Offset))
		}
	}

	// *struct -> *struct
	if src.Type.Kind() == reflect.Ptr && src.Type.Elem().Kind() == reflect.Struct &&
		dst.Type.Kind() == reflect.Ptr && dst.Type.Elem().Kind() == reflect.Struct {
		copier := c.get(dst.Type.Elem(), src.Type.Elem())

		dstSize := int(src.Type.Elem().Size())

		return func(dstPtr, srcPtr unsafe.Pointer) {
			srcFieldPtr := (**struct{})(unsafe.Pointer(uintptr(srcPtr) + src.Offset))
			if *srcFieldPtr == nil {
				return
			}
			dstFieldPtr := (**struct{})(unsafe.Pointer(uintptr(dstPtr) + dst.Offset))
			if dstFieldPtr == nil {
				*dstFieldPtr = (*struct{})(alloc(dstSize))
			}

			copier.copy(unsafe.Pointer(*dstFieldPtr), unsafe.Pointer(*srcFieldPtr))
		}
	}

	if !c.options.Skip {
		panic(fmt.Errorf(`field «%s» of type «%s» is not assignable to field «%s» of type «%s»`, src.Name, src.Type.String(), dst.Name, dst.Type.String()))
	}

	return nil
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

	srcStruct := c.cache.GetByType(src)
	dstStruct := c.cache.GetByType(dst)

	for i := 0; i < srcStruct.NumField(); i++ {
		srcField := srcStruct.Field(i)
		if dstField, ok := dstStruct.FieldByName(srcField.Name); ok {
			if f := c.fieldCopier(dstField, srcField); f != nil {
				copier.copiers = append(copier.copiers, f)
			}
		}
	}
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

// Copier fills a destination from source.
type Copier struct {
	copiers []fieldCopier
}

// Copy copies the contents of src into dst. Dst and src each must be a pointer to struct.
func (c Copier) Copy(dst, src interface{}) {
	// More safe and independent from internal structs
	// srcValue := reflect.ValueOf(src)
	// if srcValue.Kind() != reflect.Ptr {
	// 	panic("source must be pointer to struct")
	// }
	// srcPtr := unsafe.Pointer(srcValue.Pointer())
	// srcValue = srcValue.Elem()
	// dstValue := reflect.ValueOf(dst)
	// if dstValue.Kind() != reflect.Ptr {
	// 	panic("source must be pointer to struct")
	// }
	// dstPtr := unsafe.Pointer(dstValue.Pointer())
	// dstValue = dstValue.Elem()
	dstPtr := ifaceToPtr(dst)
	srcPtr := ifaceToPtr(src)

	c.copy(dstPtr, srcPtr)
}

func (c Copier) copy(dst, src unsafe.Pointer) {
	for _, c := range c.copiers {
		c(dst, src)
	}
}

// defaultCopier uses Copier with a "copy" tag.
var defaultCopier = New(Tag(defaultTagName))

// Prepare caches structures of src and dst.  Dst and src each must be a pointer to struct.
// contents is not copied. It can be used for checking ability of copying.
//
//   copy.Prepare(&dst, &src)
func Prepare(dst, src interface{}) {
	defaultCopier.Prepare(src, dst)
}

// Copy copies the contents of src into dst. Dst and src each must be a pointer to a struct.
func Copy(dst, src interface{}) {
	defaultCopier.Copy(src, dst)
}

// Get Copier for a specific destination and source.
func Get(dst, src interface{}) Copier {
	return defaultCopier.Get(src, dst)
}

// More safe and independent from internal structs
// func ifaceToPtr(i interface{}) unsafe.Pointer {
// 	v := reflect.ValueOf(i)
// 	if v.Kind() != reflect.Ptr {
// 		panic("source must be pointer to struct")
// 	}
// 	return unsafe.Pointer(v.Pointer())
// }

func ifaceToPtr(i interface{}) unsafe.Pointer {
	if i == nil {
		panic("input parameter is nil")
	}

	type iface struct {
		Type, Data unsafe.Pointer
	}

	return (*iface)(unsafe.Pointer(&i)).Data
}

func memcopy(dst, src unsafe.Pointer, size int) {
	var srcSlice []byte
	srcSH := (*reflect.SliceHeader)((unsafe.Pointer(&srcSlice)))
	srcSH.Data = uintptr(src)
	srcSH.Cap = int(size)
	srcSH.Len = int(size)

	var dstSlice []byte
	dstSH := (*reflect.SliceHeader)((unsafe.Pointer(&dstSlice)))
	dstSH.Data = uintptr(dst)
	dstSH.Cap = int(size)
	dstSH.Len = int(size)

	copy(dstSlice, srcSlice)
}

func alloc(size int) unsafe.Pointer {
	size = (size + 7) / 8 // size in int64
	return unsafe.Pointer(&(make([]int64, size)[0]))
}
