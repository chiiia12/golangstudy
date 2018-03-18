package ex05

import (
	"testing"
	"strings"
)

func TestSplit(t *testing.T) {
	for _, test := range []struct {
		str  string
		sep  string
		want int
	}{
		{
			"a:b:c",
			":",
			3,
		},
		{
			"hoge,hoge,hoge",
			",",
			3,
		},
		{
			"a aa aaa aaaa",
			" ",
			4,
		},
	} {
		words := strings.Split(test.str, test.sep)
		if got, want := len(words), test.want; got != want {
			t.Errorf("Split(%q,%q) returned %d words,want %d", test.str, test.sep, got, test.want)
		}
	}
}
