package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	var l LineCount
	fmt.Fprintf(&l, "hello hello \nhello hello")
	fmt.Println(l)
}

type LineCount int

func (l *LineCount) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		*l += LineCount(1)
	}
	return len(p), nil
}
