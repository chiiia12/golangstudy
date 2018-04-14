package main

import (
	"testing"
)

type Sample struct {
	y *Sample
}

func TestCirculation(t *testing.T) {
	a, b, c := &Sample{}, &Sample{}, &Sample{}
	a.y = b
	b.y = c
	c.y = a
	if !IsCirculation(a) {
		t.Errorf("IsCirculation(a) is not true")
	}
	d := &Sample{}
	d.y = &Sample{}
	if IsCirculation(d) {
		t.Errorf("IsCirculation(d) is not false")
	}
}
