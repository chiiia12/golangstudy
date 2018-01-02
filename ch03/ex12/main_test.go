package ex12

import (
	"testing"
)

func TestAnagramTrue(t *testing.T) {
	if !anagram("dormitory", "dirtyroom") {
		t.Error("expect is true.but anagram() return false")
	}
}
func TestAnagramDifferentLength(t *testing.T) {
	if anagram("dormitory", "dirtyroo") {
		t.Error("expect is false.but anagram() return true")
	}
}
func TestAnagramFalse(t *testing.T) {
	if anagram("dormitory", "dirtyrooa") {
		t.Error("expect is false.but anagram() return true")
	}
}

func TestAnagramJapanese(t *testing.T) {
	if !anagram("日本語", "語日本") {
		t.Error("expect is false.but anagram() return true")
	}
}
