package main

import "fmt"

type sample struct {
	a int
	b string
}

func (s sample) Method() {
	fmt.Println("Method")
}

func main() {
	var x *sample
	x = &sample{1, "hoge"}
	fmt.Printf("(x==nil) is %v\n", x == nil)

	var y *sample
	fmt.Printf("(y==nil) is %v\n", y == nil)

	var z MyInterface
	z = y
	fmt.Printf("(z==nil) is %v\n", z == nil)

}

type MyInterface interface {
	Method()
}
