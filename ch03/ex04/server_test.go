// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"net/url"
	"strings"
	"testing"
)

func TestCorner(t *testing.T) {
	var tests = []struct {
		i        int
		j        int
		expected bool
	}{
		{50, 50, false},
		{50, 49, true},
		{49, 50, true},
	}

	for _, test := range tests {
		_, _, result := corner(test.i, test.j)
		if result != test.expected {
			t.Errorf("i = %v, j = %v, result = %v", test.i, test.j, test.expected)
		}
	}
}

func TestParseQuery(t *testing.T) {
	var tests = []struct {
		url              string
		width            float64
		height           float64
		red, green, blue uint64
	}{
		{"http://test/", 600, 320, 255, 255, 255},
		{"http://test/?width=300&red=10", 300, 320, 10, 255, 255},
	}

	for _, test := range tests {
		u, _ := url.Parse(test.url)
		w, h := parsequery(u)
		if w != test.width {
			t.Errorf("Actual %v, Expected %v", w, test.width)
		}
		if h != test.height {
			t.Errorf("Actual %v, Expected %v", h, test.height)
		}
	}
}

func TestMain(t *testing.T) {
	stdout = new(bytes.Buffer) // captured output
	surface(600, 320, 255, 255, 255)
	got := stdout.(*bytes.Buffer).String()
	if strings.Contains(got, "NaN") {
		t.Errorf("SVG included NaN value.\n%s", got)
	}
}
