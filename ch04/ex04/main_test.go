package main

import (
	"testing"
)

func TestRemoveDeplicated1(t *testing.T) {
	var s = []int{0, 1, 2, 3, 4, 5, 5, 6, 7, 8}
	s = removeDepulicate(s, 0, 8)
	if l := len(s); l != 9 {
		t.Errorf("length of removeDeplicate(s[:8]) is not 9,actual %v , %v", l, s)
	}
}

func TestRemoveDeplicated2(t *testing.T) {
	var s = []int{0, 1, 2, 3, 4, 5, 5, 6, 7, 8}
	s = removeDepulicate(s, 0, 3)
	if l := len(s); l != 10 {
		t.Errorf("length of removeDeplicate(s[:3]) is not 10,actual %v", l)
	}
}

func TestRemoveDeplicated3(t *testing.T) {
	var s = []int{0, 5, 2, 3, 4, 5, 5, 6, 7, 8}
	s = removeDepulicate(s, 0, 5)
	if l := len(s); l != 10 {
		t.Errorf("length of removeDeplicate(s[:5]) is not 10,actual %v", l)
	}
}
