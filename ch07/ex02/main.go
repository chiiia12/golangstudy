package main

import (
	"io"
	"fmt"
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
	var w io.Writer
	w.Write([]byte("hogehoge"))
	i, l := CountingWriter(w)
	fmt.Printf("writer %v\n", i)
	fmt.Println("int64", *l)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var result WriterWrapper
	result.w = w
	return &result, &result.len
}

/*
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x20 pc=0x1088fe8]

goroutine 1 [running]:
main.main()
/Users/vv001292/workspace/golangstudy/ch07/ex02/main.go:21 +0x58

Process finished with exit code 2
*/
