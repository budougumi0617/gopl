package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		input    []string
		expected string
	}{
		{[]string{"wordfreq"}, "Execute with the name of text file\n"},
		{[]string{"wordfreq", "no_file"}, "Failed file open\n"},
		{[]string{"wordfreq", "foo.txt"}, "\"4time\" was found 4 times\n"},
	}

	for _, test := range tests {
		os.Args = test.input
		stdout = new(bytes.Buffer) // captured output
		main()
		got := stdout.(*bytes.Buffer).String()

		if !strings.Contains(got, test.expected) {
			t.Errorf("Result:%q, Expected:%q", got, test.expected)
		}

	}
}
