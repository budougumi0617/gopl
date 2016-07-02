// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"os"
	"testing"
)

func TestShowSha(t *testing.T) {
	var tests = []struct {
		args     []string
		expected string
	}{
		{[]string{"showsha", "x"}, "2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881\n"},
		{[]string{"showsha", "x", "x"}, "2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881\n2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881\n"},
		{[]string{"showsha", "-hash=SHA256", "x"}, "2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881\n"},
		{[]string{"showsha", "-hash", "SHA256", "x"}, "2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881\n"},
		{[]string{"showsha", "-hash", "SHA384", "x"}, "d752c2c51fba0e29aa190570a9d4253e44077a058d3297fa3a5630d5bd012622f97c28acaed313b5c83bb990caa7da85\n"},
		{[]string{"showsha", "-hash", "SHA512", "x"}, "a4abd4448c49562d828115d13a1fccea927f52b4d5459297f8b43e42da89238bc13626e43dcb38ddb082488927ec904fb42057443983e88585179d50551afe62\n"},
		{[]string{"showsha", "-hash", "SHa256", "x"}, "Invalid Option. Choose hash. SHA256, SHA384, or SHA512\n"},
	}

	for _, test := range tests {
		os.Args = test.args
		stdout = new(bytes.Buffer) // captured output
		main()
		got := stdout.(*bytes.Buffer).String()
		if got != test.expected {
			t.Errorf("Result:%q, Expected:%q", got, test.expected)
		}
	}
}
