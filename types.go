package copy

import "reflect"

type ValueKind int

const (
	UnknownKind ValueKind = 0
	StructValue ValueKind = 1
	SliceValue
	MapValue
	PtrValue       ValueKind = 0b10
	StructPtrValue ValueKind = StructValue + PtrValue
	SlicePtrValue  ValueKind = SliceValue + PtrValue
	MapPtrValue    ValueKind = MapValue + PtrValue
)

func getValueKind(t reflect.Type) ValueKind {
	var kind ValueKind
	if t.Kind() == reflect.Ptr {
		kind = PtrValue
		t = t.Elem()
	}
	switch k := t.Kind(); {
	case k == reflect.Struct:
		kind += StructValue
	case k == reflect.Slice:
		kind += SliceValue
	case k == reflect.Map && t.Key().Kind() == reflect.String && t.Elem().Kind() == reflect.Interface:
		kind += MapValue
	default:
		return UnknownKind
	}
	return kind
}

func getCopier(c *Copiers, dst, src reflect.Type) internalCopier {
	srcKind := getValueKind(src)
	dstKind := getValueKind(dst)

	switch {
	case srcKind == StructValue && dstKind == StructValue:
		return NewStructCopier(c)
	case srcKind == StructValue && dstKind == StructPtrValue:
		return NewStructToPStructCopier(c)
	case srcKind == StructPtrValue && dstKind == StructValue:
		return NewPStructToStructCopier(c)
	case srcKind == StructPtrValue && dstKind == StructPtrValue:
		return NewPStructToPStructCopier(c)
	}
	return nil
}
