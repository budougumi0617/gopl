// Copyright 2017 budougumi0617 All Rights Reserved.

package main

import "github.com/budougumi0617/gopl/ch04/ex14/github"

type LocalStore struct {
	Users      *github.Users
	Issues     *github.Issues
	Milestones *github.Milestones
}

// NewLocalStore returns new instance.
func NewLocalStore() *LocalStore {
	return &LocalStore{&github.Users{}, &github.Issues{}, &github.Milestones{}}
}

// Load get repository data, and GitHub users.
func (ls *LocalStore) Load(url string) error {
	c := github.NewClient(url)

	if err := ls.Users.GetUsers(c); err != nil {
		return err
	}
	if err := ls.Issues.GetIssues(c); err != nil {
		return err
	}
	if err := ls.Milestones.GetMilestones(c); err != nil {
		return err
	}

	return nil
}
