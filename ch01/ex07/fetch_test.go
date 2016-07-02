// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"fmt"
	"os"
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
		args []string
	}{
		{[]string{"fetch", "http://github.com/budougumi0617"}},
	}

	for _, test := range tests {
		os.Args = test.args
		stdout = new(bytes.Buffer) // captured output
		stderr = stdout
		main()
		got := stdout.(*bytes.Buffer).String()
		if len(got) == 0 {
			t.Errorf("Result = %q", got)
		}
	}
}
