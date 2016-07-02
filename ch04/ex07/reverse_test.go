// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	expected := "tsえt,tsて\n"
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if !strings.Contains(got, expected) {
		t.Errorf("Result:%q, does not conatain:%q", got, expected)
	}
}
