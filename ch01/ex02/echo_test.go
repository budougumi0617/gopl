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
		{[]string{"echo"}, "[0] echo\n"},
		{[]string{"echo", "one", "two", "three"}, "[0] echo\n[1] one\n[2] two\n[3] three\n"},
		{[]string{"echo", "a", "b", "c"}, "[0] echo\n[1] a\n[2] b\n[3] c\n"},
		{[]string{"echo", "1", "2", "3"}, "[0] echo\n[1] 1\n[2] 2\n[3] 3\n"},
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
