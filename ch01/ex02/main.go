package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var args = os.Args
	fmt.Println(echoArgs(args))
}

func echoArgs(args []string) string {
	s := ""
	for index, val := range args[0:] {
		s += strings.Join([]string{strconv.Itoa(index), val}, ":")
		s += "\n"
	}
	return s
}
