package main

import (
	"os"
	"fmt"
	"net/http"
	"io"
	"golang.org/x/net/html"
	"strings"
)

const (
	ContentType = "Content-Type"
	TextHTML    = "text/htm"
)

func main() {
	if err := Extract(os.Args[1]); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func Extract(url string) interface{} {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	host := resp.Request.URL.Host
	path := resp.Request.URL.Path
	scheme := resp.Request.URL.Scheme
	hostURL := scheme + "://" + host
	fmt.Printf("scheme=%s,host=%s,path=%s\n", scheme, host, path)
	fmt.Printf("hostURL=%s\n", hostURL)

	if err := os.Mkdir(host, os.ModePerm); err != nil {
		if os.IsExist(err) {
			fmt.Printf("%s directory exists! Plese delete it\n", host)
			return nil
		}
		fmt.Printf("os.Mkdir :%v\n", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s:%s", url, resp.Status)
	}
	return extractByType(resp, url, host+"/root.html", host, hostURL)
}
func extractByType(resp *http.Response, url, path, host, hostURL string) error {
	contentType := extractContentType(resp.Header)
	if contentType[0] != TextHTML {
		f, err := os.Create(path)
		if err != nil {
			return nil
		}
		defer f.Close()
		io.Copy(f, resp.Body)
		resp.Body.Close()
		f.Close()
		return nil
	}
	_, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML:%v", url, err)
	}
	return nil
}
func extractContentType(header http.Header) []string {
	contentType, ok := header[ContentType]
	if !ok {
		return nil
	}
	return strings.Split(contentType[0], ";")

}
