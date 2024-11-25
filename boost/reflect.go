package boost

import (
	"fmt"
	"reflect"
	"unsafe"
)

// TypeName gets type name by reflect
func TypeName(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// ElemTypeName gets type name by reflect
func ElemTypeName(v interface{}) string {
	return reflect.TypeOf(v).Elem().String()
}

// Type gets type by reflect
func Type(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}

// InterfaceByType creates a new interface of specific type
func InterfaceByType(t reflect.Type) interface{} {
	return reflect.New(t.Elem()).Interface()
}

// GetUnexportedPtr gets struct unexported field ptr
func GetUnexportedPtr(s interface{}, name string) unsafe.Pointer {
	return unsafe.Pointer(reflect.Indirect(
		reflect.ValueOf(s)).FieldByName(name).UnsafeAddr())
}

// GetValuePtr gets unsafe pointer for value
func GetValuePtr(value reflect.Value) unsafe.Pointer {
	return unsafe.Pointer(value.UnsafeAddr())
}

// GetValueFieldPtr gets unsafe pointer for field value
func GetValueFieldPtr(value reflect.Value, name string) unsafe.Pointer {
	if value.Kind() != reflect.Struct {
		panic(fmt.Errorf("GetFieldValuePtr value is not struct"))
	}
	return GetValuePtr(value.FieldByName(name))
}

// GetValueByPath gets values by path
func GetValueByPath(value reflect.Value, path []string) (reflect.Value, bool) {
	for _, field := range path {
		value = value.FieldByName(field)
		if !value.IsValid() {
			return value, false
		}
	}
	return value, true
}
