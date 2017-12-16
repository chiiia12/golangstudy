package main

import (
	"flag"
	"fmt"
	"gopl.io/ch4/github"
	"net/http"
	"net/url"
	"io/ioutil"
	"strings"
	"time"
)

var (
	command = flag.String("command", "", "option -command create/show/update")
	token   = flag.String("token", "", "option -token is github generated token")
	title   = flag.String("title", "", "option -title is issue's title")
	//stringのポインタが返ってくる
)

const (
	createIssuesURL = "https://api.github.com/repos/chiiia12/golangstudy"
)

func main() {
	flag.Parse()
	switch *command {
	case "create":
		createIssue()
	case "show":
		github.SearchIssues([]string{"repo:chiiia12/golangstudy", "is:open", "json", "decoder"})
		fmt.Println("show command selected")
	case "update":
		fmt.Println("update command selected")
	}
}

func createIssue() {
	fmt.Println("token is ", *token)
	values := url.Values{}
	values.Add("title", *title)
	values.Add("body", "sample")
	values.Add("label", "bug")
	req, err := http.NewRequest("POST", createIssuesURL, strings.NewReader(values.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer " + *token)

	client := &http.Client{Timeout: time.Duration(10 * time.Second)}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("create is failed", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("status code is not OK.status is %v", resp.Status)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("readAll is err", err)
	}
	fmt.Println("create is success", string(b))
}
