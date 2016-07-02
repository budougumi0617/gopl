// Copyright 2016 budougumi0617 All Rights Reserved.
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

	if len(got) == 0 {
		t.Errorf("no output")
	}
}

func TestJoin(t *testing.T) {
	sep := " "
	args := []string{"one", "two", "three", "four"}
	expected := strings.Join(args, sep)
	actural := join(sep, args[0], args[1], args[2], args[3])

	if expected != actural {
		t.Errorf("Expected:\n%v\nActual:\n%v", expected, actural)
	}

	args = []string{""} // no string
	expected = strings.Join(args, sep)
	actural = join(sep, args[0])

	if expected != actural {
		t.Errorf("Expected:\n%v\nActual:\n%v", expected, actural)
	}

	args = []string{} // empty string array
	expected = strings.Join(args, sep)
	actural = join(sep) // same as empty string array

	if expected != actural {
		t.Errorf("Expected:\n%v\nActual:\n%v", expected, actural)
	}
}
