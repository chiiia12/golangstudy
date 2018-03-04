// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"log"
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

//実質キャンセルしてるだけ。mainを終わらせてるからキャンセルになってるだけ。
func main() {
	for _, url := range os.Args[1:] {
		go func() {
			fetch(done, url)
		}()
	}
	for {
		select {
		case <-done:
			log.Println("done")
			return
		}
	}

}

func fetch(done chan struct{}, url string) {
	log.Println("fetch" + url)
	if canceled() {
		log.Println("cancel" + url)
		return
	}
	select {
	case <-done:
		log.Println("done in fetch")
		return
	default:
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	_, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("fetched url :%s\n", url)
	done <- struct{}{}
}
