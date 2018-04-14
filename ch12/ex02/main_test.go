package main

import (
	"testing"
)

type Link struct {
	X *Link
}

func TestDisplay(t *testing.T) {
	a, b, c := &Link{}, &Link{}, &Link{}
	a.X = b
	b.X = c
	c.X = a
	Display("a", a)
}
