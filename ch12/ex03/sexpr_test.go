package sexpr

import (
	"testing"
	"fmt"
)

func TestBoolTrueEncode(t *testing.T) {
	var isHoge bool = true
	data, err := Marshal(isHoge)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "t" {
		t.Errorf("Marshal(isHoge) is not t.actual is %v", string(data))
	}
}
func TestBoolFalseEncode(t *testing.T) {
	var isHoge bool = false
	data, err := Marshal(isHoge)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "nil" {
		t.Errorf("Marshal(isHoge) is not nil.actual is %v", string(data))
	}
}

func TestFLoatEncode(t *testing.T) {
	a := float64(1.98988)
	data, _ := Marshal(a)
	if string(data) != fmt.Sprintf("%f", 1.98988) {
		t.Errorf("Marshal(a) is not 1.98988. actual is %v", string(data))
	}
}
func TestComplexEncode(t *testing.T) {
	c := complex(1, 2)
	data, _ := Marshal(c)
	if string(data) != fmt.Sprintf("#C(%f,%f)", 1.0, 2.0) {
		t.Errorf("Marshal(c) is not #C(1.0,2.0). actual is %v", string(data))
	}
}
func TestInterface(t *testing.T) {
	type Sample struct {
		Interface interface{}
	}
	arr := Sample{
		Interface: []int{1, 2, 3},
	}
	data, _ := Marshal(arr)
	if string(data) != "((Interface (\"[]int\" (1 2 3))))" {
		t.Errorf("Marshal(arr) is not ((Interface (\"[]int\" (1 2 3)))). actual is %v", string(data))
	}
}
