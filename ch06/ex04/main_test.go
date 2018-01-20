package main

import (
	"testing"
	"reflect"
)

func TestElems(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)
	if actual := x.Elems(); !reflect.DeepEqual(actual, []int{1, 2}) {
		t.Errorf("x.Elems doesn't return []int{1,2}.actual is %v", actual)
	}
}
