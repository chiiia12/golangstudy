package main

import (
	"testing"
)

func BenchmarkArgsBuilder1(b *testing.B) {
	input := []string{"one", "two", "three"}
	for i := 0; i < b.N; i++ {
		argsBuilder1(input)
	}
}

func BenchmarkArgsBuilder2(b *testing.B) {
	input := []string{"one", "two", "three"}
	for i := 0; i < b.N; i++ {
		argsBuilder2(input)
	}
}
