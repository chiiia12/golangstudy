package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os"
)

//!-table

//!+main
func main() {
	breadthFirst(visit, os.Args[1:])
}

func visit(item string) []string {
	files, err := ioutil.ReadDir(item)
	if err != nil {
		return nil
	}
	fmt.Println(item)

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, visit(filepath.Join(item, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(item, file.Name()))
	}
	return paths
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-main
