package copy

import (
	"reflect"
	"sync"
	"unsafe"

	"github.com/gotidy/copy/cache"
)

const defaultTagName = "copy"

type funcKey struct {
	Src  reflect.Type
	Dest reflect.Type
}

// Copier is structs copier
type Copier struct {
	cache *cache.Cache

	mu      sync.RWMutex
	copiers map[funcKey][]fieldCopier
}

// New create new Copier
func New(tagName string) *Copier {
	return &Copier{cache: cache.New(tagName), copiers: make(map[funcKey][]fieldCopier)}
}

func (c *Copier) getFieldsCopiers(src, dest reflect.Type) []fieldCopier {
	c.mu.RLock()
	fc, ok := c.copiers[funcKey{Src: src, Dest: dest}]
	c.mu.RUnlock()
	if ok {
		return fc
	}

	srcStruct := c.cache.GetByType(src)
	destStruct := c.cache.GetByType(dest)
	for i := 0; i < srcStruct.NumField(); i++ {
		srcField := srcStruct.Field(i)
		if destField, ok := destStruct.FieldByName(srcField.Name); ok {
			if f := getFieldCopier(srcField, destField); f != nil {
				fc = append(fc, f)
			}
		}
	}
	c.mu.Lock()
	c.copiers[funcKey{Src: src, Dest: dest}] = fc
	c.mu.Unlock()
	return fc
}

// Prefetch caches structures of src and dest.  Dst and src each must be a pointer to struct.
// contents is not copied. It can be used for ability of coping.
//
// c := copy.New("")
// c.Prefetch(&src, &dest)
func (c *Copier) Prefetch(src, dest interface{}) {
	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() != reflect.Ptr {
		panic("source must be pointer to struct")
	}
	srcValue = srcValue.Elem()
	destValue := reflect.ValueOf(dest)
	if destValue.Kind() != reflect.Ptr {
		panic("source must be pointer to struct")
	}
	destValue = destValue.Elem()

	_ = c.getFieldsCopiers(srcValue.Type(), destValue.Type())
}

// Copy copies the contents of src into dst. Dst and src each must be a pointer to struct.
func (c *Copier) Copy(src, dest interface{}) {
	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() != reflect.Ptr {
		panic("source must be pointer to struct")
	}
	srcPtr := unsafe.Pointer(srcValue.Pointer())
	srcValue = srcValue.Elem()
	destValue := reflect.ValueOf(dest)
	if destValue.Kind() != reflect.Ptr {
		panic("source must be pointer to struct")
	}
	destPtr := unsafe.Pointer(destValue.Pointer())
	destValue = destValue.Elem()

	fc := c.getFieldsCopiers(srcValue.Type(), destValue.Type())

	for _, c := range fc {
		c(srcPtr, destPtr)
	}
}

var DefaultCopier = New(defaultTagName)

// Prefetch caches structures of src and dest.  Dst and src each must be a pointer to struct.
// contents is not copied. It can be used for ability of coping.
//
// copy.Prefetch(&src, &dest)
func Prefetch(src, dest interface{}) {
	DefaultCopier.Prefetch(src, dest)
}

// Copy copies the contents of src into dst. Dst and src each must be a pointer to struct.
func Copy(src, dest interface{}) {
	DefaultCopier.Copy(src, dest)
}
