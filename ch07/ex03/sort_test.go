package treesort

import "testing"

func TestString(t *testing.T) {

	x :=Sort([]int{1,7,29,30,88,76,23,16})
	x.String()
}
