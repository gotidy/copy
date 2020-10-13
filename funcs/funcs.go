
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
		// int to int
		{Src: typeOf(int(0)), Dest: typeOf(int(0))}:    CopyIntToInt,
		{Src: typeOfPointer(int(0)), Dest: typeOf(int(0))}:    CopyPIntToInt,
		{Src: typeOf(int(0)), Dest: typeOfPointer(int(0))}:    CopyIntToPInt,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(int(0))}:    CopyPIntToPInt, 
		// int8 to int
		{Src: typeOf(int8(0)), Dest: typeOf(int(0))}:    CopyInt8ToInt,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(int(0))}:    CopyPInt8ToInt,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(int(0))}:    CopyInt8ToPInt,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(int(0))}:    CopyPInt8ToPInt, 
		// int16 to int
		{Src: typeOf(int16(0)), Dest: typeOf(int(0))}:    CopyInt16ToInt,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(int(0))}:    CopyPInt16ToInt,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(int(0))}:    CopyInt16ToPInt,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(int(0))}:    CopyPInt16ToPInt, 
		// int32 to int
		{Src: typeOf(int32(0)), Dest: typeOf(int(0))}:    CopyInt32ToInt,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(int(0))}:    CopyPInt32ToInt,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(int(0))}:    CopyInt32ToPInt,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(int(0))}:    CopyPInt32ToPInt, 
		// int64 to int
		{Src: typeOf(int64(0)), Dest: typeOf(int(0))}:    CopyInt64ToInt,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(int(0))}:    CopyPInt64ToInt,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(int(0))}:    CopyInt64ToPInt,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(int(0))}:    CopyPInt64ToPInt, 
		// uint to int
		{Src: typeOf(uint(0)), Dest: typeOf(int(0))}:    CopyUintToInt,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(int(0))}:    CopyPUintToInt,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(int(0))}:    CopyUintToPInt,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(int(0))}:    CopyPUintToPInt, 
		// uint8 to int
		{Src: typeOf(uint8(0)), Dest: typeOf(int(0))}:    CopyUint8ToInt,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(int(0))}:    CopyPUint8ToInt,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(int(0))}:    CopyUint8ToPInt,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(int(0))}:    CopyPUint8ToPInt, 
		// uint16 to int
		{Src: typeOf(uint16(0)), Dest: typeOf(int(0))}:    CopyUint16ToInt,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(int(0))}:    CopyPUint16ToInt,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(int(0))}:    CopyUint16ToPInt,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(int(0))}:    CopyPUint16ToPInt, 
		// uint32 to int
		{Src: typeOf(uint32(0)), Dest: typeOf(int(0))}:    CopyUint32ToInt,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(int(0))}:    CopyPUint32ToInt,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(int(0))}:    CopyUint32ToPInt,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(int(0))}:    CopyPUint32ToPInt, 
		// uint64 to int
		{Src: typeOf(uint64(0)), Dest: typeOf(int(0))}:    CopyUint64ToInt,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(int(0))}:    CopyPUint64ToInt,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(int(0))}:    CopyUint64ToPInt,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(int(0))}:    CopyPUint64ToPInt, 
		// int to int8
		{Src: typeOf(int(0)), Dest: typeOf(int8(0))}:    CopyIntToInt8,
		{Src: typeOfPointer(int(0)), Dest: typeOf(int8(0))}:    CopyPIntToInt8,
		{Src: typeOf(int(0)), Dest: typeOfPointer(int8(0))}:    CopyIntToPInt8,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(int8(0))}:    CopyPIntToPInt8, 
		// int8 to int8
		{Src: typeOf(int8(0)), Dest: typeOf(int8(0))}:    CopyInt8ToInt8,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(int8(0))}:    CopyPInt8ToInt8,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(int8(0))}:    CopyInt8ToPInt8,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(int8(0))}:    CopyPInt8ToPInt8, 
		// int16 to int8
		{Src: typeOf(int16(0)), Dest: typeOf(int8(0))}:    CopyInt16ToInt8,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(int8(0))}:    CopyPInt16ToInt8,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(int8(0))}:    CopyInt16ToPInt8,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(int8(0))}:    CopyPInt16ToPInt8, 
		// int32 to int8
		{Src: typeOf(int32(0)), Dest: typeOf(int8(0))}:    CopyInt32ToInt8,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(int8(0))}:    CopyPInt32ToInt8,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(int8(0))}:    CopyInt32ToPInt8,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(int8(0))}:    CopyPInt32ToPInt8, 
		// int64 to int8
		{Src: typeOf(int64(0)), Dest: typeOf(int8(0))}:    CopyInt64ToInt8,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(int8(0))}:    CopyPInt64ToInt8,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(int8(0))}:    CopyInt64ToPInt8,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(int8(0))}:    CopyPInt64ToPInt8, 
		// uint to int8
		{Src: typeOf(uint(0)), Dest: typeOf(int8(0))}:    CopyUintToInt8,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(int8(0))}:    CopyPUintToInt8,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(int8(0))}:    CopyUintToPInt8,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(int8(0))}:    CopyPUintToPInt8, 
		// uint8 to int8
		{Src: typeOf(uint8(0)), Dest: typeOf(int8(0))}:    CopyUint8ToInt8,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(int8(0))}:    CopyPUint8ToInt8,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(int8(0))}:    CopyUint8ToPInt8,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(int8(0))}:    CopyPUint8ToPInt8, 
		// uint16 to int8
		{Src: typeOf(uint16(0)), Dest: typeOf(int8(0))}:    CopyUint16ToInt8,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(int8(0))}:    CopyPUint16ToInt8,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(int8(0))}:    CopyUint16ToPInt8,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(int8(0))}:    CopyPUint16ToPInt8, 
		// uint32 to int8
		{Src: typeOf(uint32(0)), Dest: typeOf(int8(0))}:    CopyUint32ToInt8,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(int8(0))}:    CopyPUint32ToInt8,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(int8(0))}:    CopyUint32ToPInt8,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(int8(0))}:    CopyPUint32ToPInt8, 
		// uint64 to int8
		{Src: typeOf(uint64(0)), Dest: typeOf(int8(0))}:    CopyUint64ToInt8,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(int8(0))}:    CopyPUint64ToInt8,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(int8(0))}:    CopyUint64ToPInt8,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(int8(0))}:    CopyPUint64ToPInt8, 
		// int to int16
		{Src: typeOf(int(0)), Dest: typeOf(int16(0))}:    CopyIntToInt16,
		{Src: typeOfPointer(int(0)), Dest: typeOf(int16(0))}:    CopyPIntToInt16,
		{Src: typeOf(int(0)), Dest: typeOfPointer(int16(0))}:    CopyIntToPInt16,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(int16(0))}:    CopyPIntToPInt16, 
		// int8 to int16
		{Src: typeOf(int8(0)), Dest: typeOf(int16(0))}:    CopyInt8ToInt16,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(int16(0))}:    CopyPInt8ToInt16,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(int16(0))}:    CopyInt8ToPInt16,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(int16(0))}:    CopyPInt8ToPInt16, 
		// int16 to int16
		{Src: typeOf(int16(0)), Dest: typeOf(int16(0))}:    CopyInt16ToInt16,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(int16(0))}:    CopyPInt16ToInt16,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(int16(0))}:    CopyInt16ToPInt16,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(int16(0))}:    CopyPInt16ToPInt16, 
		// int32 to int16
		{Src: typeOf(int32(0)), Dest: typeOf(int16(0))}:    CopyInt32ToInt16,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(int16(0))}:    CopyPInt32ToInt16,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(int16(0))}:    CopyInt32ToPInt16,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(int16(0))}:    CopyPInt32ToPInt16, 
		// int64 to int16
		{Src: typeOf(int64(0)), Dest: typeOf(int16(0))}:    CopyInt64ToInt16,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(int16(0))}:    CopyPInt64ToInt16,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(int16(0))}:    CopyInt64ToPInt16,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(int16(0))}:    CopyPInt64ToPInt16, 
		// uint to int16
		{Src: typeOf(uint(0)), Dest: typeOf(int16(0))}:    CopyUintToInt16,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(int16(0))}:    CopyPUintToInt16,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(int16(0))}:    CopyUintToPInt16,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(int16(0))}:    CopyPUintToPInt16, 
		// uint8 to int16
		{Src: typeOf(uint8(0)), Dest: typeOf(int16(0))}:    CopyUint8ToInt16,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(int16(0))}:    CopyPUint8ToInt16,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(int16(0))}:    CopyUint8ToPInt16,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(int16(0))}:    CopyPUint8ToPInt16, 
		// uint16 to int16
		{Src: typeOf(uint16(0)), Dest: typeOf(int16(0))}:    CopyUint16ToInt16,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(int16(0))}:    CopyPUint16ToInt16,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(int16(0))}:    CopyUint16ToPInt16,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(int16(0))}:    CopyPUint16ToPInt16, 
		// uint32 to int16
		{Src: typeOf(uint32(0)), Dest: typeOf(int16(0))}:    CopyUint32ToInt16,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(int16(0))}:    CopyPUint32ToInt16,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(int16(0))}:    CopyUint32ToPInt16,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(int16(0))}:    CopyPUint32ToPInt16, 
		// uint64 to int16
		{Src: typeOf(uint64(0)), Dest: typeOf(int16(0))}:    CopyUint64ToInt16,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(int16(0))}:    CopyPUint64ToInt16,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(int16(0))}:    CopyUint64ToPInt16,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(int16(0))}:    CopyPUint64ToPInt16, 
		// int to int32
		{Src: typeOf(int(0)), Dest: typeOf(int32(0))}:    CopyIntToInt32,
		{Src: typeOfPointer(int(0)), Dest: typeOf(int32(0))}:    CopyPIntToInt32,
		{Src: typeOf(int(0)), Dest: typeOfPointer(int32(0))}:    CopyIntToPInt32,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(int32(0))}:    CopyPIntToPInt32, 
		// int8 to int32
		{Src: typeOf(int8(0)), Dest: typeOf(int32(0))}:    CopyInt8ToInt32,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(int32(0))}:    CopyPInt8ToInt32,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(int32(0))}:    CopyInt8ToPInt32,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(int32(0))}:    CopyPInt8ToPInt32, 
		// int16 to int32
		{Src: typeOf(int16(0)), Dest: typeOf(int32(0))}:    CopyInt16ToInt32,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(int32(0))}:    CopyPInt16ToInt32,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(int32(0))}:    CopyInt16ToPInt32,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(int32(0))}:    CopyPInt16ToPInt32, 
		// int32 to int32
		{Src: typeOf(int32(0)), Dest: typeOf(int32(0))}:    CopyInt32ToInt32,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(int32(0))}:    CopyPInt32ToInt32,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(int32(0))}:    CopyInt32ToPInt32,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(int32(0))}:    CopyPInt32ToPInt32, 
		// int64 to int32
		{Src: typeOf(int64(0)), Dest: typeOf(int32(0))}:    CopyInt64ToInt32,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(int32(0))}:    CopyPInt64ToInt32,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(int32(0))}:    CopyInt64ToPInt32,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(int32(0))}:    CopyPInt64ToPInt32, 
		// uint to int32
		{Src: typeOf(uint(0)), Dest: typeOf(int32(0))}:    CopyUintToInt32,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(int32(0))}:    CopyPUintToInt32,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(int32(0))}:    CopyUintToPInt32,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(int32(0))}:    CopyPUintToPInt32, 
		// uint8 to int32
		{Src: typeOf(uint8(0)), Dest: typeOf(int32(0))}:    CopyUint8ToInt32,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(int32(0))}:    CopyPUint8ToInt32,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(int32(0))}:    CopyUint8ToPInt32,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(int32(0))}:    CopyPUint8ToPInt32, 
		// uint16 to int32
		{Src: typeOf(uint16(0)), Dest: typeOf(int32(0))}:    CopyUint16ToInt32,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(int32(0))}:    CopyPUint16ToInt32,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(int32(0))}:    CopyUint16ToPInt32,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(int32(0))}:    CopyPUint16ToPInt32, 
		// uint32 to int32
		{Src: typeOf(uint32(0)), Dest: typeOf(int32(0))}:    CopyUint32ToInt32,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(int32(0))}:    CopyPUint32ToInt32,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(int32(0))}:    CopyUint32ToPInt32,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(int32(0))}:    CopyPUint32ToPInt32, 
		// uint64 to int32
		{Src: typeOf(uint64(0)), Dest: typeOf(int32(0))}:    CopyUint64ToInt32,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(int32(0))}:    CopyPUint64ToInt32,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(int32(0))}:    CopyUint64ToPInt32,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(int32(0))}:    CopyPUint64ToPInt32, 
		// int to int64
		{Src: typeOf(int(0)), Dest: typeOf(int64(0))}:    CopyIntToInt64,
		{Src: typeOfPointer(int(0)), Dest: typeOf(int64(0))}:    CopyPIntToInt64,
		{Src: typeOf(int(0)), Dest: typeOfPointer(int64(0))}:    CopyIntToPInt64,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(int64(0))}:    CopyPIntToPInt64, 
		// int8 to int64
		{Src: typeOf(int8(0)), Dest: typeOf(int64(0))}:    CopyInt8ToInt64,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(int64(0))}:    CopyPInt8ToInt64,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(int64(0))}:    CopyInt8ToPInt64,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(int64(0))}:    CopyPInt8ToPInt64, 
		// int16 to int64
		{Src: typeOf(int16(0)), Dest: typeOf(int64(0))}:    CopyInt16ToInt64,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(int64(0))}:    CopyPInt16ToInt64,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(int64(0))}:    CopyInt16ToPInt64,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(int64(0))}:    CopyPInt16ToPInt64, 
		// int32 to int64
		{Src: typeOf(int32(0)), Dest: typeOf(int64(0))}:    CopyInt32ToInt64,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(int64(0))}:    CopyPInt32ToInt64,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(int64(0))}:    CopyInt32ToPInt64,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(int64(0))}:    CopyPInt32ToPInt64, 
		// int64 to int64
		{Src: typeOf(int64(0)), Dest: typeOf(int64(0))}:    CopyInt64ToInt64,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(int64(0))}:    CopyPInt64ToInt64,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(int64(0))}:    CopyInt64ToPInt64,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(int64(0))}:    CopyPInt64ToPInt64, 
		// uint to int64
		{Src: typeOf(uint(0)), Dest: typeOf(int64(0))}:    CopyUintToInt64,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(int64(0))}:    CopyPUintToInt64,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(int64(0))}:    CopyUintToPInt64,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(int64(0))}:    CopyPUintToPInt64, 
		// uint8 to int64
		{Src: typeOf(uint8(0)), Dest: typeOf(int64(0))}:    CopyUint8ToInt64,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(int64(0))}:    CopyPUint8ToInt64,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(int64(0))}:    CopyUint8ToPInt64,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(int64(0))}:    CopyPUint8ToPInt64, 
		// uint16 to int64
		{Src: typeOf(uint16(0)), Dest: typeOf(int64(0))}:    CopyUint16ToInt64,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(int64(0))}:    CopyPUint16ToInt64,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(int64(0))}:    CopyUint16ToPInt64,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(int64(0))}:    CopyPUint16ToPInt64, 
		// uint32 to int64
		{Src: typeOf(uint32(0)), Dest: typeOf(int64(0))}:    CopyUint32ToInt64,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(int64(0))}:    CopyPUint32ToInt64,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(int64(0))}:    CopyUint32ToPInt64,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(int64(0))}:    CopyPUint32ToPInt64, 
		// uint64 to int64
		{Src: typeOf(uint64(0)), Dest: typeOf(int64(0))}:    CopyUint64ToInt64,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(int64(0))}:    CopyPUint64ToInt64,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(int64(0))}:    CopyUint64ToPInt64,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(int64(0))}:    CopyPUint64ToPInt64, 
		// int to uint
		{Src: typeOf(int(0)), Dest: typeOf(uint(0))}:    CopyIntToUint,
		{Src: typeOfPointer(int(0)), Dest: typeOf(uint(0))}:    CopyPIntToUint,
		{Src: typeOf(int(0)), Dest: typeOfPointer(uint(0))}:    CopyIntToPUint,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(uint(0))}:    CopyPIntToPUint, 
		// int8 to uint
		{Src: typeOf(int8(0)), Dest: typeOf(uint(0))}:    CopyInt8ToUint,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(uint(0))}:    CopyPInt8ToUint,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(uint(0))}:    CopyInt8ToPUint,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(uint(0))}:    CopyPInt8ToPUint, 
		// int16 to uint
		{Src: typeOf(int16(0)), Dest: typeOf(uint(0))}:    CopyInt16ToUint,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(uint(0))}:    CopyPInt16ToUint,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(uint(0))}:    CopyInt16ToPUint,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(uint(0))}:    CopyPInt16ToPUint, 
		// int32 to uint
		{Src: typeOf(int32(0)), Dest: typeOf(uint(0))}:    CopyInt32ToUint,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(uint(0))}:    CopyPInt32ToUint,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(uint(0))}:    CopyInt32ToPUint,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(uint(0))}:    CopyPInt32ToPUint, 
		// int64 to uint
		{Src: typeOf(int64(0)), Dest: typeOf(uint(0))}:    CopyInt64ToUint,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(uint(0))}:    CopyPInt64ToUint,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(uint(0))}:    CopyInt64ToPUint,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(uint(0))}:    CopyPInt64ToPUint, 
		// uint to uint
		{Src: typeOf(uint(0)), Dest: typeOf(uint(0))}:    CopyUintToUint,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(uint(0))}:    CopyPUintToUint,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(uint(0))}:    CopyUintToPUint,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(uint(0))}:    CopyPUintToPUint, 
		// uint8 to uint
		{Src: typeOf(uint8(0)), Dest: typeOf(uint(0))}:    CopyUint8ToUint,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(uint(0))}:    CopyPUint8ToUint,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(uint(0))}:    CopyUint8ToPUint,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(uint(0))}:    CopyPUint8ToPUint, 
		// uint16 to uint
		{Src: typeOf(uint16(0)), Dest: typeOf(uint(0))}:    CopyUint16ToUint,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(uint(0))}:    CopyPUint16ToUint,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(uint(0))}:    CopyUint16ToPUint,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(uint(0))}:    CopyPUint16ToPUint, 
		// uint32 to uint
		{Src: typeOf(uint32(0)), Dest: typeOf(uint(0))}:    CopyUint32ToUint,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(uint(0))}:    CopyPUint32ToUint,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(uint(0))}:    CopyUint32ToPUint,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(uint(0))}:    CopyPUint32ToPUint, 
		// uint64 to uint
		{Src: typeOf(uint64(0)), Dest: typeOf(uint(0))}:    CopyUint64ToUint,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(uint(0))}:    CopyPUint64ToUint,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(uint(0))}:    CopyUint64ToPUint,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(uint(0))}:    CopyPUint64ToPUint, 
		// int to uint8
		{Src: typeOf(int(0)), Dest: typeOf(uint8(0))}:    CopyIntToUint8,
		{Src: typeOfPointer(int(0)), Dest: typeOf(uint8(0))}:    CopyPIntToUint8,
		{Src: typeOf(int(0)), Dest: typeOfPointer(uint8(0))}:    CopyIntToPUint8,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(uint8(0))}:    CopyPIntToPUint8, 
		// int8 to uint8
		{Src: typeOf(int8(0)), Dest: typeOf(uint8(0))}:    CopyInt8ToUint8,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(uint8(0))}:    CopyPInt8ToUint8,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(uint8(0))}:    CopyInt8ToPUint8,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(uint8(0))}:    CopyPInt8ToPUint8, 
		// int16 to uint8
		{Src: typeOf(int16(0)), Dest: typeOf(uint8(0))}:    CopyInt16ToUint8,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(uint8(0))}:    CopyPInt16ToUint8,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(uint8(0))}:    CopyInt16ToPUint8,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(uint8(0))}:    CopyPInt16ToPUint8, 
		// int32 to uint8
		{Src: typeOf(int32(0)), Dest: typeOf(uint8(0))}:    CopyInt32ToUint8,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(uint8(0))}:    CopyPInt32ToUint8,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(uint8(0))}:    CopyInt32ToPUint8,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(uint8(0))}:    CopyPInt32ToPUint8, 
		// int64 to uint8
		{Src: typeOf(int64(0)), Dest: typeOf(uint8(0))}:    CopyInt64ToUint8,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(uint8(0))}:    CopyPInt64ToUint8,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(uint8(0))}:    CopyInt64ToPUint8,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(uint8(0))}:    CopyPInt64ToPUint8, 
		// uint to uint8
		{Src: typeOf(uint(0)), Dest: typeOf(uint8(0))}:    CopyUintToUint8,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(uint8(0))}:    CopyPUintToUint8,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(uint8(0))}:    CopyUintToPUint8,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(uint8(0))}:    CopyPUintToPUint8, 
		// uint8 to uint8
		{Src: typeOf(uint8(0)), Dest: typeOf(uint8(0))}:    CopyUint8ToUint8,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(uint8(0))}:    CopyPUint8ToUint8,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(uint8(0))}:    CopyUint8ToPUint8,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(uint8(0))}:    CopyPUint8ToPUint8, 
		// uint16 to uint8
		{Src: typeOf(uint16(0)), Dest: typeOf(uint8(0))}:    CopyUint16ToUint8,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(uint8(0))}:    CopyPUint16ToUint8,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(uint8(0))}:    CopyUint16ToPUint8,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(uint8(0))}:    CopyPUint16ToPUint8, 
		// uint32 to uint8
		{Src: typeOf(uint32(0)), Dest: typeOf(uint8(0))}:    CopyUint32ToUint8,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(uint8(0))}:    CopyPUint32ToUint8,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(uint8(0))}:    CopyUint32ToPUint8,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(uint8(0))}:    CopyPUint32ToPUint8, 
		// uint64 to uint8
		{Src: typeOf(uint64(0)), Dest: typeOf(uint8(0))}:    CopyUint64ToUint8,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(uint8(0))}:    CopyPUint64ToUint8,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(uint8(0))}:    CopyUint64ToPUint8,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(uint8(0))}:    CopyPUint64ToPUint8, 
		// int to uint16
		{Src: typeOf(int(0)), Dest: typeOf(uint16(0))}:    CopyIntToUint16,
		{Src: typeOfPointer(int(0)), Dest: typeOf(uint16(0))}:    CopyPIntToUint16,
		{Src: typeOf(int(0)), Dest: typeOfPointer(uint16(0))}:    CopyIntToPUint16,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(uint16(0))}:    CopyPIntToPUint16, 
		// int8 to uint16
		{Src: typeOf(int8(0)), Dest: typeOf(uint16(0))}:    CopyInt8ToUint16,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(uint16(0))}:    CopyPInt8ToUint16,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(uint16(0))}:    CopyInt8ToPUint16,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(uint16(0))}:    CopyPInt8ToPUint16, 
		// int16 to uint16
		{Src: typeOf(int16(0)), Dest: typeOf(uint16(0))}:    CopyInt16ToUint16,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(uint16(0))}:    CopyPInt16ToUint16,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(uint16(0))}:    CopyInt16ToPUint16,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(uint16(0))}:    CopyPInt16ToPUint16, 
		// int32 to uint16
		{Src: typeOf(int32(0)), Dest: typeOf(uint16(0))}:    CopyInt32ToUint16,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(uint16(0))}:    CopyPInt32ToUint16,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(uint16(0))}:    CopyInt32ToPUint16,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(uint16(0))}:    CopyPInt32ToPUint16, 
		// int64 to uint16
		{Src: typeOf(int64(0)), Dest: typeOf(uint16(0))}:    CopyInt64ToUint16,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(uint16(0))}:    CopyPInt64ToUint16,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(uint16(0))}:    CopyInt64ToPUint16,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(uint16(0))}:    CopyPInt64ToPUint16, 
		// uint to uint16
		{Src: typeOf(uint(0)), Dest: typeOf(uint16(0))}:    CopyUintToUint16,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(uint16(0))}:    CopyPUintToUint16,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(uint16(0))}:    CopyUintToPUint16,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(uint16(0))}:    CopyPUintToPUint16, 
		// uint8 to uint16
		{Src: typeOf(uint8(0)), Dest: typeOf(uint16(0))}:    CopyUint8ToUint16,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(uint16(0))}:    CopyPUint8ToUint16,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(uint16(0))}:    CopyUint8ToPUint16,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(uint16(0))}:    CopyPUint8ToPUint16, 
		// uint16 to uint16
		{Src: typeOf(uint16(0)), Dest: typeOf(uint16(0))}:    CopyUint16ToUint16,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(uint16(0))}:    CopyPUint16ToUint16,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(uint16(0))}:    CopyUint16ToPUint16,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(uint16(0))}:    CopyPUint16ToPUint16, 
		// uint32 to uint16
		{Src: typeOf(uint32(0)), Dest: typeOf(uint16(0))}:    CopyUint32ToUint16,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(uint16(0))}:    CopyPUint32ToUint16,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(uint16(0))}:    CopyUint32ToPUint16,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(uint16(0))}:    CopyPUint32ToPUint16, 
		// uint64 to uint16
		{Src: typeOf(uint64(0)), Dest: typeOf(uint16(0))}:    CopyUint64ToUint16,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(uint16(0))}:    CopyPUint64ToUint16,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(uint16(0))}:    CopyUint64ToPUint16,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(uint16(0))}:    CopyPUint64ToPUint16, 
		// int to uint32
		{Src: typeOf(int(0)), Dest: typeOf(uint32(0))}:    CopyIntToUint32,
		{Src: typeOfPointer(int(0)), Dest: typeOf(uint32(0))}:    CopyPIntToUint32,
		{Src: typeOf(int(0)), Dest: typeOfPointer(uint32(0))}:    CopyIntToPUint32,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(uint32(0))}:    CopyPIntToPUint32, 
		// int8 to uint32
		{Src: typeOf(int8(0)), Dest: typeOf(uint32(0))}:    CopyInt8ToUint32,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(uint32(0))}:    CopyPInt8ToUint32,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(uint32(0))}:    CopyInt8ToPUint32,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(uint32(0))}:    CopyPInt8ToPUint32, 
		// int16 to uint32
		{Src: typeOf(int16(0)), Dest: typeOf(uint32(0))}:    CopyInt16ToUint32,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(uint32(0))}:    CopyPInt16ToUint32,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(uint32(0))}:    CopyInt16ToPUint32,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(uint32(0))}:    CopyPInt16ToPUint32, 
		// int32 to uint32
		{Src: typeOf(int32(0)), Dest: typeOf(uint32(0))}:    CopyInt32ToUint32,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(uint32(0))}:    CopyPInt32ToUint32,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(uint32(0))}:    CopyInt32ToPUint32,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(uint32(0))}:    CopyPInt32ToPUint32, 
		// int64 to uint32
		{Src: typeOf(int64(0)), Dest: typeOf(uint32(0))}:    CopyInt64ToUint32,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(uint32(0))}:    CopyPInt64ToUint32,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(uint32(0))}:    CopyInt64ToPUint32,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(uint32(0))}:    CopyPInt64ToPUint32, 
		// uint to uint32
		{Src: typeOf(uint(0)), Dest: typeOf(uint32(0))}:    CopyUintToUint32,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(uint32(0))}:    CopyPUintToUint32,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(uint32(0))}:    CopyUintToPUint32,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(uint32(0))}:    CopyPUintToPUint32, 
		// uint8 to uint32
		{Src: typeOf(uint8(0)), Dest: typeOf(uint32(0))}:    CopyUint8ToUint32,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(uint32(0))}:    CopyPUint8ToUint32,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(uint32(0))}:    CopyUint8ToPUint32,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(uint32(0))}:    CopyPUint8ToPUint32, 
		// uint16 to uint32
		{Src: typeOf(uint16(0)), Dest: typeOf(uint32(0))}:    CopyUint16ToUint32,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(uint32(0))}:    CopyPUint16ToUint32,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(uint32(0))}:    CopyUint16ToPUint32,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(uint32(0))}:    CopyPUint16ToPUint32, 
		// uint32 to uint32
		{Src: typeOf(uint32(0)), Dest: typeOf(uint32(0))}:    CopyUint32ToUint32,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(uint32(0))}:    CopyPUint32ToUint32,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(uint32(0))}:    CopyUint32ToPUint32,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(uint32(0))}:    CopyPUint32ToPUint32, 
		// uint64 to uint32
		{Src: typeOf(uint64(0)), Dest: typeOf(uint32(0))}:    CopyUint64ToUint32,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(uint32(0))}:    CopyPUint64ToUint32,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(uint32(0))}:    CopyUint64ToPUint32,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(uint32(0))}:    CopyPUint64ToPUint32, 
		// int to uint64
		{Src: typeOf(int(0)), Dest: typeOf(uint64(0))}:    CopyIntToUint64,
		{Src: typeOfPointer(int(0)), Dest: typeOf(uint64(0))}:    CopyPIntToUint64,
		{Src: typeOf(int(0)), Dest: typeOfPointer(uint64(0))}:    CopyIntToPUint64,
		{Src: typeOfPointer(int(0)), Dest: typeOfPointer(uint64(0))}:    CopyPIntToPUint64, 
		// int8 to uint64
		{Src: typeOf(int8(0)), Dest: typeOf(uint64(0))}:    CopyInt8ToUint64,
		{Src: typeOfPointer(int8(0)), Dest: typeOf(uint64(0))}:    CopyPInt8ToUint64,
		{Src: typeOf(int8(0)), Dest: typeOfPointer(uint64(0))}:    CopyInt8ToPUint64,
		{Src: typeOfPointer(int8(0)), Dest: typeOfPointer(uint64(0))}:    CopyPInt8ToPUint64, 
		// int16 to uint64
		{Src: typeOf(int16(0)), Dest: typeOf(uint64(0))}:    CopyInt16ToUint64,
		{Src: typeOfPointer(int16(0)), Dest: typeOf(uint64(0))}:    CopyPInt16ToUint64,
		{Src: typeOf(int16(0)), Dest: typeOfPointer(uint64(0))}:    CopyInt16ToPUint64,
		{Src: typeOfPointer(int16(0)), Dest: typeOfPointer(uint64(0))}:    CopyPInt16ToPUint64, 
		// int32 to uint64
		{Src: typeOf(int32(0)), Dest: typeOf(uint64(0))}:    CopyInt32ToUint64,
		{Src: typeOfPointer(int32(0)), Dest: typeOf(uint64(0))}:    CopyPInt32ToUint64,
		{Src: typeOf(int32(0)), Dest: typeOfPointer(uint64(0))}:    CopyInt32ToPUint64,
		{Src: typeOfPointer(int32(0)), Dest: typeOfPointer(uint64(0))}:    CopyPInt32ToPUint64, 
		// int64 to uint64
		{Src: typeOf(int64(0)), Dest: typeOf(uint64(0))}:    CopyInt64ToUint64,
		{Src: typeOfPointer(int64(0)), Dest: typeOf(uint64(0))}:    CopyPInt64ToUint64,
		{Src: typeOf(int64(0)), Dest: typeOfPointer(uint64(0))}:    CopyInt64ToPUint64,
		{Src: typeOfPointer(int64(0)), Dest: typeOfPointer(uint64(0))}:    CopyPInt64ToPUint64, 
		// uint to uint64
		{Src: typeOf(uint(0)), Dest: typeOf(uint64(0))}:    CopyUintToUint64,
		{Src: typeOfPointer(uint(0)), Dest: typeOf(uint64(0))}:    CopyPUintToUint64,
		{Src: typeOf(uint(0)), Dest: typeOfPointer(uint64(0))}:    CopyUintToPUint64,
		{Src: typeOfPointer(uint(0)), Dest: typeOfPointer(uint64(0))}:    CopyPUintToPUint64, 
		// uint8 to uint64
		{Src: typeOf(uint8(0)), Dest: typeOf(uint64(0))}:    CopyUint8ToUint64,
		{Src: typeOfPointer(uint8(0)), Dest: typeOf(uint64(0))}:    CopyPUint8ToUint64,
		{Src: typeOf(uint8(0)), Dest: typeOfPointer(uint64(0))}:    CopyUint8ToPUint64,
		{Src: typeOfPointer(uint8(0)), Dest: typeOfPointer(uint64(0))}:    CopyPUint8ToPUint64, 
		// uint16 to uint64
		{Src: typeOf(uint16(0)), Dest: typeOf(uint64(0))}:    CopyUint16ToUint64,
		{Src: typeOfPointer(uint16(0)), Dest: typeOf(uint64(0))}:    CopyPUint16ToUint64,
		{Src: typeOf(uint16(0)), Dest: typeOfPointer(uint64(0))}:    CopyUint16ToPUint64,
		{Src: typeOfPointer(uint16(0)), Dest: typeOfPointer(uint64(0))}:    CopyPUint16ToPUint64, 
		// uint32 to uint64
		{Src: typeOf(uint32(0)), Dest: typeOf(uint64(0))}:    CopyUint32ToUint64,
		{Src: typeOfPointer(uint32(0)), Dest: typeOf(uint64(0))}:    CopyPUint32ToUint64,
		{Src: typeOf(uint32(0)), Dest: typeOfPointer(uint64(0))}:    CopyUint32ToPUint64,
		{Src: typeOfPointer(uint32(0)), Dest: typeOfPointer(uint64(0))}:    CopyPUint32ToPUint64, 
		// uint64 to uint64
		{Src: typeOf(uint64(0)), Dest: typeOf(uint64(0))}:    CopyUint64ToUint64,
		{Src: typeOfPointer(uint64(0)), Dest: typeOf(uint64(0))}:    CopyPUint64ToUint64,
		{Src: typeOf(uint64(0)), Dest: typeOfPointer(uint64(0))}:    CopyUint64ToPUint64,
		{Src: typeOfPointer(uint64(0)), Dest: typeOfPointer(uint64(0))}:    CopyPUint64ToPUint64, 
		// float32 to float32
		{Src: typeOf(float32(0)), Dest: typeOf(float32(0))}:    CopyFloat32ToFloat32,
		{Src: typeOfPointer(float32(0)), Dest: typeOf(float32(0))}:    CopyPFloat32ToFloat32,
		{Src: typeOf(float32(0)), Dest: typeOfPointer(float32(0))}:    CopyFloat32ToPFloat32,
		{Src: typeOfPointer(float32(0)), Dest: typeOfPointer(float32(0))}:    CopyPFloat32ToPFloat32, 
		// float64 to float32
		{Src: typeOf(float64(0)), Dest: typeOf(float32(0))}:    CopyFloat64ToFloat32,
		{Src: typeOfPointer(float64(0)), Dest: typeOf(float32(0))}:    CopyPFloat64ToFloat32,
		{Src: typeOf(float64(0)), Dest: typeOfPointer(float32(0))}:    CopyFloat64ToPFloat32,
		{Src: typeOfPointer(float64(0)), Dest: typeOfPointer(float32(0))}:    CopyPFloat64ToPFloat32, 
		// float32 to float64
		{Src: typeOf(float32(0)), Dest: typeOf(float64(0))}:    CopyFloat32ToFloat64,
		{Src: typeOfPointer(float32(0)), Dest: typeOf(float64(0))}:    CopyPFloat32ToFloat64,
		{Src: typeOf(float32(0)), Dest: typeOfPointer(float64(0))}:    CopyFloat32ToPFloat64,
		{Src: typeOfPointer(float32(0)), Dest: typeOfPointer(float64(0))}:    CopyPFloat32ToPFloat64, 
		// float64 to float64
		{Src: typeOf(float64(0)), Dest: typeOf(float64(0))}:    CopyFloat64ToFloat64,
		{Src: typeOfPointer(float64(0)), Dest: typeOf(float64(0))}:    CopyPFloat64ToFloat64,
		{Src: typeOf(float64(0)), Dest: typeOfPointer(float64(0))}:    CopyFloat64ToPFloat64,
		{Src: typeOfPointer(float64(0)), Dest: typeOfPointer(float64(0))}:    CopyPFloat64ToPFloat64, 
		// bool to bool
		{Src: typeOf(bool(false)), Dest: typeOf(bool(false))}:    CopyBoolToBool,
		{Src: typeOfPointer(bool(false)), Dest: typeOf(bool(false))}:    CopyPBoolToBool,
		{Src: typeOf(bool(false)), Dest: typeOfPointer(bool(false))}:    CopyBoolToPBool,
		{Src: typeOfPointer(bool(false)), Dest: typeOfPointer(bool(false))}:    CopyPBoolToPBool, 
		// complex64 to complex64
		{Src: typeOf(complex64(0)), Dest: typeOf(complex64(0))}:    CopyComplex64ToComplex64,
		{Src: typeOfPointer(complex64(0)), Dest: typeOf(complex64(0))}:    CopyPComplex64ToComplex64,
		{Src: typeOf(complex64(0)), Dest: typeOfPointer(complex64(0))}:    CopyComplex64ToPComplex64,
		{Src: typeOfPointer(complex64(0)), Dest: typeOfPointer(complex64(0))}:    CopyPComplex64ToPComplex64, 
		// complex128 to complex64
		{Src: typeOf(complex128(0)), Dest: typeOf(complex64(0))}:    CopyComplex128ToComplex64,
		{Src: typeOfPointer(complex128(0)), Dest: typeOf(complex64(0))}:    CopyPComplex128ToComplex64,
		{Src: typeOf(complex128(0)), Dest: typeOfPointer(complex64(0))}:    CopyComplex128ToPComplex64,
		{Src: typeOfPointer(complex128(0)), Dest: typeOfPointer(complex64(0))}:    CopyPComplex128ToPComplex64, 
		// complex64 to complex128
		{Src: typeOf(complex64(0)), Dest: typeOf(complex128(0))}:    CopyComplex64ToComplex128,
		{Src: typeOfPointer(complex64(0)), Dest: typeOf(complex128(0))}:    CopyPComplex64ToComplex128,
		{Src: typeOf(complex64(0)), Dest: typeOfPointer(complex128(0))}:    CopyComplex64ToPComplex128,
		{Src: typeOfPointer(complex64(0)), Dest: typeOfPointer(complex128(0))}:    CopyPComplex64ToPComplex128, 
		// complex128 to complex128
		{Src: typeOf(complex128(0)), Dest: typeOf(complex128(0))}:    CopyComplex128ToComplex128,
		{Src: typeOfPointer(complex128(0)), Dest: typeOf(complex128(0))}:    CopyPComplex128ToComplex128,
		{Src: typeOf(complex128(0)), Dest: typeOfPointer(complex128(0))}:    CopyComplex128ToPComplex128,
		{Src: typeOfPointer(complex128(0)), Dest: typeOfPointer(complex128(0))}:    CopyPComplex128ToPComplex128, 
		// string to string
		{Src: typeOf(string("")), Dest: typeOf(string(""))}:    CopyStringToString,
		{Src: typeOfPointer(string("")), Dest: typeOf(string(""))}:    CopyPStringToString,
		{Src: typeOf(string("")), Dest: typeOfPointer(string(""))}:    CopyStringToPString,
		{Src: typeOfPointer(string("")), Dest: typeOfPointer(string(""))}:    CopyPStringToPString, 
		// []byte to string
		{Src: typeOf([]byte(nil)), Dest: typeOf(string(""))}:    CopyBytesToString,
		{Src: typeOfPointer([]byte(nil)), Dest: typeOf(string(""))}:    CopyPBytesToString,
		{Src: typeOf([]byte(nil)), Dest: typeOfPointer(string(""))}:    CopyBytesToPString,
		{Src: typeOfPointer([]byte(nil)), Dest: typeOfPointer(string(""))}:    CopyPBytesToPString, 
		// string to []byte
		{Src: typeOf(string("")), Dest: typeOf([]byte(nil))}:    CopyStringToBytes,
		{Src: typeOfPointer(string("")), Dest: typeOf([]byte(nil))}:    CopyPStringToBytes,
		{Src: typeOf(string("")), Dest: typeOfPointer([]byte(nil))}:    CopyStringToPBytes,
		{Src: typeOfPointer(string("")), Dest: typeOfPointer([]byte(nil))}:    CopyPStringToPBytes, 
		// []byte to []byte
		{Src: typeOf([]byte(nil)), Dest: typeOf([]byte(nil))}:    CopyBytesToBytes,
		{Src: typeOfPointer([]byte(nil)), Dest: typeOf([]byte(nil))}:    CopyPBytesToBytes,
		{Src: typeOf([]byte(nil)), Dest: typeOfPointer([]byte(nil))}:    CopyBytesToPBytes,
		{Src: typeOfPointer([]byte(nil)), Dest: typeOfPointer([]byte(nil))}:    CopyPBytesToPBytes, 
		// time.Time to time.Time
		{Src: typeOf(time.Time(time.Time{})), Dest: typeOf(time.Time(time.Time{}))}:    CopyTimeToTime,
		{Src: typeOfPointer(time.Time(time.Time{})), Dest: typeOf(time.Time(time.Time{}))}:    CopyPTimeToTime,
		{Src: typeOf(time.Time(time.Time{})), Dest: typeOfPointer(time.Time(time.Time{}))}:    CopyTimeToPTime,
		{Src: typeOfPointer(time.Time(time.Time{})), Dest: typeOfPointer(time.Time(time.Time{}))}:    CopyPTimeToPTime, 
		// time.Duration to time.Duration
		{Src: typeOf(time.Duration(0)), Dest: typeOf(time.Duration(0))}:    CopyDurationToDuration,
		{Src: typeOfPointer(time.Duration(0)), Dest: typeOf(time.Duration(0))}:    CopyPDurationToDuration,
		{Src: typeOf(time.Duration(0)), Dest: typeOfPointer(time.Duration(0))}:    CopyDurationToPDuration,
		{Src: typeOfPointer(time.Duration(0)), Dest: typeOfPointer(time.Duration(0))}:    CopyPDurationToPDuration,	
	},
}
 

