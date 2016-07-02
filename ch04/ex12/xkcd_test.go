// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	DEBUG = true
	var tests = []struct {
		input    *bytes.Buffer
		expected string
	}{
		{bytes.NewBufferString("\\q"), "Exit xkcd"},
		{bytes.NewBufferString("stick\n\\q"), "http://xkcd.com/9/"},
	}
	for _, test := range tests {
		stdout = new(bytes.Buffer) // captured output
		stdin = test.input
		main()
		got := stdout.(*bytes.Buffer).String()
		if !strings.Contains(got, test.expected) {
			t.Errorf("Actual :%q, Expected %q", got, test.expected)
		}
	}
}

func TestGetLastPage(t *testing.T) {
	if result, err := getLastPage(); err != nil || result < 0 {
		t.Errorf("num = %v, err = %v", result, err)
	}
}
