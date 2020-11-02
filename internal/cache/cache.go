package cache

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// Field info.
type Field struct {
	Type       reflect.Type
	Name       string
	Anonymous  bool
	Offset     uintptr
	ParentName string
}

// Struct fields info.
type Struct struct {
	Fields []Field
	Names  map[string]Field
}

func parseTag(tag string) (name string, omit bool) {
	if tag == "-" {
		return "", true
	}

	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], false
	}

	return tag, false
}

// NewStruct inits the new struct info.
func NewStruct(t reflect.Type, tagName string) Struct {
	s := Struct{Fields: make([]Field, 0, t.NumField()), Names: make(map[string]Field, t.NumField())}

	var traverse func(t reflect.Type, name string, offset uintptr)
	traverse = func(t reflect.Type, name string, offset uintptr) {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.PkgPath != "" {
				continue
			}

			fi := Field{
				Type:       field.Type,
				Name:       field.Name,
				Offset:     field.Offset + offset,
				Anonymous:  field.Anonymous && field.Type.Kind() == reflect.Struct,
				ParentName: name,
			}

			if tagName != "" {
				if tag, ok := field.Tag.Lookup(tagName); ok {
					s, omit := parseTag(tag)
					if omit {
						continue
					}

					if s != "" {
						fi.Name = s
					}
				}
			}

			s.Fields = append(s.Fields, fi)
			s.Names[fi.Name] = fi

			if fi.Anonymous {
				traverse(fi.Type, fi.Name+".", fi.Offset)
			}
		}
	}
	traverse(t, "", 0)

	return s
}

// Field returns a struct type's i'th field.
// It panics if i is not in the range [0, NumField()).
func (s Struct) Field(i int) Field {
	return s.Fields[i]
}

// FieldByName returns the struct field with the given name
// and a boolean indicating if the field was found.
func (s Struct) FieldByName(name string) (Field, bool) {
	f, ok := s.Names[name]
	return f, ok
}

// NumField returns a struct type's field count.
// It panics if the type's Kind is not Struct.
func (s Struct) NumField() int {
	return len(s.Fields)
}

// Cache is structs' cache.
type Cache struct {
	mu      sync.RWMutex
	tag     string
	structs map[reflect.Type]Struct
}

// New creates structs Cache.
func New(tagName string) *Cache {
	return &Cache{tag: tagName, structs: make(map[reflect.Type]Struct)}
}

// Get returns struct fields info.
func (c *Cache) Get(i interface{}) Struct {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return c.GetByType(t)
}

// GetByType returns struct fields info.
func (c *Cache) GetByType(t reflect.Type) Struct {
	c.mu.RLock()
	s, ok := c.structs[t]
	c.mu.RUnlock()
	if ok {
		return s
	}

	if t.Kind() != reflect.Struct {
		panic(fmt.Errorf("type %s is not struct", t))
	}

	s = NewStruct(t, c.tag)
	c.mu.Lock()
	c.structs[t] = s
	c.mu.Unlock()

	return s
}
