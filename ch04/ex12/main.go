package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"strings"
)

type Comic struct {
	Title      string `json:"title"`
	Url        string `json:"link"`
	Transcript string `json:"transcript"`
}

const filename = "comic.json"

func main() {
	index := []Comic{}
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("readFile error", err)
		index = initialize()
	} else {
		json.Unmarshal(raw, &index)//ポインタを渡さないとlen(index)の値は0になる
	}

	fmt.Println("len(index)is ", len(index))

	//検索
	for _, v := range os.Args[1:] {
		for _, c := range index {
			if strings.Contains(c.Title, v) {
				fmt.Printf("searchword: %v result: title : %v,transcript: %v\n\n", v, c.Title, c.Transcript)
				fmt.Println("============================================================")
			}
		}
	}

}
func initialize() []Comic {
	index := []Comic{}
	//initialize for offline download
	for i := 1; i <= 100; i++ {
		resp, err := http.Get(fmt.Sprintf("https://xkcd.com/%d/info.0.json", i))
		if err != nil {
			continue
		}
		b, err := ioutil.ReadAll(resp.Body)
		comic := new(Comic) //ポインタが返ってくる

		if err := json.Unmarshal(b, comic); err != nil {
			fmt.Println(err)
			continue
		}
		index = append(index, *comic)
	}
	b, _ := json.Marshal(index)
	ioutil.WriteFile(filename, b, os.ModePerm)
	return index
}
