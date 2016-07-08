// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import "testing"

func TestWordsCounter(t *testing.T) {
	var tests = []struct {
		words    []byte
		expected int
	}{
		{[]byte(""), 0},
		{[]byte("one"), 1},
		{[]byte("one two\nthree four"), 4},
		{[]byte("one two \nthree four"), 4},
	}
	for _, test := range tests {
		var wc WordsCounter
		wc.Write(test.words)
		if int(wc) != test.expected {
			t.Errorf("Result = %v, Expected %v", wc, test.expected)
		}
	}
}

func TestLinesCounter(t *testing.T) {
	var tests = []struct {
		lines    []byte
		expected int
	}{
		{[]byte(""), 0},
		{[]byte("one"), 1},
		{[]byte("one\n  two"), 2},
		{[]byte("one \ntwo\n\nfour"), 4},
		{[]byte("one \ntwo\n\nfour\n"), 4},
	}
	for _, test := range tests {
		var lc LinesCounter
		lc.Write(test.lines)
		if int(lc) != test.expected {
			t.Errorf("Result = %v, Expected %v", lc, test.expected)
		}
	}
}

func Example_main() {
	main()

	// Output:
	// Count 3
	// Count 4
}
