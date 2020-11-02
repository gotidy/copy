package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var data = struct {
	Types [][]string
	Sizes []int
}{
	Types: [][]string{
		{
			"int", "int8", "int16", "int32", "int64",
			"uint", "uint8", "uint16", "uint32", "uint64",
		},
		{
			"float32", "float64",
		},
		{
			"bool",
		},
		{
			"complex64", "complex128",
		},
		{
			"string", "[]byte",
		},
		{
			"time.Time",
		},
		{
			"time.Duration",
		},
	},
}

const fileTemplate = `
package funcs

import (
	"reflect"
	"sync"
	"time"
	"unsafe"
)

type funcKey struct {
	Src  reflect.Type
	Dest reflect.Type
}

func typeOf(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}

func typeOfPointer(v interface{}) reflect.Type {
	return reflect.PtrTo(reflect.TypeOf(v))
}

// CopyFuncs is the storage of functions intended for copying data.
type CopyFuncs struct {
	mu    sync.RWMutex
	funcs map[funcKey]func(dst, src unsafe.Pointer)
	sizes []func(dst, src unsafe.Pointer)
}

// Get the copy function for the pair of types, if it is not found then nil is returned.
func (t *CopyFuncs) Get(dst, src reflect.Type) func(dst, src unsafe.Pointer) {
	t.mu.RLock()
	f := t.funcs[funcKey{Src: src, Dest: dst}]
	t.mu.RUnlock()
	if f != nil {
		return f
	}

	if dst.Kind() != src.Kind() {
		return nil
	}

	if dst.Kind() == reflect.String {
		// TODO
		return nil
	}

	same := dst == src

	switch dst.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		same = same || dst.Elem() == src.Elem()
	}

	if same && dst.Size() == src.Size() && src.Size() > 0 && src.Size() <= uintptr(len(t.sizes)) {
		return t.sizes[src.Size()-1]
	}

	return nil
}

// Set the copy function for the pair of types.
func (t *CopyFuncs) Set(dst, src reflect.Type, f func(dst, src unsafe.Pointer)) {
	t.mu.Lock()
	t.funcs[funcKey{Src: src, Dest: dst}] = f
	t.mu.Unlock()
}

// Get the copy function for the pair of types, if it is not found then nil is returned.
func Get(dst, src reflect.Type) func(dst, src unsafe.Pointer) {
	return funcs.Get(dst, src)
}

// Set the copy function for the pair of types.
func Set(dst, src reflect.Type, f func(dst, src unsafe.Pointer)) {
	funcs.Set(dst, src, f)
}

var funcs = &CopyFuncs{
	funcs: map[funcKey]func(dst, src unsafe.Pointer){
		{{- range $types :=.Types}}{{range $dst := $types}}{{range $src := $types}} 
		// {{$src}} to {{$dst}}
		{Src: typeOf({{$src}}({{default $src}})), Dest: typeOf({{$dst}}({{default $dst}}))}:    copy{{title $src}}To{{title $dst}},
		{Src: typeOfPointer({{$src}}({{default $src}})), Dest: typeOf({{$dst}}({{default $dst}}))}:    copyP{{title $src}}To{{title $dst}},
		{Src: typeOf({{$src}}({{default $src}})), Dest: typeOfPointer({{$dst}}({{default $dst}}))}:    copy{{title $src}}ToP{{title $dst}},
		{Src: typeOfPointer({{$src}}({{default $src}})), Dest: typeOfPointer({{$dst}}({{default $dst}}))}:    copyP{{title $src}}ToP{{title $dst}},
		{{- end}}{{end}}{{end}}	
	},
	sizes: []func(dst, src unsafe.Pointer){
		{{range $size := $.Sizes -}}
		copy{{$size}},
		{{- end}} 
	},
}
{{range $types :=.Types}}{{range $dst := $types}}{{range $src := $types}} 

// {{$src}} to {{$dst}}

func copy{{title $src}}To{{title $dst}}(dst, src unsafe.Pointer) {
	*(*{{$dst}})(unsafe.Pointer(dst)) = {{$dst}}(*(*{{$src}})(unsafe.Pointer(src)))
}

func copyP{{title $src}}To{{title $dst}}(dst, src unsafe.Pointer) {
	var v {{$dst}}
	if p := *(**{{$src}})(unsafe.Pointer(src)); p != nil {
		v = {{$dst}}(*p)
	}
	*(*{{$dst}})(unsafe.Pointer(dst)) = v
}

func copy{{title $src}}ToP{{title $dst}}(dst, src unsafe.Pointer) {
	v := {{$dst}}(*(*{{$src}})(unsafe.Pointer(src)))
	p := (**{{$dst}})(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func copyP{{title $src}}ToP{{title $dst}}(dst, src unsafe.Pointer) {
	var v {{$dst}}
	if p := *(**{{$src}})(unsafe.Pointer(src)); p != nil {
		v = {{$dst}}(*p)
	}

	p := (**{{$dst}})(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}
	
{{- end}}{{end}}{{end}}	

// Memcopy funcs
{{- range $size := $.Sizes}}
func copy{{$size}}(dst, src unsafe.Pointer) {
	*(*[{{$size}}]byte)(unsafe.Pointer(dst)) = *(*[{{$size}}]byte)(unsafe.Pointer(src))
}
{{end}} 

`

func title(s string) string {
	if strings.HasPrefix(s, "[]") {
		s = strings.TrimPrefix(s, "[]") + "s"
	}

	if parts := strings.Split(s, "."); len(parts) > 1 {
		s = parts[len(parts)-1]
	}

	return strings.Title(s)
}

func defaultValue(t string) string {
	switch t {
	case "bool":
		return "false"
	case "time.Time":
		return "time.Time{}"
	case "string":
		return `""`
	case "[]byte":
		return "nil"
	default:
		return "0"
	}
}

func createFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf(`unable to open file "%s": %s`, path, err)
	}

	return file
}

func main() {
	for i := 1; i < 256; i++ {
		data.Sizes = append(data.Sizes, i)
	}

	if len(os.Args) == 0 {
		log.Fatal("unable to get executable path, args is empty")
	}

	root := filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0])))

	funcMap := template.FuncMap{
		"title":   title,
		"default": defaultValue,
	}

	tmpl, err := template.New("template").Funcs(funcMap).Parse(fileTemplate)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	unsafeFile := createFile(filepath.Join(root, "funcs", "funcs.go"))
	defer unsafeFile.Close()

	err = tmpl.Execute(unsafeFile, data)
	if err != nil {
		log.Fatalf("execution: %s", err) //nolint
	}
}
