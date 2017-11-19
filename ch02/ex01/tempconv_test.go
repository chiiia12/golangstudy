package tempconv

import "testing"

func TestCToK0(t *testing.T) {
	var expect Celsius = 0
	if CToK(expect) != AbsoluteZeroK {
		t.Error(`CToK(0) is not '-273.15'`)
	}
}
func TestCToK10(t *testing.T) {
	var expect Celsius = 10
	if CToK(expect) != AbsoluteZeroK+10 {
		t.Error(`CTok(10) is not '-263.15'`)
	}
}
