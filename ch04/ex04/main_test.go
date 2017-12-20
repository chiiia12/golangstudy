package main

import (
	"testing"
)

func TestRotate(t *testing.T) {
	var s = []int{0, 1, 2, 3, 4}
	result := rotate(s, 2)
	var expect = []int{2, 3, 4, 0, 1}
	for i, v := range result {
		if v != expect[i] {
			t.Errorf("actual is not expect. actual is %v, expect is %v", v, expect[i])
		}
	}
}

func TestRotate1(t *testing.T) {
	var s = []int{0, 1, 2, 3, 4}
	result := rotate(s, 4)
	var expect = []int{4, 0, 1, 2, 3}
	for i, v := range result {
		if v != expect[i] {
			t.Errorf("actual is not expect. actual is %v, expect is %v", v, expect[i])
		}
	}
}
