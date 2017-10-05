// Copyright 2017 budougumi0617 All Rights Reserved.

package ex14

import "github.com/budougumi0617/gopl/ch04/ex14/github"

type LocalStore struct {
	Users *github.Users
	Issues *github.Issues
	Milestones *github.Milestones
}

func (ls *LocalStore) Load(url string) err {

}
