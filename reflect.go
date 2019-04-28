package main

import (
	"errors"
	"reflect"
	"strings"
	"unicode"
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
		err = traverseValue(val, func(s string, v interface{}) {
			var ok bool
			if m, ok = v.(map[string]interface{}); !ok {
				err = errors.New("Type is not struct or map with string keys")
			}
		}, "", opt, 0)
		return
	}
	return nil, errors.New("Type is not struct or map with string keys")
}

func traverseValue(v reflect.Value, f func(string, interface{}), name string, opt traverseOpt, depth int) error {
	v = reflect.Indirect(v)
	depth++

	kind := v.Kind()
	switch kind {
	case reflect.Struct, reflect.Map:
		// Return map as is when key type is not string
		if kind == reflect.Map && v.Type().Key().Kind() != reflect.String {
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

		if !opt.Flat || name == "" {
			m = map[string]interface{}{}
		} else {
			addToMap = f
		}

		if opt.Depth == 0 || depth <= opt.Depth {
			switch kind {
			case reflect.Struct:
				structType := v.Type()
				for i := 0; i < structType.NumField(); i++ {
					structField := structType.Field(i)
					fieldName := structField.Name
					if tag, ok := structField.Tag.Lookup(opt.TagName); ok {
						if s, omit := parseTag(tag); !omit && (s != "") {
							fieldName = s
						}
					}
					if err := traverseValue(v.Field(i), addToMap, addFieldName(name, fieldName, opt.Flat), opt, depth); err != nil {
						return err
					}
				}
			case reflect.Map:
				iter := v.MapRange()
				for iter.Next() {
					key := iter.Key()
					v := iter.Value()
					if err := traverseValue(v, addToMap, addFieldName(name, key.String(), opt.Flat), opt, depth); err != nil {
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
				if err := traverseValue(v.Index(i), addToSlice, "", opt, depth); err != nil {
					return err
				}
			}
		}
		f(name, s)
	default:
		f(name, v.Interface())
	}

	return nil
}
