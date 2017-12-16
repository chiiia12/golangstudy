package main

import (
	"fmt"
	"log"
	"os"
	"gopl.io/ch4/github"
	"time"
)

//!+
//TODO:名前なんとかする
//TODO:キーごとに手動でfor書かなくてもよくできないか？

func main() {
	now := time.Now()
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	groupMap := make(map[string][]github.Issue)
	for _, item := range result.Items {
		if item.CreatedAt.After(now.AddDate(0, -1, 0)) {
			groupMap["MonthAfter"] = append(groupMap["MonthAfter"], *item)
		} else if item.CreatedAt.After(now.AddDate(-1, 0, 0)) {
			groupMap["YearAfter"] = append(groupMap["YearAfter"], *item)
		} else {
			groupMap["YearBefore"] = append(groupMap["YearBefore"], *item)
		}

	}
	fmt.Println("1ヶ月未満")
	for _, v := range groupMap["MonthAfter"] {
		fmt.Printf("%v #%-5d %9.9s %,55s \n", v.CreatedAt, v.Number, v.User.Login, v.Title)
	}
	fmt.Println("1年未満")
	for _, v := range groupMap["YearAfter"] {
		fmt.Printf("%v #%-5d %9.9s %,55s \n", v.CreatedAt, v.Number, v.User.Login, v.Title)
	}
	fmt.Println("1年以上")
	for _, v := range groupMap["YearBefore"] {
		fmt.Printf("%v #%-5d %9.9s %,55s \n", v.CreatedAt, v.Number, v.User.Login, v.Title)
	}

}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
