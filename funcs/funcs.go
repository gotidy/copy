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
	funcs map[funcKey]func(dst, src unsafe.Pointer)
	sizes []func(dst, src unsafe.Pointer)
}

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

func (t *CopyFuncs) Set(dst, src reflect.Type, f func(dst, src unsafe.Pointer)) {
	t.mu.Lock()
	t.funcs[funcKey{Src: src, Dest: dst}] = f
	t.mu.Unlock()
}

func Get(dst, src reflect.Type) func(dst, src unsafe.Pointer) {
	return funcs.Get(dst, src)
}

func Set(dst, src reflect.Type, f func(dst, src unsafe.Pointer)) {
	funcs.Set(dst, src, f)
}

var funcs = &CopyFuncs{
	funcs: map[funcKey]func(dst, src unsafe.Pointer){
		// int to int
		{Src: typeOf(int(0)), Dest: typeOf(int(0))}:               CopyIntToInt,
		{Src: typeOfPointer(int(0)), Dest: typeOf(int(0))}:        CopyPIntToInt,
		{Src: typeOf(int(0)), Dest: typeOfPointer(int(0))}:        CopyIntToPInt,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(int(0))}: CopyPIntToPInt,
		// int8 to int
		{Src: typeOf(int8(0)), Dest: typeOf(int(0))}:               CopyInt8ToInt,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(int(0))}:        CopyPInt8ToInt,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(int(0))}:        CopyInt8ToPInt,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(int(0))}: CopyPInt8ToPInt,
		// int16 to int
		{Src: typeOf(int16(0)), Dest: typeOf(int(0))}:               CopyInt16ToInt,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(int(0))}:        CopyPInt16ToInt,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(int(0))}:        CopyInt16ToPInt,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(int(0))}: CopyPInt16ToPInt,
		// int32 to int
		{Src: typeOf(int32(0)), Dest: typeOf(int(0))}:               CopyInt32ToInt,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(int(0))}:        CopyPInt32ToInt,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(int(0))}:        CopyInt32ToPInt,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(int(0))}: CopyPInt32ToPInt,
		// int64 to int
		{Src: typeOf(int64(0)), Dest: typeOf(int(0))}:               CopyInt64ToInt,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(int(0))}:        CopyPInt64ToInt,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(int(0))}:        CopyInt64ToPInt,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(int(0))}: CopyPInt64ToPInt,
		// uint to int
		{Src: typeOf(uint(0)), Dest: typeOf(int(0))}:               CopyUintToInt,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(int(0))}:        CopyPUintToInt,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(int(0))}:        CopyUintToPInt,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(int(0))}: CopyPUintToPInt,
		// uint8 to int
		{Src: typeOf(uint8(0)), Dest: typeOf(int(0))}:               CopyUint8ToInt,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(int(0))}:        CopyPUint8ToInt,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(int(0))}:        CopyUint8ToPInt,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(int(0))}: CopyPUint8ToPInt,
		// uint16 to int
		{Src: typeOf(uint16(0)), Dest: typeOf(int(0))}:               CopyUint16ToInt,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(int(0))}:        CopyPUint16ToInt,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(int(0))}:        CopyUint16ToPInt,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(int(0))}: CopyPUint16ToPInt,
		// uint32 to int
		{Src: typeOf(uint32(0)), Dest: typeOf(int(0))}:               CopyUint32ToInt,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(int(0))}:        CopyPUint32ToInt,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(int(0))}:        CopyUint32ToPInt,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(int(0))}: CopyPUint32ToPInt,
		// uint64 to int
		{Src: typeOf(uint64(0)), Dest: typeOf(int(0))}:               CopyUint64ToInt,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(int(0))}:        CopyPUint64ToInt,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(int(0))}:        CopyUint64ToPInt,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(int(0))}: CopyPUint64ToPInt,
		// int to int8
		{Src: typeOf(int(0)), Dest: typeOf(int8(0))}:               CopyIntToInt8,
		{Src: typeOfPointer(int(0)), Dest: typeOf(int8(0))}:        CopyPIntToInt8,
		{Src: typeOf(int(0)), Dest: typeOfPointer(int8(0))}:        CopyIntToPInt8,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(int8(0))}: CopyPIntToPInt8,
		// int8 to int8
		{Src: typeOf(int8(0)), Dest: typeOf(int8(0))}:               CopyInt8ToInt8,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(int8(0))}:        CopyPInt8ToInt8,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(int8(0))}:        CopyInt8ToPInt8,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(int8(0))}: CopyPInt8ToPInt8,
		// int16 to int8
		{Src: typeOf(int16(0)), Dest: typeOf(int8(0))}:               CopyInt16ToInt8,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(int8(0))}:        CopyPInt16ToInt8,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(int8(0))}:        CopyInt16ToPInt8,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(int8(0))}: CopyPInt16ToPInt8,
		// int32 to int8
		{Src: typeOf(int32(0)), Dest: typeOf(int8(0))}:               CopyInt32ToInt8,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(int8(0))}:        CopyPInt32ToInt8,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(int8(0))}:        CopyInt32ToPInt8,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(int8(0))}: CopyPInt32ToPInt8,
		// int64 to int8
		{Src: typeOf(int64(0)), Dest: typeOf(int8(0))}:               CopyInt64ToInt8,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(int8(0))}:        CopyPInt64ToInt8,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(int8(0))}:        CopyInt64ToPInt8,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(int8(0))}: CopyPInt64ToPInt8,
		// uint to int8
		{Src: typeOf(uint(0)), Dest: typeOf(int8(0))}:               CopyUintToInt8,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(int8(0))}:        CopyPUintToInt8,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(int8(0))}:        CopyUintToPInt8,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(int8(0))}: CopyPUintToPInt8,
		// uint8 to int8
		{Src: typeOf(uint8(0)), Dest: typeOf(int8(0))}:               CopyUint8ToInt8,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(int8(0))}:        CopyPUint8ToInt8,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(int8(0))}:        CopyUint8ToPInt8,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(int8(0))}: CopyPUint8ToPInt8,
		// uint16 to int8
		{Src: typeOf(uint16(0)), Dest: typeOf(int8(0))}:               CopyUint16ToInt8,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(int8(0))}:        CopyPUint16ToInt8,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(int8(0))}:        CopyUint16ToPInt8,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(int8(0))}: CopyPUint16ToPInt8,
		// uint32 to int8
		{Src: typeOf(uint32(0)), Dest: typeOf(int8(0))}:               CopyUint32ToInt8,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(int8(0))}:        CopyPUint32ToInt8,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(int8(0))}:        CopyUint32ToPInt8,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(int8(0))}: CopyPUint32ToPInt8,
		// uint64 to int8
		{Src: typeOf(uint64(0)), Dest: typeOf(int8(0))}:               CopyUint64ToInt8,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(int8(0))}:        CopyPUint64ToInt8,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(int8(0))}:        CopyUint64ToPInt8,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(int8(0))}: CopyPUint64ToPInt8,
		// int to int16
		{Src: typeOf(int(0)), Dest: typeOf(int16(0))}:               CopyIntToInt16,
		{Src: typeOfPointer(int(0)), Dest: typeOf(int16(0))}:        CopyPIntToInt16,
		{Src: typeOf(int(0)), Dest: typeOfPointer(int16(0))}:        CopyIntToPInt16,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(int16(0))}: CopyPIntToPInt16,
		// int8 to int16
		{Src: typeOf(int8(0)), Dest: typeOf(int16(0))}:               CopyInt8ToInt16,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(int16(0))}:        CopyPInt8ToInt16,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(int16(0))}:        CopyInt8ToPInt16,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(int16(0))}: CopyPInt8ToPInt16,
		// int16 to int16
		{Src: typeOf(int16(0)), Dest: typeOf(int16(0))}:               CopyInt16ToInt16,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(int16(0))}:        CopyPInt16ToInt16,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(int16(0))}:        CopyInt16ToPInt16,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(int16(0))}: CopyPInt16ToPInt16,
		// int32 to int16
		{Src: typeOf(int32(0)), Dest: typeOf(int16(0))}:               CopyInt32ToInt16,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(int16(0))}:        CopyPInt32ToInt16,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(int16(0))}:        CopyInt32ToPInt16,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(int16(0))}: CopyPInt32ToPInt16,
		// int64 to int16
		{Src: typeOf(int64(0)), Dest: typeOf(int16(0))}:               CopyInt64ToInt16,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(int16(0))}:        CopyPInt64ToInt16,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(int16(0))}:        CopyInt64ToPInt16,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(int16(0))}: CopyPInt64ToPInt16,
		// uint to int16
		{Src: typeOf(uint(0)), Dest: typeOf(int16(0))}:               CopyUintToInt16,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(int16(0))}:        CopyPUintToInt16,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(int16(0))}:        CopyUintToPInt16,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(int16(0))}: CopyPUintToPInt16,
		// uint8 to int16
		{Src: typeOf(uint8(0)), Dest: typeOf(int16(0))}:               CopyUint8ToInt16,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(int16(0))}:        CopyPUint8ToInt16,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(int16(0))}:        CopyUint8ToPInt16,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(int16(0))}: CopyPUint8ToPInt16,
		// uint16 to int16
		{Src: typeOf(uint16(0)), Dest: typeOf(int16(0))}:               CopyUint16ToInt16,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(int16(0))}:        CopyPUint16ToInt16,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(int16(0))}:        CopyUint16ToPInt16,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(int16(0))}: CopyPUint16ToPInt16,
		// uint32 to int16
		{Src: typeOf(uint32(0)), Dest: typeOf(int16(0))}:               CopyUint32ToInt16,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(int16(0))}:        CopyPUint32ToInt16,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(int16(0))}:        CopyUint32ToPInt16,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(int16(0))}: CopyPUint32ToPInt16,
		// uint64 to int16
		{Src: typeOf(uint64(0)), Dest: typeOf(int16(0))}:               CopyUint64ToInt16,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(int16(0))}:        CopyPUint64ToInt16,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(int16(0))}:        CopyUint64ToPInt16,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(int16(0))}: CopyPUint64ToPInt16,
		// int to int32
		{Src: typeOf(int(0)), Dest: typeOf(int32(0))}:               CopyIntToInt32,
		{Src: typeOfPointer(int(0)), Dest: typeOf(int32(0))}:        CopyPIntToInt32,
		{Src: typeOf(int(0)), Dest: typeOfPointer(int32(0))}:        CopyIntToPInt32,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(int32(0))}: CopyPIntToPInt32,
		// int8 to int32
		{Src: typeOf(int8(0)), Dest: typeOf(int32(0))}:               CopyInt8ToInt32,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(int32(0))}:        CopyPInt8ToInt32,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(int32(0))}:        CopyInt8ToPInt32,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(int32(0))}: CopyPInt8ToPInt32,
		// int16 to int32
		{Src: typeOf(int16(0)), Dest: typeOf(int32(0))}:               CopyInt16ToInt32,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(int32(0))}:        CopyPInt16ToInt32,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(int32(0))}:        CopyInt16ToPInt32,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(int32(0))}: CopyPInt16ToPInt32,
		// int32 to int32
		{Src: typeOf(int32(0)), Dest: typeOf(int32(0))}:               CopyInt32ToInt32,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(int32(0))}:        CopyPInt32ToInt32,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(int32(0))}:        CopyInt32ToPInt32,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(int32(0))}: CopyPInt32ToPInt32,
		// int64 to int32
		{Src: typeOf(int64(0)), Dest: typeOf(int32(0))}:               CopyInt64ToInt32,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(int32(0))}:        CopyPInt64ToInt32,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(int32(0))}:        CopyInt64ToPInt32,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(int32(0))}: CopyPInt64ToPInt32,
		// uint to int32
		{Src: typeOf(uint(0)), Dest: typeOf(int32(0))}:               CopyUintToInt32,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(int32(0))}:        CopyPUintToInt32,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(int32(0))}:        CopyUintToPInt32,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(int32(0))}: CopyPUintToPInt32,
		// uint8 to int32
		{Src: typeOf(uint8(0)), Dest: typeOf(int32(0))}:               CopyUint8ToInt32,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(int32(0))}:        CopyPUint8ToInt32,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(int32(0))}:        CopyUint8ToPInt32,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(int32(0))}: CopyPUint8ToPInt32,
		// uint16 to int32
		{Src: typeOf(uint16(0)), Dest: typeOf(int32(0))}:               CopyUint16ToInt32,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(int32(0))}:        CopyPUint16ToInt32,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(int32(0))}:        CopyUint16ToPInt32,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(int32(0))}: CopyPUint16ToPInt32,
		// uint32 to int32
		{Src: typeOf(uint32(0)), Dest: typeOf(int32(0))}:               CopyUint32ToInt32,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(int32(0))}:        CopyPUint32ToInt32,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(int32(0))}:        CopyUint32ToPInt32,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(int32(0))}: CopyPUint32ToPInt32,
		// uint64 to int32
		{Src: typeOf(uint64(0)), Dest: typeOf(int32(0))}:               CopyUint64ToInt32,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(int32(0))}:        CopyPUint64ToInt32,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(int32(0))}:        CopyUint64ToPInt32,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(int32(0))}: CopyPUint64ToPInt32,
		// int to int64
		{Src: typeOf(int(0)), Dest: typeOf(int64(0))}:               CopyIntToInt64,
		{Src: typeOfPointer(int(0)), Dest: typeOf(int64(0))}:        CopyPIntToInt64,
		{Src: typeOf(int(0)), Dest: typeOfPointer(int64(0))}:        CopyIntToPInt64,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(int64(0))}: CopyPIntToPInt64,
		// int8 to int64
		{Src: typeOf(int8(0)), Dest: typeOf(int64(0))}:               CopyInt8ToInt64,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(int64(0))}:        CopyPInt8ToInt64,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(int64(0))}:        CopyInt8ToPInt64,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(int64(0))}: CopyPInt8ToPInt64,
		// int16 to int64
		{Src: typeOf(int16(0)), Dest: typeOf(int64(0))}:               CopyInt16ToInt64,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(int64(0))}:        CopyPInt16ToInt64,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(int64(0))}:        CopyInt16ToPInt64,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(int64(0))}: CopyPInt16ToPInt64,
		// int32 to int64
		{Src: typeOf(int32(0)), Dest: typeOf(int64(0))}:               CopyInt32ToInt64,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(int64(0))}:        CopyPInt32ToInt64,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(int64(0))}:        CopyInt32ToPInt64,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(int64(0))}: CopyPInt32ToPInt64,
		// int64 to int64
		{Src: typeOf(int64(0)), Dest: typeOf(int64(0))}:               CopyInt64ToInt64,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(int64(0))}:        CopyPInt64ToInt64,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(int64(0))}:        CopyInt64ToPInt64,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(int64(0))}: CopyPInt64ToPInt64,
		// uint to int64
		{Src: typeOf(uint(0)), Dest: typeOf(int64(0))}:               CopyUintToInt64,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(int64(0))}:        CopyPUintToInt64,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(int64(0))}:        CopyUintToPInt64,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(int64(0))}: CopyPUintToPInt64,
		// uint8 to int64
		{Src: typeOf(uint8(0)), Dest: typeOf(int64(0))}:               CopyUint8ToInt64,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(int64(0))}:        CopyPUint8ToInt64,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(int64(0))}:        CopyUint8ToPInt64,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(int64(0))}: CopyPUint8ToPInt64,
		// uint16 to int64
		{Src: typeOf(uint16(0)), Dest: typeOf(int64(0))}:               CopyUint16ToInt64,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(int64(0))}:        CopyPUint16ToInt64,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(int64(0))}:        CopyUint16ToPInt64,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(int64(0))}: CopyPUint16ToPInt64,
		// uint32 to int64
		{Src: typeOf(uint32(0)), Dest: typeOf(int64(0))}:               CopyUint32ToInt64,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(int64(0))}:        CopyPUint32ToInt64,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(int64(0))}:        CopyUint32ToPInt64,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(int64(0))}: CopyPUint32ToPInt64,
		// uint64 to int64
		{Src: typeOf(uint64(0)), Dest: typeOf(int64(0))}:               CopyUint64ToInt64,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(int64(0))}:        CopyPUint64ToInt64,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(int64(0))}:        CopyUint64ToPInt64,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(int64(0))}: CopyPUint64ToPInt64,
		// int to uint
		{Src: typeOf(int(0)), Dest: typeOf(uint(0))}:               CopyIntToUint,
		{Src: typeOfPointer(int(0)), Dest: typeOf(uint(0))}:        CopyPIntToUint,
		{Src: typeOf(int(0)), Dest: typeOfPointer(uint(0))}:        CopyIntToPUint,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(uint(0))}: CopyPIntToPUint,
		// int8 to uint
		{Src: typeOf(int8(0)), Dest: typeOf(uint(0))}:               CopyInt8ToUint,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(uint(0))}:        CopyPInt8ToUint,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(uint(0))}:        CopyInt8ToPUint,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(uint(0))}: CopyPInt8ToPUint,
		// int16 to uint
		{Src: typeOf(int16(0)), Dest: typeOf(uint(0))}:               CopyInt16ToUint,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(uint(0))}:        CopyPInt16ToUint,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(uint(0))}:        CopyInt16ToPUint,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(uint(0))}: CopyPInt16ToPUint,
		// int32 to uint
		{Src: typeOf(int32(0)), Dest: typeOf(uint(0))}:               CopyInt32ToUint,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(uint(0))}:        CopyPInt32ToUint,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(uint(0))}:        CopyInt32ToPUint,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(uint(0))}: CopyPInt32ToPUint,
		// int64 to uint
		{Src: typeOf(int64(0)), Dest: typeOf(uint(0))}:               CopyInt64ToUint,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(uint(0))}:        CopyPInt64ToUint,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(uint(0))}:        CopyInt64ToPUint,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(uint(0))}: CopyPInt64ToPUint,
		// uint to uint
		{Src: typeOf(uint(0)), Dest: typeOf(uint(0))}:               CopyUintToUint,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(uint(0))}:        CopyPUintToUint,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(uint(0))}:        CopyUintToPUint,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(uint(0))}: CopyPUintToPUint,
		// uint8 to uint
		{Src: typeOf(uint8(0)), Dest: typeOf(uint(0))}:               CopyUint8ToUint,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(uint(0))}:        CopyPUint8ToUint,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(uint(0))}:        CopyUint8ToPUint,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(uint(0))}: CopyPUint8ToPUint,
		// uint16 to uint
		{Src: typeOf(uint16(0)), Dest: typeOf(uint(0))}:               CopyUint16ToUint,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(uint(0))}:        CopyPUint16ToUint,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(uint(0))}:        CopyUint16ToPUint,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(uint(0))}: CopyPUint16ToPUint,
		// uint32 to uint
		{Src: typeOf(uint32(0)), Dest: typeOf(uint(0))}:               CopyUint32ToUint,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(uint(0))}:        CopyPUint32ToUint,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(uint(0))}:        CopyUint32ToPUint,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(uint(0))}: CopyPUint32ToPUint,
		// uint64 to uint
		{Src: typeOf(uint64(0)), Dest: typeOf(uint(0))}:               CopyUint64ToUint,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(uint(0))}:        CopyPUint64ToUint,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(uint(0))}:        CopyUint64ToPUint,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(uint(0))}: CopyPUint64ToPUint,
		// int to uint8
		{Src: typeOf(int(0)), Dest: typeOf(uint8(0))}:               CopyIntToUint8,
		{Src: typeOfPointer(int(0)), Dest: typeOf(uint8(0))}:        CopyPIntToUint8,
		{Src: typeOf(int(0)), Dest: typeOfPointer(uint8(0))}:        CopyIntToPUint8,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(uint8(0))}: CopyPIntToPUint8,
		// int8 to uint8
		{Src: typeOf(int8(0)), Dest: typeOf(uint8(0))}:               CopyInt8ToUint8,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(uint8(0))}:        CopyPInt8ToUint8,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(uint8(0))}:        CopyInt8ToPUint8,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(uint8(0))}: CopyPInt8ToPUint8,
		// int16 to uint8
		{Src: typeOf(int16(0)), Dest: typeOf(uint8(0))}:               CopyInt16ToUint8,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(uint8(0))}:        CopyPInt16ToUint8,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(uint8(0))}:        CopyInt16ToPUint8,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(uint8(0))}: CopyPInt16ToPUint8,
		// int32 to uint8
		{Src: typeOf(int32(0)), Dest: typeOf(uint8(0))}:               CopyInt32ToUint8,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(uint8(0))}:        CopyPInt32ToUint8,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(uint8(0))}:        CopyInt32ToPUint8,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(uint8(0))}: CopyPInt32ToPUint8,
		// int64 to uint8
		{Src: typeOf(int64(0)), Dest: typeOf(uint8(0))}:               CopyInt64ToUint8,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(uint8(0))}:        CopyPInt64ToUint8,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(uint8(0))}:        CopyInt64ToPUint8,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(uint8(0))}: CopyPInt64ToPUint8,
		// uint to uint8
		{Src: typeOf(uint(0)), Dest: typeOf(uint8(0))}:               CopyUintToUint8,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(uint8(0))}:        CopyPUintToUint8,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(uint8(0))}:        CopyUintToPUint8,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(uint8(0))}: CopyPUintToPUint8,
		// uint8 to uint8
		{Src: typeOf(uint8(0)), Dest: typeOf(uint8(0))}:               CopyUint8ToUint8,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(uint8(0))}:        CopyPUint8ToUint8,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(uint8(0))}:        CopyUint8ToPUint8,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(uint8(0))}: CopyPUint8ToPUint8,
		// uint16 to uint8
		{Src: typeOf(uint16(0)), Dest: typeOf(uint8(0))}:               CopyUint16ToUint8,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(uint8(0))}:        CopyPUint16ToUint8,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(uint8(0))}:        CopyUint16ToPUint8,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(uint8(0))}: CopyPUint16ToPUint8,
		// uint32 to uint8
		{Src: typeOf(uint32(0)), Dest: typeOf(uint8(0))}:               CopyUint32ToUint8,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(uint8(0))}:        CopyPUint32ToUint8,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(uint8(0))}:        CopyUint32ToPUint8,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(uint8(0))}: CopyPUint32ToPUint8,
		// uint64 to uint8
		{Src: typeOf(uint64(0)), Dest: typeOf(uint8(0))}:               CopyUint64ToUint8,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(uint8(0))}:        CopyPUint64ToUint8,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(uint8(0))}:        CopyUint64ToPUint8,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(uint8(0))}: CopyPUint64ToPUint8,
		// int to uint16
		{Src: typeOf(int(0)), Dest: typeOf(uint16(0))}:               CopyIntToUint16,
		{Src: typeOfPointer(int(0)), Dest: typeOf(uint16(0))}:        CopyPIntToUint16,
		{Src: typeOf(int(0)), Dest: typeOfPointer(uint16(0))}:        CopyIntToPUint16,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(uint16(0))}: CopyPIntToPUint16,
		// int8 to uint16
		{Src: typeOf(int8(0)), Dest: typeOf(uint16(0))}:               CopyInt8ToUint16,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(uint16(0))}:        CopyPInt8ToUint16,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(uint16(0))}:        CopyInt8ToPUint16,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(uint16(0))}: CopyPInt8ToPUint16,
		// int16 to uint16
		{Src: typeOf(int16(0)), Dest: typeOf(uint16(0))}:               CopyInt16ToUint16,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(uint16(0))}:        CopyPInt16ToUint16,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(uint16(0))}:        CopyInt16ToPUint16,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(uint16(0))}: CopyPInt16ToPUint16,
		// int32 to uint16
		{Src: typeOf(int32(0)), Dest: typeOf(uint16(0))}:               CopyInt32ToUint16,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(uint16(0))}:        CopyPInt32ToUint16,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(uint16(0))}:        CopyInt32ToPUint16,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(uint16(0))}: CopyPInt32ToPUint16,
		// int64 to uint16
		{Src: typeOf(int64(0)), Dest: typeOf(uint16(0))}:               CopyInt64ToUint16,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(uint16(0))}:        CopyPInt64ToUint16,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(uint16(0))}:        CopyInt64ToPUint16,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(uint16(0))}: CopyPInt64ToPUint16,
		// uint to uint16
		{Src: typeOf(uint(0)), Dest: typeOf(uint16(0))}:               CopyUintToUint16,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(uint16(0))}:        CopyPUintToUint16,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(uint16(0))}:        CopyUintToPUint16,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(uint16(0))}: CopyPUintToPUint16,
		// uint8 to uint16
		{Src: typeOf(uint8(0)), Dest: typeOf(uint16(0))}:               CopyUint8ToUint16,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(uint16(0))}:        CopyPUint8ToUint16,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(uint16(0))}:        CopyUint8ToPUint16,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(uint16(0))}: CopyPUint8ToPUint16,
		// uint16 to uint16
		{Src: typeOf(uint16(0)), Dest: typeOf(uint16(0))}:               CopyUint16ToUint16,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(uint16(0))}:        CopyPUint16ToUint16,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(uint16(0))}:        CopyUint16ToPUint16,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(uint16(0))}: CopyPUint16ToPUint16,
		// uint32 to uint16
		{Src: typeOf(uint32(0)), Dest: typeOf(uint16(0))}:               CopyUint32ToUint16,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(uint16(0))}:        CopyPUint32ToUint16,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(uint16(0))}:        CopyUint32ToPUint16,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(uint16(0))}: CopyPUint32ToPUint16,
		// uint64 to uint16
		{Src: typeOf(uint64(0)), Dest: typeOf(uint16(0))}:               CopyUint64ToUint16,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(uint16(0))}:        CopyPUint64ToUint16,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(uint16(0))}:        CopyUint64ToPUint16,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(uint16(0))}: CopyPUint64ToPUint16,
		// int to uint32
		{Src: typeOf(int(0)), Dest: typeOf(uint32(0))}:               CopyIntToUint32,
		{Src: typeOfPointer(int(0)), Dest: typeOf(uint32(0))}:        CopyPIntToUint32,
		{Src: typeOf(int(0)), Dest: typeOfPointer(uint32(0))}:        CopyIntToPUint32,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(uint32(0))}: CopyPIntToPUint32,
		// int8 to uint32
		{Src: typeOf(int8(0)), Dest: typeOf(uint32(0))}:               CopyInt8ToUint32,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(uint32(0))}:        CopyPInt8ToUint32,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(uint32(0))}:        CopyInt8ToPUint32,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(uint32(0))}: CopyPInt8ToPUint32,
		// int16 to uint32
		{Src: typeOf(int16(0)), Dest: typeOf(uint32(0))}:               CopyInt16ToUint32,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(uint32(0))}:        CopyPInt16ToUint32,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(uint32(0))}:        CopyInt16ToPUint32,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(uint32(0))}: CopyPInt16ToPUint32,
		// int32 to uint32
		{Src: typeOf(int32(0)), Dest: typeOf(uint32(0))}:               CopyInt32ToUint32,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(uint32(0))}:        CopyPInt32ToUint32,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(uint32(0))}:        CopyInt32ToPUint32,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(uint32(0))}: CopyPInt32ToPUint32,
		// int64 to uint32
		{Src: typeOf(int64(0)), Dest: typeOf(uint32(0))}:               CopyInt64ToUint32,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(uint32(0))}:        CopyPInt64ToUint32,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(uint32(0))}:        CopyInt64ToPUint32,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(uint32(0))}: CopyPInt64ToPUint32,
		// uint to uint32
		{Src: typeOf(uint(0)), Dest: typeOf(uint32(0))}:               CopyUintToUint32,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(uint32(0))}:        CopyPUintToUint32,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(uint32(0))}:        CopyUintToPUint32,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(uint32(0))}: CopyPUintToPUint32,
		// uint8 to uint32
		{Src: typeOf(uint8(0)), Dest: typeOf(uint32(0))}:               CopyUint8ToUint32,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(uint32(0))}:        CopyPUint8ToUint32,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(uint32(0))}:        CopyUint8ToPUint32,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(uint32(0))}: CopyPUint8ToPUint32,
		// uint16 to uint32
		{Src: typeOf(uint16(0)), Dest: typeOf(uint32(0))}:               CopyUint16ToUint32,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(uint32(0))}:        CopyPUint16ToUint32,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(uint32(0))}:        CopyUint16ToPUint32,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(uint32(0))}: CopyPUint16ToPUint32,
		// uint32 to uint32
		{Src: typeOf(uint32(0)), Dest: typeOf(uint32(0))}:               CopyUint32ToUint32,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(uint32(0))}:        CopyPUint32ToUint32,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(uint32(0))}:        CopyUint32ToPUint32,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(uint32(0))}: CopyPUint32ToPUint32,
		// uint64 to uint32
		{Src: typeOf(uint64(0)), Dest: typeOf(uint32(0))}:               CopyUint64ToUint32,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(uint32(0))}:        CopyPUint64ToUint32,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(uint32(0))}:        CopyUint64ToPUint32,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(uint32(0))}: CopyPUint64ToPUint32,
		// int to uint64
		{Src: typeOf(int(0)), Dest: typeOf(uint64(0))}:               CopyIntToUint64,
		{Src: typeOfPointer(int(0)), Dest: typeOf(uint64(0))}:        CopyPIntToUint64,
		{Src: typeOf(int(0)), Dest: typeOfPointer(uint64(0))}:        CopyIntToPUint64,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(uint64(0))}: CopyPIntToPUint64,
		// int8 to uint64
		{Src: typeOf(int8(0)), Dest: typeOf(uint64(0))}:               CopyInt8ToUint64,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(uint64(0))}:        CopyPInt8ToUint64,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(uint64(0))}:        CopyInt8ToPUint64,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(uint64(0))}: CopyPInt8ToPUint64,
		// int16 to uint64
		{Src: typeOf(int16(0)), Dest: typeOf(uint64(0))}:               CopyInt16ToUint64,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(uint64(0))}:        CopyPInt16ToUint64,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(uint64(0))}:        CopyInt16ToPUint64,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(uint64(0))}: CopyPInt16ToPUint64,
		// int32 to uint64
		{Src: typeOf(int32(0)), Dest: typeOf(uint64(0))}:               CopyInt32ToUint64,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(uint64(0))}:        CopyPInt32ToUint64,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(uint64(0))}:        CopyInt32ToPUint64,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(uint64(0))}: CopyPInt32ToPUint64,
		// int64 to uint64
		{Src: typeOf(int64(0)), Dest: typeOf(uint64(0))}:               CopyInt64ToUint64,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(uint64(0))}:        CopyPInt64ToUint64,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(uint64(0))}:        CopyInt64ToPUint64,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(uint64(0))}: CopyPInt64ToPUint64,
		// uint to uint64
		{Src: typeOf(uint(0)), Dest: typeOf(uint64(0))}:               CopyUintToUint64,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(uint64(0))}:        CopyPUintToUint64,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(uint64(0))}:        CopyUintToPUint64,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(uint64(0))}: CopyPUintToPUint64,
		// uint8 to uint64
		{Src: typeOf(uint8(0)), Dest: typeOf(uint64(0))}:               CopyUint8ToUint64,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(uint64(0))}:        CopyPUint8ToUint64,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(uint64(0))}:        CopyUint8ToPUint64,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(uint64(0))}: CopyPUint8ToPUint64,
		// uint16 to uint64
		{Src: typeOf(uint16(0)), Dest: typeOf(uint64(0))}:               CopyUint16ToUint64,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(uint64(0))}:        CopyPUint16ToUint64,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(uint64(0))}:        CopyUint16ToPUint64,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(uint64(0))}: CopyPUint16ToPUint64,
		// uint32 to uint64
		{Src: typeOf(uint32(0)), Dest: typeOf(uint64(0))}:               CopyUint32ToUint64,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(uint64(0))}:        CopyPUint32ToUint64,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(uint64(0))}:        CopyUint32ToPUint64,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(uint64(0))}: CopyPUint32ToPUint64,
		// uint64 to uint64
		{Src: typeOf(uint64(0)), Dest: typeOf(uint64(0))}:               CopyUint64ToUint64,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(uint64(0))}:        CopyPUint64ToUint64,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(uint64(0))}:        CopyUint64ToPUint64,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(uint64(0))}: CopyPUint64ToPUint64,
		// float32 to float32
		{Src: typeOf(float32(0)), Dest: typeOf(float32(0))}:               CopyFloat32ToFloat32,
		{Src: typeOfPointer(float32(0)), Dest: typeOf(float32(0))}:        CopyPFloat32ToFloat32,
		{Src: typeOf(float32(0)), Dest: typeOfPointer(float32(0))}:        CopyFloat32ToPFloat32,
		{Src: typeOfPointer(float32(0)), Dest: typeOfPointer(float32(0))}: CopyPFloat32ToPFloat32,
		// float64 to float32
		{Src: typeOf(float64(0)), Dest: typeOf(float32(0))}:               CopyFloat64ToFloat32,
		{Src: typeOfPointer(float64(0)), Dest: typeOf(float32(0))}:        CopyPFloat64ToFloat32,
		{Src: typeOf(float64(0)), Dest: typeOfPointer(float32(0))}:        CopyFloat64ToPFloat32,
		{Src: typeOfPointer(float64(0)), Dest: typeOfPointer(float32(0))}: CopyPFloat64ToPFloat32,
		// float32 to float64
		{Src: typeOf(float32(0)), Dest: typeOf(float64(0))}:               CopyFloat32ToFloat64,
		{Src: typeOfPointer(float32(0)), Dest: typeOf(float64(0))}:        CopyPFloat32ToFloat64,
		{Src: typeOf(float32(0)), Dest: typeOfPointer(float64(0))}:        CopyFloat32ToPFloat64,
		{Src: typeOfPointer(float32(0)), Dest: typeOfPointer(float64(0))}: CopyPFloat32ToPFloat64,
		// float64 to float64
		{Src: typeOf(float64(0)), Dest: typeOf(float64(0))}:               CopyFloat64ToFloat64,
		{Src: typeOfPointer(float64(0)), Dest: typeOf(float64(0))}:        CopyPFloat64ToFloat64,
		{Src: typeOf(float64(0)), Dest: typeOfPointer(float64(0))}:        CopyFloat64ToPFloat64,
		{Src: typeOfPointer(float64(0)), Dest: typeOfPointer(float64(0))}: CopyPFloat64ToPFloat64,
		// bool to bool
		{Src: typeOf(bool(false)), Dest: typeOf(bool(false))}:               CopyBoolToBool,
		{Src: typeOfPointer(bool(false)), Dest: typeOf(bool(false))}:        CopyPBoolToBool,
		{Src: typeOf(bool(false)), Dest: typeOfPointer(bool(false))}:        CopyBoolToPBool,
		{Src: typeOfPointer(bool(false)), Dest: typeOfPointer(bool(false))}: CopyPBoolToPBool,
		// complex64 to complex64
		{Src: typeOf(complex64(0)), Dest: typeOf(complex64(0))}:               CopyComplex64ToComplex64,
		{Src: typeOfPointer(complex64(0)), Dest: typeOf(complex64(0))}:        CopyPComplex64ToComplex64,
		{Src: typeOf(complex64(0)), Dest: typeOfPointer(complex64(0))}:        CopyComplex64ToPComplex64,
		{Src: typeOfPointer(complex64(0)), Dest: typeOfPointer(complex64(0))}: CopyPComplex64ToPComplex64,
		// complex128 to complex64
		{Src: typeOf(complex128(0)), Dest: typeOf(complex64(0))}:               CopyComplex128ToComplex64,
		{Src: typeOfPointer(complex128(0)), Dest: typeOf(complex64(0))}:        CopyPComplex128ToComplex64,
		{Src: typeOf(complex128(0)), Dest: typeOfPointer(complex64(0))}:        CopyComplex128ToPComplex64,
		{Src: typeOfPointer(complex128(0)), Dest: typeOfPointer(complex64(0))}: CopyPComplex128ToPComplex64,
		// complex64 to complex128
		{Src: typeOf(complex64(0)), Dest: typeOf(complex128(0))}:               CopyComplex64ToComplex128,
		{Src: typeOfPointer(complex64(0)), Dest: typeOf(complex128(0))}:        CopyPComplex64ToComplex128,
		{Src: typeOf(complex64(0)), Dest: typeOfPointer(complex128(0))}:        CopyComplex64ToPComplex128,
		{Src: typeOfPointer(complex64(0)), Dest: typeOfPointer(complex128(0))}: CopyPComplex64ToPComplex128,
		// complex128 to complex128
		{Src: typeOf(complex128(0)), Dest: typeOf(complex128(0))}:               CopyComplex128ToComplex128,
		{Src: typeOfPointer(complex128(0)), Dest: typeOf(complex128(0))}:        CopyPComplex128ToComplex128,
		{Src: typeOf(complex128(0)), Dest: typeOfPointer(complex128(0))}:        CopyComplex128ToPComplex128,
		{Src: typeOfPointer(complex128(0)), Dest: typeOfPointer(complex128(0))}: CopyPComplex128ToPComplex128,
		// string to string
		{Src: typeOf(string("")), Dest: typeOf(string(""))}:               CopyStringToString,
		{Src: typeOfPointer(string("")), Dest: typeOf(string(""))}:        CopyPStringToString,
		{Src: typeOf(string("")), Dest: typeOfPointer(string(""))}:        CopyStringToPString,
		{Src: typeOfPointer(string("")), Dest: typeOfPointer(string(""))}: CopyPStringToPString,
		// []byte to string
		{Src: typeOf([]byte(nil)), Dest: typeOf(string(""))}:               CopyBytesToString,
		{Src: typeOfPointer([]byte(nil)), Dest: typeOf(string(""))}:        CopyPBytesToString,
		{Src: typeOf([]byte(nil)), Dest: typeOfPointer(string(""))}:        CopyBytesToPString,
		{Src: typeOfPointer([]byte(nil)), Dest: typeOfPointer(string(""))}: CopyPBytesToPString,
		// string to []byte
		{Src: typeOf(string("")), Dest: typeOf([]byte(nil))}:               CopyStringToBytes,
		{Src: typeOfPointer(string("")), Dest: typeOf([]byte(nil))}:        CopyPStringToBytes,
		{Src: typeOf(string("")), Dest: typeOfPointer([]byte(nil))}:        CopyStringToPBytes,
		{Src: typeOfPointer(string("")), Dest: typeOfPointer([]byte(nil))}: CopyPStringToPBytes,
		// []byte to []byte
		{Src: typeOf([]byte(nil)), Dest: typeOf([]byte(nil))}:               CopyBytesToBytes,
		{Src: typeOfPointer([]byte(nil)), Dest: typeOf([]byte(nil))}:        CopyPBytesToBytes,
		{Src: typeOf([]byte(nil)), Dest: typeOfPointer([]byte(nil))}:        CopyBytesToPBytes,
		{Src: typeOfPointer([]byte(nil)), Dest: typeOfPointer([]byte(nil))}: CopyPBytesToPBytes,
		// time.Time to time.Time
		{Src: typeOf(time.Time(time.Time{})), Dest: typeOf(time.Time(time.Time{}))}:               CopyTimeToTime,
		{Src: typeOfPointer(time.Time(time.Time{})), Dest: typeOf(time.Time(time.Time{}))}:        CopyPTimeToTime,
		{Src: typeOf(time.Time(time.Time{})), Dest: typeOfPointer(time.Time(time.Time{}))}:        CopyTimeToPTime,
		{Src: typeOfPointer(time.Time(time.Time{})), Dest: typeOfPointer(time.Time(time.Time{}))}: CopyPTimeToPTime,
		// time.Duration to time.Duration
		{Src: typeOf(time.Duration(0)), Dest: typeOf(time.Duration(0))}:               CopyDurationToDuration,
		{Src: typeOfPointer(time.Duration(0)), Dest: typeOf(time.Duration(0))}:        CopyPDurationToDuration,
		{Src: typeOf(time.Duration(0)), Dest: typeOfPointer(time.Duration(0))}:        CopyDurationToPDuration,
		{Src: typeOfPointer(time.Duration(0)), Dest: typeOfPointer(time.Duration(0))}: CopyPDurationToPDuration,
	},
	sizes: []func(dst, src unsafe.Pointer){
		Copy1, Copy2, Copy3, Copy4, Copy5, Copy6, Copy7, Copy8, Copy9, Copy10, Copy11, Copy12, Copy13, Copy14, Copy15, Copy16, Copy17, Copy18, Copy19, Copy20, Copy21, Copy22, Copy23, Copy24, Copy25, Copy26, Copy27, Copy28, Copy29, Copy30, Copy31, Copy32, Copy33, Copy34, Copy35, Copy36, Copy37, Copy38, Copy39, Copy40, Copy41, Copy42, Copy43, Copy44, Copy45, Copy46, Copy47, Copy48, Copy49, Copy50, Copy51, Copy52, Copy53, Copy54, Copy55, Copy56, Copy57, Copy58, Copy59, Copy60, Copy61, Copy62, Copy63, Copy64, Copy65, Copy66, Copy67, Copy68, Copy69, Copy70, Copy71, Copy72, Copy73, Copy74, Copy75, Copy76, Copy77, Copy78, Copy79, Copy80, Copy81, Copy82, Copy83, Copy84, Copy85, Copy86, Copy87, Copy88, Copy89, Copy90, Copy91, Copy92, Copy93, Copy94, Copy95, Copy96, Copy97, Copy98, Copy99, Copy100, Copy101, Copy102, Copy103, Copy104, Copy105, Copy106, Copy107, Copy108, Copy109, Copy110, Copy111, Copy112, Copy113, Copy114, Copy115, Copy116, Copy117, Copy118, Copy119, Copy120, Copy121, Copy122, Copy123, Copy124, Copy125, Copy126, Copy127, Copy128, Copy129, Copy130, Copy131, Copy132, Copy133, Copy134, Copy135, Copy136, Copy137, Copy138, Copy139, Copy140, Copy141, Copy142, Copy143, Copy144, Copy145, Copy146, Copy147, Copy148, Copy149, Copy150, Copy151, Copy152, Copy153, Copy154, Copy155, Copy156, Copy157, Copy158, Copy159, Copy160, Copy161, Copy162, Copy163, Copy164, Copy165, Copy166, Copy167, Copy168, Copy169, Copy170, Copy171, Copy172, Copy173, Copy174, Copy175, Copy176, Copy177, Copy178, Copy179, Copy180, Copy181, Copy182, Copy183, Copy184, Copy185, Copy186, Copy187, Copy188, Copy189, Copy190, Copy191, Copy192, Copy193, Copy194, Copy195, Copy196, Copy197, Copy198, Copy199, Copy200, Copy201, Copy202, Copy203, Copy204, Copy205, Copy206, Copy207, Copy208, Copy209, Copy210, Copy211, Copy212, Copy213, Copy214, Copy215, Copy216, Copy217, Copy218, Copy219, Copy220, Copy221, Copy222, Copy223, Copy224, Copy225, Copy226, Copy227, Copy228, Copy229, Copy230, Copy231, Copy232, Copy233, Copy234, Copy235, Copy236, Copy237, Copy238, Copy239, Copy240, Copy241, Copy242, Copy243, Copy244, Copy245, Copy246, Copy247, Copy248, Copy249, Copy250, Copy251, Copy252, Copy253, Copy254, Copy255,
	},
}

