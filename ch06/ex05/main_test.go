package ex05

import "testing"

func TestAdd(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)
	if actual := x.String(); actual != "{1 2}" {
		t.Errorf("x add result is not {1 2} actual is %v", actual)
	}
}
