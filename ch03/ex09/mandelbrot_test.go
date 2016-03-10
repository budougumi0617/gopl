package main

import (
	"net/url"
	"testing"
)

func TestParseQuery(t *testing.T) {
	var tests = []struct {
		url       string
		expectedx float64
		expectedy float64
		expecteds float64
	}{
		{"http://test/", 0, 0, 2},
		{"http://test/?x=1&y=2&scale=5", 1, 2, 5},
	}

	for _, test := range tests {
		u, _ := url.Parse(test.url)
		x, y, s := parsequery(u)
		if x != test.expectedx {
			t.Errorf("Actual %v, Expected %v", x, test.expectedx)
		}
		if y != test.expectedy {
			t.Errorf("Actual %v, Expected %v", y, test.expectedy)
		}
		if s != test.expecteds {
			t.Errorf("Actual %v, Expected %v", s, test.expecteds)
		}
	}
}

func TestMain(t *testing.T) {

	img := createImage(0, 0, 0)

	if b := img.Bounds(); b.Dx() != 1024 || b.Dy() != 1024 {
		t.Errorf("Actual dx %v, Expected %v", b.Dx(), 1024)
		t.Errorf("Actual dy %v, Expected %v", b.Dy(), 1024)
	}
}
