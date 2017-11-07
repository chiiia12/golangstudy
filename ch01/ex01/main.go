package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var args = os.Args
	fmt.Println(echoArgs(args))
}

func echoArgs(args []string) string {
	return strings.Join(args[0:], " ")
}
