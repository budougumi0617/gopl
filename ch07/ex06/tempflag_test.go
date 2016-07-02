// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"os"
	"testing"
)

func TestLimitReader(t *testing.T) {
	var tests = []struct {
		args     []string
		expected string
	}{
		{[]string{"dumy", "-temp", "0K"}, "-273.15°C\n"},
		{[]string{"dumy", "-temp", "0°C"}, "0°C\n"},
		{[]string{"dumy", "-temp", "32F"}, "0°C\n"},
	}
	for _, test := range tests {
		os.Args = test.args
		stdout = new(bytes.Buffer) // captured output
		main()
		got := stdout.(*bytes.Buffer).String()
		if got != test.expected {
			t.Errorf("Result = %s, Expected %s", got, test.expected)
		}

	}
}
