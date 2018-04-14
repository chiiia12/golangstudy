package main

import (
	"unsafe"
	"fmt"
)

var x struct {
	a bool
	b int16
	c []int
}

func main() {
	//pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	//*pb = 42
	//fmt.Println(x.b)

	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb := (*int16)(unsafe.Pointer(tmp))
	*pb = 42
	fmt.Println(x.b)

}