// int to int

func CopyIntToInt(dst, src unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dst)) = int(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dst)) = v
}

func CopyIntToPInt(dst, src unsafe.Pointer) {
	v := int(*(*int)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int8 to int

func CopyInt8ToInt(dst, src unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dst)) = int(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dst)) = v
}

func CopyInt8ToPInt(dst, src unsafe.Pointer) {
	v := int(*(*int8)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int16 to int

func CopyInt16ToInt(dst, src unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dst)) = int(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dst)) = v
}

func CopyInt16ToPInt(dst, src unsafe.Pointer) {
	v := int(*(*int16)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int32 to int

func CopyInt32ToInt(dst, src unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dst)) = int(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dst)) = v
}

func CopyInt32ToPInt(dst, src unsafe.Pointer) {
	v := int(*(*int32)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int64 to int

func CopyInt64ToInt(dst, src unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dst)) = int(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dst)) = v
}

func CopyInt64ToPInt(dst, src unsafe.Pointer) {
	v := int(*(*int64)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint to int

func CopyUintToInt(dst, src unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dst)) = int(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dst)) = v
}

func CopyUintToPInt(dst, src unsafe.Pointer) {
	v := int(*(*uint)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint8 to int

func CopyUint8ToInt(dst, src unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dst)) = int(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dst)) = v
}

