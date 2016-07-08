// Copyright 2016 budougumi0617 All Rights Reserved.
// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"
)

var stdout io.Writer = os.Stdout // modified during testing

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": DaysAgo}).
	Parse(templ))

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	splitByYears(result)
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	//
	//   req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	//   if err != nil {
	//       return nil, err
	//   }
	//   req.Header.Set(
	//       "Accept", "application/vnd.github.v3.text-match+json")
	//   resp, err := http.DefaultClient.Do(req)

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DaysAgo convert time to days
func DaysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func splitByYears(result *IssuesSearchResult) {
	var month IssuesSearchResult
	var year IssuesSearchResult
	var other IssuesSearchResult
	for _, issue := range result.Items {
		d := DaysAgo(issue.CreatedAt)
		switch {
		case d < 31:
			(month.TotalCount)++
			month.Items = append(month.Items, issue)
		case d < 365:
			(year.TotalCount)++
			year.Items = append(year.Items, issue)
		default:
			(other.TotalCount)++
			other.Items = append(other.Items, issue)
		}
	}
	fmt.Fprint(stdout, "Within a single month:\t")
	if err := report.Execute(stdout, month); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(stdout, "\nWithin a year:\t")
	if err := report.Execute(stdout, year); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(stdout, "\nOver a year ago:\t")
	if err := report.Execute(stdout, other); err != nil {
		log.Fatal(err)
	}
}
