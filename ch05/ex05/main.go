package main

import (
	"net/http"
	"golang.org/x/net/html"
	"strings"
	"fmt"
)

func main() {
	words, images, err := CountWordsImages("https://golang.org")
	//words, images, err := CountWordsImages("https://github.com/")
	fmt.Printf("words is %v,images %v,error is %v", words, images, err)
}
func CountWordsImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	z := html.NewTokenizer(resp.Body)
	words, images = countWordsAndImages(z)
	resp.Body.Close()
	return
}
func countWordsAndImages(z *html.Tokenizer) (words, images int) {
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
			b, _, _ := z.TagAttr()
			if (string(b) == "src") {
				images++
			}
		case html.EndTagToken:
			isSkip = false
		case html.TextToken:
			if isSkip {
				continue
			}
			text := strings.Fields(string(z.Text()))
			words += len(text)
		}
	}

}