// int to int

func CopyIntToInt(src, dest unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dest)) = int(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dest)) = v
}

func CopyIntToPInt(src, dest unsafe.Pointer) {
	v := int(*(*int)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int8 to int

func CopyInt8ToInt(src, dest unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dest)) = int(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dest)) = v
}

func CopyInt8ToPInt(src, dest unsafe.Pointer) {
	v := int(*(*int8)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int16 to int

func CopyInt16ToInt(src, dest unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dest)) = int(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dest)) = v
}

func CopyInt16ToPInt(src, dest unsafe.Pointer) {
	v := int(*(*int16)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int32 to int

func CopyInt32ToInt(src, dest unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dest)) = int(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dest)) = v
}

func CopyInt32ToPInt(src, dest unsafe.Pointer) {
	v := int(*(*int32)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int64 to int

func CopyInt64ToInt(src, dest unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dest)) = int(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dest)) = v
}

func CopyInt64ToPInt(src, dest unsafe.Pointer) {
	v := int(*(*int64)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint to int

func CopyUintToInt(src, dest unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dest)) = int(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dest)) = v
}

func CopyUintToPInt(src, dest unsafe.Pointer) {
	v := int(*(*uint)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint8 to int

func CopyUint8ToInt(src, dest unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dest)) = int(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dest)) = v
}

