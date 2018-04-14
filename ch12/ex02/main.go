package main

import (
	"reflect"
	"fmt"
	"strconv"
	"github.com/chiiia12/golangstudy/ch07/ex13"
)

type Sample struct {
	X int
}

func main() {
	e, _ := eval.Parse("sqrt(A/pi)")
	Display("e", e)

	m := make(map[Sample]string)
	m[Sample{3}] = "one"
	Display("m", m)

	a := []string{"hoge", "fuga", "piyo"}
	Display("a", a)
}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

var displayCount = 0;

func display(path string, v reflect.Value) {
	if displayCount > 10 {
		return
	}
	displayCount++

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Struct:
		var str string
		str += v.Type().Name()
		for i := 0; i < v.NumField(); i++ {
			str += fmt.Sprintf(" %s:%s", v.Type().Field(i).Name, formatAtom(v.Field(i)))
		}
		return str
	case reflect.Array:
		var str string
		for i := 0; i < v.Len(); i++ {
			str += fmt.Sprintf("index[%v]:%v\n", v.Index(i).Type().String());
		}
		return str
	default: //reflect.Interface
		return v.Type().String() + " value"
	}
}
