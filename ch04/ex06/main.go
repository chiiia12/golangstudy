package main

import (
	"unicode"
)

func convertToASCII(str []byte) []byte {
	spaceCount := 0
	strlen := len(str)
	for i, v := range str {
		if unicode.IsSpace(rune(v)) {
			spaceCount++
		} else {
			spaceCount = 0
		}

		if spaceCount == 2 {
			for j := i; j < len(str)-1; j++ {
				str[j] = str[j+1]
			}
			strlen--
			spaceCount = 0
		}
	}
	return str[:strlen]
}
