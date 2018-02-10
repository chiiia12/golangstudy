package main

import (
	"encoding/xml"
	"os"
	"fmt"
	"io"
	"bytes"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e *Element) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("<%s", e.Type.Local))
	for _, v := range e.Attr {
		buf.WriteString(fmt.Sprintf(" %s=\"%s\"", v.Name.Local, v.Value))
	}
	buf.WriteString(">\n")
	for _, v := range e.Children {
		switch n := v.(type) {
		case *Element:
			buf.WriteString(n.String())
		case CharData:
			buf.WriteString(string(n))
		}
	}
	buf.WriteString(fmt.Sprintf("\n</%s>\n", e.Type.Local))
	return buf.String()
}

func main() {
	e := buildTree(os.Stdin)
	fmt.Printf("%s\n", e)
}

func buildTree(r io.Reader) *Element {
	dec := xml.NewDecoder(r)
	var stack []*Element

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			e := &Element{tok.Name, tok.Attr, nil}
			if len(stack) > 1 {
				index := len(stack) - 1
				stack[index].Children = append(stack[index].Children, e)
			}
			stack = append(stack, e)
		case xml.EndElement:
			switch {
			case len(stack) == 1:
				return stack[0]
			case len(stack) > 1:
				stack = stack[:len(stack)-1]
			}
		case xml.CharData:
			index := len(stack) - 1
			stack[index].Children = append(stack[index].Children, CharData(tok))

		}
	}
	return nil
}
