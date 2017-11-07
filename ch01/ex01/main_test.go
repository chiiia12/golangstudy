package main

import "testing"

func TestEchoArgs(t *testing.T) {
	var args = []string{"main", "aaa"}
	if echoArgs(args) != "main aaa" {
		t.Error(`echoArgs('main aaa') is not 'main aaa'`)
	}
}
