package main

import (
	"testing"
	"fmt"
)

func TestMin(t *testing.T) {
	actual := min(1, 2, 3, 4, 5)
	if actual != 1 {
		t.Errorf("actual is not 1")
	}
}
func TestMinNoParam(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	min()
	t.Errorf("TestMinNoPram did not panic")
}

func TestMax(t *testing.T) {
	actual := max(1, 4, 6, 24, 6)
	if actual != 24 {
		t.Errorf("actual is not 24")
	}
}

func TestMaxNoParam(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	max()
	t.Errorf("TestMaxNoPram did not panic")
}
