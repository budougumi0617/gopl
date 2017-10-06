// Copyright 2017 budougumi0617 All Rights Reserved.
// Package github provides a Go API for the GitHub issue tracker.
package github

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

// Issue issue information
// https://developer.github.com/v3/issues/#list-issues-for-a-repository
// https://api.github.com/repos/golang/go/issues
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// Issues manages milestone data for GitHub repository.
type Issues struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// GetIssues gets issues from web.
func (i *Issues) GetIssues(c *Client) error {
	req, _ := c.NewRequest(context.Background(), "GET", c.URL+"/issues", nil)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Get issues from %v\n", c.URL)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status is not ok: %v", resp.StatusCode)
	}

	if err := DecodeBody(resp, &(i.Items)); err != nil {
		return err
	}
	return nil
}

// RenderIssues writes HTML table.
func (i *Issues) RenderIssues(w io.Writer) {
	issueList := template.Must(template.New("issueList").Parse(`
	<h1>GitHub Issues</h1>
	<table>
	  <tr><th>No.</th><th>Title</th><th>Status</th></tr>
	  {{range .Items}}
	  <tr>
	    <td>{{.Number}}</td><td>{{.Title}}</td><td>{{.State}}</td>
	  </tr>
	  {{end}}
	</table>
	`))

	if err := issueList.Execute(w, &i); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}
}
