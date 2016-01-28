package main

import (
	"bytes"
	"os"
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	input := make([]string, 1000)
	for i := range input {
		if (i % 2) == 0 {
			input[i] = "foo"
		} else {
			input[i] = "bar"
		}

	}
	for i := 0; i < b.N; i++ {
		echo1(input)
	}
}

func BenchmarkEcho2(b *testing.B) {
	input := make([]string, 1000)
	for i := range input {
		if (i % 2) == 0 {
			input[i] = "foo"
		} else {
			input[i] = "bar"
		}

	}
	for i := 0; i < b.N; i++ {
		echo2(input)
	}
}

func BenchmarkEcho3(b *testing.B) {
	input := make([]string, 1000)
	for i := range input {
		if (i % 2) == 0 {
			input[i] = "foo"
		} else {
			input[i] = "bar"
		}

	}
	for i := 0; i < b.N; i++ {
		echo3(input)
	}
}

func BenchmarkEcho4(b *testing.B) {
	input := make([]string, 1000)
	for i := range input {
		if (i % 2) == 0 {
			input[i] = "foo"
		} else {
			input[i] = "bar"
		}

	}
	for i := 0; i < b.N; i++ {
		echo4(input)
	}
}

func TestEchos(t *testing.T) {
	var tests = []struct {
		args     []string
		expected string
	}{
		{[]string{"echo"}, "\n\n\n[]\n"},
		{[]string{"echo", "foo", "bar"}, "foo bar\nfoo bar\nfoo bar\n[foo bar]\n"},
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
