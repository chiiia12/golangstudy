package main

import (
	"testing"
	"fmt"
)

func TestConvertASCIISpace1(t *testing.T) {
	str := []byte("あい  うえお") //len is 17
	fmt.Printf("str before is %v\n", str)
	str = convertToASCII(str)
	fmt.Printf("str result is %v\n", str)
	if actual := len(str); actual != 16 {
		t.Errorf("len(str) is not 16.actual is %v", actual)
	}
}

func TestConvertASCIISpace2(t *testing.T) {
	str := []byte("あいうえお") //len is 15
	fmt.Printf("str before is %v\n", str)
	str = convertToASCII(str)
	fmt.Printf("str result is %v\n", str)
	if actual := len(str); actual != 15 {
		t.Errorf("len(str) is not 15.actual is %v", actual)
	}
}

func TestConvertASCIISpace3(t *testing.T) {
	str := []byte("あ  いうえ  お") //len is 19
	fmt.Printf("str before is %v\n", str)
	str = convertToASCII(str)
	fmt.Printf("str result is %v\n", str)
	if actual := len(str); actual != 17 {
		t.Errorf("len(str) is not 17.actual is %v", actual)
	}
}
