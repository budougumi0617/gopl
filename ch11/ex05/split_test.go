// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"strings"
	"testing"
)

type Inputs struct {
	s   string
	sep string
}

func TestSplit(t *testing.T) {
	var tests = []struct {
		inputs Inputs
		want   int
	}{
		{Inputs{"a:b:c", ":"}, 3},
		{Inputs{"a:b:c", ","}, 1},
		{Inputs{"a,b,c", ","}, 3},
		{Inputs{"a,b,c", ":"}, 1},
	}
	for _, test := range tests {
		words := strings.Split(test.inputs.s, test.inputs.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.inputs.s, test.inputs.sep, got, test.want)
		}
	}
}