func CopyUint8ToPInt(src, dest unsafe.Pointer) {
	v := int(*(*uint8)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint16 to int

func CopyUint16ToInt(src, dest unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dest)) = int(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dest)) = v
}

func CopyUint16ToPInt(src, dest unsafe.Pointer) {
	v := int(*(*uint16)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint32 to int

func CopyUint32ToInt(src, dest unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dest)) = int(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dest)) = v
}

func CopyUint32ToPInt(src, dest unsafe.Pointer) {
	v := int(*(*uint32)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint64 to int

func CopyUint64ToInt(src, dest unsafe.Pointer) {
	*(*int)(unsafe.Pointer(dest)) = int(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}
	*(*int)(unsafe.Pointer(dest)) = v
}

func CopyUint64ToPInt(src, dest unsafe.Pointer) {
	v := int(*(*uint64)(unsafe.Pointer(src)))
	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPInt(src, dest unsafe.Pointer) {
	var v int
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int(*p)
	}

	p := (**int)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int to int8

func CopyIntToInt8(src, dest unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dest)) = int8(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dest)) = v
}

func CopyIntToPInt8(src, dest unsafe.Pointer) {
	v := int8(*(*int)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int8 to int8

func CopyInt8ToInt8(src, dest unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dest)) = int8(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dest)) = v
}

