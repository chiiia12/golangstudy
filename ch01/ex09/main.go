package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	urlPrefix = "http://"
)

func main() {
	url := os.Args[1:]
	fetchUrl(os.Stdout, url)
}

func fetchUrl(w io.Writer, urllist []string) {
	for _, url := range urllist {
		if !strings.HasPrefix(url, urlPrefix) {
			url = urlPrefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(w, resp.Body)
		fmt.Fprintln(w, resp.Status)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}

}
