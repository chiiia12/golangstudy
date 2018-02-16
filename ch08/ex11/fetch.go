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

func main() {
	done := make(chan struct{})
	responses := make(chan string, 1)
	for _, url := range os.Args[1:] {
		go func() {
			select {
			case <-done:
				log.Println("done has received")
				close(responses)

			case responses <- fetch(url):
				log.Println("responses has received")
				close(done)
			}
		}()
	}
}

func fetch(url string) string {
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
	//fmt.Printf("%s", b)
	fmt.Printf("fetched url :%s\n", url)
	return url
}
