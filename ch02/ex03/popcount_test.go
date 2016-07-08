// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args = []string{"popcount", "0x1234567890ABCDEF"}
	stdout = new(bytes.Buffer) // captured output
	expected := "result 32\n"
	main()
	got := stdout.(*bytes.Buffer).String()
	if got != expected {
		t.Errorf("Result = %q, Expected %q", got, expected)
	}
}

func TestPopCountByLoop(t *testing.T) {
	var tests = []struct {
		args     uint64
		expected int
	}{
		{0, 0},
		{1, 1},
		{10, 2},
	}

	for _, test := range tests {
		original := PopCount(test.args)
		got := PopCountByLoop(test.args)
		if original != got || got != test.expected {
			t.Errorf("Result = %v, Expected %v", got, test.expected)
		}
	}
}
