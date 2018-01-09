package main

import (
	"testing"
	"strings"
	"golang.org/x/net/html"
	"bytes"
)

func Test(t *testing.T) {
	parameters := []struct {
		input, expected string
	}{
		{"<div id='hoge'>aaaaa</div>", "<html>\n" +
			"  <head/>\n" +
			"  <body>\n" +
			"    <div id='hoge'>\n" +
			"  </body>\n" +
			"</html>\n"},
		{"<div id='hoge'>aaaaa</div><div>bbb</div>",
			"<html>\n" +
				"  <head/>\n" +
				"  <body>\n" +
				"    <div id='hoge'>\n" +
				"  </body>\n" +
				"</html>\n"},
	}
	for i := range parameters {
		t.Run(parameters[i].input, func(t *testing.T) {
			r := strings.NewReader(parameters[i].input)
			doc, _ := html.Parse(r)
			buf := &bytes.Buffer{}
			ElementById(doc, "hoge", buf)
			outputString := buf.String()
			if outputString != parameters[i].expected {
				t.Errorf("outputString is %v", outputString)
			}
		})
	}
}
