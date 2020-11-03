package funcs

import (
	"bytes"
	"crypto/rand"
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/gotidy/ptr"
)

func checkEqual(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Not equal. Actual: %v; expected: %v", actual, expected)
	}
}

func TestBytesCopiers(t *testing.T) {
	src := make([]byte, 100)
	_, err := rand.Read(src)
	if err != nil {
		t.Fatalf("rand: %s", err)
	}
	for i := 1; i <= len(src); i++ {
		dst := make([]byte, i)
		src := src[:i]
		copier := Get(reflect.TypeOf(dst), reflect.TypeOf(src))
		copier(unsafe.Pointer(reflect.ValueOf(&dst).Pointer()), unsafe.Pointer(reflect.ValueOf(&src).Pointer()))
		if !bytes.Equal(dst, src) {
			t.Fatalf("a destination with len %d is not equal a source", i)
		}
	}
}

func TestMemCopiers(t *testing.T) {
	src := make([]byte, len(funcs.sizes))
	_, err := rand.Read(src)
	if err != nil {
		t.Fatalf("rand: %s", err)
	}
	for i := 1; i <= len(src); i++ {
		dst := make([]byte, i)
		src := src[:i]
		typ := reflect.ArrayOf(i, reflect.TypeOf(uint8(0)))
		copier := Get(typ, typ)
		copier(unsafe.Pointer(reflect.ValueOf(dst).Pointer()), unsafe.Pointer(reflect.ValueOf(src).Pointer()))
		if !bytes.Equal(dst, src) {
			t.Fatalf("a destination with len %d is not equal a source", i)
		}
	}
}

func TestTypesCopiers(t *testing.T) {
	b := []byte("COVID-21")
	testValues := make(map[reflect.Type]reflect.Value)
	for _, v := range []interface{}{
		ptr.Int(10),
		ptr.Int8(10),
		ptr.Int16(10),
		ptr.Int32(10),
		ptr.Int64(10),
		ptr.UInt(10),
		ptr.UInt8(10),
		ptr.UInt16(10),
		ptr.UInt32(10),
		ptr.UInt64(10),
		ptr.Float32(10.0),
		ptr.Float64(10.0),
		ptr.Complex64(10.0),
		ptr.Complex128(10.0),
		ptr.Bool(true),
		ptr.Duration(time.Second),
		ptr.Time(time.Date(2021, 2, 18, 16, 0, 1, 0, time.UTC)),
		ptr.String("COVID-21"),
		&b,
	} {
		val := reflect.ValueOf(v)
		testValues[val.Elem().Type()] = val

		valP := reflect.New(val.Type())
		valP.Elem().Set(val)
		testValues[valP.Elem().Type()] = valP
	}

	for key, copier := range funcs.funcs {
		// var dst reflect.Value

		src, ok := testValues[key.Src]
		if !ok {
			continue
		}

		dst := reflect.New(key.Dest)

		t.Logf("Src key «%s»; src «%#v»", key.Src, src)
		t.Logf("Dst key «%s»; dst «%#v»", key.Dest, dst)

		set := func(dst, src reflect.Value) {
			copier(unsafe.Pointer(dst.Pointer()), unsafe.Pointer(src.Pointer()))

			dst = dst.Elem()
			src = src.Elem()
			t.Logf("src «%#v»", src)
			t.Logf("dst «%#v»", dst)

			if key.Dest.Kind() == reflect.Ptr {
				dst = dst.Elem()
			}
			dstI := dst.Interface()

			if key.Src.Kind() == reflect.Ptr {
				src = src.Elem()
			}
			src = src.Convert(dst.Type())
			srcI := src.Interface()

			if !reflect.DeepEqual(dstI, srcI) {
				t.Errorf("want «%v» got «%v»", srcI, dstI)
			}
		}

		set(dst, src)
		if elem := dst.Elem(); elem.Kind() == reflect.Ptr {
			elem.Set(reflect.New(elem.Type().Elem()))
			set(dst, src)
		}
	}
}

func TestSet(t *testing.T) {
	Set(reflect.TypeOf(int(0)), reflect.TypeOf(int(0)), func(dst, src unsafe.Pointer) {
		*(*int)(unsafe.Pointer(dst)) = int(*(*int)(unsafe.Pointer(src)))
	})
}

func TestGet(t *testing.T) {
	copier := Get(reflect.TypeOf(struct{ I int }{I: 1}), reflect.TypeOf([]byte{2}))
	if copier != nil {
		t.Error("should return nil when types are incompatible")
	}

	copier = Get(reflect.TypeOf(map[string]string{}), reflect.TypeOf(map[string]string{}))
	if copier == nil {
		t.Error("Get(map, map) should not return nil")
	}
}
