package main

import "testing"

func TestMapIntSetAdd(t *testing.T) {
	var x MapIntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(200)
	if actual := x.String(); actual != "[1 2 3 200]" {
		t.Errorf("x.String() = %v. want [1 2 3 200]", actual)
	}
}
func TestMapIntSetHas(t *testing.T) {
	var x MapIntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(200)
	if actual := x.Has(1); !actual {
		t.Errorf("x.Has(1) = false")
	}
}

func TestMapIntSetUnionWith(t *testing.T) {
	var x, y MapIntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)

	y.Add(2)
	y.Add(3)
	y.Add(4)
	x.UnionWith(&y)

	if actual := x.String(); actual != "[1 2 3 4]" {
		t.Errorf("x.String() = %v. want [1 2 3 4]", actual)
	}
}
