package main

import (
	"golang.org/x/net/html"
	"net/http"
	"fmt"
)

func main() {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	images := ElementsByTagName(doc, "img")
	heading := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	fmt.Println(images, heading)
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	//TODO:動かない気がする。つづきから
	var d *html.Node
	for _, v := range name {
		if doc.Data == v {
			d = append(d, doc)
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		ElementsByTagName(c, name)
	}
	return d

}
