package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

var linkMap map[string]string

func init() {
	linkMap = map[string]string{
		"a":    "href",
		"img":  "src",
		"link": "href",
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		links = appendlink(links, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func appendlink(links []string, n *html.Node) []string {
	if v := linkMap[n.Data]; v != "" {
		for _, a := range n.Attr {
			if a.Key == v {
				links = append(links, a.Val)
			}
		}
	}
	return links
}
