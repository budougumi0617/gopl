// Copyright 2017 budougumi0617 All Rights Reserved.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/budougumi0617/gopl/ch04/ex11/github"
)

var (
	issueNo                                    int
	title, body                                string
	createFlag, closeFlag, editFlag, printFlag bool
	editorFlag                                 bool
)

func main() {

	// Set flags
	flag.IntVar(&issueNo, "number", 0, "issue number")
	flag.IntVar(&issueNo, "n", 0, "issue number")
	flag.StringVar(&title, "title", "", "issue title")
	flag.StringVar(&title, "t", "", "issue title")
	flag.StringVar(&body, "body", "", "issue body")
	flag.StringVar(&body, "b", "", "issue body")
	flag.BoolVar(&editorFlag, "vim", false, "wirte body by vim")
	flag.BoolVar(&createFlag, "create", false, "create an issue")
	flag.BoolVar(&createFlag, "cr", false, "create an issue")
	flag.BoolVar(&closeFlag, "cl", false, "close an issue")
	flag.BoolVar(&closeFlag, "close", false, "close an issue")
	flag.BoolVar(&editFlag, "edit", false, "edit an issue")
	flag.BoolVar(&editFlag, "e", false, "edit an issue")
	flag.BoolVar(&printFlag, "print", false, "print an issue")
	flag.BoolVar(&printFlag, "p", false, "print an issue")
	flag.Parse()

	validateFlag()

	c := github.NewClient(flag.Arg(0))
	c.Query()

	var issue *github.Issue
	var err error
	switch {
	case createFlag:
		if editorFlag {
			body = executeVim()
		}
		issue, err = c.CreateIssue(title, body)
	case closeFlag:
		if issueNo < 1 {
			fmt.Print("Need to set issue number by \"-number\" or \"-n\"\n")
			os.Exit(1)
		}
		issue, err = c.CloseIssue(issueNo)
	case editFlag:
		if issueNo < 1 {
			fmt.Print("Need to set issue number by \"-number\" or \"-n\"\n")
			os.Exit(1)
		}
		if editorFlag {
			body = executeVim()
		}
		issue, err = c.EditIssue(title, body, issueNo)
	case printFlag:
		if issueNo < 1 {
			fmt.Print("Need to set issue number by \"-number\" or \"-n\"\n")
			os.Exit(1)
		}
		issue, err = c.GetIssue(issueNo)
	}
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", issue)
}

func validateFlag() {
	flags := []bool{createFlag, closeFlag, editFlag, printFlag}

	trueCount := 0
	for _, f := range flags {
		if f {
			trueCount++
		}
	}
	if trueCount != 1 {
		fmt.Print("Need to set operation flag only 1.")
		os.Exit(1)
	}
}

func executeVim() string {
	f, err := ioutil.TempFile("", ".tmp")
	if err != nil {
		panic(fmt.Errorf("Cannot create tmp file: %v", err))
	}

	name := f.Name()
	f.Close()

	cmd := exec.Command("vim", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(fmt.Errorf("Cannot execute vim: %v", err))
	}

	b, err := ioutil.ReadFile(name)
	if err != nil {
		panic(fmt.Errorf("Cannot read temp file: %v", err))
	}
	return string(b)
}
