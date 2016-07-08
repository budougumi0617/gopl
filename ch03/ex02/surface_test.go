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
	if strings.Contains(got, "NaN") {
		t.Errorf("SVG included NaN value.\n%s", got)
	}
}
