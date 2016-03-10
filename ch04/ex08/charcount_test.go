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
		{"tEst 13 差\tƐ\n", "rune\tcount\n'\\t'\t1\n'\\n'\t1\n' '\t2\n'1'\t1\n'3'\t1\n'E'\t1\n's'\t1\n't'\t2\n'Ɛ'\t1\n'差'\t1\n\nlen\tcount\n1\t10\n2\t1\n3\t1\n4\t0\n\nCategory\tcount\nControl\t2\nDigit\t2\nGraphic\t10\nLetter\t6\nLower\t3\nNotPrintable\t2\nNumber\t2\nPrintable\t10\nSpace\t4\nUpper\t2\n"},
	}

	for _, test := range tests {
		stdin = bytes.NewBufferString(test.input)
		stdout = new(bytes.Buffer) // captured output
		main()
		got := stdout.(*bytes.Buffer).String()

		if got != test.expected {
			t.Errorf("Result = %q, Expected %q", got, test.expected)
		}
	}
}
