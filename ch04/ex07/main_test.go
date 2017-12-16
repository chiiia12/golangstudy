package main

import (
	"testing"
	"fmt"
)

func TestReverse(t *testing.T) {
	str := []byte("あいうえお") //len 15
	reversed := reverse(str)
	fmt.Println(string(reversed))
	for i, v := range reversed {
		if v != str[len(str)-1-i] {
			t.Errorf("reversed[i] and str[len(str])-i] is different actual reversed[i] is %v str[len(str)-i]is %v", v, str[len(str)-1-i])
		}
	}
}
