package main

import (
	"io"
	"fmt"
	"os"
)

type WriterWrapper struct {
	len int64
	w   io.Writer
}

func (w *WriterWrapper) Write(p []byte) (n int, err error) {
	n, err = w.w.Write(p)
	w.len += int64(n)
	return
}

func main() {
	std := os.Stdout
	i, l := CountingWriter(std)
	i.Write([]byte("hogehoge"))
	fmt.Printf("writer %v\n", i)
	fmt.Println("int64", *l)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var result WriterWrapper
	result.w = w
	return &result, &result.len
}
