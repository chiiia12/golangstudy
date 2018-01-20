package main

import "testing"

func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll(1, 2, 144)
	if actual := x.String(); actual != "{1 2 144}" {
		t.Errorf("addall result is not {1 2 144} %v", actual)
	}
}
