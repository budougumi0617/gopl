// Copyright 2017 budougumi0617 All Rights Reserved.

// Package github is GitHub API.
package github

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"syscall"

	"encoding/json"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"runtime"
)

var userAgent = fmt.Sprintf("MyGoClient (%s)", runtime.Version())

// Client is base struct for GitHub API.
type Client struct {
	URL        string
	HTTPClient *http.Client

	Username, Password string
}

func NewClient() *Client {
	return &Client{URL: GitHubAPIURL + "budougumi0617/gopl/", HTTPClient: http.DefaultClient}
}

// Query sets username and password.
func (c *Client) Query() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	c.Username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	bytesPassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("\nPassword error %v\n", err)
	}
	c.Password = string(bytesPassword)
	fmt.Print("\n")
}

func (c *Client) newRequest(ctx context.Context, method, apiUrl string, body io.Reader) (*http.Request, error) {
	url := c.URL + apiUrl
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func (c *Client) GetIssues() {
	req, _ := c.newRequest(context.Background(), "GET", "issues/15", nil)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Errorf("Do failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("create failed: %s", resp.Status)
	}

	var issue Issue
	if err := decodeBody(resp, &issue); err != nil {
		fmt.Errorf("create failed: %v", err)
	}
	fmt.Printf("issue %v\n", issue)
}
