package main

import (
	"flag"
	"fmt"
	"gopl.io/ch4/github"
	"net/http"
	"io/ioutil"
	"time"
	"bytes"
)

var (
	command = flag.String("command", "", "option -command create/show/update")
	token   = flag.String("token", "", "option -token is github generated token")
	title   = flag.String("title", "", "option -title is issue's title")
	issue   = flag.String("issue", "", "option -issue is number of issue ")
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
	case "update":
		fmt.Println("update command selected")
	case "close":
		fmt.Println("close command selected")
		closeIssue()
	}
}
func closeIssue() {
	fmt.Println("token is ", *token)
	strJson := `{
		"state":"close"
	}`

	req, err := http.NewRequest("PATCH", createIssuesURL + "/issues/" + *issue, bytes.NewBuffer([]byte(strJson)))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
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
	fmt.Println("close is success", string(b))

}

func createIssue() {
	fmt.Println("token is ", *token)
	strJson := `{
	"title": "Found a bug",
		"body": "I'm having a problem with this.",
		"name":"repo name",
		"labels": [
		]
	}`

	req, err := http.NewRequest("POST", createIssuesURL, bytes.NewBuffer([]byte(strJson)))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
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
