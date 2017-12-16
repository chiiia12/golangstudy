package main

import (
	"testing"
	"fmt"
)

func TestReverse1(t *testing.T) {
	str := []byte("あいうえお") //len 15
	reversed := reverse(str)
	fmt.Println(string(reversed))
	if actual := string(reversed); actual != "おえういあ" {
		t.Errorf("string(reversed) is おえういあ. actual is %v", actual)
	}
}

func TestReverse2(t *testing.T) {
	str := []byte("abcde") //len 15
	reversed := reverse(str)
	fmt.Println(string(reversed))
	if actual := string(reversed); actual != "edcba" {
		t.Errorf("string(reversed) is edcba. actual is %v", actual)
	}
}
