package main

import (
	"testing"
)

func TestExpandContain(t *testing.T) {
	s := "hogehoge$foohogehoge$foo"
	act := expand(s, func(s string) string {
		return "<<" + s + ">>"
	})
	exp := "hogehoge<<foo>>hogehoge<<foo>>"
	if act != exp {
		t.Errorf("expect is %v but actual is %v", exp, act)
	}

}

func TestExpandNotContain(t *testing.T) {
	s := "hogehogehogehoge"
	act := expand(s, func(s string) string {
		return "<<" + s + ">"
	})
	exp := "hogehogehogehoge"
	if act != exp {
		t.Errorf("expect is %v but actual is %v", exp, act)
	}

}
