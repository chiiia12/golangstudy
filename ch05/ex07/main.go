package main

import (
	"golang.org/x/net/html"
	"fmt"
	"os"
	"net/http"
	"io"
)

func main() {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.Parse error: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement, os.Stdout)
}

func forEachNode(n *html.Node, pre func(n *html.Node, isOmit bool, out io.Writer), post func(n *html.Node, out io.Writer), out io.Writer) {

	if n.FirstChild == nil {
		pre(n, true, out)
		return
	}
	if pre != nil {
		pre(n, false, out)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, out)
	}
	if post != nil {
		post(n, out)
	}
}

var depth int

func startElement(n *html.Node, isOmit bool, out io.Writer) {
	str := ""
	for _, a := range n.Attr {
		str += " " + a.Key + "='" + a.Val + "'"
	}
	switch(n.Type) {
	case html.ElementNode:
		if isOmit {
			fmt.Fprintf(out, "%*s<%s%s/>\n", depth*2, "", n.Data, str)
		} else {
			fmt.Fprintf(out, "%*s<%s%s>\n", depth*2, "", n.Data, str)
			depth++
		}
	case html.CommentNode:
		fmt.Fprintf(out, "%*s<!--%s-->\n", depth*2, "", n.Data)
	case html.TextNode:
		if (n.Data != "") {
			fmt.Fprintf(out, "%*s%s\n", depth*2, "", n.Data)
		}

	}
}
func endElement(n *html.Node, out io.Writer) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
	}
}
