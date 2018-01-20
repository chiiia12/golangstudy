package main

import (
	"fmt"
	"bytes"
)

func main() {
}

type IntSet struct {
	words []uint
}

func (s *IntSet) Elems() []int {
	elem := []int{}
	for _, v := range s.words {
		var j uint;
		tmp := v
		for tmp != 0 {
			if tmp&(1<<j) == 1<<j {
				elem = append(elem, int(j))
				tmp = tmp &^ (1 << j)
			}
			j++
		}
	}
	return elem
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
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
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
