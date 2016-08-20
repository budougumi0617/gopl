// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"testing"
)

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"t", "rune\tcount\n't'\t1\n\nlen\tcount\n1\t1\n2\t0\n3\t0\n4\t0\n"},
		{"f4", "rune\tcount\n'4'\t1\n'f'\t1\n\nlen\tcount\n1\t2\n2\t0\n3\t0\n4\t0\n"},
	}

	for _, test := range tests {
		stdin = bytes.NewBufferString(test.input)
		stdout = new(bytes.Buffer) // captured output
		main()
		got := stdout.(*bytes.Buffer).String()

		if got != test.expected {
			t.Errorf("Result = \n%q\n\nExpected = \n%q", got, test.expected)
		}
	}
}
