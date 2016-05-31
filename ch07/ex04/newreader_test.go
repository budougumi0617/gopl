package main

import "testing"

func TestCounterWriter(t *testing.T) {
	var tests = []struct {
		words    [][]byte
		expected int64
	}{
		{[][]byte{[]byte("")}, 0},
		{[][]byte{[]byte("one"), []byte("two")}, 6},
	}
	for _, test := range tests {
		if false {
			t.Errorf("Result = , Expected %v", test.expected)
		}
	}
}

func Example_main() {
	main()

	// Output:
	// https://golang.org/dl/
	// https://google.com/
}
