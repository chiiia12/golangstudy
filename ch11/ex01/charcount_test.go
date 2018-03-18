package main

import (
	"testing"
	"strings"
)

func TestCharCount(t *testing.T) {
	for _, test := range []struct {
		samplestr string
		counts    map[rune]int
		utflen    []int
		invalid   int
	}{
		{
			"aaabbbccc",
			map[rune]int{'a': 3, 'b': 3, 'c': 3},
			[]int{0, 9, 0, 0, 0},
			0,
		},
		{
			"こんにちは世界",
			map[rune]int{'こ': 1, 'ん': 1, 'に': 1, 'ち': 1, 'は': 1, '世': 1, '界': 1},
			[]int{0, 0, 0, 7, 0},
			0,
		},
		{
			"こんにちは世界\300", //invalid number
			map[rune]int{'こ': 1, 'ん': 1, 'に': 1, 'ち': 1, 'は': 1, '世': 1, '界': 1},
			[]int{0, 0, 0, 7, 0},
			1,
		},
	} {

		reader := strings.NewReader(test.samplestr)
		counts, utflen, invalid := charcount(reader)
		if len(counts) != len(test.counts) {
			t.Errorf("len(counts)=%v. want %v", len(counts), len(test.counts))
		}
		for k, v := range test.counts {
			if counts[k] != v {
				t.Errorf("counts[%v]=%v want %v", k, v, counts[k])
			}
		}
		if len(utflen) != len(test.utflen) {
			t.Errorf("len(utflen)=%v want %v ", len(utflen), len(test.utflen))
		}
		for i, v := range test.utflen {
			if utflen[i] != v {
				t.Errorf("utflen[%v] = %v want %v", i, v, utflen[i])
			}
		}
		if invalid != test.invalid {
			t.Errorf("invalid = %v want %v", invalid, test.invalid)
		}
	}
}