func CopyInt8ToPInt8(src, dest unsafe.Pointer) {
	v := int8(*(*int8)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int16 to int8

func CopyInt16ToInt8(src, dest unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dest)) = int8(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dest)) = v
}

func CopyInt16ToPInt8(src, dest unsafe.Pointer) {
	v := int8(*(*int16)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int32 to int8

func CopyInt32ToInt8(src, dest unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dest)) = int8(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dest)) = v
}

func CopyInt32ToPInt8(src, dest unsafe.Pointer) {
	v := int8(*(*int32)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int64 to int8

func CopyInt64ToInt8(src, dest unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dest)) = int8(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dest)) = v
}

func CopyInt64ToPInt8(src, dest unsafe.Pointer) {
	v := int8(*(*int64)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint to int8

func CopyUintToInt8(src, dest unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dest)) = int8(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dest)) = v
}

func CopyUintToPInt8(src, dest unsafe.Pointer) {
	v := int8(*(*uint)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint8 to int8

func CopyUint8ToInt8(src, dest unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dest)) = int8(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dest)) = v
}

func CopyUint8ToPInt8(src, dest unsafe.Pointer) {
	v := int8(*(*uint8)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint16 to int8

func CopyUint16ToInt8(src, dest unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dest)) = int8(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dest)) = v
}

