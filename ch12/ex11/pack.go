package main

import (
	"reflect"
	"log"
	"strings"
	"fmt"
)

func Pack(ptr interface{}) string {
	var result []string
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		structField := v.Type().Field(i)
		tag := structField.Tag
		name := tag.Get("http")
		log.Println(name)
		result = append(result, toParam(name, v.Field(i)))
	}
	return strings.Join(result, "&")
}
func toParam(name string, v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return fmt.Sprintf("%s=%s", name, v.String())
	case reflect.Int:
		return fmt.Sprintf("%s=%v", name, v.Int())
	case reflect.Bool:
		return fmt.Sprintf("%s=%v", name, v.Bool())
	case reflect.Array, reflect.Slice:
		var str []string
		for i := 0; i < v.Len(); i++ {
			str = append(str, toParam(name, v.Index(i)))
		}
		return strings.Join(str, "&")
	default:
		return ""
	}
}
