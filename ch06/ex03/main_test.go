package main

import "testing"

func TestIntersectWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(200)
	y.Add(3)
	y.Add(4)
	y.Add(5)
	y.Add(1000)
	x.IntersectWith(&y)
	if actual := x.String(); actual != "{3}" {
		t.Errorf("intersectWith result isnot {3},actual %v", actual)
	}
}
func TestDifferenceWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	y.Add(3)
	y.Add(4)
	y.Add(5)
	x.DifferenceWith(&y)
	if actual := x.String(); actual != "{1 2}" {
		t.Errorf("differenceWith result isnot {1 2},actual %v", actual)
	}
}
func TestSymmetricDifference(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	y.Add(3)
	y.Add(4)
	y.Add(5)
	x.SymmetricDifference(&y)
	if actual := x.String(); actual != "{1 2 4 5}" {
		t.Errorf("SymmetricDifference result isnot {1 2 4 5},actual %v", actual)
	}

}
