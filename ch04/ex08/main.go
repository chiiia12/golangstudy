package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := map[string]int{
		"control": 0,
		"digit":   0,
		"graphic": 0,
		"letter":  0,
		"lower":   0,
		"mark":    0,
		"number":  0,
		"print":   0,
		"punct":   0,
		"space":   0,
		"symbol":  0,
		"title":   0,
		"upper":   0,
	}
	//counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsControl(r) {
			counts["control"]++
			continue
		}
		if unicode.IsDigit(r) {
			counts["digit"]++
			continue
		}
		if unicode.IsGraphic(r) {
			counts["graphic"]++
			continue
		}
		if unicode.IsLower(r) {
			counts["lower"]++
			continue
		}
		if unicode.IsMark(r) {
			counts["mark"]++
			continue
		}
		if unicode.IsLower(r) {
			counts["lower"]++
			continue
		}
		if unicode.IsMark(r) {
			counts["mark"]++
			continue
		}
		if unicode.IsNumber(r) {
			counts["number"]++
			continue
		}
		if unicode.IsPrint(r) {
			counts["print"]++
			continue
		}
		if unicode.IsPunct(r) {
			counts["punct"]++
			continue
		}
		if unicode.IsSpace(r) {
			counts["space"]++
			continue
		}
		if unicode.IsSymbol(r) {
			counts["symbol"]++
			continue
		}
		if unicode.IsTitle(r) {
			counts["title"]++
			continue
		}
		if unicode.IsUpper(r) {
			counts["upper"]++
			continue
		}

		utflen[n]++
	}
	for i, v := range counts {
		fmt.Println(i, v)
	}
}
