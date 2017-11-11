package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestFetchUrl(t *testing.T) {
	buf := &bytes.Buffer{}
	url := []string{"http://gopl.io"}
	fetchUrl(buf, url)
	outputString := buf.String()
	if !strings.Contains(outputString, "The Go Programming Language") {
		t.Error(`outputString does not contains "The Go Programming Language"`)
	}
}
