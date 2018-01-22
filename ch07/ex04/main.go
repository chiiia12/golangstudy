package main

import (
	"io"
	"fmt"
	"os"
	"golang.org/x/net/html"
)

func main() {
	r := NewReader("<a href='http://google.com>hoge</a>")
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.parse: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("doc is %v\n", doc)
}

func NewReader(s string) io.Reader {
	var result SampleReader
	result.s = s
	return &result
}

type SampleReader struct {
	s string
}

func (r *SampleReader) Read(p []byte) (n int, err error) {
	return len(p), nil
}
