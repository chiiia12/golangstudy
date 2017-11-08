package main

import (
	"testing"
)

func TestEchoArgs(t *testing.T) {
	var args = []string{"zero", "one"}
	if output := echoArgs(args); output != "0:zero\n1:one\n" {
		t.Error(`echoArgs("args") is not '0:zero\n1:one\n' actual is %q`, output)
	}
}
