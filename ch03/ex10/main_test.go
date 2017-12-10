package main

import (
	"testing"
)

func TestComma1(t *testing.T) {
	if actual := comma("123");actual !="123"{
	t.Error("comma(123) is not 123",actual)
	}
}

func TestComma2(t *testing.T) {
	if actual := comma("1234");actual !="1,234"{
		t.Error("comma(1234) is not 1,234",actual)
	}
}

func TestComma3(t *testing.T) {
	if actual := comma("1234567");actual !="1,234,567"{
		t.Error("comma(1234567) is not 1,234,567",actual)
	}
}
