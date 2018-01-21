package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	var w WordCount
	name := "Dolly Dolly"
	fmt.Fprintf(&w, "hello, %s", name)
	fmt.Println(w)
}

type WordCount int

func (w *WordCount) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		*w += WordCount(1)
	}
	return len(p), nil
}
