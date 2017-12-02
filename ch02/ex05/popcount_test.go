package popcount

import (
	"testing"
	"fmt"
	"gopl.io/ch2/popcount"
)

func BenchmarkPopCountEx05(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountEx05(63)
	}
}
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(63)
	}
}
func TestPopCount(t *testing.T) {
	fmt.Println(`popcount is `, PopCountEx05(16))
}
