// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	exit := m.Run()
	if exit != 0 {
		//os.Exit(exit)
		fmt.Printf("result:%v", exit)
	}
}

func TestFetch(t *testing.T) {
	var tests = []struct {
		args     []string
		expected string
	}{
		{[]string{"fetch", "https://google.com"}, "https://google.com"},
		{[]string{"fetch", "https://"}, "no Host in request URL"},
	}

	for _, test := range tests {
		os.Args = test.args

		stdout = new(bytes.Buffer) // captured output
		main()
		got := stdout.(*bytes.Buffer).String()
		if !strings.Contains(got, test.expected) {
			t.Errorf("Result = %q, Expected %q", got, test.expected)
		}
	}
}
