package main

import (
	"bytes"
	"testing"
)

func TestUnit(t *testing.T) {
	expected := "ZB/EB = 1000\n"
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if got != expected {
		t.Errorf("Result = %q, Expected %q", got, expected)
	}

}
