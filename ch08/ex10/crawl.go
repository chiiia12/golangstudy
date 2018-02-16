package main

import (
	"os"
	"fmt"
	"log"
	"./links"
)

var done = make(chan struct{})

func canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		log.Println("something inputted")
		close(done)
	}()
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() {
		select {
		case <-done:
			return
		case worklist <- os.Args[1:]:
		}
	}()

	for i := 0; i < 20; i++ {
		go func() {
			if canceled() {
				return
			}
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					select {
					case <-done:
						return
					case worklist <- foundLinks:
					}
				}()
			}
		}()
	}
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				select {
				case unseenLinks <- link:
				case <-done:
					log.Println("done get in list loop")
					return
				}
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	if canceled() {
		return nil
	}
	fmt.Println(url)
	select {
	case tokens <- struct{}{}:
	case <-done:
		return nil
	}
	list, err := links.Extract(done, url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
