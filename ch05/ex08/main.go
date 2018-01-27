package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
	"golang.org/x/net/html"
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
	n := ElementById(doc, "lowframe", os.Stdout)
	fmt.Println(n.Data, n.Attr)

}

func ElementById(doc *html.Node, id string, out io.Writer) *html.Node {

	//forEachNodeまでid渡さなくてもよかったかも。クロージャーで保持できる
	forEachNode(doc, id, startElementById, endElement, out)
	return doc

}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, out io.Writer, id string) bool, out io.Writer) bool {

	if pre != nil {
		if !pre(n, out, id) {
			//os.Exit(0)
			return false
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, id, pre, post, out) {
			return false
		}
	}
	if post != nil {
		post(n, out, "")
	}
	return true
}

var depth int

func endElement(n *html.Node, out io.Writer, id string) bool {
	if n.Type == html.ElementNode {
		depth--
		fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
	}
	return true
}

func startElementById(n *html.Node, out io.Writer, id string) bool {
	isOmit := false
	if n.FirstChild == nil {
		isOmit = true;
	}
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
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				depth--
				return false
			}
		}
	case html.CommentNode:
		fmt.Fprintf(out, "%*s<!--%s-->\n", depth*2, "", n.Data)
	case html.TextNode:
		if (n.Data != "") {
			fmt.Fprintf(out, "%*s%s\n", depth*2, "", n.Data)
		}

	}
	return true
}
