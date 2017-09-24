// Copyright 2017 budougumi0617 All Rights Reserved.

// Package github is GitHub API.
package github

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// Client is base struct for GitHub API.
type Client struct {
	URL        *url.URL
	HTTPClient *http.Client

	Username, Password string
}

// Query sets username and passward.
func (c *Cient) Query() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	c.Username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	bytesPassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("\nPassword error %v\n", err)
	}
	c.Password = bytePassword
}
