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

func TestFetchUrlNotHasPrefix(t *testing.T) {
	buf := &bytes.Buffer{}
	url := []string{"gopl.io"}
	fetchUrl(buf, url)
	outputString := buf.String()
	if !strings.Contains(outputString, "The Go Programming Language") {
		t.Error(`outputString does not contains "The Go Programming Language"`)
	}
	if !strings.Contains(outputString, "200 OK") {
		t.Error(`outputString does not contains "200 OK"`)
	}
}
