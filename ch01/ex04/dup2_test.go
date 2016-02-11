package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args     []string
		input    string
		expected string
	}{
		{[]string{"dup2", "test_file1.txt", "test_file2.txt"}, "", "ccc\t2\n\n	test_file1.txt\n\ttest_file1.txt\n"},
	}

	for _, test := range tests {
		os.Args = test.args
		in = bytes.NewBufferString(test.input)
		out = new(bytes.Buffer) // captured output
		main()
		got := out.(*bytes.Buffer).String()
		if strings.Contains(got, test.expected) {
			t.Errorf("Result = %q, Expected %q", got, test.expected)
		}
	}
}
