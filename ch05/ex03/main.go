package main

import (
	"golang.org/x/net/html"
	"os"
	"fmt"
)

func main() {
	z := html.NewTokenizer(os.Stdin)
	isSkip := false
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			return
		}
		if tt == html.StartTagToken {
			tagname, _ := z.TagName()
			if a := string(tagname); a == "script" || a == "style" {
				isSkip = true
				continue
			}
		}
		if tt == html.EndTagToken {
			isSkip = false
		}
		if tt == html.TextToken {

			if isSkip {
				continue
			}
			fmt.Println(string(z.Text()))
		}
	}
}
