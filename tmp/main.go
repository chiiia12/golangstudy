package main

import (
	"fmt"
)

type Hoge struct{}

func (h *Hoge) M() {
	fmt.Println(h)
}

func main() {
	var h *Hoge
	h.M()
}
