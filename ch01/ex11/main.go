package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	file = "output.txt"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	args := os.Args[1:]
	outputString := ""
	for _, url := range args {
		go fetch(url, ch)
	}
	for range args {
		outputString += <-ch + "\n"
	}
	for _, url := range args {
		go fetch(url, ch)
	}
	for range args {
		outputString += <-ch + "\n"
	}
	ioutil.WriteFile(file, []byte(outputString), os.ModePerm)
	fmt.Printf(outputString)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
