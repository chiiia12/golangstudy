package ex05

import (
	"fmt"
	"bytes"
)

const ENV = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func (s *IntSet) Add(x int) {
	word, bit := x/ENV, uint(x%ENV)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < ENV; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", ENV*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
