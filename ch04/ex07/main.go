package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := []byte("あいうえお")
	reverse(str)
}
func reverse(str []byte) []byte {
	fmt.Println("string length is ", utf8.RuneCountInString(string(str)))
	fmt.Println(str)
	byteCount := len(str) / utf8.RuneCountInString(string(str))
	for i := 0; i < len(str)/byteCount/2; i ++ {
		for j := 0; j < byteCount; j++ {
			tmp := str[i*byteCount+j]
			str[i*byteCount+j] = str[len(str)-((i+1)*byteCount)+j]
			str[len(str)-((i+1)*byteCount)+j] = tmp
			fmt.Printf("str[i+j] is %v str[len(str)-1-(i*byteCount)+j] is %v\n", i*byteCount+j, len(str)-((i+1)*byteCount)+j)
		}
	}
	fmt.Println(str)
	return str
}
