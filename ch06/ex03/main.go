package main

import (
	"fmt"
	"bytes"
)

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(5)
	y.Add(1)
	y.Add(2)
	y.Add(3)
	y.Add(4)

	fmt.Println(x.String())
	fmt.Println(y.String())
	//x.UnionWith(&y)
	//x.IntersectWith(&y)
	x.DifferenceWith(&y)
	fmt.Println(x.String())
}

type IntSet struct {
	words []uint64
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//積集合
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, v := range s.words {
		s.words[i] = v & t.words[i]
	}
}

//差集合
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, v := range s.words {
		intersect := v & t.words[i]
		s.words[i] = v & ^intersect
	}
}

//対照差
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, v := range s.words {
		tmp := v | t.words[i]
		s.words[i] = tmp &^ (v & t.words[i])
	}
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
