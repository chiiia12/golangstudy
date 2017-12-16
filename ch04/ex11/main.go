package main

import (
	"flag"
	"fmt"
	"gopl.io/ch4/github"
	"net/http"
	"net/url"
	"os"
	"io/ioutil"
)

var (
	command = flag.String("command", "", "option -command create/show/update")
	//stringのポインタが返ってくる
)

const (
	createIssuesURL = "https://api.github.com/repos/chiiia12/golangstudy"
)

func main() {
	flag.Parse()
	switch *command {
	case "create":
		createIssue(os.Args[2])
	case "show":
		github.SearchIssues([]string{"repo:chiiia12/golangstudy", "is:open", "json", "decoder"})
		fmt.Println("show command selected")
	case "update":
		fmt.Println("update command selected")
	}
}

func createIssue(title string) {
	//TODO:headerにtokenとcontenttypeつける
	//TODO:respをmappingするstructつくる
	values := url.Values{}
	values.Add("title", title)
	resp, err := http.PostForm(createIssuesURL, values)
	fmt.Println("postForm")
	if err != nil {
		fmt.Println("create is failed", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Errorf("status code is not OK.status is %v", resp.Status)
	}
	b, err := ioutil.ReadAll(resp.Body)
	fmt.Println("create is success", b)
}
