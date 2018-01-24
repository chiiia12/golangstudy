package main

import (
	"fmt"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]map[string]int{
	"algorithms": {"data structures": 1},
	"calculus":   {"linear algebra": 1},

	"compilers": {
		"data structures":       1,
		"formal languages":      1,
		"computer organization": 1,
	},

	"data structures":       {"discrete math": 1},
	"databases":             {"data structures": 1},
	"discrete math":         {"intro to programming": 1},
	"formal languages":      {"discrete math": 1},
	"networks":              {"operating systems": 1},
	"operating systems":     {"data structures": 1, "computer organization": 1},
	"programming languages": {"data structures": 1, "computer organization": 1},
}

//!-table

//!+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]int) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]int)

	visitAll = func(items map[string]int) {
		for val, _ := range items {
			if !seen[val] {
				seen[val] = true
				visitAll(m[val])
				order = append(order, val)
			}
		}
	}

	keys := make(map[string]int)
	for key := range m {
		keys[key] = 1
	}

	//sort.Strings(keys)
	visitAll(keys)
	return order
}

//!-main
