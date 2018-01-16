package main

import "testing"

func TestMultiStringJoin(t *testing.T) {
	actual := join(",", "hoge", "fuga", "piyo")
	if actual != "hoge,fuga,piyo" {
		t.Errorf("actual is \"hoge,fuga,piyo\"")
	}
}
