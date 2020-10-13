package copy

type internal struct {
	I int
}
type testStruct struct {
	S  string
	I  int
	BB []bool
	V  internal
}

// func Test_getFieldCopier(t *testing.T) {
// 	src := testStruct{
// 		S:  "string",
// 		I:  10,
// 		BB: []bool{true, false},
// 		V:  internal{I: 5},
// 	}
// 	dest := testStruct{}
// 	srcType := reflect.TypeOf(src)
// 	destType := reflect.TypeOf(dest)
// 	for i := 0; i < srcType.NumField(); i++ {
// 		srcField := srcType.Field(i)
// 		if srcField.PkgPath != "" {
// 			continue
// 		}
// 		if destField, ok := destType.FieldByName(srcField.Name); ok {
// 			if f := getFieldCopier(srcField, destField); f != nil {
// 				f(unsafe.Pointer(&src), unsafe.Pointer(&dest))
// 			}
// 		}
// 	}
// 	if !reflect.DeepEqual(dest, src) {
// 		t.Errorf("getFieldCopier() = %v, want %v", dest, src)
// 	}
// 	// t.Logf("%#v", dest)
// }
