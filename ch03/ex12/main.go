package ex12

import "strings"

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, r := range s1 {
		if !strings.Contains(s2, string(r)) {
			return false
		}
	}
	return true
}
