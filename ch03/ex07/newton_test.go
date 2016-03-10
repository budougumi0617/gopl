package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if len(got) < 0 {
		t.Error("Cannot create image")
	}
}
