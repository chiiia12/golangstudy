package main

import (
	"fmt"
	"os"
	"bytes"
	"math"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	l := len(s);
	if index := strings.Index(s, "."); index > 0 {
		l =len(s[:index])
	}
	m := int(math.Mod(float64(l), 3))
	for i, _ := range s {
		buf.WriteString(string(s[i]))
		if (i+1)%3 == m && i+1 < l {
			buf.WriteString(",")
		}
	}

	return buf.String()
}
