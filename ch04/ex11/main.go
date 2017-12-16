package main

import (
	"flag"
	"fmt"
)

var (
	command = flag.String("command", "", "option -command create/show/update")
	//stringのポインタが返ってくる
)

func main() {
	flag.Parse()
	switch *command {
	case "create":
		fmt.Println("create command selected")
	case "show":
		fmt.Println("show command selected")
	case "update":
		fmt.Println("update command selected")
	}
}
