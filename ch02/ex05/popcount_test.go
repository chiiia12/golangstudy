package popcount

import (
	"testing"
	"gopl.io/ch2/popcount"
)

func BenchmarkPopCount64(b *testing.B) {
	popcount.PopCount(64)
}
func TestPopCount64(t *testing.T) {
	PopCount64(64)
}
