package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(200)
	if actual := x.String(); actual != "{1 2 3 200}" {
		t.Errorf("x.String() = %v. want {1 2 3 200}", actual)
	}
}
func TestHas(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(200)
	if actual := x.Has(1); !actual {
		t.Errorf("x.Has(1) = false")
	}
}

func TestUnionWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)

	y.Add(2)
	y.Add(3)
	y.Add(4)
	x.UnionWith(&y)

	if actual := x.String(); actual != "{1 2 3 4}" {
		t.Errorf("x.String() = %v. want {1 2 3 4}", actual)
	}
}
