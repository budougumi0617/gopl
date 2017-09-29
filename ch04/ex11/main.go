// Copyright 2017 budougumi0617 All Rights Reserved.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/budougumi0617/gopl/ch04/ex11/github"
)

func main() {
	var (
		issueNo                                    int
		title, body                                string
		createFlag, closeFlag, editFlag, printFlag bool
	)
	flag.IntVar(&issueNo, "number", 0, "issue number")
	flag.IntVar(&issueNo, "n", 0, "issue number")
	flag.StringVar(&title, "title", "", "issue title")
	flag.StringVar(&title, "t", "", "issue title")
	flag.StringVar(&body, "body", "", "issue body")
	flag.StringVar(&body, "b", "", "issue body")
	flag.BoolVar(&createFlag, "create", false, "create an issue")
	flag.BoolVar(&createFlag, "cr", false, "create an issue")
	flag.BoolVar(&closeFlag, "cl", false, "close an issue")
	flag.BoolVar(&closeFlag, "close", false, "close an issue")
	flag.BoolVar(&editFlag, "edit", false, "edit an issue")
	flag.BoolVar(&editFlag, "e", false, "edit an issue")
	flag.BoolVar(&printFlag, "print", false, "print an issue")
	flag.BoolVar(&printFlag, "p", false, "print an issue")
	flag.Parse()

	c := github.NewClient()
	c.Query()

	switch {
	case createFlag:
		b := body
		fmt.Printf("body = %v\n", b)
	case closeFlag:
	case editFlag:
	case printFlag:
		if issueNo < 1 {
			fmt.Print("Need to set issue number by \"-number\" or \"-n\"\n")
			os.Exit(1)
		}
		issue, err := c.GetIssue(issueNo)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%v", issue)
	}
}

func define(name string) (defined bool) {
	flag.VisitAll(func(f *flag.Flag) {
		if f.Name == name {
			defined = true
		}
	})
	return
}
