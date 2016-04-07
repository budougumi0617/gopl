package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCorner(t *testing.T) {
	var tests = []struct {
		input, expected string
		f               func(string) string
	}{
		{"$The $Go Programming $Language", "THE GO Programming LANGUAGE", strings.ToUpper},
		{"$The $Go Programming $Language", "the go Programming language", strings.ToLower},
	}

	for _, test := range tests {
		got := expand(test.input, test.f)
		if got != test.expected {
			t.Errorf("Expected:\n%q\nActual:\n%q\n", test.expected, got)
		}
	}
}

func TestMain(t *testing.T) {
	expected := "$The $Go Programming $Language\nTHE GO Programming LANGUAGE\n"
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if got != expected {
		t.Errorf("Expected:\n%q\nActual:\n%q\n", expected, got)
	}
}
