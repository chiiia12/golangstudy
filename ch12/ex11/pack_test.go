package main

import "testing"

func TestPack(t *testing.T) {
	var data struct {
		Labels     []string `http:"label"`
		MaxResults int      `http:"maxvalue"`
		Exact      bool     `http:"exact"`
	}
	data.Labels = []string{"hoge", "fuga"}
	data.MaxResults = 4
	data.Exact = true
	result := Pack(&data)
	if result != "label=hoge&label=fuga&maxvalue=4&exact=true" {
		t.Errorf("result isnot equal %v", result)
	}
}
