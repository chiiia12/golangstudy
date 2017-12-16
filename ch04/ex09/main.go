package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {

	files := os.Args[1:]
	countWords(files)
}

func countWords(files []string) map[string]int {
	wordCount := make(map[string]int)
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err is %v", err)
			continue
		}
		input := bufio.NewScanner(f)
		input.Split(bufio.ScanWords)
		for input.Scan() {
			wordCount[input.Text()]++
		}
	}
	return wordCount
}
