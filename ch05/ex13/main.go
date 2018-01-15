package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	nu "net/url"
	"os"
	"net/http"
	"io/ioutil"
)

func main() {
	//breadthFirst(crawl, os.Args[1:])
	breadthFirst(crawl, []string{"https://golang.org"})
}
func crawl(url string) []string {
	u, _ := nu.Parse(url)
	if err := os.Mkdir(u.Host, 0777); err != nil {
		fmt.Println(err)
	}
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	for _, v := range list {
		u2, _ := nu.Parse(v)
		if u2.Host == u.Host {
			if err := os.Mkdir(u2.Host+u2.Path, 0777); err != nil {
				fmt.Println(err)
				var b []byte
				if b, err = downloadHtml(v); err != nil {
					fmt.Println(err)
				}
				//TODO:ディレクトリの移動または指定したところにファイル作れてない。
				//TODO:file existsをなんとかする
				os.Rename(u2.Host, u2.Host+"/"+u2.Path)
				file, _ := os.Create("index.html")
				file.Write(b)
			}

		}
	}
	return list
}

func downloadHtml(url string) (byteArray []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	byteArray, _ = ioutil.ReadAll(resp.Body)
	return
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}