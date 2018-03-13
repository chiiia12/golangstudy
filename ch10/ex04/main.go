package main

import (
	"os/exec"
	"log"
	"flag"
	"encoding/json"
	"fmt"
)

type ImportData struct {
	Deps []string
}

var pac = flag.String("package", "", "input package name. ex) -package hash")

var dependencyMap = make(map[string]struct{})

func main() {
	getGoList(*pac)
	for k, _ := range dependencyMap {
		getGoList(k)
	}
	for k, _ := range dependencyMap {
		fmt.Printf("%v\n", k)
	}
}
func getGoList(pacName string) {
	out, err := exec.Command("go", "list", "-json", pacName).Output()
	if err != nil {
		log.Println(err)
	}
	data := new(ImportData)
	json.Unmarshal(out, data)
	for _, d := range data.Deps {
		dependencyMap[d] = struct{}{}
	}
}
