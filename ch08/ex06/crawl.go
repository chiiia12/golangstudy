package main

import (
	"fmt"
	"log"
	"gopl.io/ch5/links"
	"flag"
)

var depth = flag.Int("depth", 3, "input find links depth")

type leveledUrlInfo struct {
	url   string
	depth int
}

func main() {
	flag.Parse()
	fmt.Printf("depth is :%v flag.Args is :%v\n", &depth, flag.Args())
	worklist := make(chan []*leveledUrlInfo)
	unseenLinks := make(chan string)

	go func() {
		var tmp []*leveledUrlInfo
		for _, v := range flag.Args() {
			tmp = append(tmp, &leveledUrlInfo{v, 0})
		}
		worklist <- tmp
	}()

	for i := 0; i < 20; i++ {
		//20個goroutineを立ち上げる
		go func() {
			num := 0
			for link := range unseenLinks {
				//fmt.Println(link)
				foundLinks := crawl(link, num)
				num++
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link.url
			}
		}
	}
}

func crawl(url string, d int) []*leveledUrlInfo {
	if d > *depth {
		return []*leveledUrlInfo{}
	}
	fmt.Printf("url is :%v depth is :%v\n", url, d)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	var resultlist []*leveledUrlInfo
	for _, v := range list {
		resultlist = append(resultlist, &leveledUrlInfo{v, d + 1})
	}
	return resultlist
}
