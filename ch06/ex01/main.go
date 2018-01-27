package main

import (
	"fmt"
	"bytes"
)

func main() {
}

type IntSet struct {
	words []uint64
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

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

//要素数を返します
func (s *IntSet) Len() int {
	count := 0
	for _, v := range s.words {
		tmp := v
		var j uint
		for tmp != 0 {
			//0かどうかを比較すればよいよ
			if v&(1<<j) == (1 << j) {
				count++
			}
			tmp = tmp &^ (1 << j)
			j++
		}
	}
	return count
}

//セットからｘを取り除きます
//存在しない数字いれると落ちる
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if s.words[word]&(1<<bit) == 1<<bit {
		s.words[word] &^= 1 << bit
	}
}

//セットからすべての要素を取り除きます
func (s *IntSet) Clear() {
	s.words = nil
}

//セットのコピーを返します
func (s *IntSet) Copy() *IntSet {
	var copy IntSet
	for _, v := range s.words {
		copy.words = append(copy.words, v)
	}
	return &copy
}
