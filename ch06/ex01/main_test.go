package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	if actual := x.String(); actual != "{1 144}" {
		t.Errorf("add result is not {1 144} %v", actual)
	}
}
func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	if actual := x.String(); actual != "{1 144}" {
		t.Errorf("add result is not {1 144} %v", actual)
	}
	x.Remove(1)
	if actual := x.String(); actual != "{144}" {
		t.Errorf("remove result is not {144} %v", actual)
	}
}

func TestClear(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	if actual := x.String(); actual != "{1 144}" {
		t.Errorf("add result is not {1 144} %v", actual)
	}
	x.Clear()
	if actual := x.String(); actual != "{}" {
		t.Errorf("clear result is not {} %v", actual)
	}

}
func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	if actual := x.String(); actual != "{1 144}" {
		t.Errorf("add result is not {1 144} %v", actual)
	}
	z := x.Copy()
	z.Add(2)
	if actual := x.String(); actual != "{1 144}" {
		t.Errorf("copy result x is not {1 144} %v", actual)
	}
	if actual := z.String(); actual != "{1 2 144}" {
		t.Errorf("copy result z is not {1 2 144} %v", actual)
	}
}
func TestLen(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	if actual := x.String(); actual != "{1 144}" {
		t.Errorf("add result is not {1 144} %v", actual)
	}
	if actual := x.Len(); actual != 2 {
		t.Errorf("x.Len is not 2 %v", actual)
	}
}
