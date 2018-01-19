package main

import (
	"net/http"
	"os"
	"path"
	"io"
	"fmt"
)

func main() {
	fmt.Println(fetch("https://golang.org/doc"))
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "./index.html"
	}
	fmt.Println(local)
	f, err := os.Create(local)
	if err != nil {
		fmt.Println(err)
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	defer func() {
		closeErr := f.Close()
		fmt.Println("close")
		if err == nil {
			err = closeErr
		}
	}()
	return local, n, err
}
