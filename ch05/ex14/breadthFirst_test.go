package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestBreadthFirst(t *testing.T) {
	var tests = []struct {
		corse    string
		expected []string
	}{
		{"discrete math", []string{"intro to programming"}},
		{"programming languages", []string{"data structures", "computer organization", "discrete math", "intro to programming"}},
	}

	for _, test := range tests {
		stdout = new(bytes.Buffer) // captured output
		breadthFirst(depend, prereqs[test.corse])
		got := stdout.(*bytes.Buffer).String()
		for _, c := range test.expected {
			if !strings.Contains(got, c) {
				t.Errorf("Result = %q, Expected %q", got, c)
			}
		}
	}

}

func TestMain(t *testing.T) {
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if len(got) == 0 {
		t.Errorf("No output")
	}
}
