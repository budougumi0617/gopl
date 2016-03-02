package main

import (
	"bytes"
	"os"
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		args     []string
		expected string
	}{
		{[]string{"comma", "1", "12", "123", "1234", "1234567890"}, "  1\n  12\n  123\n  1,234\n  1,234,567,890\n"},
		{[]string{"comma", "123456"}, "  123,456\n"},
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
