package main

import (
	"bytes"
	"os"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	var tests = []struct {
		args     []string
		expected string
	}{
		{[]string{"isanagram", "foo", "bar"}, "foo and bar are not anagram\n"},
		{[]string{"isanagram", "foo", "foo"}, "foo and foo are not anagram\n"},
		{[]string{"isanagram", "asdf", "fdsa"}, "asdf and fdsa are anagram\n"},
		{[]string{"isanagram", "asdf", "fdsa", "fdsa"}, "Invalid args\n"},
	}

	for _, test := range tests {
		os.Args = test.args
		stdout = new(bytes.Buffer) // captured output
		main()
		got := stdout.(*bytes.Buffer).String()
		if got != test.expected {
			t.Errorf("Result = %q, Expected %q", got, test.expected)
		}
	}
}