func CopyUint16ToPInt8(src, dest unsafe.Pointer) {
	v := int8(*(*uint16)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint32 to int8

func CopyUint32ToInt8(src, dest unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dest)) = int8(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dest)) = v
}

func CopyUint32ToPInt8(src, dest unsafe.Pointer) {
	v := int8(*(*uint32)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint64 to int8

func CopyUint64ToInt8(src, dest unsafe.Pointer) {
	*(*int8)(unsafe.Pointer(dest)) = int8(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}
	*(*int8)(unsafe.Pointer(dest)) = v
}

func CopyUint64ToPInt8(src, dest unsafe.Pointer) {
	v := int8(*(*uint64)(unsafe.Pointer(src)))
	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPInt8(src, dest unsafe.Pointer) {
	var v int8
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int8(*p)
	}

	p := (**int8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int to int16

func CopyIntToInt16(src, dest unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dest)) = int16(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dest)) = v
}

func CopyIntToPInt16(src, dest unsafe.Pointer) {
	v := int16(*(*int)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int8 to int16

func CopyInt8ToInt16(src, dest unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dest)) = int16(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dest)) = v
}

func CopyInt8ToPInt16(src, dest unsafe.Pointer) {
	v := int16(*(*int8)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int16 to int16

func CopyInt16ToInt16(src, dest unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dest)) = int16(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dest)) = v
}

func CopyInt16ToPInt16(src, dest unsafe.Pointer) {
	v := int16(*(*int16)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int32 to int16

func CopyInt32ToInt16(src, dest unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dest)) = int16(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dest)) = v
}

func CopyInt32ToPInt16(src, dest unsafe.Pointer) {
	v := int16(*(*int32)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int64 to int16

func CopyInt64ToInt16(src, dest unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dest)) = int16(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dest)) = v
}

func CopyInt64ToPInt16(src, dest unsafe.Pointer) {
	v := int16(*(*int64)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint to int16

func CopyUintToInt16(src, dest unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dest)) = int16(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dest)) = v
}

func CopyUintToPInt16(src, dest unsafe.Pointer) {
	v := int16(*(*uint)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint8 to int16

func CopyUint8ToInt16(src, dest unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dest)) = int16(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dest)) = v
}

func CopyUint8ToPInt16(src, dest unsafe.Pointer) {
	v := int16(*(*uint8)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint16 to int16

func CopyUint16ToInt16(src, dest unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dest)) = int16(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dest)) = v
}

func CopyUint16ToPInt16(src, dest unsafe.Pointer) {
	v := int16(*(*uint16)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint32 to int16

func CopyUint32ToInt16(src, dest unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dest)) = int16(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dest)) = v
}

func CopyUint32ToPInt16(src, dest unsafe.Pointer) {
	v := int16(*(*uint32)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint64 to int16

func CopyUint64ToInt16(src, dest unsafe.Pointer) {
	*(*int16)(unsafe.Pointer(dest)) = int16(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}
	*(*int16)(unsafe.Pointer(dest)) = v
}

func CopyUint64ToPInt16(src, dest unsafe.Pointer) {
	v := int16(*(*uint64)(unsafe.Pointer(src)))
	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPInt16(src, dest unsafe.Pointer) {
	var v int16
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int16(*p)
	}

	p := (**int16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int to int32

func CopyIntToInt32(src, dest unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dest)) = int32(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dest)) = v
}

func CopyIntToPInt32(src, dest unsafe.Pointer) {
	v := int32(*(*int)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int8 to int32

func CopyInt8ToInt32(src, dest unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dest)) = int32(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dest)) = v
}

func CopyInt8ToPInt32(src, dest unsafe.Pointer) {
	v := int32(*(*int8)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int16 to int32

func CopyInt16ToInt32(src, dest unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dest)) = int32(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dest)) = v
}

func CopyInt16ToPInt32(src, dest unsafe.Pointer) {
	v := int32(*(*int16)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int32 to int32

func CopyInt32ToInt32(src, dest unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dest)) = int32(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dest)) = v
}

func CopyInt32ToPInt32(src, dest unsafe.Pointer) {
	v := int32(*(*int32)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int64 to int32

func CopyInt64ToInt32(src, dest unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dest)) = int32(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dest)) = v
}

func CopyInt64ToPInt32(src, dest unsafe.Pointer) {
	v := int32(*(*int64)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint to int32

func CopyUintToInt32(src, dest unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dest)) = int32(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dest)) = v
}

