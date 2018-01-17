package main

import (
	"golang.org/x/net/html"
	"net/http"
	"fmt"
)

func main() {
	resp, err := http.Get("https://recruit-tech.co.jp/")
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	images := ElementsByTagName(doc, "img")
	heading := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	for _, v := range images {
		fmt.Println(v.Data, v.Attr)
	}
	for _, v := range heading {
		fmt.Println(v.Data, v.Attr)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var d []*html.Node
	if doc.Type == html.ElementNode {
		for _, v := range name {
			//fmt.Println(doc.Data, v)
			if doc.Data == v {
				d = append(d, doc)
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		d = append(d, ElementsByTagName(c, name...)...)
		//fmt.Println(ElementsByTagName(c, name...))
	}
	return d
}
