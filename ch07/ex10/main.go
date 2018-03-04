package main

import (
	"sort"
)

func main() {
	s := sortPalindrome("hogeegoh")
	IsPalindrome(&s)
}

type sortPalindrome string

func (s *sortPalindrome) Len() int {
	//swapではruneにしてるけどlen([]rune(string(*s)))でしてないから日本語通らなそう
	return len(*s)
}
func (s *sortPalindrome) Less(i, j int) bool {
	rs := []rune(string(*s))
	return rs[i] < rs[j]
}
func (s *sortPalindrome) Swap(i, j int) {
	rs := []rune(string(*s))
	rs[i], rs[j] = rs[j], rs[i]
	n := sortPalindrome(string(rs))
	s = &n
}

func IsPalindrome(s sort.Interface) bool {
	len := s.Len() - 1
	for i, j := 0, len; i < len/2; {
		//s.Less(i,j)||s.Less(j,i)でもいける
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
		i++
		j--
	}
	return true
}
