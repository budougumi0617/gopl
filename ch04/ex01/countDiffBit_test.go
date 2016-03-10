package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	stdout = new(bytes.Buffer) // captured output
	expected := "Different bit count is 125\n"
	main()
	got := stdout.(*bytes.Buffer).String()
	if got != expected {
		t.Errorf("Result = %q, Expected %q", got, expected)
	}
}

func TestPopCountDiffBit(t *testing.T) {
	var tests = []struct {
		sha1     [32]uint8
		sha2     [32]uint8
		expected int
	}{
		{[...]uint8{31: 1}, [...]uint8{31: 0}, 1},
		{[...]uint8{31: 5}, [...]uint8{31: 1}, 1},
		{[...]uint8{0: 11, 30: 2, 31: 5}, [...]uint8{31: 1}, 5},
	}

	for _, test := range tests {
		got := countDiffBit(test.sha1, test.sha2)
		if got != test.expected {
			t.Errorf("Result = %v, Expected %v", got, test.expected)
		}
	}
}