func CopyUint8ToPInt(dst, src unsafe.Pointer) {
	v := int(*(*uint8)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint16 to int

func CopyUint16ToInt(dst, src unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dst)) = int(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dst)) = v
}

func CopyUint16ToPInt(dst, src unsafe.Pointer) {
	v := int(*(*uint16)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint32 to int

func CopyUint32ToInt(dst, src unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dst)) = int(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dst)) = v
}

func CopyUint32ToPInt(dst, src unsafe.Pointer) {
	v := int(*(*uint32)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint64 to int

func CopyUint64ToInt(dst, src unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dst)) = int(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dst)) = v
}

func CopyUint64ToPInt(dst, src unsafe.Pointer) {
	v := int(*(*uint64)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPInt(dst, src unsafe.Pointer) {
	var v int
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int to int8

func CopyIntToInt8(dst, src unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dst)) = int8(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dst)) = v
}

func CopyIntToPInt8(dst, src unsafe.Pointer) {
	v := int8(*(*int)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int8 to int8

func CopyInt8ToInt8(dst, src unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dst)) = int8(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dst)) = v
}

func CopyInt8ToPInt8(dst, src unsafe.Pointer) {
	v := int8(*(*int8)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int16 to int8

func CopyInt16ToInt8(dst, src unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dst)) = int8(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dst)) = v
}

func CopyInt16ToPInt8(dst, src unsafe.Pointer) {
	v := int8(*(*int16)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int32 to int8

func CopyInt32ToInt8(dst, src unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dst)) = int8(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dst)) = v
}

func CopyInt32ToPInt8(dst, src unsafe.Pointer) {
	v := int8(*(*int32)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int64 to int8

func CopyInt64ToInt8(dst, src unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dst)) = int8(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dst)) = v
}

