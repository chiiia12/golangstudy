package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	var w WordCount
	w.Write([]byte("hello hello hello"))
	fmt.Println("wordcount is", w)
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
