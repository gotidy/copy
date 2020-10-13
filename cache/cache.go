package cache

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type Field struct {
	Type reflect.Type
	Name string
	// Use    bool
	Offset uintptr
}

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

func NewStructInfo(t reflect.Type, tagName string) Struct {
	s := Struct{Fields: make([]Field, 0, t.NumField()), Names: make(map[string]Field, t.NumField())}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}
		fi := Field{
			Type: field.Type,
			Name: field.Name,
			// Use:    true,
			Offset: field.Offset,
		}
		if tagName != "" {
			if tag, ok := field.Tag.Lookup(tagName); ok {
				s, omit := parseTag(tag)
				if omit {
					// fi.Use = false
					continue
				}
				if s != "" {
					fi.Name = s
				}
			}
		}
		s.Fields = append(s.Fields, fi)
		s.Names[fi.Name] = fi
	}
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

type Cache struct {
	mu      sync.RWMutex
	tag     string
	structs map[reflect.Type]Struct
}

func New(tagName string) *Cache {
	return &Cache{tag: tagName, structs: make(map[reflect.Type]Struct)}
}

func (c *Cache) Get(i interface{}) Struct {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return c.GetByType(t)
}

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

	s = NewStructInfo(t, c.tag)
	c.mu.Lock()
	c.structs[t] = s
	c.mu.Unlock()
	return s
}
