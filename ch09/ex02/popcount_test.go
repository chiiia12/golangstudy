package ex02

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	for i := 0; i < 20; i++ {
		go func() {
			PopCount(64)
		}()
	}
}