func CopyInt64ToPInt8(dst, src unsafe.Pointer) {
	v := int8(*(*int64)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint to int8

func CopyUintToInt8(dst, src unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dst)) = int8(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dst)) = v
}

func CopyUintToPInt8(dst, src unsafe.Pointer) {
	v := int8(*(*uint)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint8 to int8

func CopyUint8ToInt8(dst, src unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dst)) = int8(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dst)) = v
}

func CopyUint8ToPInt8(dst, src unsafe.Pointer) {
	v := int8(*(*uint8)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint16 to int8

func CopyUint16ToInt8(dst, src unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dst)) = int8(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dst)) = v
}

func CopyUint16ToPInt8(dst, src unsafe.Pointer) {
	v := int8(*(*uint16)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint32 to int8

func CopyUint32ToInt8(dst, src unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dst)) = int8(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dst)) = v
}

func CopyUint32ToPInt8(dst, src unsafe.Pointer) {
	v := int8(*(*uint32)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint64 to int8

func CopyUint64ToInt8(dst, src unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dst)) = int8(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dst)) = v
}

func CopyUint64ToPInt8(dst, src unsafe.Pointer) {
	v := int8(*(*uint64)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPInt8(dst, src unsafe.Pointer) {
	var v int8
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int to int16

func CopyIntToInt16(dst, src unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dst)) = int16(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dst)) = v
}

func CopyIntToPInt16(dst, src unsafe.Pointer) {
	v := int16(*(*int)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int8 to int16

func CopyInt8ToInt16(dst, src unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dst)) = int16(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dst)) = v
}

func CopyInt8ToPInt16(dst, src unsafe.Pointer) {
	v := int16(*(*int8)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int16 to int16

func CopyInt16ToInt16(dst, src unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dst)) = int16(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dst)) = v
}

func CopyInt16ToPInt16(dst, src unsafe.Pointer) {
	v := int16(*(*int16)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int32 to int16

func CopyInt32ToInt16(dst, src unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dst)) = int16(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dst)) = v
}

func CopyInt32ToPInt16(dst, src unsafe.Pointer) {
	v := int16(*(*int32)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int64 to int16

func CopyInt64ToInt16(dst, src unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dst)) = int16(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dst)) = v
}

func CopyInt64ToPInt16(dst, src unsafe.Pointer) {
	v := int16(*(*int64)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint to int16

func CopyUintToInt16(dst, src unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dst)) = int16(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dst)) = v
}

func CopyUintToPInt16(dst, src unsafe.Pointer) {
	v := int16(*(*uint)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint8 to int16

func CopyUint8ToInt16(dst, src unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dst)) = int16(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dst)) = v
}

func CopyUint8ToPInt16(dst, src unsafe.Pointer) {
	v := int16(*(*uint8)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint16 to int16

func CopyUint16ToInt16(dst, src unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dst)) = int16(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dst)) = v
}

func CopyUint16ToPInt16(dst, src unsafe.Pointer) {
	v := int16(*(*uint16)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint32 to int16

func CopyUint32ToInt16(dst, src unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dst)) = int16(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dst)) = v
}

func CopyUint32ToPInt16(dst, src unsafe.Pointer) {
	v := int16(*(*uint32)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint64 to int16

func CopyUint64ToInt16(dst, src unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dst)) = int16(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dst)) = v
}

func CopyUint64ToPInt16(dst, src unsafe.Pointer) {
	v := int16(*(*uint64)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPInt16(dst, src unsafe.Pointer) {
	var v int16
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int to int32

func CopyIntToInt32(dst, src unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dst)) = int32(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dst)) = v
}

func CopyIntToPInt32(dst, src unsafe.Pointer) {
	v := int32(*(*int)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int8 to int32

func CopyInt8ToInt32(dst, src unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dst)) = int32(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dst)) = v
}

func CopyInt8ToPInt32(dst, src unsafe.Pointer) {
	v := int32(*(*int8)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int16 to int32

func CopyInt16ToInt32(dst, src unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dst)) = int32(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dst)) = v
}

func CopyInt16ToPInt32(dst, src unsafe.Pointer) {
	v := int32(*(*int16)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int32 to int32

func CopyInt32ToInt32(dst, src unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dst)) = int32(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dst)) = v
}

func CopyInt32ToPInt32(dst, src unsafe.Pointer) {
	v := int32(*(*int32)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int64 to int32

func CopyInt64ToInt32(dst, src unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dst)) = int32(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dst)) = v
}

func CopyInt64ToPInt32(dst, src unsafe.Pointer) {
	v := int32(*(*int64)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint to int32

func CopyUintToInt32(dst, src unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dst)) = int32(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dst)) = v
}

func CopyUintToPInt32(dst, src unsafe.Pointer) {
	v := int32(*(*uint)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint8 to int32

func CopyUint8ToInt32(dst, src unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dst)) = int32(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dst)) = v
}

func CopyUint8ToPInt32(dst, src unsafe.Pointer) {
	v := int32(*(*uint8)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint16 to int32

func CopyUint16ToInt32(dst, src unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dst)) = int32(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dst)) = v
}

func CopyUint16ToPInt32(dst, src unsafe.Pointer) {
	v := int32(*(*uint16)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint32 to int32

func CopyUint32ToInt32(dst, src unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dst)) = int32(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dst)) = v
}

func CopyUint32ToPInt32(dst, src unsafe.Pointer) {
	v := int32(*(*uint32)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint64 to int32

func CopyUint64ToInt32(dst, src unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dst)) = int32(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dst)) = v
}

func CopyUint64ToPInt32(dst, src unsafe.Pointer) {
	v := int32(*(*uint64)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPInt32(dst, src unsafe.Pointer) {
	var v int32
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int to int64

func CopyIntToInt64(dst, src unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dst)) = int64(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dst)) = v
}

func CopyIntToPInt64(dst, src unsafe.Pointer) {
	v := int64(*(*int)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int8 to int64

func CopyInt8ToInt64(dst, src unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dst)) = int64(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dst)) = v
}

func CopyInt8ToPInt64(dst, src unsafe.Pointer) {
	v := int64(*(*int8)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int16 to int64

func CopyInt16ToInt64(dst, src unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dst)) = int64(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dst)) = v
}

func CopyInt16ToPInt64(dst, src unsafe.Pointer) {
	v := int64(*(*int16)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int32 to int64

func CopyInt32ToInt64(dst, src unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dst)) = int64(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dst)) = v
}

func CopyInt32ToPInt64(dst, src unsafe.Pointer) {
	v := int64(*(*int32)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int64 to int64

func CopyInt64ToInt64(dst, src unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dst)) = int64(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dst)) = v
}

func CopyInt64ToPInt64(dst, src unsafe.Pointer) {
	v := int64(*(*int64)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint to int64

func CopyUintToInt64(dst, src unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dst)) = int64(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dst)) = v
}

func CopyUintToPInt64(dst, src unsafe.Pointer) {
	v := int64(*(*uint)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint8 to int64

func CopyUint8ToInt64(dst, src unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dst)) = int64(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dst)) = v
}

func CopyUint8ToPInt64(dst, src unsafe.Pointer) {
	v := int64(*(*uint8)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint16 to int64

func CopyUint16ToInt64(dst, src unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dst)) = int64(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dst)) = v
}

func CopyUint16ToPInt64(dst, src unsafe.Pointer) {
	v := int64(*(*uint16)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint32 to int64

func CopyUint32ToInt64(dst, src unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dst)) = int64(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dst)) = v
}

func CopyUint32ToPInt64(dst, src unsafe.Pointer) {
	v := int64(*(*uint32)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint64 to int64

func CopyUint64ToInt64(dst, src unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dst)) = int64(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dst)) = v
}

func CopyUint64ToPInt64(dst, src unsafe.Pointer) {
	v := int64(*(*uint64)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPInt64(dst, src unsafe.Pointer) {
	var v int64
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int to uint

func CopyIntToUint(dst, src unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dst)) = uint(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dst)) = v
}

func CopyIntToPUint(dst, src unsafe.Pointer) {
	v := uint(*(*int)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int8 to uint

func CopyInt8ToUint(dst, src unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dst)) = uint(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dst)) = v
}

func CopyInt8ToPUint(dst, src unsafe.Pointer) {
	v := uint(*(*int8)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int16 to uint

func CopyInt16ToUint(dst, src unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dst)) = uint(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dst)) = v
}

func CopyInt16ToPUint(dst, src unsafe.Pointer) {
	v := uint(*(*int16)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int32 to uint

func CopyInt32ToUint(dst, src unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dst)) = uint(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dst)) = v
}

func CopyInt32ToPUint(dst, src unsafe.Pointer) {
	v := uint(*(*int32)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int64 to uint

func CopyInt64ToUint(dst, src unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dst)) = uint(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dst)) = v
}

func CopyInt64ToPUint(dst, src unsafe.Pointer) {
	v := uint(*(*int64)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint to uint

func CopyUintToUint(dst, src unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dst)) = uint(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dst)) = v
}

func CopyUintToPUint(dst, src unsafe.Pointer) {
	v := uint(*(*uint)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint8 to uint

func CopyUint8ToUint(dst, src unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dst)) = uint(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dst)) = v
}

func CopyUint8ToPUint(dst, src unsafe.Pointer) {
	v := uint(*(*uint8)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint16 to uint

func CopyUint16ToUint(dst, src unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dst)) = uint(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dst)) = v
}

func CopyUint16ToPUint(dst, src unsafe.Pointer) {
	v := uint(*(*uint16)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint32 to uint

func CopyUint32ToUint(dst, src unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dst)) = uint(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dst)) = v
}

func CopyUint32ToPUint(dst, src unsafe.Pointer) {
	v := uint(*(*uint32)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint64 to uint

func CopyUint64ToUint(dst, src unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dst)) = uint(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dst)) = v
}

func CopyUint64ToPUint(dst, src unsafe.Pointer) {
	v := uint(*(*uint64)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPUint(dst, src unsafe.Pointer) {
	var v uint
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int to uint8

func CopyIntToUint8(dst, src unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dst)) = uint8(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dst)) = v
}

func CopyIntToPUint8(dst, src unsafe.Pointer) {
	v := uint8(*(*int)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int8 to uint8

func CopyInt8ToUint8(dst, src unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dst)) = uint8(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dst)) = v
}

func CopyInt8ToPUint8(dst, src unsafe.Pointer) {
	v := uint8(*(*int8)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int16 to uint8

func CopyInt16ToUint8(dst, src unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dst)) = uint8(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dst)) = v
}

func CopyInt16ToPUint8(dst, src unsafe.Pointer) {
	v := uint8(*(*int16)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int32 to uint8

func CopyInt32ToUint8(dst, src unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dst)) = uint8(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dst)) = v
}

func CopyInt32ToPUint8(dst, src unsafe.Pointer) {
	v := uint8(*(*int32)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int64 to uint8

func CopyInt64ToUint8(dst, src unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dst)) = uint8(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dst)) = v
}

func CopyInt64ToPUint8(dst, src unsafe.Pointer) {
	v := uint8(*(*int64)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint to uint8

func CopyUintToUint8(dst, src unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dst)) = uint8(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dst)) = v
}

func CopyUintToPUint8(dst, src unsafe.Pointer) {
	v := uint8(*(*uint)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint8 to uint8

func CopyUint8ToUint8(dst, src unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dst)) = uint8(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dst)) = v
}

func CopyUint8ToPUint8(dst, src unsafe.Pointer) {
	v := uint8(*(*uint8)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint16 to uint8

func CopyUint16ToUint8(dst, src unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dst)) = uint8(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dst)) = v
}

func CopyUint16ToPUint8(dst, src unsafe.Pointer) {
	v := uint8(*(*uint16)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint32 to uint8

func CopyUint32ToUint8(dst, src unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dst)) = uint8(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dst)) = v
}

func CopyUint32ToPUint8(dst, src unsafe.Pointer) {
	v := uint8(*(*uint32)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint64 to uint8

func CopyUint64ToUint8(dst, src unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dst)) = uint8(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dst)) = v
}

func CopyUint64ToPUint8(dst, src unsafe.Pointer) {
	v := uint8(*(*uint64)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPUint8(dst, src unsafe.Pointer) {
	var v uint8
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int to uint16

func CopyIntToUint16(dst, src unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dst)) = uint16(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dst)) = v
}

func CopyIntToPUint16(dst, src unsafe.Pointer) {
	v := uint16(*(*int)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int8 to uint16

func CopyInt8ToUint16(dst, src unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dst)) = uint16(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dst)) = v
}

func CopyInt8ToPUint16(dst, src unsafe.Pointer) {
	v := uint16(*(*int8)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int16 to uint16

func CopyInt16ToUint16(dst, src unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dst)) = uint16(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dst)) = v
}

func CopyInt16ToPUint16(dst, src unsafe.Pointer) {
	v := uint16(*(*int16)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int32 to uint16

func CopyInt32ToUint16(dst, src unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dst)) = uint16(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dst)) = v
}

func CopyInt32ToPUint16(dst, src unsafe.Pointer) {
	v := uint16(*(*int32)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int64 to uint16

func CopyInt64ToUint16(dst, src unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dst)) = uint16(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dst)) = v
}

func CopyInt64ToPUint16(dst, src unsafe.Pointer) {
	v := uint16(*(*int64)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint to uint16

func CopyUintToUint16(dst, src unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dst)) = uint16(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dst)) = v
}

func CopyUintToPUint16(dst, src unsafe.Pointer) {
	v := uint16(*(*uint)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint8 to uint16

func CopyUint8ToUint16(dst, src unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dst)) = uint16(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dst)) = v
}

func CopyUint8ToPUint16(dst, src unsafe.Pointer) {
	v := uint16(*(*uint8)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint16 to uint16

func CopyUint16ToUint16(dst, src unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dst)) = uint16(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dst)) = v
}

func CopyUint16ToPUint16(dst, src unsafe.Pointer) {
	v := uint16(*(*uint16)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint32 to uint16

func CopyUint32ToUint16(dst, src unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dst)) = uint16(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dst)) = v
}

func CopyUint32ToPUint16(dst, src unsafe.Pointer) {
	v := uint16(*(*uint32)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint64 to uint16

func CopyUint64ToUint16(dst, src unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dst)) = uint16(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dst)) = v
}

func CopyUint64ToPUint16(dst, src unsafe.Pointer) {
	v := uint16(*(*uint64)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPUint16(dst, src unsafe.Pointer) {
	var v uint16
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int to uint32

func CopyIntToUint32(dst, src unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dst)) = uint32(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dst)) = v
}

func CopyIntToPUint32(dst, src unsafe.Pointer) {
	v := uint32(*(*int)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int8 to uint32

func CopyInt8ToUint32(dst, src unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dst)) = uint32(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dst)) = v
}

func CopyInt8ToPUint32(dst, src unsafe.Pointer) {
	v := uint32(*(*int8)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int16 to uint32

func CopyInt16ToUint32(dst, src unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dst)) = uint32(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dst)) = v
}

func CopyInt16ToPUint32(dst, src unsafe.Pointer) {
	v := uint32(*(*int16)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int32 to uint32

func CopyInt32ToUint32(dst, src unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dst)) = uint32(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dst)) = v
}

func CopyInt32ToPUint32(dst, src unsafe.Pointer) {
	v := uint32(*(*int32)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int64 to uint32

func CopyInt64ToUint32(dst, src unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dst)) = uint32(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dst)) = v
}

func CopyInt64ToPUint32(dst, src unsafe.Pointer) {
	v := uint32(*(*int64)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint to uint32

func CopyUintToUint32(dst, src unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dst)) = uint32(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dst)) = v
}

