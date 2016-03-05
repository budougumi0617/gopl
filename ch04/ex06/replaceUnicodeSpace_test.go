package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	expected := "Hello　世界\nHello 世界\n"
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if got != expected {
		t.Errorf("Result:%q, Expected:%q", got, expected)
	}
}

func TestReplaceSpace(t *testing.T) {
	var tests = []struct {
		input    []byte
		expected []byte
	}{
		{[]byte("Hello　world"), []byte("Hello world")},
		{[]byte("　world"), []byte(" world")},
		{[]byte("Hello　"), []byte("Hello ")},
	}

	for _, test := range tests {
		got := replaceSpace(test.input)
		for i := range got {
			if got[i] != test.expected[i] {
				t.Errorf("Result:%q, Expected:%q", got, test.expected)
			}
		}
	}
}
