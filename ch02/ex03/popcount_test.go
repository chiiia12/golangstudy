package popcount

import (
	"testing"
)

func BenchmarkPopCountFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(64)
	}
}
func BenchmarkPopCountNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountNormal(64)
	}
}
