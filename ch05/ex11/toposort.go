// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var stdout io.Writer = os.Stdout // modified during testing

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"linear algebra": {"calculus"}, // Add new prerequisit
	"calculus":       {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Fprintf(stdout, "%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(string, []string)

	visitAll = func(corse string, items []string) {
		for _, item := range items {
			//fmt.Printf("Check %q prerequest %q\n", corse, m[item])
			for _, prerq := range m[item] {
				if corse == prerq {
					s := "Found Cycles between \"" + item + "\" and \"" + corse + "\""
					order = append(order, s)
				}
			}

			if !seen[item] {
				seen[item] = true
				visitAll(item, m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll("", keys)
	return order
}
