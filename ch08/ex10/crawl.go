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
		//log.Println("canceled is done")
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
			log.Println("list loop")
		}
		log.Println("worklist loop")
	}
	log.Println("worklist loop done")
	//worklist loopが終わらなくて下のforではdoneをキャッチできない。→main関数が終わらない
	for {
		log.Println("channel for")
		select {
		case <-done:
			log.Println("main's loop is done")
			for range worklist {
				//channelの呼び出し
			}
			for range unseenLinks {
				//channelの呼び出し
			}
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
	case <-done:
		log.Print("in crawl methoed :done ")
		return nil
	}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
