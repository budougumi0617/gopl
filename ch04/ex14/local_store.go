// Copyright 2017 budougumi0617 All Rights Reserved.

package main

import "github.com/budougumi0617/gopl/ch04/ex14/github"

type LocalStore struct {
	github.Users
	github.Issues
	github.Milestones
}

// NewLocalStore returns new instance.
func NewLocalStore() *LocalStore {
	return &LocalStore{}
}

// Load get repository data, and GitHub users.
func (ls *LocalStore) Load(url string) error {
	c := github.NewClient(url)

	if err := ls.GetUsers(c); err != nil {
		return err
	}
	if err := ls.GetIssues(c); err != nil {
		return err
	}
	if err := ls.GetMilestones(c); err != nil {
		return err
	}

	return nil
}
