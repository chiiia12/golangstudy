package main

import (
	"flag"
	"fmt"
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
	createIssuesURL = "https://api.github.com/repos/chiiia12/golangstudy/issues/"
)

type Param struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	State string `json:"state,omitempty"`
	Name string `json:"name,omitempty"`
}

func main() {
	flag.Parse()
	switch *command {
	case "create":
		createIssue(Param{Title:*title,Body:*body})
	case "show":
		showIssue()
	case "update":
		updateIssue(Param{Title: *title, Body: *body})
	case "close":
		updateIssue(Param{State: "close"})
	}
}
func showIssue() {
	resp ,err:=http.Get("https://api.github.com/repos/chiiia12/golangstudy/issues")
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
func updateIssue(param Param) {
	fmt.Println("token is ", *token)
	strJson, _ := json.Marshal(param)
	fmt.Println(param,string(strJson))

	req, err := http.NewRequest("PATCH", createIssuesURL+ *issue, bytes.NewBuffer(strJson))
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

func createIssue(param Param) {
	fmt.Println("token is ", *token)
	strJson, _ := json.Marshal(param)

	req, err := http.NewRequest("POST", createIssuesURL, bytes.NewBuffer(strJson))
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
