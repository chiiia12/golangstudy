package main

import (
	"fmt"
	"log"
	"gopl.io/ch5/links"
	"flag"
)

var depth = flag.Int("depth", 3, "input find links depth")

func main() {
	flag.Parse()
	fmt.Printf("depth is :%v flag.Args is :%v\n", &depth, flag.Args())
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- flag.Args() }()

	for i := 0; i < 20; i++ {
		go func() {
			num := 0
			for link := range unseenLinks {
				foundLinks := crawl(link, num)
				num++
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl(url string, d int) []string {
	if d > *depth {
		return []string{}
	}
	fmt.Printf("url is :%v depth is :%v\n", url, d)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
