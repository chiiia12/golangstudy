package main

import (
	"strings"
	"testing"
)

func TestGatherContainsFiles(t *testing.T) {
	var files = []string{"test1.txt", "test2.txt"}
	duplicatefiles := make(map[string][]string)
	gatherContainsFiles(files, duplicatefiles)
	if actual := duplicatefiles["a"]; strings.Join(actual[0:], ",") != "test1.txt" {
		t.Error(`duplicatefiles["a"] is not "test1.txt" actual is %d`, strings.Join(actual[0:], " "))
	}

	if actual := duplicatefiles["b"]; strings.Join(actual[0:], ",") != "test1.txt,test2.txt" {
		t.Error(`duplicatefiles["b"] is not "test1.txt","test2.txt" actual is %d`, strings.Join(actual[0:], " "))
	}

	if actual := duplicatefiles["c"]; strings.Join(actual[0:], ",") != "test2.txt" {
		t.Error(`duplicatefiles["c"] is not "test2.txt" actual is %d`, strings.Join(actual[0:], " "))
	}
}
