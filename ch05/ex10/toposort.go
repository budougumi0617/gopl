package main

import (
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout // modified during testing

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]map[string]bool{
	"algorithms": map[string]bool{"data structures": false},
	"calculus":   map[string]bool{"linear algebra": false},

	"compilers": map[string]bool{
		"data structures":       false,
		"formal languages":      false,
		"computer organization": false,
	},

	"data structures":       map[string]bool{"discrete math": false},
	"databases":             map[string]bool{"data structures": false},
	"discrete math":         map[string]bool{"intro to programming": false},
	"formal languages":      map[string]bool{"discrete math": false},
	"networks":              map[string]bool{"operating systems": false},
	"operating systems":     map[string]bool{"data structures": false, "computer organization": false},
	"programming languages": map[string]bool{"data structures": false, "computer organization": false},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Fprintf(stdout, "%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}

	// sort.Strings(keys) // Eliminate the initial sort.
	visitAll(keys)
	return order
}
