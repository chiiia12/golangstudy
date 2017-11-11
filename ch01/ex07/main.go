package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1:]
	fetchUrl(os.Stdout, url)
}

func fetchUrl(w io.Writer, urllist []string) {
	for _, url := range urllist {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		_, errcopy := io.Copy(w, resp.Body)
		resp.Body.Close()
		if errcopy != nil {
			fmt.Fprintf(os.Stderr, "fetch:reading %s: %v\n", url, errcopy)
			os.Exit(1)
		}
	}

}
