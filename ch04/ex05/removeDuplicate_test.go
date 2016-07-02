// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	expected := "[test1 sss sss sss a ttt ttt test2]\n[test1 sss a ttt test2]\n"
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if got != expected {
		t.Errorf("Result:%q, Expected:%q", got, expected)
	}
}

func TestRemoveRuplicate(t *testing.T) {
	var tests = []struct {
		input    []string
		expected []string
	}{
		{[]string{"foo", "foo", "foo", "test", "test", "foo"}, []string{"foo", "test", "foo"}},
	}

	for _, test := range tests {
		got := removeDuplicate(test.input)
		for i := range got {
			if got[i] != test.expected[i] {
				t.Errorf("Result:%q, Expected:%q", got, test.expected)
			}
		}
	}
}
