package main

import (
	"testing"
	"strings"
	"bytes"
)

func TestLimitReader(t *testing.T) {
	r := strings.NewReader("Hello,World")
	buf := new(bytes.Buffer)
	buf.ReadFrom(LimitReader(r, 5))
	if actual := buf.String(); actual != "Hello" {
		t.Errorf("LimitReader doesn't return Hello actual is [%v]", actual)
	}
}
