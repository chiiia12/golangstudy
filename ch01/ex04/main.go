package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	duplicatefiles := make(map[string][]string)
	files := os.Args[1:]
	gatherContainsFiles(files, duplicatefiles)
	for line, n := range duplicatefiles {
		fmt.Printf("%s\t%s\n", line, strings.Join(n[0:], " "))
	}
}

func gatherContainsFiles(files []string, duplicatefiles map[string][]string) {
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			duplicatefiles[input.Text()] = append(duplicatefiles[input.Text()], arg)
		}
		f.Close()
	}
}
