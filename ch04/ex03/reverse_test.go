// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	expected := "[0 1 2 3 4 5]\n[5 4 3 2 1 0]\n"
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if got != expected {
		t.Errorf("Result:%q, Expected:%q", got, expected)
	}
}
