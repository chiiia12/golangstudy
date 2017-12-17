package main

import (
	"flag"
	"fmt"
	"gopl.io/ch4/github"
	"net/http"
	"io/ioutil"
	"time"
	"bytes"
	"encoding/json"
)

var (
	command = flag.String("command", "", "option -command create/show/update")
	token   = flag.String("token", "", "option -token is github generated token")
	title   = flag.String("title", "", "option -title is issue's title")
	issue   = flag.String("issue", "", "option -issue is number of issue ")
	body    = flag.String("body", "", "option -body is body of issue for update")
	//stringのポインタが返ってくる
)

const (
	createIssuesURL = "https://api.github.com/repos/chiiia12/golangstudy"
)

type UpdateParam struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	State string `json:"state,omitempty"`
}

func main() {
	flag.Parse()
	switch *command {
	case "create":
		createIssue()
	case "show":
		github.SearchIssues([]string{"repo:chiiia12/golangstudy", "is:open", "json", "decoder"})
	case "update":
		updateIssue(UpdateParam{Title: *title, Body: *body})
	case "close":
		updateIssue(UpdateParam{State: "close"})
	}
}
func updateIssue(param UpdateParam) {
	fmt.Println("token is ", *token)
	strJson, _ := json.Marshal(param)
	fmt.Println(param,string(strJson))

	req, err := http.NewRequest("PATCH", createIssuesURL + "/issues/" + *issue, bytes.NewBuffer(strJson))
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
