package main

import "testing"

func TestCountWords(t *testing.T) {
	wordCount := countWords([]string{"test1.txt"})
	if actual := wordCount["This"]; actual != 2 {
		t.Errorf("wordCount[\"This\"] is not 2. actual is %v", actual)
	}
	if actual := wordCount["is"]; actual != 2 {
		t.Errorf("wordCount[\"is\"] is not 2. actual is %v", actual)
	}
	if actual := wordCount["sample"]; actual != 1 {
		t.Errorf("wordCount[\"sample\"] is not 1. actual is %v", actual)
	}
	if actual := wordCount["sample2"]; actual != 1 {
		t.Errorf("wordCount[\"sample\"] is not 1. actual is %v", actual)
	}
}
