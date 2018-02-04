package main

import (
	"encoding/xml"
	"os"
	"io"
	"fmt"
	"strings"
)

type Element struct {
	name string
	attr []xml.Attr
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []Element

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			element := Element{tok.Name.Local, tok.Attr}
			stack = append(stack, element)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				arr := []string{}
				for _, v := range stack {
					arr = append(arr, v.name)
				}
				fmt.Printf("%s: %s\n", strings.Join(arr, " "), tok)
			}
		}
	}
}
func containsAll(x []Element, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0].name == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
