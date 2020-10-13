package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Types struct {
	zero string
	list []string
}

var types = [][]string{
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

type CopyFuncs struct {
	mu    sync.RWMutex
	funcs map[funcKey]func(src, dest unsafe.Pointer)
}

func (t *CopyFuncs) Get(src, dest reflect.Type) func(src, dest unsafe.Pointer) {
	t.mu.RLock()
	f := t.funcs[funcKey{Src: src, Dest: dest}]
	t.mu.RUnlock()
	return f
}

func (t *CopyFuncs) Set(src, dest reflect.Type, f func(src, dest unsafe.Pointer)) {
	t.mu.Lock()
	t.funcs[funcKey{Src: src, Dest: dest}] = f
	t.mu.Unlock()
}

func Get(src, dest reflect.Type) func(src, dest unsafe.Pointer) {
	return funcs.Get(src, dest)
}

func Set(src, dest reflect.Type, f func(src, dest unsafe.Pointer)) {
	funcs.Set(src, dest, f)
}

var funcs = &CopyFuncs{
	funcs: map[funcKey]func(src, dest unsafe.Pointer){
		{{- range $types :=.}}{{range $dest := $types}}{{range $src := $types}} 
		// {{$src}} to {{$dest}}
		{Src: typeOf({{$src}}({{default $src}})), Dest: typeOf({{$dest}}({{default $dest}}))}:    Copy{{title $src}}To{{title $dest}},
		{Src: typeOfPointer({{$src}}({{default $src}})), Dest: typeOf({{$dest}}({{default $dest}}))}:    CopyP{{title $src}}To{{title $dest}},
		{Src: typeOf({{$src}}({{default $src}})), Dest: typeOfPointer({{$dest}}({{default $dest}}))}:    Copy{{title $src}}ToP{{title $dest}},
		{Src: typeOfPointer({{$src}}({{default $src}})), Dest: typeOfPointer({{$dest}}({{default $dest}}))}:    CopyP{{title $src}}ToP{{title $dest}},
		{{- end}}{{end}}{{end}}	
	},
}
{{range $types :=.}}{{range $dest := $types}}{{range $src := $types}} 

// {{$src}} to {{$dest}}

func Copy{{title $src}}To{{title $dest}}(src, dest unsafe.Pointer) {
	*(*{{$dest}})(unsafe.Pointer(dest)) = {{$dest}}(*(*{{$src}})(unsafe.Pointer(src)))
}

func CopyP{{title $src}}To{{title $dest}}(src, dest unsafe.Pointer) {
	var v {{$dest}}
	if p := *(**{{$src}})(unsafe.Pointer(src)); p != nil {
		v = {{$dest}}(*p)
	}
	*(*{{$dest}})(unsafe.Pointer(dest)) = v
}

func Copy{{title $src}}ToP{{title $dest}}(src, dest unsafe.Pointer) {
	v := {{$dest}}(*(*{{$src}})(unsafe.Pointer(src)))
	p := (**{{$dest}})(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyP{{title $src}}ToP{{title $dest}}(src, dest unsafe.Pointer) {
	var v {{$dest}}
	if p := *(**{{$src}})(unsafe.Pointer(src)); p != nil {
		v = {{$dest}}(*p)
	}

	p := (**{{$dest}})(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}
	
{{- end}}{{end}}{{end}}	

`

func Title(s string) string {
	if strings.HasPrefix(s, "[]") {
		s = strings.TrimPrefix(s, "[]") + "s"
	}
	if parts := strings.Split(s, "."); len(parts) > 1 {
		s = parts[len(parts)-1]
	}
	return strings.Title(s)
}

func Default(t string) string {
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

func Create(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf(`unable to open file "%s": %s`, path, err)
	}
	return file
}

func main() {
	if len(os.Args) == 0 {
		log.Fatal("unable to get executable path, args is empty")
	}
	root := filepath.Dir(filepath.Dir(os.Args[0]))
	unsafeFile := Create(filepath.Join(root, "funcs", "funcs.go"))
	defer unsafeFile.Close()

	funcMap := template.FuncMap{
		"title":   Title,
		"default": Default,
	}

	tmpl, err := template.New("template").Funcs(funcMap).Parse(fileTemplate)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	err = tmpl.Execute(unsafeFile, types)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

}
