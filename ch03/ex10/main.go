// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.

// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"fmt"
	"os"
	"bytes"
	"math"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	m := int(math.Mod(float64(len(s)), 3))
	for i, _ := range s {
		buf.WriteString(string(s[i]))
		if (i+1)%3 == m && i+1 < len(s) {
			buf.WriteString(",")
		}
	}

	return buf.String()
}
