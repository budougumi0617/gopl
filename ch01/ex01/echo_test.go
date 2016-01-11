package main

import (
	"bytes"
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args     []string
		expected string
	}{
		{[]string{"echo"}, "echo\n"},
		{[]string{"echo", "one", "two", "three"}, "echo one two three\n"},
		{[]string{"echo", "a", "b", "c"}, "echo a b c\n"},
		{[]string{"echo", "1", "2", "3"}, "echo 1 2 3\n"},
	}

	for _, test := range tests {
		os.Args = test.args
		out = new(bytes.Buffer) // captured output
		main()
		got := out.(*bytes.Buffer).String()
		if got != test.expected {
			t.Errorf("Result = %q, Expected %q", got, test.expected)
		}
	}
}
