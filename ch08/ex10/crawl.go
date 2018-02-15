package main

import (
	"os"
	"fmt"
	"gopl.io/ch5/links"
	"log"
)

var done = make(chan struct{})

func canceled() bool {
	select {
	case <-done:
		log.Println("canceled is done")
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

	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
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
	for {
		select {
		case <-done:
			for range worklist {

			}
			log.Println("main's loop is done")
			return
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
	//case <-done:
	//	log.Print("in crawl methoed :done ")
	//	return nil
	}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
