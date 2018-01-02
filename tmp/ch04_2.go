package main

import (
	"fmt"
	"strconv"
	"log"
)

func main() {
	x := 100
	str := "12345"
	if len(str) != 0 {
		x, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(x)
}
