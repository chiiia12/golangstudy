package main

import (
	"os/exec"
	"log"
	"fmt"
	"flag"
	"encoding/json"
)

type ImportData struct {
	//Dir         string
	//ImportPath  string
	//Name        string
	//Stale       bool
	//StaleReason string
	//GoFiles     []string
	//Imports     []string
	Deps []string
}

var pac = flag.String("package", "", "input package name. ex) -package hash")

func main() {
	out, err := exec.Command("go", "list", "-json", *pac).Output()
	if err != nil {
		log.Println(err)
	}
	//TODO:setにappendしてかぶりない方がよさそう。
	//TODO:関数に切り出すなどしたほうがよさそう
	data := new(ImportData)
	json.Unmarshal(out, data)
	for _, d := range data.Deps {
		fmt.Printf("%v\n", d)
	}
	for _, d := range data.Deps {
		out, err := exec.Command("go", "list", "-json", d).Output()
		if err != nil {
			log.Println(err)
		}
		data := new(ImportData)
		json.Unmarshal(out, data)
		for _, e := range data.Deps {
			fmt.Printf("%v\n", e)
		}

	}
}
