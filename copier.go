package copy

import "reflect"

type TypeInfo struct {
	Type Type
	Name string
}

func NewTypeInfo(typ reflect.Type) TypeInfo {
	p := reflect.New(typ)
	return TypeInfo{Type: TypeOf(p.Interface()), Name: p.String()}
}

func (t TypeInfo) Check(typ Type) bool {
	return t.Type != typ
}

type BaseCopier struct {
	*Copiers

	dst TypeInfo
	src TypeInfo
}

func NewBaseCopier(c *Copiers) BaseCopier {
	return BaseCopier{Copiers: c}
}

func (b *BaseCopier) init(dst, src reflect.Type) {
	b.dst = NewTypeInfo(dst)
	b.src = NewTypeInfo(src)
}
