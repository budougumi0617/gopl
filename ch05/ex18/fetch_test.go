// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"os"
	"testing"
)

func TestFetch(t *testing.T) {
	var tests = []struct {
		url string
	}{
		{"https://github.com/budougumi0617/"},
		{"https://github.com/budougumi0617/GoTraining/404status"},
	}

	for _, test := range tests {
		local, n, err := fetch(test.url)
		nlocal, nn, nerr := newfetch(test.url)
		if local != nlocal || n != n {
			t.Errorf("fetch\tlocal\t%v\nnewfetch\tlocal\t%v\n\nfetch\tn\t%v\nnewfetch\tn\t%v\n\n", local, nlocal, n, nn)
		}
		if err != nil && err != nerr {
			t.Errorf("fetch\terr\t%v\nnewfetch\terr\t%v\n", err, nerr)
		}
	}
}

func TestMain(t *testing.T) {
	os.Args = []string{"fetch", "https://github.com/budougumi0617/"}
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if len(got) == 0 {
		t.Errorf("No output")
	}
}
