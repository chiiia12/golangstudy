package main

import (
	"net/http"
	"log"
	"html/template"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Data struct {
	IssueList []Issue
}
type Issue struct {
	Url       string
	Number    int
	Title     string
	User      User
	Milestone Milestone
}
type User struct {
	Login string
	Url   string
}
type Milestone struct {
	Url    string
	Title  string
	Number int
}
type IssueDetail struct {
	Number int
	Title  string
	Body   string
	State  string
}
type UserDetail struct {
	Login    string
	Id       int
	Name     string
	Company  string
	Blog     string
	Location string
	Email    string
}
type MilestoneDetail struct {
	Id          int
	Number      int
	Title       string
	Description string
	State       string
}

var issueDetailMap = make(map[string]IssueDetail)
var userDetailMap = make(map[string]UserDetail)
var milestoneDetailMap = make(map[string]MilestoneDetail)
var issueList []Issue
var tmpl = template.Must(template.New("issuelist").Parse(`
<table>
<tr style='text-align: left'>
  <th>Number</th>
  <th>User</th>
  <th>Milestones</th>
  <th>Title</th>
</tr>
{{range .IssueList}}
<tr>
  <td><a href='/issue?number={{.Number}}'>{{.Number}}</a></td>
  <td><a href='/user?number={{.User.Login}}'>{{.User.Login}}</a></td>
  <td><a href='/milestone?number={{.Milestone.Number}}'>{{.Milestone.Title}}</a></td>
  <td><a href=''>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/issue", issueHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/milestone", milestoneHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func milestoneHandler(w http.ResponseWriter, r *http.Request) {
	initialize()
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	number := r.Form["number"][0]
	fmt.Fprint(w, number)
	fmt.Fprint(w, milestoneDetailMap[number])

}
func userHandler(w http.ResponseWriter, r *http.Request) {
	initialize()
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	number := r.Form["number"][0]
	fmt.Fprint(w, number)
	fmt.Fprint(w, userDetailMap[number])
}

func issueHandler(w http.ResponseWriter, r *http.Request) {
	initialize()
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	number := r.Form["number"][0]
	fmt.Fprint(w, number)
	fmt.Fprint(w, issueDetailMap[number])
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	initialize()
	var data = *new(Data)
	data.IssueList = issueList
	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal(err)
	}

}

func initialize() {
	if len(issueList) > 0 {
		return
	}
	resp, err := http.Get("https://api.github.com/repos/golang/go/issues")
	//場合によってclient_id/client_secretを付与する
	if err != nil {
		fmt.Println("create is failed", err)
		return
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&issueList); err != nil {
		fmt.Println("issueList newDecoder ", err)
	}
	//データ作成
	makeDetailMap(issueList)

}

func makeDetailMap(issueList []Issue) {
	for _, v := range issueList {
		//issueDetailを詰める
		resp, err := http.Get(v.Url)
		//場合によってclient_id/client_secretを付与する
		if err != nil {
			fmt.Println("makeIssueDetailMap1", err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		var detail = new(IssueDetail)
		json.Unmarshal(b, detail)
		issueDetailMap[strconv.Itoa(detail.Number)] = *detail

		//userDetailを詰める
		resp, err = http.Get(v.User.Url)
		//場合によってclient_id/client_secretを付与する
		if err != nil {
			fmt.Println(err)
		}
		b, err = ioutil.ReadAll(resp.Body)
		var detailUser = new(UserDetail)
		json.Unmarshal(b, detailUser)
		userDetailMap[detailUser.Login] = *detailUser
		fmt.Println(detailUser.Login, userDetailMap[detailUser.Login].Name)

		//milestoneDetailを詰める
		if v.Milestone.Url == "" {
			continue
		}
		resp, err = http.Get(v.Milestone.Url)
		//場合によってclient_id/client_secretを付与する
		if err != nil {
			fmt.Println(err)
		}
		b, err = ioutil.ReadAll(resp.Body)
		var detailMilestone = new(MilestoneDetail)
		json.Unmarshal(b, detailMilestone)
		milestoneDetailMap[strconv.Itoa(detailMilestone.Number)] = *detailMilestone
	}
	fmt.Printf("issueDetailMap length is %v\n", len(issueDetailMap))
	fmt.Printf("userDetailMap length is %v\n", len(userDetailMap))
}
