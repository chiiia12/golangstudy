package main

import (
	"testing"
	"crypto/sha256"
)

func TestPopCount1(t *testing.T) {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	if result := popcount(c1, c2); result <= 0 {
		t.Errorf("\"x\" and \"X\" popcount is 0 .actual is %v", result)
	}

}
func TestPopCount2(t *testing.T) {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("x"))
	if result := popcount(c1, c2); result != 0 {
		t.Errorf("\"x\" and \"X\" popcount is not 0 .actual is %v", result)
	}

}
