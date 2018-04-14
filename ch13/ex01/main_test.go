package main

import "testing"

func TestEqual(t *testing.T) {
	var a, b float32
	a, b = 1, 1+1E-10
	if !Equal(a, b) {
		t.Error("Equal(a,b) is not true")
	}
	var c, d float64
	c, d = 1, 1+1E-10
	if !Equal(c, d) {
		t.Error("Equal(c,d) is not true")
	}
	//complexは比較できない？
	//var e, f complex64
	//e, f = complex64(1+1i), complex64(1.0+1.0E-10+1i)
	//if !Equal(e, f) {
	//	t.Error("Equal(e,f) is not true")
	//}

	var h, g complex128
	h, g = complex(1, 1), complex(1+1E-10, 1)
	if !Equal(h, g) {
		t.Error("Equal(h,g) is not true")
	}
}
