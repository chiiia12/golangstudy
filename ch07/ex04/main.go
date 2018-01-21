package main

import "io"

func main() {

}
func NewReader(s string) io.Reader {
	var result SampleReader
	return &result
}

type SampleReader int

func (r *SampleReader) Read(p []byte) (n int, err error) {

	return len(p), nil
}
