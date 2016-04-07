package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if !strings.Contains(got, "Found Cycles between \"linear algebra\" and \"calculus\"") {
		t.Errorf("Could not find cycles: %s", got)
	}
}
