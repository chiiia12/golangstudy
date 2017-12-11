package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%T %T", c1, c2)
	popcount(c1, c2)
}

func popcount(c1, c2 [32]uint8) int {
	var popcount int
	for i, _ := range c1 {
		fmt.Printf("%v,%v\n", c1[i], c2[i])
		a := c1[i] ^ c2[i]
		for a > 0 {
			a = a & (a - 1)
			popcount++
		}
	}
	fmt.Println("popcount is {}", popcount)
	return popcount
}
