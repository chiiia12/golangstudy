package main

import (
	"os"
	"strings"
)

func main() {
	var args = os.Args
	argsBuilder1(args)
	argsBuilder2(args)
}

func argsBuilder1(args []string) string {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func argsBuilder2(args []string) string {
	return strings.Join(args[0:], " ")
}
