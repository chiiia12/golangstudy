package main

import (
	"testing"
)

func TestComma(t *testing.T) {
	parameters := []struct {
		input, expected string
	}{
		{"123", "123"}, {"1234", "1,234"}, {"1234567", "1,234,567"},
		{"123.123", "123.123"}, {"1234.1234", "1,234.1234"}, {"-1234.1234", "-1,234.1234"},
	}
	for i := range parameters {
		actual := comma(parameters[i].input)
		if actual != parameters[i].expected {
			t.Errorf("comma[%s] is not %s actual is %s", parameters[i].input, parameters[i].expected, actual)
		}
	}
}
