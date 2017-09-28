// Copyright 2017 budougumi0617 All Rights Reserved.

package main

import (
	"flag"

	"fmt"
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

	issue, err := c.GetIssue()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", issue)
}
