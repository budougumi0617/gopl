package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	expected := "[0 1 2 3 4 5]\n[3 4 5 0 1 2]\n"
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if got != expected {
		t.Errorf("Result:%q, Expected:%q", got, expected)
	}
}

func TestRotate(t *testing.T) {
	var tests = []struct {
		input    []int
		order    int
		expected []int
	}{
		{[]int{0, 1, 2}, -1, []int{0, 1, 2}},
		{[]int{0, 1, 2}, 0, []int{0, 1, 2}},
		{[]int{0, 1, 2}, 1, []int{1, 2, 0}},
		{[]int{0, 1, 2}, 2, []int{2, 0, 1}},
		{[]int{0, 1, 2}, 3, []int{0, 1, 2}},
		{[]int{0, 1, 2}, 4, []int{0, 1, 2}},
	}

	for _, test := range tests {
		s := test.input
		rotate(s, test.order)
		if !reflect.DeepEqual(s, test.expected) {
			t.Errorf("Order = %v, Result = %v, Expected %v", test.order, s, test.expected)
		}
	}
}
