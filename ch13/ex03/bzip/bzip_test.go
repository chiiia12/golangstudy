package bzip

import (
	"testing"
	"io"
	"bytes"
	"compress/bzip2"
	"sync"
)

func TestBzip2(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := NewWriter(&compressed)

	tee := io.MultiWriter(w, &uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hello")
	}
	if err := w.Close(); err != nil {
		t.Errorf(err.Error())
	}
	if actual, expect := compressed.Len(), 255; actual != expect {
		t.Errorf("actual is not expect.actual: %v expect: %v", actual, expect)
	}
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Errorf("bytes of uncompressed is not equal to bytes of decompressed. uncompressed %v. decompressed %v.", uncompressed, decompressed)
	}
}

func TestBzip2ConCurrent(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := NewWriter(&compressed)

	tee := io.MultiWriter(&uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hello")
	}
	tee = io.MultiWriter(w)
	var wg sync.WaitGroup
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			io.WriteString(tee, "hello")
			wg.Done()
		}()
	}
	wg.Wait()

	if err := w.Close(); err != nil {
		t.Errorf(err.Error())
	}
	if actual, expect := compressed.Len(), 255; actual != expect {
		t.Errorf("actual is not expect.actual: %v expect: %v", actual, expect)
	}
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Errorf("bytes of uncompressed is not equal to bytes of decompressed. uncompressed %v. decompressed %v.", uncompressed, decompressed)
	}
}
