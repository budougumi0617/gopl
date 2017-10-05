// Copyright 2017 budougumi0617 All Rights Reserved.
// Package github provides a Go API for the GitHub issue tracker.
package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
	"io"
)

// IssuesSearchResult search result
// https://developer.github.com/v3/issues/#list-issues-for-a-repository
// https://api.github.com/repos/golang/go/issues
type Issues struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue issue information
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

func (i *Issues) GetIssues(url string) err {
	req, _ := c.newRequest(context.Background(), "GET", GitHubAPIURL+url+"/issues", nil)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status is not ok: %v", resp.StatusCode)
	}

	if err := decodeBody(resp, &(i.Items)); err != nil {
		return err
	}
	return nil
}

func (i *Issues) RenderIssues(w io.Writer) {
	issueList := template.Must(template.New("issueList").Parse(`
	<h1>GitHub Issues</h1>
	<table>
	  <th>Name</th><th>URL</>
	  {{range .Items}}
	  <tr>
	    <td>{{.Number}}</td><td>{{.Title}}</td><td>{{.Title}}</td><td>{{.State}}</td>
	  </tr>
	  {{end}}
	</table>
	`))

	if err := issueList.Execute(w, &i); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}
}