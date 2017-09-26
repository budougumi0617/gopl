// Copyright 2017 budougumi0617 All Rights Reserved.

package main

import (
	"github.com/budougumi0617/gopl/ch04/ex11/github"
)

func main() {
	c := github.NewClient()
	c.Query()

	c.CreateIssue("Client test", "Created from CLI.")
}
