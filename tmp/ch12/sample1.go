package main

import (
	"strconv"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(Sprint(1))
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)
}

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return "???"
	}
}
