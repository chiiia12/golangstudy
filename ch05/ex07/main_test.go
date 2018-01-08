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
		{"<div>aaaaa</div>", "<html>\n" +
			"  <head/>\n" +
			"  <body>\n" +
			"    <div>\n" +
			"      aaaaa\n" +
			"    </div>\n" +
			"  </body>\n" +
			"</html>\n"},
		{"<div></div>", "<html>\n" +
			"  <head/>\n" +
			"  <body>\n" +
			"    <div/>\n" +
			"  </body>\n" +
			"</html>\n"},
		{"<a href='http://hoge.com'>", "<html>\n" +
			"  <head/>\n" +
			"  <body>\n" +
			"    <a href='http://hoge.com'/>\n" +
			"  </body>\n" +
			"</html>\n"},
	}
	for i := range parameters {
		t.Run(parameters[i].input, func(t *testing.T) {
			r := strings.NewReader(parameters[i].input)
			doc, _ := html.Parse(r)
			buf := &bytes.Buffer{}
			forEachNode(doc, startElement, endElement, buf)
			outputString := buf.String()
			if outputString != parameters[i].expected {
				t.Errorf("outputString is %v", outputString)
			}
		})
	}
}