func CopyUintToPUint32(dst, src unsafe.Pointer) {
	v := uint32(*(*uint)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint8 to uint32

func CopyUint8ToUint32(dst, src unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dst)) = uint32(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dst)) = v
}

func CopyUint8ToPUint32(dst, src unsafe.Pointer) {
	v := uint32(*(*uint8)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint16 to uint32

func CopyUint16ToUint32(dst, src unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dst)) = uint32(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dst)) = v
}

func CopyUint16ToPUint32(dst, src unsafe.Pointer) {
	v := uint32(*(*uint16)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint32 to uint32

func CopyUint32ToUint32(dst, src unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dst)) = uint32(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dst)) = v
}

func CopyUint32ToPUint32(dst, src unsafe.Pointer) {
	v := uint32(*(*uint32)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint64 to uint32

func CopyUint64ToUint32(dst, src unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dst)) = uint32(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dst)) = v
}

func CopyUint64ToPUint32(dst, src unsafe.Pointer) {
	v := uint32(*(*uint64)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPUint32(dst, src unsafe.Pointer) {
	var v uint32
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int to uint64

func CopyIntToUint64(dst, src unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dst)) = uint64(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dst)) = v
}

func CopyIntToPUint64(dst, src unsafe.Pointer) {
	v := uint64(*(*int)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int8 to uint64

func CopyInt8ToUint64(dst, src unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dst)) = uint64(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dst)) = v
}

func CopyInt8ToPUint64(dst, src unsafe.Pointer) {
	v := uint64(*(*int8)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int16 to uint64

func CopyInt16ToUint64(dst, src unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dst)) = uint64(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dst)) = v
}

func CopyInt16ToPUint64(dst, src unsafe.Pointer) {
	v := uint64(*(*int16)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int32 to uint64

func CopyInt32ToUint64(dst, src unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dst)) = uint64(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dst)) = v
}

func CopyInt32ToPUint64(dst, src unsafe.Pointer) {
	v := uint64(*(*int32)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// int64 to uint64

func CopyInt64ToUint64(dst, src unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dst)) = uint64(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dst)) = v
}

func CopyInt64ToPUint64(dst, src unsafe.Pointer) {
	v := uint64(*(*int64)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint to uint64

func CopyUintToUint64(dst, src unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dst)) = uint64(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dst)) = v
}

func CopyUintToPUint64(dst, src unsafe.Pointer) {
	v := uint64(*(*uint)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint8 to uint64

func CopyUint8ToUint64(dst, src unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dst)) = uint64(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dst)) = v
}

func CopyUint8ToPUint64(dst, src unsafe.Pointer) {
	v := uint64(*(*uint8)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint16 to uint64

func CopyUint16ToUint64(dst, src unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dst)) = uint64(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dst)) = v
}

func CopyUint16ToPUint64(dst, src unsafe.Pointer) {
	v := uint64(*(*uint16)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint32 to uint64

func CopyUint32ToUint64(dst, src unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dst)) = uint64(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dst)) = v
}

func CopyUint32ToPUint64(dst, src unsafe.Pointer) {
	v := uint64(*(*uint32)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// uint64 to uint64

func CopyUint64ToUint64(dst, src unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dst)) = uint64(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dst)) = v
}

