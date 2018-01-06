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
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			t, _ := z.TagName()
			if string(t) == "script" || string(t) == "style" {
				isSkip = true
			}
		case html.EndTagToken:
			isSkip = false
		case html.TextToken:
			if isSkip {
				continue
			}
			fmt.Printf(string(z.Text()))
		}
	}
}