func CopyUintToPInt32(src, dest unsafe.Pointer) {
	v := int32(*(*uint)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint8 to int32

func CopyUint8ToInt32(src, dest unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dest)) = int32(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dest)) = v
}

func CopyUint8ToPInt32(src, dest unsafe.Pointer) {
	v := int32(*(*uint8)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint16 to int32

func CopyUint16ToInt32(src, dest unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dest)) = int32(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dest)) = v
}

func CopyUint16ToPInt32(src, dest unsafe.Pointer) {
	v := int32(*(*uint16)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint32 to int32

func CopyUint32ToInt32(src, dest unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dest)) = int32(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dest)) = v
}

func CopyUint32ToPInt32(src, dest unsafe.Pointer) {
	v := int32(*(*uint32)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint64 to int32

func CopyUint64ToInt32(src, dest unsafe.Pointer) {
	*(*int32)(unsafe.Pointer(dest)) = int32(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}
	*(*int32)(unsafe.Pointer(dest)) = v
}

func CopyUint64ToPInt32(src, dest unsafe.Pointer) {
	v := int32(*(*uint64)(unsafe.Pointer(src)))
	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPInt32(src, dest unsafe.Pointer) {
	var v int32
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int32(*p)
	}

	p := (**int32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int to int64

func CopyIntToInt64(src, dest unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dest)) = int64(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dest)) = v
}

func CopyIntToPInt64(src, dest unsafe.Pointer) {
	v := int64(*(*int)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int8 to int64

func CopyInt8ToInt64(src, dest unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dest)) = int64(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dest)) = v
}

func CopyInt8ToPInt64(src, dest unsafe.Pointer) {
	v := int64(*(*int8)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int16 to int64

func CopyInt16ToInt64(src, dest unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dest)) = int64(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dest)) = v
}

func CopyInt16ToPInt64(src, dest unsafe.Pointer) {
	v := int64(*(*int16)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int32 to int64

func CopyInt32ToInt64(src, dest unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dest)) = int64(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dest)) = v
}

func CopyInt32ToPInt64(src, dest unsafe.Pointer) {
	v := int64(*(*int32)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int64 to int64

func CopyInt64ToInt64(src, dest unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dest)) = int64(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dest)) = v
}

func CopyInt64ToPInt64(src, dest unsafe.Pointer) {
	v := int64(*(*int64)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint to int64

func CopyUintToInt64(src, dest unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dest)) = int64(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dest)) = v
}

func CopyUintToPInt64(src, dest unsafe.Pointer) {
	v := int64(*(*uint)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint8 to int64

func CopyUint8ToInt64(src, dest unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dest)) = int64(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dest)) = v
}

func CopyUint8ToPInt64(src, dest unsafe.Pointer) {
	v := int64(*(*uint8)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint16 to int64

func CopyUint16ToInt64(src, dest unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dest)) = int64(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dest)) = v
}

func CopyUint16ToPInt64(src, dest unsafe.Pointer) {
	v := int64(*(*uint16)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint32 to int64

func CopyUint32ToInt64(src, dest unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dest)) = int64(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dest)) = v
}

func CopyUint32ToPInt64(src, dest unsafe.Pointer) {
	v := int64(*(*uint32)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint64 to int64

func CopyUint64ToInt64(src, dest unsafe.Pointer) {
	*(*int64)(unsafe.Pointer(dest)) = int64(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}
	*(*int64)(unsafe.Pointer(dest)) = v
}

func CopyUint64ToPInt64(src, dest unsafe.Pointer) {
	v := int64(*(*uint64)(unsafe.Pointer(src)))
	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPInt64(src, dest unsafe.Pointer) {
	var v int64
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = int64(*p)
	}

	p := (**int64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int to uint

func CopyIntToUint(src, dest unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dest)) = uint(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dest)) = v
}

func CopyIntToPUint(src, dest unsafe.Pointer) {
	v := uint(*(*int)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int8 to uint

func CopyInt8ToUint(src, dest unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dest)) = uint(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dest)) = v
}

func CopyInt8ToPUint(src, dest unsafe.Pointer) {
	v := uint(*(*int8)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int16 to uint

func CopyInt16ToUint(src, dest unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dest)) = uint(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dest)) = v
}

func CopyInt16ToPUint(src, dest unsafe.Pointer) {
	v := uint(*(*int16)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int32 to uint

func CopyInt32ToUint(src, dest unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dest)) = uint(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dest)) = v
}

func CopyInt32ToPUint(src, dest unsafe.Pointer) {
	v := uint(*(*int32)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int64 to uint

func CopyInt64ToUint(src, dest unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dest)) = uint(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dest)) = v
}

func CopyInt64ToPUint(src, dest unsafe.Pointer) {
	v := uint(*(*int64)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint to uint

func CopyUintToUint(src, dest unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dest)) = uint(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dest)) = v
}

func CopyUintToPUint(src, dest unsafe.Pointer) {
	v := uint(*(*uint)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint8 to uint

func CopyUint8ToUint(src, dest unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dest)) = uint(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dest)) = v
}

func CopyUint8ToPUint(src, dest unsafe.Pointer) {
	v := uint(*(*uint8)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint16 to uint

func CopyUint16ToUint(src, dest unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dest)) = uint(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dest)) = v
}

func CopyUint16ToPUint(src, dest unsafe.Pointer) {
	v := uint(*(*uint16)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint32 to uint

func CopyUint32ToUint(src, dest unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dest)) = uint(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dest)) = v
}

func CopyUint32ToPUint(src, dest unsafe.Pointer) {
	v := uint(*(*uint32)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint64 to uint

func CopyUint64ToUint(src, dest unsafe.Pointer) {
	*(*uint)(unsafe.Pointer(dest)) = uint(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}
	*(*uint)(unsafe.Pointer(dest)) = v
}

func CopyUint64ToPUint(src, dest unsafe.Pointer) {
	v := uint(*(*uint64)(unsafe.Pointer(src)))
	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPUint(src, dest unsafe.Pointer) {
	var v uint
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint(*p)
	}

	p := (**uint)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int to uint8

func CopyIntToUint8(src, dest unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dest)) = uint8(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dest)) = v
}

func CopyIntToPUint8(src, dest unsafe.Pointer) {
	v := uint8(*(*int)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int8 to uint8

func CopyInt8ToUint8(src, dest unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dest)) = uint8(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dest)) = v
}

func CopyInt8ToPUint8(src, dest unsafe.Pointer) {
	v := uint8(*(*int8)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int16 to uint8

func CopyInt16ToUint8(src, dest unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dest)) = uint8(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dest)) = v
}

func CopyInt16ToPUint8(src, dest unsafe.Pointer) {
	v := uint8(*(*int16)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int32 to uint8

func CopyInt32ToUint8(src, dest unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dest)) = uint8(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dest)) = v
}

func CopyInt32ToPUint8(src, dest unsafe.Pointer) {
	v := uint8(*(*int32)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int64 to uint8

func CopyInt64ToUint8(src, dest unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dest)) = uint8(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dest)) = v
}

func CopyInt64ToPUint8(src, dest unsafe.Pointer) {
	v := uint8(*(*int64)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint to uint8

func CopyUintToUint8(src, dest unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dest)) = uint8(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dest)) = v
}

func CopyUintToPUint8(src, dest unsafe.Pointer) {
	v := uint8(*(*uint)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint8 to uint8

func CopyUint8ToUint8(src, dest unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dest)) = uint8(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dest)) = v
}

func CopyUint8ToPUint8(src, dest unsafe.Pointer) {
	v := uint8(*(*uint8)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint16 to uint8

func CopyUint16ToUint8(src, dest unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dest)) = uint8(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dest)) = v
}

func CopyUint16ToPUint8(src, dest unsafe.Pointer) {
	v := uint8(*(*uint16)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint32 to uint8

func CopyUint32ToUint8(src, dest unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dest)) = uint8(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dest)) = v
}

func CopyUint32ToPUint8(src, dest unsafe.Pointer) {
	v := uint8(*(*uint32)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint64 to uint8

func CopyUint64ToUint8(src, dest unsafe.Pointer) {
	*(*uint8)(unsafe.Pointer(dest)) = uint8(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}
	*(*uint8)(unsafe.Pointer(dest)) = v
}

func CopyUint64ToPUint8(src, dest unsafe.Pointer) {
	v := uint8(*(*uint64)(unsafe.Pointer(src)))
	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPUint8(src, dest unsafe.Pointer) {
	var v uint8
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint8(*p)
	}

	p := (**uint8)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int to uint16

func CopyIntToUint16(src, dest unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dest)) = uint16(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dest)) = v
}

func CopyIntToPUint16(src, dest unsafe.Pointer) {
	v := uint16(*(*int)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int8 to uint16

func CopyInt8ToUint16(src, dest unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dest)) = uint16(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dest)) = v
}

func CopyInt8ToPUint16(src, dest unsafe.Pointer) {
	v := uint16(*(*int8)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int16 to uint16

func CopyInt16ToUint16(src, dest unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dest)) = uint16(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dest)) = v
}

func CopyInt16ToPUint16(src, dest unsafe.Pointer) {
	v := uint16(*(*int16)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int32 to uint16

func CopyInt32ToUint16(src, dest unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dest)) = uint16(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dest)) = v
}

func CopyInt32ToPUint16(src, dest unsafe.Pointer) {
	v := uint16(*(*int32)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int64 to uint16

func CopyInt64ToUint16(src, dest unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dest)) = uint16(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dest)) = v
}

func CopyInt64ToPUint16(src, dest unsafe.Pointer) {
	v := uint16(*(*int64)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint to uint16

func CopyUintToUint16(src, dest unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dest)) = uint16(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dest)) = v
}

func CopyUintToPUint16(src, dest unsafe.Pointer) {
	v := uint16(*(*uint)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint8 to uint16

func CopyUint8ToUint16(src, dest unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dest)) = uint16(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dest)) = v
}

func CopyUint8ToPUint16(src, dest unsafe.Pointer) {
	v := uint16(*(*uint8)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint16 to uint16

func CopyUint16ToUint16(src, dest unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dest)) = uint16(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dest)) = v
}

