package main

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestSearchIssues(t *testing.T) {
	terms := []string{"budougumi0617/gotraining"}
	if _, err := SearchIssues(terms); err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestSplitByYears(t *testing.T) {
	var tests = []struct {
		input     *IssuesSearchResult
		expected1 string
		expected2 string
		expected3 string
	}{
		{new(IssuesSearchResult), "Within a single month:\t1 issues:\n----------------------------------------\nNumber: 2", "Within a year:\t0 issues:",
			"Over a year ago:\t1 issues:\n----------------------------------------\nNumber: 1"},
	}
	isr := tests[0].input
	isr.TotalCount = 3
	isr.Items = append(isr.Items, &Issue{Number: 1, Title: "test1", User: &User{Login: "foo"}, CreatedAt: time.Date(2001, 5, 20, 23, 59, 59, 0, time.UTC)})
	isr.Items = append(isr.Items, &Issue{Number: 2, Title: "test2", User: &User{Login: "bar"}, CreatedAt: time.Now()})
	for _, test := range tests {
		stdout = new(bytes.Buffer) // captured output
		splitByYears(test.input)
		got := stdout.(*bytes.Buffer).String()
		if !strings.Contains(got, test.expected1) {
			t.Errorf("Result:%q, Expected:%q", got, test.expected1)
		} else if !strings.Contains(got, test.expected2) {
			t.Errorf("Result:%q, Expected:%q", got, test.expected2)
		} else if !strings.Contains(got, test.expected2) {
			t.Errorf("Result:%q, Expected:%q", got, test.expected3)
		}

	}
}
