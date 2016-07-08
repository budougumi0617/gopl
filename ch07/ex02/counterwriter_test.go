// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"testing"
)

func TestCounterWriter(t *testing.T) {
	var tests = []struct {
		words    [][]byte
		expected int64
	}{
		{[][]byte{[]byte("")}, 0},
		{[][]byte{[]byte("one"), []byte("two")}, 6},
	}
	for _, test := range tests {
		w, c := CounterWriter(new(bytes.Buffer))
		for _, b := range test.words {
			w.Write(b)
		}
		if *c != test.expected {
			t.Errorf("Result = %v, Expected %v", *c, test.expected)
		}
	}
}

func Example_main() {
	main()

	// Output:
	// test input
	// test input
	// Byte count 22
}
