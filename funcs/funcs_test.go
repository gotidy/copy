package funcs

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/gotidy/ptr"
)

func checkEqual(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Not equal. Actual: %v; expected: %v", actual, expected)
	}
}

func TestCopiers(t *testing.T) {
	// var i int

	type args struct {
		Src  interface{}
		Dest interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// int
		{name: "CopyIntToInt", args: args{Src: ptr.Int(10), Dest: ptr.Int(0)}, want: int(10)},
		{name: "CopyInt8ToInt", args: args{Src: ptr.Int8(10), Dest: ptr.Int(0)}, want: int(10)},
		{name: "CopyInt16ToInt", args: args{Src: ptr.Int16(10), Dest: ptr.Int(0)}, want: int(10)},
		{name: "CopyInt32ToInt", args: args{Src: ptr.Int32(10), Dest: ptr.Int(0)}, want: int(10)},
		{name: "CopyInt64ToInt", args: args{Src: ptr.Int64(10), Dest: ptr.Int(0)}, want: int(10)},
		{name: "CopyUIntToInt", args: args{Src: ptr.UInt(10), Dest: ptr.Int(0)}, want: int(10)},
		{name: "CopyUInt8ToInt", args: args{Src: ptr.UInt8(10), Dest: ptr.Int(0)}, want: int(10)},
		{name: "CopyUInt16ToInt", args: args{Src: ptr.UInt16(10), Dest: ptr.Int(0)}, want: int(10)},
		{name: "CopyUInt32ToInt", args: args{Src: ptr.UInt32(10), Dest: ptr.Int(0)}, want: int(10)},
		{name: "CopyUInt64ToInt", args: args{Src: ptr.UInt64(10), Dest: ptr.Int(0)}, want: int(10)},
		// int8
		{name: "CopyIntToInt8", args: args{Src: ptr.Int(10), Dest: ptr.Int8(0)}, want: int8(10)},
		{name: "CopyInt8ToInt8", args: args{Src: ptr.Int8(10), Dest: ptr.Int8(0)}, want: int8(10)},
		{name: "CopyInt16ToInt8", args: args{Src: ptr.Int16(10), Dest: ptr.Int8(0)}, want: int8(10)},
		{name: "CopyInt32ToInt8", args: args{Src: ptr.Int32(10), Dest: ptr.Int8(0)}, want: int8(10)},
		{name: "CopyInt64ToInt8", args: args{Src: ptr.Int64(10), Dest: ptr.Int8(0)}, want: int8(10)},
		{name: "CopyUIntToInt8", args: args{Src: ptr.UInt(10), Dest: ptr.Int8(0)}, want: int8(10)},
		{name: "CopyUInt8ToInt8", args: args{Src: ptr.UInt8(10), Dest: ptr.Int8(0)}, want: int8(10)},
		{name: "CopyUInt16ToInt8", args: args{Src: ptr.UInt16(10), Dest: ptr.Int8(0)}, want: int8(10)},
		{name: "CopyUInt32ToInt8", args: args{Src: ptr.UInt32(10), Dest: ptr.Int8(0)}, want: int8(10)},
		{name: "CopyUInt64ToInt8", args: args{Src: ptr.UInt64(10), Dest: ptr.Int8(0)}, want: int8(10)},
		// int16
		{name: "CopyIntToInt16", args: args{Src: ptr.Int(10), Dest: ptr.Int16(0)}, want: int16(10)},
		{name: "CopyInt8ToInt16", args: args{Src: ptr.Int8(10), Dest: ptr.Int16(0)}, want: int16(10)},
		{name: "CopyInt16ToInt16", args: args{Src: ptr.Int16(10), Dest: ptr.Int16(0)}, want: int16(10)},
		{name: "CopyInt32ToInt16", args: args{Src: ptr.Int32(10), Dest: ptr.Int16(0)}, want: int16(10)},
		{name: "CopyInt64ToInt16", args: args{Src: ptr.Int64(10), Dest: ptr.Int16(0)}, want: int16(10)},
		{name: "CopyUIntToInt16", args: args{Src: ptr.UInt(10), Dest: ptr.Int16(0)}, want: int16(10)},
		{name: "CopyUInt8ToInt16", args: args{Src: ptr.UInt8(10), Dest: ptr.Int16(0)}, want: int16(10)},
		{name: "CopyUInt16ToInt16", args: args{Src: ptr.UInt16(10), Dest: ptr.Int16(0)}, want: int16(10)},
		{name: "CopyUInt32ToInt16", args: args{Src: ptr.UInt32(10), Dest: ptr.Int16(0)}, want: int16(10)},
		{name: "CopyUInt64ToInt16", args: args{Src: ptr.UInt64(10), Dest: ptr.Int16(0)}, want: int16(10)},
		// int32
		{name: "CopyIntToInt32", args: args{Src: ptr.Int(10), Dest: ptr.Int32(0)}, want: int32(10)},
		{name: "CopyInt8ToInt32", args: args{Src: ptr.Int8(10), Dest: ptr.Int32(0)}, want: int32(10)},
		{name: "CopyInt16ToInt32", args: args{Src: ptr.Int16(10), Dest: ptr.Int32(0)}, want: int32(10)},
		{name: "CopyInt32ToInt32", args: args{Src: ptr.Int32(10), Dest: ptr.Int32(0)}, want: int32(10)},
		{name: "CopyInt64ToInt32", args: args{Src: ptr.Int64(10), Dest: ptr.Int32(0)}, want: int32(10)},
		{name: "CopyUIntToInt32", args: args{Src: ptr.UInt(10), Dest: ptr.Int32(0)}, want: int32(10)},
		{name: "CopyUInt8ToInt32", args: args{Src: ptr.UInt8(10), Dest: ptr.Int32(0)}, want: int32(10)},
		{name: "CopyUInt16ToInt32", args: args{Src: ptr.UInt16(10), Dest: ptr.Int32(0)}, want: int32(10)},
		{name: "CopyUInt32ToInt32", args: args{Src: ptr.UInt32(10), Dest: ptr.Int32(0)}, want: int32(10)},
		{name: "CopyUInt64ToInt32", args: args{Src: ptr.UInt64(10), Dest: ptr.Int32(0)}, want: int32(10)},
		// int64
		{name: "CopyIntToInt64", args: args{Src: ptr.Int(10), Dest: ptr.Int64(0)}, want: int64(10)},
		{name: "CopyInt8ToInt64", args: args{Src: ptr.Int8(10), Dest: ptr.Int64(0)}, want: int64(10)},
		{name: "CopyInt16ToInt64", args: args{Src: ptr.Int16(10), Dest: ptr.Int64(0)}, want: int64(10)},
		{name: "CopyInt32ToInt64", args: args{Src: ptr.Int32(10), Dest: ptr.Int64(0)}, want: int64(10)},
		{name: "CopyInt64ToInt64", args: args{Src: ptr.Int64(10), Dest: ptr.Int64(0)}, want: int64(10)},
		{name: "CopyUIntToInt64", args: args{Src: ptr.UInt(10), Dest: ptr.Int64(0)}, want: int64(10)},
		{name: "CopyUInt8ToInt64", args: args{Src: ptr.UInt8(10), Dest: ptr.Int64(0)}, want: int64(10)},
		{name: "CopyUInt16ToInt64", args: args{Src: ptr.UInt16(10), Dest: ptr.Int64(0)}, want: int64(10)},
		{name: "CopyUInt32ToInt64", args: args{Src: ptr.UInt32(10), Dest: ptr.Int64(0)}, want: int64(10)},
		{name: "CopyUInt64ToInt64", args: args{Src: ptr.UInt64(10), Dest: ptr.Int64(0)}, want: int64(10)},
		// uint
		{name: "CopyIntToInt", args: args{Src: ptr.Int(10), Dest: ptr.UInt(0)}, want: uint(10)},
		{name: "CopyInt8ToInt", args: args{Src: ptr.Int8(10), Dest: ptr.UInt(0)}, want: uint(10)},
		{name: "CopyInt16ToInt", args: args{Src: ptr.Int16(10), Dest: ptr.UInt(0)}, want: uint(10)},
		{name: "CopyInt32ToInt", args: args{Src: ptr.Int32(10), Dest: ptr.UInt(0)}, want: uint(10)},
		{name: "CopyInt64ToInt", args: args{Src: ptr.Int64(10), Dest: ptr.UInt(0)}, want: uint(10)},
		{name: "CopyUIntToInt", args: args{Src: ptr.UInt(10), Dest: ptr.UInt(0)}, want: uint(10)},
		{name: "CopyUInt8ToInt", args: args{Src: ptr.UInt8(10), Dest: ptr.UInt(0)}, want: uint(10)},
		{name: "CopyUInt16ToInt", args: args{Src: ptr.UInt16(10), Dest: ptr.UInt(0)}, want: uint(10)},
		{name: "CopyUInt32ToInt", args: args{Src: ptr.UInt32(10), Dest: ptr.UInt(0)}, want: uint(10)},
		{name: "CopyUInt64ToInt", args: args{Src: ptr.UInt64(10), Dest: ptr.UInt(0)}, want: uint(10)},
		// uint8
		{name: "CopyIntToInt8", args: args{Src: ptr.Int(10), Dest: ptr.UInt8(0)}, want: uint8(10)},
		{name: "CopyInt8ToInt8", args: args{Src: ptr.Int8(10), Dest: ptr.UInt8(0)}, want: uint8(10)},
		{name: "CopyInt16ToInt8", args: args{Src: ptr.Int16(10), Dest: ptr.UInt8(0)}, want: uint8(10)},
		{name: "CopyInt32ToInt8", args: args{Src: ptr.Int32(10), Dest: ptr.UInt8(0)}, want: uint8(10)},
		{name: "CopyInt64ToInt8", args: args{Src: ptr.Int64(10), Dest: ptr.UInt8(0)}, want: uint8(10)},
		{name: "CopyUIntToInt8", args: args{Src: ptr.UInt(10), Dest: ptr.UInt8(0)}, want: uint8(10)},
		{name: "CopyUInt8ToInt8", args: args{Src: ptr.UInt8(10), Dest: ptr.UInt8(0)}, want: uint8(10)},
		{name: "CopyUInt16ToInt8", args: args{Src: ptr.UInt16(10), Dest: ptr.UInt8(0)}, want: uint8(10)},
		{name: "CopyUInt32ToInt8", args: args{Src: ptr.UInt32(10), Dest: ptr.UInt8(0)}, want: uint8(10)},
		{name: "CopyUInt64ToInt8", args: args{Src: ptr.UInt64(10), Dest: ptr.UInt8(0)}, want: uint8(10)},
		// uint16
		{name: "CopyIntToInt16", args: args{Src: ptr.Int(10), Dest: ptr.UInt16(0)}, want: uint16(10)},
		{name: "CopyInt8ToInt16", args: args{Src: ptr.Int8(10), Dest: ptr.UInt16(0)}, want: uint16(10)},
		{name: "CopyInt16ToInt16", args: args{Src: ptr.Int16(10), Dest: ptr.UInt16(0)}, want: uint16(10)},
		{name: "CopyInt32ToInt16", args: args{Src: ptr.Int32(10), Dest: ptr.UInt16(0)}, want: uint16(10)},
		{name: "CopyInt64ToInt16", args: args{Src: ptr.Int64(10), Dest: ptr.UInt16(0)}, want: uint16(10)},
		{name: "CopyUIntToInt16", args: args{Src: ptr.UInt(10), Dest: ptr.UInt16(0)}, want: uint16(10)},
		{name: "CopyUInt8ToInt16", args: args{Src: ptr.UInt8(10), Dest: ptr.UInt16(0)}, want: uint16(10)},
		{name: "CopyUInt16ToInt16", args: args{Src: ptr.UInt16(10), Dest: ptr.UInt16(0)}, want: uint16(10)},
		{name: "CopyUInt32ToInt16", args: args{Src: ptr.UInt32(10), Dest: ptr.UInt16(0)}, want: uint16(10)},
		{name: "CopyUInt64ToInt16", args: args{Src: ptr.UInt64(10), Dest: ptr.UInt16(0)}, want: uint16(10)},
		// uint32
		{name: "CopyIntToInt32", args: args{Src: ptr.Int(10), Dest: ptr.UInt32(0)}, want: uint32(10)},
		{name: "CopyInt8ToInt32", args: args{Src: ptr.Int8(10), Dest: ptr.UInt32(0)}, want: uint32(10)},
		{name: "CopyInt16ToInt32", args: args{Src: ptr.Int16(10), Dest: ptr.UInt32(0)}, want: uint32(10)},
		{name: "CopyInt32ToInt32", args: args{Src: ptr.Int32(10), Dest: ptr.UInt32(0)}, want: uint32(10)},
		{name: "CopyInt64ToInt32", args: args{Src: ptr.Int64(10), Dest: ptr.UInt32(0)}, want: uint32(10)},
		{name: "CopyUIntToInt32", args: args{Src: ptr.UInt(10), Dest: ptr.UInt32(0)}, want: uint32(10)},
		{name: "CopyUInt8ToInt32", args: args{Src: ptr.UInt8(10), Dest: ptr.UInt32(0)}, want: uint32(10)},
		{name: "CopyUInt16ToInt32", args: args{Src: ptr.UInt16(10), Dest: ptr.UInt32(0)}, want: uint32(10)},
		{name: "CopyUInt32ToInt32", args: args{Src: ptr.UInt32(10), Dest: ptr.UInt32(0)}, want: uint32(10)},
		{name: "CopyUInt64ToInt32", args: args{Src: ptr.UInt64(10), Dest: ptr.UInt32(0)}, want: uint32(10)},
		// uint64
		{name: "CopyIntToInt64", args: args{Src: ptr.Int(10), Dest: ptr.UInt64(0)}, want: uint64(10)},
		{name: "CopyInt8ToInt64", args: args{Src: ptr.Int8(10), Dest: ptr.UInt64(0)}, want: uint64(10)},
		{name: "CopyInt16ToInt64", args: args{Src: ptr.Int16(10), Dest: ptr.UInt64(0)}, want: uint64(10)},
		{name: "CopyInt32ToInt64", args: args{Src: ptr.Int32(10), Dest: ptr.UInt64(0)}, want: uint64(10)},
		{name: "CopyInt64ToInt64", args: args{Src: ptr.Int64(10), Dest: ptr.UInt64(0)}, want: uint64(10)},
		{name: "CopyUIntToInt64", args: args{Src: ptr.UInt(10), Dest: ptr.UInt64(0)}, want: uint64(10)},
		{name: "CopyUInt8ToInt64", args: args{Src: ptr.UInt8(10), Dest: ptr.UInt64(0)}, want: uint64(10)},
		{name: "CopyUInt16ToInt64", args: args{Src: ptr.UInt16(10), Dest: ptr.UInt64(0)}, want: uint64(10)},
		{name: "CopyUInt32ToInt64", args: args{Src: ptr.UInt32(10), Dest: ptr.UInt64(0)}, want: uint64(10)},
		{name: "CopyUInt64ToInt64", args: args{Src: ptr.UInt64(10), Dest: ptr.UInt64(0)}, want: uint64(10)},
		// float
		{name: "CopyFloat32ToFloat64", args: args{Src: ptr.Float32(10.0), Dest: ptr.Float64(0)}, want: float64(10.0)},
		{name: "CopyFloat64ToFloat64", args: args{Src: ptr.Float64(10.0), Dest: ptr.Float64(0)}, want: float64(10.0)},
		{name: "CopyFloat32ToFloat32", args: args{Src: ptr.Float32(10.0), Dest: ptr.Float32(0)}, want: float32(10.0)},
		{name: "CopyFloat64ToFloat32", args: args{Src: ptr.Float64(10.0), Dest: ptr.Float32(0)}, want: float32(10.0)},
		// bool
		{name: "CopyBoolToBool", args: args{Src: ptr.Bool(true), Dest: ptr.Bool(false)}, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srcPtr := reflect.ValueOf(tt.args.Src)
			src := reflect.Indirect(srcPtr)
			dstPtr := reflect.ValueOf(tt.args.Dest)
			dst := reflect.Indirect(dstPtr)
			copier := Get(dst.Type(), src.Type())
			copier(unsafe.Pointer(dstPtr.Pointer()), unsafe.Pointer(srcPtr.Pointer()))
			checkEqual(t, dst.Interface(), tt.want)
		})
	}
}
