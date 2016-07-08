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
		{[]string{"fetch", "https://github.com/budougumi0617"}, "resp.status 200 OK\n"},
		{[]string{"fetch", "https://github.com/budougumi0617/GoTraining/404status"}, "resp.status 404 Not Found\n"},
		//{[]string{"fetch", "https://"}, "\n"},
	}

	for _, test := range tests {
		os.Args = test.args
		stdout = new(bytes.Buffer) // captured output
		stderr = stdout
		main()
		got := stdout.(*bytes.Buffer).String()
		if !strings.Contains(got, test.expected) {
			t.Errorf("Result = %q, Expected %q", got, test.expected)
		}
	}
}
