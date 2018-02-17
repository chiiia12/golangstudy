package ch09

import (
	"net/http"
	"io/ioutil"
	"testing"
	"sync"
	"time"
	"log"
	"fmt"
)

//!+httpRequestBody
func httpGetBody(url string, done chan struct{}) (interface{}, error) {
	select {
	case <-done:
		return nil, fmt.Errorf("httpGetBody has cancelled")
	default:
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	}
}

//!-httpRequestBody

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

func TestMemo(t *testing.T) {
	done := make(chan struct{})
	m := New(httpGetBody, done)
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url, done)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}
func TestMemoCancel(t *testing.T) {
	done := make(chan struct{})
	m := New(httpGetBody, done)
	var n sync.WaitGroup
	go func() {
		time.Sleep(time.Second / 100)
		close(done)
		log.Println("done channel has closed")
	}()
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url, done)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}
