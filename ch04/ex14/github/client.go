// Copyright 2017 budougumi0617 All Rights Reserved.

// Package github provides a Go API for the GitHub data.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"context"
	"fmt"
	"net/http"

	"encoding/json"
	"io"
	"runtime"
)

// GitHubAPIURL URL of API
const GitHubAPIURL = "https://api.github.com/repos/"

var userAgent = fmt.Sprintf("MyGoClient (%s)", runtime.Version())

// Client is base struct for GitHub API.
type Client struct {
	URL        string
	HTTPClient *http.Client
}

// NewClient returns new Client
func NewClient(repo string) *Client {
	return &Client{URL: GitHubAPIURL + repo, HTTPClient: http.DefaultClient}
}

// NewRequest returns new request.
func (c *Client) NewRequest(ctx context.Context, method, apiURL string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, apiURL, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}

// DecodeBody parses json.
func DecodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
