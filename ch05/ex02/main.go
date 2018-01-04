package main

import (
	"golang.org/x/net/html"
	"os"
	"fmt"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countNode: %v\n", err)
		os.Exit(1)
	}
	count := make(map[string]int)
	for i, v := range countNode(count, doc) {
		fmt.Println(i, v)
	}
}

func countNode(count map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	////子を探る
	if n.FirstChild != nil {
		count = countNode(count, n.FirstChild)
	}
	if n.NextSibling != nil {
		count = countNode(count, n.NextSibling)
	}
	return count
}
