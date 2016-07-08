// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args = []string{"crawl", "https://raw.githubusercontent.com/budougumi0617/GoTraining/master/LICENSE"}
	stdout = new(bytes.Buffer) // captured output
	main()

	got := stdout.(*bytes.Buffer).String()
	if len(got) == 0 {
		t.Errorf("no output")
	}

}