func CopyUint16ToPUint16(src, dest unsafe.Pointer) {
	v := uint16(*(*uint16)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint32 to uint16

func CopyUint32ToUint16(src, dest unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dest)) = uint16(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dest)) = v
}

func CopyUint32ToPUint16(src, dest unsafe.Pointer) {
	v := uint16(*(*uint32)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint64 to uint16

func CopyUint64ToUint16(src, dest unsafe.Pointer) {
	*(*uint16)(unsafe.Pointer(dest)) = uint16(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}
	*(*uint16)(unsafe.Pointer(dest)) = v
}

func CopyUint64ToPUint16(src, dest unsafe.Pointer) {
	v := uint16(*(*uint64)(unsafe.Pointer(src)))
	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPUint16(src, dest unsafe.Pointer) {
	var v uint16
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint16(*p)
	}

	p := (**uint16)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int to uint32

func CopyIntToUint32(src, dest unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dest)) = uint32(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dest)) = v
}

func CopyIntToPUint32(src, dest unsafe.Pointer) {
	v := uint32(*(*int)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int8 to uint32

func CopyInt8ToUint32(src, dest unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dest)) = uint32(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dest)) = v
}

func CopyInt8ToPUint32(src, dest unsafe.Pointer) {
	v := uint32(*(*int8)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int16 to uint32

func CopyInt16ToUint32(src, dest unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dest)) = uint32(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dest)) = v
}

func CopyInt16ToPUint32(src, dest unsafe.Pointer) {
	v := uint32(*(*int16)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int32 to uint32

func CopyInt32ToUint32(src, dest unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dest)) = uint32(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dest)) = v
}

func CopyInt32ToPUint32(src, dest unsafe.Pointer) {
	v := uint32(*(*int32)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int64 to uint32

func CopyInt64ToUint32(src, dest unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dest)) = uint32(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dest)) = v
}

func CopyInt64ToPUint32(src, dest unsafe.Pointer) {
	v := uint32(*(*int64)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint to uint32

func CopyUintToUint32(src, dest unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dest)) = uint32(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dest)) = v
}

func CopyUintToPUint32(src, dest unsafe.Pointer) {
	v := uint32(*(*uint)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint8 to uint32

func CopyUint8ToUint32(src, dest unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dest)) = uint32(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dest)) = v
}

func CopyUint8ToPUint32(src, dest unsafe.Pointer) {
	v := uint32(*(*uint8)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint16 to uint32

func CopyUint16ToUint32(src, dest unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dest)) = uint32(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dest)) = v
}

func CopyUint16ToPUint32(src, dest unsafe.Pointer) {
	v := uint32(*(*uint16)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint32 to uint32

func CopyUint32ToUint32(src, dest unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dest)) = uint32(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dest)) = v
}

func CopyUint32ToPUint32(src, dest unsafe.Pointer) {
	v := uint32(*(*uint32)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint64 to uint32

func CopyUint64ToUint32(src, dest unsafe.Pointer) {
	*(*uint32)(unsafe.Pointer(dest)) = uint32(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}
	*(*uint32)(unsafe.Pointer(dest)) = v
}

func CopyUint64ToPUint32(src, dest unsafe.Pointer) {
	v := uint32(*(*uint64)(unsafe.Pointer(src)))
	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPUint32(src, dest unsafe.Pointer) {
	var v uint32
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint32(*p)
	}

	p := (**uint32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int to uint64

func CopyIntToUint64(src, dest unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dest)) = uint64(*(*int)(unsafe.Pointer(src)))
}

func CopyPIntToUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dest)) = v
}

func CopyIntToPUint64(src, dest unsafe.Pointer) {
	v := uint64(*(*int)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPIntToPUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**int)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int8 to uint64

func CopyInt8ToUint64(src, dest unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dest)) = uint64(*(*int8)(unsafe.Pointer(src)))
}

func CopyPInt8ToUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dest)) = v
}

func CopyInt8ToPUint64(src, dest unsafe.Pointer) {
	v := uint64(*(*int8)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt8ToPUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**int8)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int16 to uint64

func CopyInt16ToUint64(src, dest unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dest)) = uint64(*(*int16)(unsafe.Pointer(src)))
}

func CopyPInt16ToUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dest)) = v
}

func CopyInt16ToPUint64(src, dest unsafe.Pointer) {
	v := uint64(*(*int16)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt16ToPUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**int16)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int32 to uint64

func CopyInt32ToUint64(src, dest unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dest)) = uint64(*(*int32)(unsafe.Pointer(src)))
}

func CopyPInt32ToUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dest)) = v
}

func CopyInt32ToPUint64(src, dest unsafe.Pointer) {
	v := uint64(*(*int32)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt32ToPUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**int32)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// int64 to uint64

func CopyInt64ToUint64(src, dest unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dest)) = uint64(*(*int64)(unsafe.Pointer(src)))
}

func CopyPInt64ToUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dest)) = v
}

func CopyInt64ToPUint64(src, dest unsafe.Pointer) {
	v := uint64(*(*int64)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPInt64ToPUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**int64)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint to uint64

func CopyUintToUint64(src, dest unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dest)) = uint64(*(*uint)(unsafe.Pointer(src)))
}

func CopyPUintToUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dest)) = v
}

func CopyUintToPUint64(src, dest unsafe.Pointer) {
	v := uint64(*(*uint)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUintToPUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**uint)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint8 to uint64

func CopyUint8ToUint64(src, dest unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dest)) = uint64(*(*uint8)(unsafe.Pointer(src)))
}

func CopyPUint8ToUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dest)) = v
}

func CopyUint8ToPUint64(src, dest unsafe.Pointer) {
	v := uint64(*(*uint8)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint8ToPUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**uint8)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint16 to uint64

func CopyUint16ToUint64(src, dest unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dest)) = uint64(*(*uint16)(unsafe.Pointer(src)))
}

func CopyPUint16ToUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dest)) = v
}

func CopyUint16ToPUint64(src, dest unsafe.Pointer) {
	v := uint64(*(*uint16)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint16ToPUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**uint16)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint32 to uint64

func CopyUint32ToUint64(src, dest unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dest)) = uint64(*(*uint32)(unsafe.Pointer(src)))
}

func CopyPUint32ToUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dest)) = v
}

func CopyUint32ToPUint64(src, dest unsafe.Pointer) {
	v := uint64(*(*uint32)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint32ToPUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**uint32)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// uint64 to uint64

func CopyUint64ToUint64(src, dest unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(dest)) = uint64(*(*uint64)(unsafe.Pointer(src)))
}

func CopyPUint64ToUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}
	*(*uint64)(unsafe.Pointer(dest)) = v
}