func CopyUint64ToPUint64(dst, src unsafe.Pointer) {
	v := uint64(*(*uint64)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPUint64(dst, src unsafe.Pointer) {
	var v uint64
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// float32 to float32

func CopyFloat32ToFloat32(dst, src unsafe.Pointer) {
	*(*float32)(unsafe.Pointer(dst)) = float32(*(*float32)(unsafe.Pointer(src)))
}

func CopyPFloat32ToFloat32(dst, src unsafe.Pointer) {
	var v float32
	if p := *(**float32)(unsafe.Pointer(src)); p != nil {
		v = float32(*p)
	}
	*(*float32)(unsafe.Pointer(dst)) = v
}

func CopyFloat32ToPFloat32(dst, src unsafe.Pointer) {
	v := float32(*(*float32)(unsafe.Pointer(src)))
	p := (**float32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPFloat32ToPFloat32(dst, src unsafe.Pointer) {
	var v float32
	if p := *(**float32)(unsafe.Pointer(src)); p != nil {
		v = float32(*p)
	}

	p := (**float32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// float64 to float32

func CopyFloat64ToFloat32(dst, src unsafe.Pointer) {
	*(*float32)(unsafe.Pointer(dst)) = float32(*(*float64)(unsafe.Pointer(src)))
}

func CopyPFloat64ToFloat32(dst, src unsafe.Pointer) {
	var v float32
	if p := *(**float64)(unsafe.Pointer(src)); p != nil {
		v = float32(*p)
	}
	*(*float32)(unsafe.Pointer(dst)) = v
}

func CopyFloat64ToPFloat32(dst, src unsafe.Pointer) {
	v := float32(*(*float64)(unsafe.Pointer(src)))
	p := (**float32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPFloat64ToPFloat32(dst, src unsafe.Pointer) {
	var v float32
	if p := *(**float64)(unsafe.Pointer(src)); p != nil {
		v = float32(*p)
	}

	p := (**float32)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// float32 to float64

func CopyFloat32ToFloat64(dst, src unsafe.Pointer) {
	*(*float64)(unsafe.Pointer(dst)) = float64(*(*float32)(unsafe.Pointer(src)))
}

func CopyPFloat32ToFloat64(dst, src unsafe.Pointer) {
	var v float64
	if p := *(**float32)(unsafe.Pointer(src)); p != nil {
		v = float64(*p)
	}
	*(*float64)(unsafe.Pointer(dst)) = v
}

func CopyFloat32ToPFloat64(dst, src unsafe.Pointer) {
	v := float64(*(*float32)(unsafe.Pointer(src)))
	p := (**float64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPFloat32ToPFloat64(dst, src unsafe.Pointer) {
	var v float64
	if p := *(**float32)(unsafe.Pointer(src)); p != nil {
		v = float64(*p)
	}

	p := (**float64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// float64 to float64

func CopyFloat64ToFloat64(dst, src unsafe.Pointer) {
	*(*float64)(unsafe.Pointer(dst)) = float64(*(*float64)(unsafe.Pointer(src)))
}

func CopyPFloat64ToFloat64(dst, src unsafe.Pointer) {
	var v float64
	if p := *(**float64)(unsafe.Pointer(src)); p != nil {
		v = float64(*p)
	}
	*(*float64)(unsafe.Pointer(dst)) = v
}

func CopyFloat64ToPFloat64(dst, src unsafe.Pointer) {
	v := float64(*(*float64)(unsafe.Pointer(src)))
	p := (**float64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPFloat64ToPFloat64(dst, src unsafe.Pointer) {
	var v float64
	if p := *(**float64)(unsafe.Pointer(src)); p != nil {
		v = float64(*p)
	}

	p := (**float64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// bool to bool

func CopyBoolToBool(dst, src unsafe.Pointer) {
	*(*bool)(unsafe.Pointer(dst)) = bool(*(*bool)(unsafe.Pointer(src)))
}

func CopyPBoolToBool(dst, src unsafe.Pointer) {
	var v bool
	if p := *(**bool)(unsafe.Pointer(src)); p != nil {
		v = bool(*p)
	}
	*(*bool)(unsafe.Pointer(dst)) = v
}

func CopyBoolToPBool(dst, src unsafe.Pointer) {
	v := bool(*(*bool)(unsafe.Pointer(src)))
	p := (**bool)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPBoolToPBool(dst, src unsafe.Pointer) {
	var v bool
	if p := *(**bool)(unsafe.Pointer(src)); p != nil {
		v = bool(*p)
	}

	p := (**bool)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// complex64 to complex64

func CopyComplex64ToComplex64(dst, src unsafe.Pointer) {
	*(*complex64)(unsafe.Pointer(dst)) = complex64(*(*complex64)(unsafe.Pointer(src)))
}

func CopyPComplex64ToComplex64(dst, src unsafe.Pointer) {
	var v complex64
	if p := *(**complex64)(unsafe.Pointer(src)); p != nil {
		v = complex64(*p)
	}
	*(*complex64)(unsafe.Pointer(dst)) = v
}

func CopyComplex64ToPComplex64(dst, src unsafe.Pointer) {
	v := complex64(*(*complex64)(unsafe.Pointer(src)))
	p := (**complex64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPComplex64ToPComplex64(dst, src unsafe.Pointer) {
	var v complex64
	if p := *(**complex64)(unsafe.Pointer(src)); p != nil {
		v = complex64(*p)
	}

	p := (**complex64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// complex128 to complex64

func CopyComplex128ToComplex64(dst, src unsafe.Pointer) {
	*(*complex64)(unsafe.Pointer(dst)) = complex64(*(*complex128)(unsafe.Pointer(src)))
}

func CopyPComplex128ToComplex64(dst, src unsafe.Pointer) {
	var v complex64
	if p := *(**complex128)(unsafe.Pointer(src)); p != nil {
		v = complex64(*p)
	}
	*(*complex64)(unsafe.Pointer(dst)) = v
}

func CopyComplex128ToPComplex64(dst, src unsafe.Pointer) {
	v := complex64(*(*complex128)(unsafe.Pointer(src)))
	p := (**complex64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPComplex128ToPComplex64(dst, src unsafe.Pointer) {
	var v complex64
	if p := *(**complex128)(unsafe.Pointer(src)); p != nil {
		v = complex64(*p)
	}

	p := (**complex64)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// complex64 to complex128

func CopyComplex64ToComplex128(dst, src unsafe.Pointer) {
	*(*complex128)(unsafe.Pointer(dst)) = complex128(*(*complex64)(unsafe.Pointer(src)))
}

func CopyPComplex64ToComplex128(dst, src unsafe.Pointer) {
	var v complex128
	if p := *(**complex64)(unsafe.Pointer(src)); p != nil {
		v = complex128(*p)
	}
	*(*complex128)(unsafe.Pointer(dst)) = v
}

func CopyComplex64ToPComplex128(dst, src unsafe.Pointer) {
	v := complex128(*(*complex64)(unsafe.Pointer(src)))
	p := (**complex128)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPComplex64ToPComplex128(dst, src unsafe.Pointer) {
	var v complex128
	if p := *(**complex64)(unsafe.Pointer(src)); p != nil {
		v = complex128(*p)
	}

	p := (**complex128)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// complex128 to complex128

func CopyComplex128ToComplex128(dst, src unsafe.Pointer) {
	*(*complex128)(unsafe.Pointer(dst)) = complex128(*(*complex128)(unsafe.Pointer(src)))
}

func CopyPComplex128ToComplex128(dst, src unsafe.Pointer) {
	var v complex128
	if p := *(**complex128)(unsafe.Pointer(src)); p != nil {
		v = complex128(*p)
	}
	*(*complex128)(unsafe.Pointer(dst)) = v
}

func CopyComplex128ToPComplex128(dst, src unsafe.Pointer) {
	v := complex128(*(*complex128)(unsafe.Pointer(src)))
	p := (**complex128)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPComplex128ToPComplex128(dst, src unsafe.Pointer) {
	var v complex128
	if p := *(**complex128)(unsafe.Pointer(src)); p != nil {
		v = complex128(*p)
	}

	p := (**complex128)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// string to string

func CopyStringToString(dst, src unsafe.Pointer) {
	*(*string)(unsafe.Pointer(dst)) = string(*(*string)(unsafe.Pointer(src)))
}

func CopyPStringToString(dst, src unsafe.Pointer) {
	var v string
	if p := *(**string)(unsafe.Pointer(src)); p != nil {
		v = string(*p)
	}
	*(*string)(unsafe.Pointer(dst)) = v
}

func CopyStringToPString(dst, src unsafe.Pointer) {
	v := string(*(*string)(unsafe.Pointer(src)))
	p := (**string)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPStringToPString(dst, src unsafe.Pointer) {
	var v string
	if p := *(**string)(unsafe.Pointer(src)); p != nil {
		v = string(*p)
	}

	p := (**string)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// []byte to string

func CopyBytesToString(dst, src unsafe.Pointer) {
	*(*string)(unsafe.Pointer(dst)) = string(*(*[]byte)(unsafe.Pointer(src)))
}

func CopyPBytesToString(dst, src unsafe.Pointer) {
	var v string
	if p := *(**[]byte)(unsafe.Pointer(src)); p != nil {
		v = string(*p)
	}
	*(*string)(unsafe.Pointer(dst)) = v
}

func CopyBytesToPString(dst, src unsafe.Pointer) {
	v := string(*(*[]byte)(unsafe.Pointer(src)))
	p := (**string)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPBytesToPString(dst, src unsafe.Pointer) {
	var v string
	if p := *(**[]byte)(unsafe.Pointer(src)); p != nil {
		v = string(*p)
	}

	p := (**string)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// string to []byte

func CopyStringToBytes(dst, src unsafe.Pointer) {
	*(*[]byte)(unsafe.Pointer(dst)) = []byte(*(*string)(unsafe.Pointer(src)))
}

func CopyPStringToBytes(dst, src unsafe.Pointer) {
	var v []byte
	if p := *(**string)(unsafe.Pointer(src)); p != nil {
		v = []byte(*p)
	}
	*(*[]byte)(unsafe.Pointer(dst)) = v
}

func CopyStringToPBytes(dst, src unsafe.Pointer) {
	v := []byte(*(*string)(unsafe.Pointer(src)))
	p := (**[]byte)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPStringToPBytes(dst, src unsafe.Pointer) {
	var v []byte
	if p := *(**string)(unsafe.Pointer(src)); p != nil {
		v = []byte(*p)
	}

	p := (**[]byte)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// []byte to []byte

func CopyBytesToBytes(dst, src unsafe.Pointer) {
	*(*[]byte)(unsafe.Pointer(dst)) = []byte(*(*[]byte)(unsafe.Pointer(src)))
}

func CopyPBytesToBytes(dst, src unsafe.Pointer) {
	var v []byte
	if p := *(**[]byte)(unsafe.Pointer(src)); p != nil {
		v = []byte(*p)
	}
	*(*[]byte)(unsafe.Pointer(dst)) = v
}

func CopyBytesToPBytes(dst, src unsafe.Pointer) {
	v := []byte(*(*[]byte)(unsafe.Pointer(src)))
	p := (**[]byte)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPBytesToPBytes(dst, src unsafe.Pointer) {
	var v []byte
	if p := *(**[]byte)(unsafe.Pointer(src)); p != nil {
		v = []byte(*p)
	}

	p := (**[]byte)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// time.Time to time.Time

func CopyTimeToTime(dst, src unsafe.Pointer) {
	*(*time.Time)(unsafe.Pointer(dst)) = time.Time(*(*time.Time)(unsafe.Pointer(src)))
}

func CopyPTimeToTime(dst, src unsafe.Pointer) {
	var v time.Time
	if p := *(**time.Time)(unsafe.Pointer(src)); p != nil {
		v = time.Time(*p)
	}
	*(*time.Time)(unsafe.Pointer(dst)) = v
}

func CopyTimeToPTime(dst, src unsafe.Pointer) {
	v := time.Time(*(*time.Time)(unsafe.Pointer(src)))
	p := (**time.Time)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPTimeToPTime(dst, src unsafe.Pointer) {
	var v time.Time
	if p := *(**time.Time)(unsafe.Pointer(src)); p != nil {
		v = time.Time(*p)
	}

	p := (**time.Time)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// time.Duration to time.Duration

func CopyDurationToDuration(dst, src unsafe.Pointer) {
	*(*time.Duration)(unsafe.Pointer(dst)) = time.Duration(*(*time.Duration)(unsafe.Pointer(src)))
}

func CopyPDurationToDuration(dst, src unsafe.Pointer) {
	var v time.Duration
	if p := *(**time.Duration)(unsafe.Pointer(src)); p != nil {
		v = time.Duration(*p)
	}
	*(*time.Duration)(unsafe.Pointer(dst)) = v
}

func CopyDurationToPDuration(dst, src unsafe.Pointer) {
	v := time.Duration(*(*time.Duration)(unsafe.Pointer(src)))
	p := (**time.Duration)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPDurationToPDuration(dst, src unsafe.Pointer) {
	var v time.Duration
	if p := *(**time.Duration)(unsafe.Pointer(src)); p != nil {
		v = time.Duration(*p)
	}

	p := (**time.Duration)(unsafe.Pointer(dst))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

// Memcopy funcs
func Copy1(dst, src unsafe.Pointer) {
	*(*[1]byte)(unsafe.Pointer(dst)) = *(*[1]byte)(unsafe.Pointer(src))
}

func Copy2(dst, src unsafe.Pointer) {
	*(*[2]byte)(unsafe.Pointer(dst)) = *(*[2]byte)(unsafe.Pointer(src))
}

func Copy3(dst, src unsafe.Pointer) {
	*(*[3]byte)(unsafe.Pointer(dst)) = *(*[3]byte)(unsafe.Pointer(src))
}

func Copy4(dst, src unsafe.Pointer) {
	*(*[4]byte)(unsafe.Pointer(dst)) = *(*[4]byte)(unsafe.Pointer(src))
}

func Copy5(dst, src unsafe.Pointer) {
	*(*[5]byte)(unsafe.Pointer(dst)) = *(*[5]byte)(unsafe.Pointer(src))
}

func Copy6(dst, src unsafe.Pointer) {
	*(*[6]byte)(unsafe.Pointer(dst)) = *(*[6]byte)(unsafe.Pointer(src))
}

func Copy7(dst, src unsafe.Pointer) {
	*(*[7]byte)(unsafe.Pointer(dst)) = *(*[7]byte)(unsafe.Pointer(src))
}

func Copy8(dst, src unsafe.Pointer) {
	*(*[8]byte)(unsafe.Pointer(dst)) = *(*[8]byte)(unsafe.Pointer(src))
}

func Copy9(dst, src unsafe.Pointer) {
	*(*[9]byte)(unsafe.Pointer(dst)) = *(*[9]byte)(unsafe.Pointer(src))
}

func Copy10(dst, src unsafe.Pointer) {
	*(*[10]byte)(unsafe.Pointer(dst)) = *(*[10]byte)(unsafe.Pointer(src))
}

func Copy11(dst, src unsafe.Pointer) {
	*(*[11]byte)(unsafe.Pointer(dst)) = *(*[11]byte)(unsafe.Pointer(src))
}

func Copy12(dst, src unsafe.Pointer) {
	*(*[12]byte)(unsafe.Pointer(dst)) = *(*[12]byte)(unsafe.Pointer(src))
}

func Copy13(dst, src unsafe.Pointer) {
	*(*[13]byte)(unsafe.Pointer(dst)) = *(*[13]byte)(unsafe.Pointer(src))
}

func Copy14(dst, src unsafe.Pointer) {
	*(*[14]byte)(unsafe.Pointer(dst)) = *(*[14]byte)(unsafe.Pointer(src))
}

func Copy15(dst, src unsafe.Pointer) {
	*(*[15]byte)(unsafe.Pointer(dst)) = *(*[15]byte)(unsafe.Pointer(src))
}

func Copy16(dst, src unsafe.Pointer) {
	*(*[16]byte)(unsafe.Pointer(dst)) = *(*[16]byte)(unsafe.Pointer(src))
}

func Copy17(dst, src unsafe.Pointer) {
	*(*[17]byte)(unsafe.Pointer(dst)) = *(*[17]byte)(unsafe.Pointer(src))
}

func Copy18(dst, src unsafe.Pointer) {
	*(*[18]byte)(unsafe.Pointer(dst)) = *(*[18]byte)(unsafe.Pointer(src))
}

func Copy19(dst, src unsafe.Pointer) {
	*(*[19]byte)(unsafe.Pointer(dst)) = *(*[19]byte)(unsafe.Pointer(src))
}

func Copy20(dst, src unsafe.Pointer) {
	*(*[20]byte)(unsafe.Pointer(dst)) = *(*[20]byte)(unsafe.Pointer(src))
}

func Copy21(dst, src unsafe.Pointer) {
	*(*[21]byte)(unsafe.Pointer(dst)) = *(*[21]byte)(unsafe.Pointer(src))
}

func Copy22(dst, src unsafe.Pointer) {
	*(*[22]byte)(unsafe.Pointer(dst)) = *(*[22]byte)(unsafe.Pointer(src))
}

func Copy23(dst, src unsafe.Pointer) {
	*(*[23]byte)(unsafe.Pointer(dst)) = *(*[23]byte)(unsafe.Pointer(src))
}

func Copy24(dst, src unsafe.Pointer) {
	*(*[24]byte)(unsafe.Pointer(dst)) = *(*[24]byte)(unsafe.Pointer(src))
}

func Copy25(dst, src unsafe.Pointer) {
	*(*[25]byte)(unsafe.Pointer(dst)) = *(*[25]byte)(unsafe.Pointer(src))
}

func Copy26(dst, src unsafe.Pointer) {
	*(*[26]byte)(unsafe.Pointer(dst)) = *(*[26]byte)(unsafe.Pointer(src))
}

func Copy27(dst, src unsafe.Pointer) {
	*(*[27]byte)(unsafe.Pointer(dst)) = *(*[27]byte)(unsafe.Pointer(src))
}

func Copy28(dst, src unsafe.Pointer) {
	*(*[28]byte)(unsafe.Pointer(dst)) = *(*[28]byte)(unsafe.Pointer(src))
}

func Copy29(dst, src unsafe.Pointer) {
	*(*[29]byte)(unsafe.Pointer(dst)) = *(*[29]byte)(unsafe.Pointer(src))
}

func Copy30(dst, src unsafe.Pointer) {
	*(*[30]byte)(unsafe.Pointer(dst)) = *(*[30]byte)(unsafe.Pointer(src))
}

func Copy31(dst, src unsafe.Pointer) {
	*(*[31]byte)(unsafe.Pointer(dst)) = *(*[31]byte)(unsafe.Pointer(src))
}

func Copy32(dst, src unsafe.Pointer) {
	*(*[32]byte)(unsafe.Pointer(dst)) = *(*[32]byte)(unsafe.Pointer(src))
}

func Copy33(dst, src unsafe.Pointer) {
	*(*[33]byte)(unsafe.Pointer(dst)) = *(*[33]byte)(unsafe.Pointer(src))
}

func Copy34(dst, src unsafe.Pointer) {
	*(*[34]byte)(unsafe.Pointer(dst)) = *(*[34]byte)(unsafe.Pointer(src))
}

func Copy35(dst, src unsafe.Pointer) {
	*(*[35]byte)(unsafe.Pointer(dst)) = *(*[35]byte)(unsafe.Pointer(src))
}

func Copy36(dst, src unsafe.Pointer) {
	*(*[36]byte)(unsafe.Pointer(dst)) = *(*[36]byte)(unsafe.Pointer(src))
}

func Copy37(dst, src unsafe.Pointer) {
	*(*[37]byte)(unsafe.Pointer(dst)) = *(*[37]byte)(unsafe.Pointer(src))
}

func Copy38(dst, src unsafe.Pointer) {
	*(*[38]byte)(unsafe.Pointer(dst)) = *(*[38]byte)(unsafe.Pointer(src))
}

func Copy39(dst, src unsafe.Pointer) {
	*(*[39]byte)(unsafe.Pointer(dst)) = *(*[39]byte)(unsafe.Pointer(src))
}

func Copy40(dst, src unsafe.Pointer) {
	*(*[40]byte)(unsafe.Pointer(dst)) = *(*[40]byte)(unsafe.Pointer(src))
}

func Copy41(dst, src unsafe.Pointer) {
	*(*[41]byte)(unsafe.Pointer(dst)) = *(*[41]byte)(unsafe.Pointer(src))
}

func Copy42(dst, src unsafe.Pointer) {
	*(*[42]byte)(unsafe.Pointer(dst)) = *(*[42]byte)(unsafe.Pointer(src))
}

func Copy43(dst, src unsafe.Pointer) {
	*(*[43]byte)(unsafe.Pointer(dst)) = *(*[43]byte)(unsafe.Pointer(src))
}

func Copy44(dst, src unsafe.Pointer) {
	*(*[44]byte)(unsafe.Pointer(dst)) = *(*[44]byte)(unsafe.Pointer(src))
}

func Copy45(dst, src unsafe.Pointer) {
	*(*[45]byte)(unsafe.Pointer(dst)) = *(*[45]byte)(unsafe.Pointer(src))
}

func Copy46(dst, src unsafe.Pointer) {
	*(*[46]byte)(unsafe.Pointer(dst)) = *(*[46]byte)(unsafe.Pointer(src))
}

func Copy47(dst, src unsafe.Pointer) {
	*(*[47]byte)(unsafe.Pointer(dst)) = *(*[47]byte)(unsafe.Pointer(src))
}

func Copy48(dst, src unsafe.Pointer) {
	*(*[48]byte)(unsafe.Pointer(dst)) = *(*[48]byte)(unsafe.Pointer(src))
}

func Copy49(dst, src unsafe.Pointer) {
	*(*[49]byte)(unsafe.Pointer(dst)) = *(*[49]byte)(unsafe.Pointer(src))
}

func Copy50(dst, src unsafe.Pointer) {
	*(*[50]byte)(unsafe.Pointer(dst)) = *(*[50]byte)(unsafe.Pointer(src))
}

func Copy51(dst, src unsafe.Pointer) {
	*(*[51]byte)(unsafe.Pointer(dst)) = *(*[51]byte)(unsafe.Pointer(src))
}

func Copy52(dst, src unsafe.Pointer) {
	*(*[52]byte)(unsafe.Pointer(dst)) = *(*[52]byte)(unsafe.Pointer(src))
}

func Copy53(dst, src unsafe.Pointer) {
	*(*[53]byte)(unsafe.Pointer(dst)) = *(*[53]byte)(unsafe.Pointer(src))
}

func Copy54(dst, src unsafe.Pointer) {
	*(*[54]byte)(unsafe.Pointer(dst)) = *(*[54]byte)(unsafe.Pointer(src))
}

func Copy55(dst, src unsafe.Pointer) {
	*(*[55]byte)(unsafe.Pointer(dst)) = *(*[55]byte)(unsafe.Pointer(src))
}

func Copy56(dst, src unsafe.Pointer) {
	*(*[56]byte)(unsafe.Pointer(dst)) = *(*[56]byte)(unsafe.Pointer(src))
}

func Copy57(dst, src unsafe.Pointer) {
	*(*[57]byte)(unsafe.Pointer(dst)) = *(*[57]byte)(unsafe.Pointer(src))
}

func Copy58(dst, src unsafe.Pointer) {
	*(*[58]byte)(unsafe.Pointer(dst)) = *(*[58]byte)(unsafe.Pointer(src))
}

func Copy59(dst, src unsafe.Pointer) {
	*(*[59]byte)(unsafe.Pointer(dst)) = *(*[59]byte)(unsafe.Pointer(src))
}

func Copy60(dst, src unsafe.Pointer) {
	*(*[60]byte)(unsafe.Pointer(dst)) = *(*[60]byte)(unsafe.Pointer(src))
}

func Copy61(dst, src unsafe.Pointer) {
	*(*[61]byte)(unsafe.Pointer(dst)) = *(*[61]byte)(unsafe.Pointer(src))
}

func Copy62(dst, src unsafe.Pointer) {
	*(*[62]byte)(unsafe.Pointer(dst)) = *(*[62]byte)(unsafe.Pointer(src))
}

func Copy63(dst, src unsafe.Pointer) {
	*(*[63]byte)(unsafe.Pointer(dst)) = *(*[63]byte)(unsafe.Pointer(src))
}

func Copy64(dst, src unsafe.Pointer) {
	*(*[64]byte)(unsafe.Pointer(dst)) = *(*[64]byte)(unsafe.Pointer(src))
}

func Copy65(dst, src unsafe.Pointer) {
	*(*[65]byte)(unsafe.Pointer(dst)) = *(*[65]byte)(unsafe.Pointer(src))
}

func Copy66(dst, src unsafe.Pointer) {
	*(*[66]byte)(unsafe.Pointer(dst)) = *(*[66]byte)(unsafe.Pointer(src))
}

func Copy67(dst, src unsafe.Pointer) {
	*(*[67]byte)(unsafe.Pointer(dst)) = *(*[67]byte)(unsafe.Pointer(src))
}

func Copy68(dst, src unsafe.Pointer) {
	*(*[68]byte)(unsafe.Pointer(dst)) = *(*[68]byte)(unsafe.Pointer(src))
}

func Copy69(dst, src unsafe.Pointer) {
	*(*[69]byte)(unsafe.Pointer(dst)) = *(*[69]byte)(unsafe.Pointer(src))
}

func Copy70(dst, src unsafe.Pointer) {
	*(*[70]byte)(unsafe.Pointer(dst)) = *(*[70]byte)(unsafe.Pointer(src))
}

func Copy71(dst, src unsafe.Pointer) {
	*(*[71]byte)(unsafe.Pointer(dst)) = *(*[71]byte)(unsafe.Pointer(src))
}

func Copy72(dst, src unsafe.Pointer) {
	*(*[72]byte)(unsafe.Pointer(dst)) = *(*[72]byte)(unsafe.Pointer(src))
}

func Copy73(dst, src unsafe.Pointer) {
	*(*[73]byte)(unsafe.Pointer(dst)) = *(*[73]byte)(unsafe.Pointer(src))
}

func Copy74(dst, src unsafe.Pointer) {
	*(*[74]byte)(unsafe.Pointer(dst)) = *(*[74]byte)(unsafe.Pointer(src))
}

func Copy75(dst, src unsafe.Pointer) {
	*(*[75]byte)(unsafe.Pointer(dst)) = *(*[75]byte)(unsafe.Pointer(src))
}

func Copy76(dst, src unsafe.Pointer) {
	*(*[76]byte)(unsafe.Pointer(dst)) = *(*[76]byte)(unsafe.Pointer(src))
}

func Copy77(dst, src unsafe.Pointer) {
	*(*[77]byte)(unsafe.Pointer(dst)) = *(*[77]byte)(unsafe.Pointer(src))
}

func Copy78(dst, src unsafe.Pointer) {
	*(*[78]byte)(unsafe.Pointer(dst)) = *(*[78]byte)(unsafe.Pointer(src))
}

func Copy79(dst, src unsafe.Pointer) {
	*(*[79]byte)(unsafe.Pointer(dst)) = *(*[79]byte)(unsafe.Pointer(src))
}

func Copy80(dst, src unsafe.Pointer) {
	*(*[80]byte)(unsafe.Pointer(dst)) = *(*[80]byte)(unsafe.Pointer(src))
}

func Copy81(dst, src unsafe.Pointer) {
	*(*[81]byte)(unsafe.Pointer(dst)) = *(*[81]byte)(unsafe.Pointer(src))
}

func Copy82(dst, src unsafe.Pointer) {
	*(*[82]byte)(unsafe.Pointer(dst)) = *(*[82]byte)(unsafe.Pointer(src))
}

func Copy83(dst, src unsafe.Pointer) {
	*(*[83]byte)(unsafe.Pointer(dst)) = *(*[83]byte)(unsafe.Pointer(src))
}

func Copy84(dst, src unsafe.Pointer) {
	*(*[84]byte)(unsafe.Pointer(dst)) = *(*[84]byte)(unsafe.Pointer(src))
}

func Copy85(dst, src unsafe.Pointer) {
	*(*[85]byte)(unsafe.Pointer(dst)) = *(*[85]byte)(unsafe.Pointer(src))
}

func Copy86(dst, src unsafe.Pointer) {
	*(*[86]byte)(unsafe.Pointer(dst)) = *(*[86]byte)(unsafe.Pointer(src))
}

func Copy87(dst, src unsafe.Pointer) {
	*(*[87]byte)(unsafe.Pointer(dst)) = *(*[87]byte)(unsafe.Pointer(src))
}

func Copy88(dst, src unsafe.Pointer) {
	*(*[88]byte)(unsafe.Pointer(dst)) = *(*[88]byte)(unsafe.Pointer(src))
}

func Copy89(dst, src unsafe.Pointer) {
	*(*[89]byte)(unsafe.Pointer(dst)) = *(*[89]byte)(unsafe.Pointer(src))
}

func Copy90(dst, src unsafe.Pointer) {
	*(*[90]byte)(unsafe.Pointer(dst)) = *(*[90]byte)(unsafe.Pointer(src))
}

func Copy91(dst, src unsafe.Pointer) {
	*(*[91]byte)(unsafe.Pointer(dst)) = *(*[91]byte)(unsafe.Pointer(src))
}

func Copy92(dst, src unsafe.Pointer) {
	*(*[92]byte)(unsafe.Pointer(dst)) = *(*[92]byte)(unsafe.Pointer(src))
}

func Copy93(dst, src unsafe.Pointer) {
	*(*[93]byte)(unsafe.Pointer(dst)) = *(*[93]byte)(unsafe.Pointer(src))
}

func Copy94(dst, src unsafe.Pointer) {
	*(*[94]byte)(unsafe.Pointer(dst)) = *(*[94]byte)(unsafe.Pointer(src))
}

func Copy95(dst, src unsafe.Pointer) {
	*(*[95]byte)(unsafe.Pointer(dst)) = *(*[95]byte)(unsafe.Pointer(src))
}

func Copy96(dst, src unsafe.Pointer) {
	*(*[96]byte)(unsafe.Pointer(dst)) = *(*[96]byte)(unsafe.Pointer(src))
}

func Copy97(dst, src unsafe.Pointer) {
	*(*[97]byte)(unsafe.Pointer(dst)) = *(*[97]byte)(unsafe.Pointer(src))
}

func Copy98(dst, src unsafe.Pointer) {
	*(*[98]byte)(unsafe.Pointer(dst)) = *(*[98]byte)(unsafe.Pointer(src))
}

func Copy99(dst, src unsafe.Pointer) {
	*(*[99]byte)(unsafe.Pointer(dst)) = *(*[99]byte)(unsafe.Pointer(src))
}

func Copy100(dst, src unsafe.Pointer) {
	*(*[100]byte)(unsafe.Pointer(dst)) = *(*[100]byte)(unsafe.Pointer(src))
}

func Copy101(dst, src unsafe.Pointer) {
	*(*[101]byte)(unsafe.Pointer(dst)) = *(*[101]byte)(unsafe.Pointer(src))
}

func Copy102(dst, src unsafe.Pointer) {
	*(*[102]byte)(unsafe.Pointer(dst)) = *(*[102]byte)(unsafe.Pointer(src))
}

func Copy103(dst, src unsafe.Pointer) {
	*(*[103]byte)(unsafe.Pointer(dst)) = *(*[103]byte)(unsafe.Pointer(src))
}

func Copy104(dst, src unsafe.Pointer) {
	*(*[104]byte)(unsafe.Pointer(dst)) = *(*[104]byte)(unsafe.Pointer(src))
}

func Copy105(dst, src unsafe.Pointer) {
	*(*[105]byte)(unsafe.Pointer(dst)) = *(*[105]byte)(unsafe.Pointer(src))
}

func Copy106(dst, src unsafe.Pointer) {
	*(*[106]byte)(unsafe.Pointer(dst)) = *(*[106]byte)(unsafe.Pointer(src))
}

func Copy107(dst, src unsafe.Pointer) {
	*(*[107]byte)(unsafe.Pointer(dst)) = *(*[107]byte)(unsafe.Pointer(src))
}

func Copy108(dst, src unsafe.Pointer) {
	*(*[108]byte)(unsafe.Pointer(dst)) = *(*[108]byte)(unsafe.Pointer(src))
}

func Copy109(dst, src unsafe.Pointer) {
	*(*[109]byte)(unsafe.Pointer(dst)) = *(*[109]byte)(unsafe.Pointer(src))
}

func Copy110(dst, src unsafe.Pointer) {
	*(*[110]byte)(unsafe.Pointer(dst)) = *(*[110]byte)(unsafe.Pointer(src))
}

func Copy111(dst, src unsafe.Pointer) {
	*(*[111]byte)(unsafe.Pointer(dst)) = *(*[111]byte)(unsafe.Pointer(src))
}

func Copy112(dst, src unsafe.Pointer) {
	*(*[112]byte)(unsafe.Pointer(dst)) = *(*[112]byte)(unsafe.Pointer(src))
}

func Copy113(dst, src unsafe.Pointer) {
	*(*[113]byte)(unsafe.Pointer(dst)) = *(*[113]byte)(unsafe.Pointer(src))
}

func Copy114(dst, src unsafe.Pointer) {
	*(*[114]byte)(unsafe.Pointer(dst)) = *(*[114]byte)(unsafe.Pointer(src))
}

func Copy115(dst, src unsafe.Pointer) {
	*(*[115]byte)(unsafe.Pointer(dst)) = *(*[115]byte)(unsafe.Pointer(src))
}

func Copy116(dst, src unsafe.Pointer) {
	*(*[116]byte)(unsafe.Pointer(dst)) = *(*[116]byte)(unsafe.Pointer(src))
}

func Copy117(dst, src unsafe.Pointer) {
	*(*[117]byte)(unsafe.Pointer(dst)) = *(*[117]byte)(unsafe.Pointer(src))
}

func Copy118(dst, src unsafe.Pointer) {
	*(*[118]byte)(unsafe.Pointer(dst)) = *(*[118]byte)(unsafe.Pointer(src))
}

func Copy119(dst, src unsafe.Pointer) {
	*(*[119]byte)(unsafe.Pointer(dst)) = *(*[119]byte)(unsafe.Pointer(src))
}

func Copy120(dst, src unsafe.Pointer) {
	*(*[120]byte)(unsafe.Pointer(dst)) = *(*[120]byte)(unsafe.Pointer(src))
}

func Copy121(dst, src unsafe.Pointer) {
	*(*[121]byte)(unsafe.Pointer(dst)) = *(*[121]byte)(unsafe.Pointer(src))
}

func Copy122(dst, src unsafe.Pointer) {
	*(*[122]byte)(unsafe.Pointer(dst)) = *(*[122]byte)(unsafe.Pointer(src))
}

func Copy123(dst, src unsafe.Pointer) {
	*(*[123]byte)(unsafe.Pointer(dst)) = *(*[123]byte)(unsafe.Pointer(src))
}

func Copy124(dst, src unsafe.Pointer) {
	*(*[124]byte)(unsafe.Pointer(dst)) = *(*[124]byte)(unsafe.Pointer(src))
}

func Copy125(dst, src unsafe.Pointer) {
	*(*[125]byte)(unsafe.Pointer(dst)) = *(*[125]byte)(unsafe.Pointer(src))
}

func Copy126(dst, src unsafe.Pointer) {
	*(*[126]byte)(unsafe.Pointer(dst)) = *(*[126]byte)(unsafe.Pointer(src))
}

func Copy127(dst, src unsafe.Pointer) {
	*(*[127]byte)(unsafe.Pointer(dst)) = *(*[127]byte)(unsafe.Pointer(src))
}

func Copy128(dst, src unsafe.Pointer) {
	*(*[128]byte)(unsafe.Pointer(dst)) = *(*[128]byte)(unsafe.Pointer(src))
}

func Copy129(dst, src unsafe.Pointer) {
	*(*[129]byte)(unsafe.Pointer(dst)) = *(*[129]byte)(unsafe.Pointer(src))
}

func Copy130(dst, src unsafe.Pointer) {
	*(*[130]byte)(unsafe.Pointer(dst)) = *(*[130]byte)(unsafe.Pointer(src))
}

func Copy131(dst, src unsafe.Pointer) {
	*(*[131]byte)(unsafe.Pointer(dst)) = *(*[131]byte)(unsafe.Pointer(src))
}

func Copy132(dst, src unsafe.Pointer) {
	*(*[132]byte)(unsafe.Pointer(dst)) = *(*[132]byte)(unsafe.Pointer(src))
}

func Copy133(dst, src unsafe.Pointer) {
	*(*[133]byte)(unsafe.Pointer(dst)) = *(*[133]byte)(unsafe.Pointer(src))
}

func Copy134(dst, src unsafe.Pointer) {
	*(*[134]byte)(unsafe.Pointer(dst)) = *(*[134]byte)(unsafe.Pointer(src))
}

func Copy135(dst, src unsafe.Pointer) {
	*(*[135]byte)(unsafe.Pointer(dst)) = *(*[135]byte)(unsafe.Pointer(src))
}

func Copy136(dst, src unsafe.Pointer) {
	*(*[136]byte)(unsafe.Pointer(dst)) = *(*[136]byte)(unsafe.Pointer(src))
}

func Copy137(dst, src unsafe.Pointer) {
	*(*[137]byte)(unsafe.Pointer(dst)) = *(*[137]byte)(unsafe.Pointer(src))
}

func Copy138(dst, src unsafe.Pointer) {
	*(*[138]byte)(unsafe.Pointer(dst)) = *(*[138]byte)(unsafe.Pointer(src))
}

func Copy139(dst, src unsafe.Pointer) {
	*(*[139]byte)(unsafe.Pointer(dst)) = *(*[139]byte)(unsafe.Pointer(src))
}

func Copy140(dst, src unsafe.Pointer) {
	*(*[140]byte)(unsafe.Pointer(dst)) = *(*[140]byte)(unsafe.Pointer(src))
}

func Copy141(dst, src unsafe.Pointer) {
	*(*[141]byte)(unsafe.Pointer(dst)) = *(*[141]byte)(unsafe.Pointer(src))
}

func Copy142(dst, src unsafe.Pointer) {
	*(*[142]byte)(unsafe.Pointer(dst)) = *(*[142]byte)(unsafe.Pointer(src))
}

func Copy143(dst, src unsafe.Pointer) {
	*(*[143]byte)(unsafe.Pointer(dst)) = *(*[143]byte)(unsafe.Pointer(src))
}

func Copy144(dst, src unsafe.Pointer) {
	*(*[144]byte)(unsafe.Pointer(dst)) = *(*[144]byte)(unsafe.Pointer(src))
}

func Copy145(dst, src unsafe.Pointer) {
	*(*[145]byte)(unsafe.Pointer(dst)) = *(*[145]byte)(unsafe.Pointer(src))
}

func Copy146(dst, src unsafe.Pointer) {
	*(*[146]byte)(unsafe.Pointer(dst)) = *(*[146]byte)(unsafe.Pointer(src))
}

func Copy147(dst, src unsafe.Pointer) {
	*(*[147]byte)(unsafe.Pointer(dst)) = *(*[147]byte)(unsafe.Pointer(src))
}

func Copy148(dst, src unsafe.Pointer) {
	*(*[148]byte)(unsafe.Pointer(dst)) = *(*[148]byte)(unsafe.Pointer(src))
}

func Copy149(dst, src unsafe.Pointer) {
	*(*[149]byte)(unsafe.Pointer(dst)) = *(*[149]byte)(unsafe.Pointer(src))
}

func Copy150(dst, src unsafe.Pointer) {
	*(*[150]byte)(unsafe.Pointer(dst)) = *(*[150]byte)(unsafe.Pointer(src))
}

func Copy151(dst, src unsafe.Pointer) {
	*(*[151]byte)(unsafe.Pointer(dst)) = *(*[151]byte)(unsafe.Pointer(src))
}

func Copy152(dst, src unsafe.Pointer) {
	*(*[152]byte)(unsafe.Pointer(dst)) = *(*[152]byte)(unsafe.Pointer(src))
}

func Copy153(dst, src unsafe.Pointer) {
	*(*[153]byte)(unsafe.Pointer(dst)) = *(*[153]byte)(unsafe.Pointer(src))
}

func Copy154(dst, src unsafe.Pointer) {
	*(*[154]byte)(unsafe.Pointer(dst)) = *(*[154]byte)(unsafe.Pointer(src))
}

func Copy155(dst, src unsafe.Pointer) {
	*(*[155]byte)(unsafe.Pointer(dst)) = *(*[155]byte)(unsafe.Pointer(src))
}

func Copy156(dst, src unsafe.Pointer) {
	*(*[156]byte)(unsafe.Pointer(dst)) = *(*[156]byte)(unsafe.Pointer(src))
}

func Copy157(dst, src unsafe.Pointer) {
	*(*[157]byte)(unsafe.Pointer(dst)) = *(*[157]byte)(unsafe.Pointer(src))
}

func Copy158(dst, src unsafe.Pointer) {
	*(*[158]byte)(unsafe.Pointer(dst)) = *(*[158]byte)(unsafe.Pointer(src))
}

func Copy159(dst, src unsafe.Pointer) {
	*(*[159]byte)(unsafe.Pointer(dst)) = *(*[159]byte)(unsafe.Pointer(src))
}

func Copy160(dst, src unsafe.Pointer) {
	*(*[160]byte)(unsafe.Pointer(dst)) = *(*[160]byte)(unsafe.Pointer(src))
}

func Copy161(dst, src unsafe.Pointer) {
	*(*[161]byte)(unsafe.Pointer(dst)) = *(*[161]byte)(unsafe.Pointer(src))
}

func Copy162(dst, src unsafe.Pointer) {
	*(*[162]byte)(unsafe.Pointer(dst)) = *(*[162]byte)(unsafe.Pointer(src))
}

func Copy163(dst, src unsafe.Pointer) {
	*(*[163]byte)(unsafe.Pointer(dst)) = *(*[163]byte)(unsafe.Pointer(src))
}

func Copy164(dst, src unsafe.Pointer) {
	*(*[164]byte)(unsafe.Pointer(dst)) = *(*[164]byte)(unsafe.Pointer(src))
}

func Copy165(dst, src unsafe.Pointer) {
	*(*[165]byte)(unsafe.Pointer(dst)) = *(*[165]byte)(unsafe.Pointer(src))
}

func Copy166(dst, src unsafe.Pointer) {
	*(*[166]byte)(unsafe.Pointer(dst)) = *(*[166]byte)(unsafe.Pointer(src))
}

func Copy167(dst, src unsafe.Pointer) {
	*(*[167]byte)(unsafe.Pointer(dst)) = *(*[167]byte)(unsafe.Pointer(src))
}

func Copy168(dst, src unsafe.Pointer) {
	*(*[168]byte)(unsafe.Pointer(dst)) = *(*[168]byte)(unsafe.Pointer(src))
}

func Copy169(dst, src unsafe.Pointer) {
	*(*[169]byte)(unsafe.Pointer(dst)) = *(*[169]byte)(unsafe.Pointer(src))
}

func Copy170(dst, src unsafe.Pointer) {
	*(*[170]byte)(unsafe.Pointer(dst)) = *(*[170]byte)(unsafe.Pointer(src))
}

func Copy171(dst, src unsafe.Pointer) {
	*(*[171]byte)(unsafe.Pointer(dst)) = *(*[171]byte)(unsafe.Pointer(src))
}

func Copy172(dst, src unsafe.Pointer) {
	*(*[172]byte)(unsafe.Pointer(dst)) = *(*[172]byte)(unsafe.Pointer(src))
}

func Copy173(dst, src unsafe.Pointer) {
	*(*[173]byte)(unsafe.Pointer(dst)) = *(*[173]byte)(unsafe.Pointer(src))
}

func Copy174(dst, src unsafe.Pointer) {
	*(*[174]byte)(unsafe.Pointer(dst)) = *(*[174]byte)(unsafe.Pointer(src))
}

func Copy175(dst, src unsafe.Pointer) {
	*(*[175]byte)(unsafe.Pointer(dst)) = *(*[175]byte)(unsafe.Pointer(src))
}

func Copy176(dst, src unsafe.Pointer) {
	*(*[176]byte)(unsafe.Pointer(dst)) = *(*[176]byte)(unsafe.Pointer(src))
}

func Copy177(dst, src unsafe.Pointer) {
	*(*[177]byte)(unsafe.Pointer(dst)) = *(*[177]byte)(unsafe.Pointer(src))
}

func Copy178(dst, src unsafe.Pointer) {
	*(*[178]byte)(unsafe.Pointer(dst)) = *(*[178]byte)(unsafe.Pointer(src))
}

func Copy179(dst, src unsafe.Pointer) {
	*(*[179]byte)(unsafe.Pointer(dst)) = *(*[179]byte)(unsafe.Pointer(src))
}

func Copy180(dst, src unsafe.Pointer) {
	*(*[180]byte)(unsafe.Pointer(dst)) = *(*[180]byte)(unsafe.Pointer(src))
}

func Copy181(dst, src unsafe.Pointer) {
	*(*[181]byte)(unsafe.Pointer(dst)) = *(*[181]byte)(unsafe.Pointer(src))
}

func Copy182(dst, src unsafe.Pointer) {
	*(*[182]byte)(unsafe.Pointer(dst)) = *(*[182]byte)(unsafe.Pointer(src))
}

func Copy183(dst, src unsafe.Pointer) {
	*(*[183]byte)(unsafe.Pointer(dst)) = *(*[183]byte)(unsafe.Pointer(src))
}

func Copy184(dst, src unsafe.Pointer) {
	*(*[184]byte)(unsafe.Pointer(dst)) = *(*[184]byte)(unsafe.Pointer(src))
}

func Copy185(dst, src unsafe.Pointer) {
	*(*[185]byte)(unsafe.Pointer(dst)) = *(*[185]byte)(unsafe.Pointer(src))
}

func Copy186(dst, src unsafe.Pointer) {
	*(*[186]byte)(unsafe.Pointer(dst)) = *(*[186]byte)(unsafe.Pointer(src))
}

func Copy187(dst, src unsafe.Pointer) {
	*(*[187]byte)(unsafe.Pointer(dst)) = *(*[187]byte)(unsafe.Pointer(src))
}

func Copy188(dst, src unsafe.Pointer) {
	*(*[188]byte)(unsafe.Pointer(dst)) = *(*[188]byte)(unsafe.Pointer(src))
}

func Copy189(dst, src unsafe.Pointer) {
	*(*[189]byte)(unsafe.Pointer(dst)) = *(*[189]byte)(unsafe.Pointer(src))
}

func Copy190(dst, src unsafe.Pointer) {
	*(*[190]byte)(unsafe.Pointer(dst)) = *(*[190]byte)(unsafe.Pointer(src))
}

func Copy191(dst, src unsafe.Pointer) {
	*(*[191]byte)(unsafe.Pointer(dst)) = *(*[191]byte)(unsafe.Pointer(src))
}

func Copy192(dst, src unsafe.Pointer) {
	*(*[192]byte)(unsafe.Pointer(dst)) = *(*[192]byte)(unsafe.Pointer(src))
}

func Copy193(dst, src unsafe.Pointer) {
	*(*[193]byte)(unsafe.Pointer(dst)) = *(*[193]byte)(unsafe.Pointer(src))
}

func Copy194(dst, src unsafe.Pointer) {
	*(*[194]byte)(unsafe.Pointer(dst)) = *(*[194]byte)(unsafe.Pointer(src))
}

func Copy195(dst, src unsafe.Pointer) {
	*(*[195]byte)(unsafe.Pointer(dst)) = *(*[195]byte)(unsafe.Pointer(src))
}

func Copy196(dst, src unsafe.Pointer) {
	*(*[196]byte)(unsafe.Pointer(dst)) = *(*[196]byte)(unsafe.Pointer(src))
}

func Copy197(dst, src unsafe.Pointer) {
	*(*[197]byte)(unsafe.Pointer(dst)) = *(*[197]byte)(unsafe.Pointer(src))
}

func Copy198(dst, src unsafe.Pointer) {
	*(*[198]byte)(unsafe.Pointer(dst)) = *(*[198]byte)(unsafe.Pointer(src))
}

func Copy199(dst, src unsafe.Pointer) {
	*(*[199]byte)(unsafe.Pointer(dst)) = *(*[199]byte)(unsafe.Pointer(src))
}

func Copy200(dst, src unsafe.Pointer) {
	*(*[200]byte)(unsafe.Pointer(dst)) = *(*[200]byte)(unsafe.Pointer(src))
}

func Copy201(dst, src unsafe.Pointer) {
	*(*[201]byte)(unsafe.Pointer(dst)) = *(*[201]byte)(unsafe.Pointer(src))
}

func Copy202(dst, src unsafe.Pointer) {
	*(*[202]byte)(unsafe.Pointer(dst)) = *(*[202]byte)(unsafe.Pointer(src))
}

func Copy203(dst, src unsafe.Pointer) {
	*(*[203]byte)(unsafe.Pointer(dst)) = *(*[203]byte)(unsafe.Pointer(src))
}

func Copy204(dst, src unsafe.Pointer) {
	*(*[204]byte)(unsafe.Pointer(dst)) = *(*[204]byte)(unsafe.Pointer(src))
}

func Copy205(dst, src unsafe.Pointer) {
	*(*[205]byte)(unsafe.Pointer(dst)) = *(*[205]byte)(unsafe.Pointer(src))
}

func Copy206(dst, src unsafe.Pointer) {
	*(*[206]byte)(unsafe.Pointer(dst)) = *(*[206]byte)(unsafe.Pointer(src))
}

func Copy207(dst, src unsafe.Pointer) {
	*(*[207]byte)(unsafe.Pointer(dst)) = *(*[207]byte)(unsafe.Pointer(src))
}

func Copy208(dst, src unsafe.Pointer) {
	*(*[208]byte)(unsafe.Pointer(dst)) = *(*[208]byte)(unsafe.Pointer(src))
}

func Copy209(dst, src unsafe.Pointer) {
	*(*[209]byte)(unsafe.Pointer(dst)) = *(*[209]byte)(unsafe.Pointer(src))
}

func Copy210(dst, src unsafe.Pointer) {
	*(*[210]byte)(unsafe.Pointer(dst)) = *(*[210]byte)(unsafe.Pointer(src))
}

func Copy211(dst, src unsafe.Pointer) {
	*(*[211]byte)(unsafe.Pointer(dst)) = *(*[211]byte)(unsafe.Pointer(src))
}

func Copy212(dst, src unsafe.Pointer) {
	*(*[212]byte)(unsafe.Pointer(dst)) = *(*[212]byte)(unsafe.Pointer(src))
}

func Copy213(dst, src unsafe.Pointer) {
	*(*[213]byte)(unsafe.Pointer(dst)) = *(*[213]byte)(unsafe.Pointer(src))
}

func Copy214(dst, src unsafe.Pointer) {
	*(*[214]byte)(unsafe.Pointer(dst)) = *(*[214]byte)(unsafe.Pointer(src))
}

func Copy215(dst, src unsafe.Pointer) {
	*(*[215]byte)(unsafe.Pointer(dst)) = *(*[215]byte)(unsafe.Pointer(src))
}

func Copy216(dst, src unsafe.Pointer) {
	*(*[216]byte)(unsafe.Pointer(dst)) = *(*[216]byte)(unsafe.Pointer(src))
}

func Copy217(dst, src unsafe.Pointer) {
	*(*[217]byte)(unsafe.Pointer(dst)) = *(*[217]byte)(unsafe.Pointer(src))
}

func Copy218(dst, src unsafe.Pointer) {
	*(*[218]byte)(unsafe.Pointer(dst)) = *(*[218]byte)(unsafe.Pointer(src))
}

func Copy219(dst, src unsafe.Pointer) {
	*(*[219]byte)(unsafe.Pointer(dst)) = *(*[219]byte)(unsafe.Pointer(src))
}

func Copy220(dst, src unsafe.Pointer) {
	*(*[220]byte)(unsafe.Pointer(dst)) = *(*[220]byte)(unsafe.Pointer(src))
}

func Copy221(dst, src unsafe.Pointer) {
	*(*[221]byte)(unsafe.Pointer(dst)) = *(*[221]byte)(unsafe.Pointer(src))
}

func Copy222(dst, src unsafe.Pointer) {
	*(*[222]byte)(unsafe.Pointer(dst)) = *(*[222]byte)(unsafe.Pointer(src))
}

func Copy223(dst, src unsafe.Pointer) {
	*(*[223]byte)(unsafe.Pointer(dst)) = *(*[223]byte)(unsafe.Pointer(src))
}

func Copy224(dst, src unsafe.Pointer) {
	*(*[224]byte)(unsafe.Pointer(dst)) = *(*[224]byte)(unsafe.Pointer(src))
}

func Copy225(dst, src unsafe.Pointer) {
	*(*[225]byte)(unsafe.Pointer(dst)) = *(*[225]byte)(unsafe.Pointer(src))
}

func Copy226(dst, src unsafe.Pointer) {
	*(*[226]byte)(unsafe.Pointer(dst)) = *(*[226]byte)(unsafe.Pointer(src))
}

func Copy227(dst, src unsafe.Pointer) {
	*(*[227]byte)(unsafe.Pointer(dst)) = *(*[227]byte)(unsafe.Pointer(src))
}

func Copy228(dst, src unsafe.Pointer) {
	*(*[228]byte)(unsafe.Pointer(dst)) = *(*[228]byte)(unsafe.Pointer(src))
}

func Copy229(dst, src unsafe.Pointer) {
	*(*[229]byte)(unsafe.Pointer(dst)) = *(*[229]byte)(unsafe.Pointer(src))
}

func Copy230(dst, src unsafe.Pointer) {
	*(*[230]byte)(unsafe.Pointer(dst)) = *(*[230]byte)(unsafe.Pointer(src))
}

func Copy231(dst, src unsafe.Pointer) {
	*(*[231]byte)(unsafe.Pointer(dst)) = *(*[231]byte)(unsafe.Pointer(src))
}

func Copy232(dst, src unsafe.Pointer) {
	*(*[232]byte)(unsafe.Pointer(dst)) = *(*[232]byte)(unsafe.Pointer(src))
}

func Copy233(dst, src unsafe.Pointer) {
	*(*[233]byte)(unsafe.Pointer(dst)) = *(*[233]byte)(unsafe.Pointer(src))
}

func Copy234(dst, src unsafe.Pointer) {
	*(*[234]byte)(unsafe.Pointer(dst)) = *(*[234]byte)(unsafe.Pointer(src))
}

func Copy235(dst, src unsafe.Pointer) {
	*(*[235]byte)(unsafe.Pointer(dst)) = *(*[235]byte)(unsafe.Pointer(src))
}

func Copy236(dst, src unsafe.Pointer) {
	*(*[236]byte)(unsafe.Pointer(dst)) = *(*[236]byte)(unsafe.Pointer(src))
}

func Copy237(dst, src unsafe.Pointer) {
	*(*[237]byte)(unsafe.Pointer(dst)) = *(*[237]byte)(unsafe.Pointer(src))
}

func Copy238(dst, src unsafe.Pointer) {
	*(*[238]byte)(unsafe.Pointer(dst)) = *(*[238]byte)(unsafe.Pointer(src))
}

func Copy239(dst, src unsafe.Pointer) {
	*(*[239]byte)(unsafe.Pointer(dst)) = *(*[239]byte)(unsafe.Pointer(src))
}

func Copy240(dst, src unsafe.Pointer) {
	*(*[240]byte)(unsafe.Pointer(dst)) = *(*[240]byte)(unsafe.Pointer(src))
}

func Copy241(dst, src unsafe.Pointer) {
	*(*[241]byte)(unsafe.Pointer(dst)) = *(*[241]byte)(unsafe.Pointer(src))
}

func Copy242(dst, src unsafe.Pointer) {
	*(*[242]byte)(unsafe.Pointer(dst)) = *(*[242]byte)(unsafe.Pointer(src))
}

func Copy243(dst, src unsafe.Pointer) {
	*(*[243]byte)(unsafe.Pointer(dst)) = *(*[243]byte)(unsafe.Pointer(src))
}

func Copy244(dst, src unsafe.Pointer) {
	*(*[244]byte)(unsafe.Pointer(dst)) = *(*[244]byte)(unsafe.Pointer(src))
}

func Copy245(dst, src unsafe.Pointer) {
	*(*[245]byte)(unsafe.Pointer(dst)) = *(*[245]byte)(unsafe.Pointer(src))
}

func Copy246(dst, src unsafe.Pointer) {
	*(*[246]byte)(unsafe.Pointer(dst)) = *(*[246]byte)(unsafe.Pointer(src))
}

func Copy247(dst, src unsafe.Pointer) {
	*(*[247]byte)(unsafe.Pointer(dst)) = *(*[247]byte)(unsafe.Pointer(src))
}

func Copy248(dst, src unsafe.Pointer) {
	*(*[248]byte)(unsafe.Pointer(dst)) = *(*[248]byte)(unsafe.Pointer(src))
}

func Copy249(dst, src unsafe.Pointer) {
	*(*[249]byte)(unsafe.Pointer(dst)) = *(*[249]byte)(unsafe.Pointer(src))
}

func Copy250(dst, src unsafe.Pointer) {
	*(*[250]byte)(unsafe.Pointer(dst)) = *(*[250]byte)(unsafe.Pointer(src))
}

func Copy251(dst, src unsafe.Pointer) {
	*(*[251]byte)(unsafe.Pointer(dst)) = *(*[251]byte)(unsafe.Pointer(src))
}

func Copy252(dst, src unsafe.Pointer) {
	*(*[252]byte)(unsafe.Pointer(dst)) = *(*[252]byte)(unsafe.Pointer(src))
}

func Copy253(dst, src unsafe.Pointer) {
	*(*[253]byte)(unsafe.Pointer(dst)) = *(*[253]byte)(unsafe.Pointer(src))
}

func Copy254(dst, src unsafe.Pointer) {
	*(*[254]byte)(unsafe.Pointer(dst)) = *(*[254]byte)(unsafe.Pointer(src))
}

func Copy255(dst, src unsafe.Pointer) {
	*(*[255]byte)(unsafe.Pointer(dst)) = *(*[255]byte)(unsafe.Pointer(src))
}
