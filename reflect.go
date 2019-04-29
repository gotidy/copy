package main

import (
	"errors"
	"reflect"
	"strings"
	"sync"
	"unicode"

	null "gopkg.in/guregu/null.v3"
)

func addFieldName(path, field string, flat bool) string {
	if path == "" || !flat {
		return field
	}
	return path + "." + field
}

func isValidTag(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		switch {
		case strings.ContainsRune("!#$%&()*+-./:<=>?@[]^_{|}~ ", c):
		case !unicode.IsLetter(c) && !unicode.IsDigit(c):
			return false
		}
	}
	return true
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

type fieldInfo struct {
	Name string
	Use  bool
	ID   bool
}

type fields []fieldInfo

type structsInfo struct {
	mu sync.RWMutex
	s  map[reflect.Type]fields
}

func (s *structsInfo) getStructFields(t reflect.Type, tagName string) (f fields) {
	s.mu.RLock()
	f, ok := s.s[t]
	s.mu.RUnlock()
	if ok {
		return f
	}
	f = make(fields, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}
		f[i].Name = field.Name
		f[i].Use = true
		if tag, ok := field.Tag.Lookup(tagName); ok {
			s, omit := parseTag(tag)
			if s != "" {
				f[i].Name = s
			}
			f[i].Use = !omit
		}
	}
	s.mu.Lock()
	s.s[t] = f
	s.mu.Unlock()
	return f
}

var structsInfoCache = &structsInfo{s: map[reflect.Type]fields{}}

var valuesProcessor = map[reflect.Type]func(reflect.Value) interface{}{
	reflect.TypeOf(null.Int{}): func(v reflect.Value) interface{} {
		val, _ := (v.Interface()).(null.Int)
		return val.Int64
	},
	reflect.TypeOf(null.Float{}): func(v reflect.Value) interface{} {
		val, _ := (v.Interface()).(null.Float)
		return val.Float64
	},
	reflect.TypeOf(null.Bool{}): func(v reflect.Value) interface{} {
		val, _ := (v.Interface()).(null.Bool)
		return val.Bool
	},
	reflect.TypeOf(null.String{}): func(v reflect.Value) interface{} {
		val, _ := (v.Interface()).(null.String)
		return val.String
	},
	reflect.TypeOf(null.Time{}): func(v reflect.Value) interface{} {
		val, _ := (v.Interface()).(null.Time)
		return val.Time
	},
}

type traverseOpt struct {
	Flat      bool
	Depth     int
	SkipSlice bool
	TagName   string
}

func Traverse(v interface{}, opt traverseOpt) (m map[string]interface{}, err error) {
	val := reflect.ValueOf(v)
	if kind := val.Kind(); kind == reflect.Struct || kind == reflect.Map && val.Type().Key().Kind() == reflect.String {
		//m = map[string]interface{}{}
		err = traverse(val, func(s string, v interface{}) {
			var ok bool
			if m, ok = v.(map[string]interface{}); !ok {
				err = errors.New("Type is not struct or map with string keys")
			}
		}, "", opt, 0)
		return
	}
	return nil, errors.New("Type is not struct or map with string keys")
}

func traverse(v reflect.Value, f func(string, interface{}), name string, opt traverseOpt, depth int) error {
	v = reflect.Indirect(v)
	depth++

	vType := v.Type()
	if p, ok := valuesProcessor[vType]; ok {
		f(name, p(v))
		return nil
	}

	kind := v.Kind()
	switch kind {
	case reflect.Struct, reflect.Map:
		// Return map as is when key type is not string
		if kind == reflect.Map && vType.Key().Kind() != reflect.String {
			// m := reflect.
			// iter := v.MapRange()
			// for iter.Next() {
			// 	key := iter.Key()
			// 	v := iter.Value()
			// 	if err := traverseValue(v, addToMap, addFieldName(name, key.String(), opt.Flat), opt); err != nil {
			// 		return err
			// 	}
			// }

			//f(name, m.Interface())
			return nil
		}

		var m map[string]interface{}
		addToMap := func(name string, v interface{}) {
			m[name] = v
		}

		// if !opt.Flat || name == "" {
		// 	m = make(map[string]interface{}, 0)
		// } else {
		// 	addToMap = f
		// }

		if opt.Depth == 0 || depth <= opt.Depth {
			switch kind {
			case reflect.Struct:
				fields := structsInfoCache.getStructFields(vType, opt.TagName)
				size := len(fields)
				if !opt.Flat || name == "" {
					if opt.Flat && depth <= 1 {
						size += size / 2
					}
					m = make(map[string]interface{}, size)
				} else {
					addToMap = f
				}
				for i := 0; i < len(fields); i++ {
					if fields[i].Use {
						vF := v.Field(i)
						if err := traverse(vF, addToMap, addFieldName(name, fields[i].Name, opt.Flat), opt, depth); err != nil {
							return err
						}
					}
				}
				// structType := v.Type()
				// for i := 0; i < structType.NumField(); i++ {
				// 	structField := structType.Field(i)
				// 	fieldName := structField.Name
				// 	if tag, ok := structField.Tag.Lookup(opt.TagName); ok {
				// 		if s, omit := parseTag(tag); !omit && (s != "") {
				// 			fieldName = s
				// 		}
				// 	}
				// 	if err := traverseValue(v.Field(i), addToMap, addFieldName(name, fieldName, opt.Flat), opt, depth); err != nil {
				// 		return err
				// 	}
				// }
			case reflect.Map:
				if !opt.Flat || name == "" {
					m = make(map[string]interface{}, v.Len()) //map[string]interface{}{}
				} else {
					addToMap = f
				}
				iter := v.MapRange()
				for iter.Next() {
					key := iter.Key()
					v := iter.Value()
					if err := traverse(v, addToMap, addFieldName(name, key.String(), opt.Flat), opt, depth); err != nil {
						return err
					}
				}
			}
		}

		if m != nil {
			f(name, m)
		}
	case reflect.Slice, reflect.Array:
		if opt.SkipSlice {
			return nil
		}

		length := v.Len()
		s := make([]interface{}, 0, length)
		addToSlice := func(name string, v interface{}) {
			s = append(s, v)
		}
		if opt.Depth == 0 || depth <= opt.Depth {
			for i := 0; i < length; i++ {
				if err := traverse(v.Index(i), addToSlice, "", opt, depth); err != nil {
					return err
				}
			}
		}
		f(name, s)
	default:
		vI := v.Interface()
		f(name, vI)
	}

	return nil
}
