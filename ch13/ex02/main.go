package main

import (
	"reflect"
	"unsafe"
	"fmt"
)

var seen = make(map[unsafe.Pointer]bool)

func IsCirculation(x interface{}) bool {
	return isCirculation(reflect.ValueOf(x))
}
func isCirculation(x reflect.Value) bool {
	if x.CanAddr() &&
		x.Kind() != reflect.Struct &&
		x.Kind() != reflect.Array &&
		x.Kind() != reflect.Map &&
		x.Kind() != reflect.Slice {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		if seen[xptr] {
			return true
		}
		seen[xptr] = true
	}
	fmt.Printf("x.Kind() is %v\n", x.Kind())
	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return isCirculation(x.Elem())
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if isCirculation(x.Field(i)) {
				return true
			}
		}
		return false
	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if isCirculation(x.Field(i)) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if isCirculation(x.MapIndex(k)) {
				return true
			}
		}
		return false
	}
	return false
}