func CopyUint64ToPUint64(src, dest unsafe.Pointer) {
	v := uint64(*(*uint64)(unsafe.Pointer(src)))
	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPUint64ToPUint64(src, dest unsafe.Pointer) {
	var v uint64
	if p := *(**uint64)(unsafe.Pointer(src)); p != nil {
		v = uint64(*p)
	}

	p := (**uint64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// float32 to float32

func CopyFloat32ToFloat32(src, dest unsafe.Pointer) {
	*(*float32)(unsafe.Pointer(dest)) = float32(*(*float32)(unsafe.Pointer(src)))
}

func CopyPFloat32ToFloat32(src, dest unsafe.Pointer) {
	var v float32
	if p := *(**float32)(unsafe.Pointer(src)); p != nil {
		v = float32(*p)
	}
	*(*float32)(unsafe.Pointer(dest)) = v
}

func CopyFloat32ToPFloat32(src, dest unsafe.Pointer) {
	v := float32(*(*float32)(unsafe.Pointer(src)))
	p := (**float32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPFloat32ToPFloat32(src, dest unsafe.Pointer) {
	var v float32
	if p := *(**float32)(unsafe.Pointer(src)); p != nil {
		v = float32(*p)
	}

	p := (**float32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// float64 to float32

func CopyFloat64ToFloat32(src, dest unsafe.Pointer) {
	*(*float32)(unsafe.Pointer(dest)) = float32(*(*float64)(unsafe.Pointer(src)))
}

func CopyPFloat64ToFloat32(src, dest unsafe.Pointer) {
	var v float32
	if p := *(**float64)(unsafe.Pointer(src)); p != nil {
		v = float32(*p)
	}
	*(*float32)(unsafe.Pointer(dest)) = v
}

func CopyFloat64ToPFloat32(src, dest unsafe.Pointer) {
	v := float32(*(*float64)(unsafe.Pointer(src)))
	p := (**float32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPFloat64ToPFloat32(src, dest unsafe.Pointer) {
	var v float32
	if p := *(**float64)(unsafe.Pointer(src)); p != nil {
		v = float32(*p)
	}

	p := (**float32)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// float32 to float64

func CopyFloat32ToFloat64(src, dest unsafe.Pointer) {
	*(*float64)(unsafe.Pointer(dest)) = float64(*(*float32)(unsafe.Pointer(src)))
}

func CopyPFloat32ToFloat64(src, dest unsafe.Pointer) {
	var v float64
	if p := *(**float32)(unsafe.Pointer(src)); p != nil {
		v = float64(*p)
	}
	*(*float64)(unsafe.Pointer(dest)) = v
}

func CopyFloat32ToPFloat64(src, dest unsafe.Pointer) {
	v := float64(*(*float32)(unsafe.Pointer(src)))
	p := (**float64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPFloat32ToPFloat64(src, dest unsafe.Pointer) {
	var v float64
	if p := *(**float32)(unsafe.Pointer(src)); p != nil {
		v = float64(*p)
	}

	p := (**float64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// float64 to float64

func CopyFloat64ToFloat64(src, dest unsafe.Pointer) {
	*(*float64)(unsafe.Pointer(dest)) = float64(*(*float64)(unsafe.Pointer(src)))
}

func CopyPFloat64ToFloat64(src, dest unsafe.Pointer) {
	var v float64
	if p := *(**float64)(unsafe.Pointer(src)); p != nil {
		v = float64(*p)
	}
	*(*float64)(unsafe.Pointer(dest)) = v
}

func CopyFloat64ToPFloat64(src, dest unsafe.Pointer) {
	v := float64(*(*float64)(unsafe.Pointer(src)))
	p := (**float64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPFloat64ToPFloat64(src, dest unsafe.Pointer) {
	var v float64
	if p := *(**float64)(unsafe.Pointer(src)); p != nil {
		v = float64(*p)
	}

	p := (**float64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// bool to bool

func CopyBoolToBool(src, dest unsafe.Pointer) {
	*(*bool)(unsafe.Pointer(dest)) = bool(*(*bool)(unsafe.Pointer(src)))
}

func CopyPBoolToBool(src, dest unsafe.Pointer) {
	var v bool
	if p := *(**bool)(unsafe.Pointer(src)); p != nil {
		v = bool(*p)
	}
	*(*bool)(unsafe.Pointer(dest)) = v
}

func CopyBoolToPBool(src, dest unsafe.Pointer) {
	v := bool(*(*bool)(unsafe.Pointer(src)))
	p := (**bool)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPBoolToPBool(src, dest unsafe.Pointer) {
	var v bool
	if p := *(**bool)(unsafe.Pointer(src)); p != nil {
		v = bool(*p)
	}

	p := (**bool)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// complex64 to complex64

func CopyComplex64ToComplex64(src, dest unsafe.Pointer) {
	*(*complex64)(unsafe.Pointer(dest)) = complex64(*(*complex64)(unsafe.Pointer(src)))
}

func CopyPComplex64ToComplex64(src, dest unsafe.Pointer) {
	var v complex64
	if p := *(**complex64)(unsafe.Pointer(src)); p != nil {
		v = complex64(*p)
	}
	*(*complex64)(unsafe.Pointer(dest)) = v
}

func CopyComplex64ToPComplex64(src, dest unsafe.Pointer) {
	v := complex64(*(*complex64)(unsafe.Pointer(src)))
	p := (**complex64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPComplex64ToPComplex64(src, dest unsafe.Pointer) {
	var v complex64
	if p := *(**complex64)(unsafe.Pointer(src)); p != nil {
		v = complex64(*p)
	}

	p := (**complex64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// complex128 to complex64

func CopyComplex128ToComplex64(src, dest unsafe.Pointer) {
	*(*complex64)(unsafe.Pointer(dest)) = complex64(*(*complex128)(unsafe.Pointer(src)))
}

func CopyPComplex128ToComplex64(src, dest unsafe.Pointer) {
	var v complex64
	if p := *(**complex128)(unsafe.Pointer(src)); p != nil {
		v = complex64(*p)
	}
	*(*complex64)(unsafe.Pointer(dest)) = v
}

func CopyComplex128ToPComplex64(src, dest unsafe.Pointer) {
	v := complex64(*(*complex128)(unsafe.Pointer(src)))
	p := (**complex64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPComplex128ToPComplex64(src, dest unsafe.Pointer) {
	var v complex64
	if p := *(**complex128)(unsafe.Pointer(src)); p != nil {
		v = complex64(*p)
	}

	p := (**complex64)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// complex64 to complex128

func CopyComplex64ToComplex128(src, dest unsafe.Pointer) {
	*(*complex128)(unsafe.Pointer(dest)) = complex128(*(*complex64)(unsafe.Pointer(src)))
}

func CopyPComplex64ToComplex128(src, dest unsafe.Pointer) {
	var v complex128
	if p := *(**complex64)(unsafe.Pointer(src)); p != nil {
		v = complex128(*p)
	}
	*(*complex128)(unsafe.Pointer(dest)) = v
}

func CopyComplex64ToPComplex128(src, dest unsafe.Pointer) {
	v := complex128(*(*complex64)(unsafe.Pointer(src)))
	p := (**complex128)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPComplex64ToPComplex128(src, dest unsafe.Pointer) {
	var v complex128
	if p := *(**complex64)(unsafe.Pointer(src)); p != nil {
		v = complex128(*p)
	}

	p := (**complex128)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// complex128 to complex128

func CopyComplex128ToComplex128(src, dest unsafe.Pointer) {
	*(*complex128)(unsafe.Pointer(dest)) = complex128(*(*complex128)(unsafe.Pointer(src)))
}

func CopyPComplex128ToComplex128(src, dest unsafe.Pointer) {
	var v complex128
	if p := *(**complex128)(unsafe.Pointer(src)); p != nil {
		v = complex128(*p)
	}
	*(*complex128)(unsafe.Pointer(dest)) = v
}

func CopyComplex128ToPComplex128(src, dest unsafe.Pointer) {
	v := complex128(*(*complex128)(unsafe.Pointer(src)))
	p := (**complex128)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPComplex128ToPComplex128(src, dest unsafe.Pointer) {
	var v complex128
	if p := *(**complex128)(unsafe.Pointer(src)); p != nil {
		v = complex128(*p)
	}

	p := (**complex128)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// string to string

func CopyStringToString(src, dest unsafe.Pointer) {
	*(*string)(unsafe.Pointer(dest)) = string(*(*string)(unsafe.Pointer(src)))
}

func CopyPStringToString(src, dest unsafe.Pointer) {
	var v string
	if p := *(**string)(unsafe.Pointer(src)); p != nil {
		v = string(*p)
	}
	*(*string)(unsafe.Pointer(dest)) = v
}

func CopyStringToPString(src, dest unsafe.Pointer) {
	v := string(*(*string)(unsafe.Pointer(src)))
	p := (**string)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPStringToPString(src, dest unsafe.Pointer) {
	var v string
	if p := *(**string)(unsafe.Pointer(src)); p != nil {
		v = string(*p)
	}

	p := (**string)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// []byte to string

func CopyBytesToString(src, dest unsafe.Pointer) {
	*(*string)(unsafe.Pointer(dest)) = string(*(*[]byte)(unsafe.Pointer(src)))
}

func CopyPBytesToString(src, dest unsafe.Pointer) {
	var v string
	if p := *(**[]byte)(unsafe.Pointer(src)); p != nil {
		v = string(*p)
	}
	*(*string)(unsafe.Pointer(dest)) = v
}

func CopyBytesToPString(src, dest unsafe.Pointer) {
	v := string(*(*[]byte)(unsafe.Pointer(src)))
	p := (**string)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPBytesToPString(src, dest unsafe.Pointer) {
	var v string
	if p := *(**[]byte)(unsafe.Pointer(src)); p != nil {
		v = string(*p)
	}

	p := (**string)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// string to []byte

func CopyStringToBytes(src, dest unsafe.Pointer) {
	*(*[]byte)(unsafe.Pointer(dest)) = []byte(*(*string)(unsafe.Pointer(src)))
}

func CopyPStringToBytes(src, dest unsafe.Pointer) {
	var v []byte
	if p := *(**string)(unsafe.Pointer(src)); p != nil {
		v = []byte(*p)
	}
	*(*[]byte)(unsafe.Pointer(dest)) = v
}

func CopyStringToPBytes(src, dest unsafe.Pointer) {
	v := []byte(*(*string)(unsafe.Pointer(src)))
	p := (**[]byte)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPStringToPBytes(src, dest unsafe.Pointer) {
	var v []byte
	if p := *(**string)(unsafe.Pointer(src)); p != nil {
		v = []byte(*p)
	}

	p := (**[]byte)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// []byte to []byte

func CopyBytesToBytes(src, dest unsafe.Pointer) {
	*(*[]byte)(unsafe.Pointer(dest)) = []byte(*(*[]byte)(unsafe.Pointer(src)))
}

func CopyPBytesToBytes(src, dest unsafe.Pointer) {
	var v []byte
	if p := *(**[]byte)(unsafe.Pointer(src)); p != nil {
		v = []byte(*p)
	}
	*(*[]byte)(unsafe.Pointer(dest)) = v
}

func CopyBytesToPBytes(src, dest unsafe.Pointer) {
	v := []byte(*(*[]byte)(unsafe.Pointer(src)))
	p := (**[]byte)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPBytesToPBytes(src, dest unsafe.Pointer) {
	var v []byte
	if p := *(**[]byte)(unsafe.Pointer(src)); p != nil {
		v = []byte(*p)
	}

	p := (**[]byte)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// time.Time to time.Time

func CopyTimeToTime(src, dest unsafe.Pointer) {
	*(*time.Time)(unsafe.Pointer(dest)) = time.Time(*(*time.Time)(unsafe.Pointer(src)))
}

func CopyPTimeToTime(src, dest unsafe.Pointer) {
	var v time.Time
	if p := *(**time.Time)(unsafe.Pointer(src)); p != nil {
		v = time.Time(*p)
	}
	*(*time.Time)(unsafe.Pointer(dest)) = v
}

func CopyTimeToPTime(src, dest unsafe.Pointer) {
	v := time.Time(*(*time.Time)(unsafe.Pointer(src)))
	p := (**time.Time)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPTimeToPTime(src, dest unsafe.Pointer) {
	var v time.Time
	if p := *(**time.Time)(unsafe.Pointer(src)); p != nil {
		v = time.Time(*p)
	}

	p := (**time.Time)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
} 

// time.Duration to time.Duration

func CopyDurationToDuration(src, dest unsafe.Pointer) {
	*(*time.Duration)(unsafe.Pointer(dest)) = time.Duration(*(*time.Duration)(unsafe.Pointer(src)))
}

func CopyPDurationToDuration(src, dest unsafe.Pointer) {
	var v time.Duration
	if p := *(**time.Duration)(unsafe.Pointer(src)); p != nil {
		v = time.Duration(*p)
	}
	*(*time.Duration)(unsafe.Pointer(dest)) = v
}

func CopyDurationToPDuration(src, dest unsafe.Pointer) {
	v := time.Duration(*(*time.Duration)(unsafe.Pointer(src)))
	p := (**time.Duration)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}

func CopyPDurationToPDuration(src, dest unsafe.Pointer) {
	var v time.Duration
	if p := *(**time.Duration)(unsafe.Pointer(src)); p != nil {
		v = time.Duration(*p)
	}

	p := (**time.Duration)(unsafe.Pointer(dest))
	if p := *p; p != nil {
		*p = v
		return
	}
	*p = &v
}	

