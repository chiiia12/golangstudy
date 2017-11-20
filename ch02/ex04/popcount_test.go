package popcount

import (
	"testing"
)

func BenchmarkPopCount64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount64(10000)
	}
}
func BenchmarkPopCountNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountNormal(10000)
	}
}
