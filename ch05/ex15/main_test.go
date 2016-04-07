package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if len(got) == 0 {
		t.Errorf("Fail result")
	}
}

func TestMax(t *testing.T) {
	var tests = []struct {
		values   []int
		expected int
	}{
		{[]int{1, 2, 30, 40}, 40},
		{[]int{-1, 3}, 3},
	}

	for _, test := range tests {
		if r := max(test.values...); r != test.expected {
			t.Errorf("Result = %v, Expected %v", r, test.expected)
		}
	}
}

func TestMaxPanic(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			if p != "No input" {
				t.Errorf("Result = %v", p)
			}
		} else {
			t.Errorf("Not occurs panic")
		}
	}()
	max()
}

func TestMin(t *testing.T) {
	var tests = []struct {
		values   []int
		expected int
	}{
		{[]int{1, 2, 30, 40}, 1},
		{[]int{3, -3}, -3},
	}

	for _, test := range tests {
		if r := min(test.values...); r != test.expected {
			t.Errorf("Result = %v, Expected %v", r, test.expected)
		}
	}
}

func TestMinPanic(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			if p != "No input" {
				t.Errorf("Result = %v", p)
			}
		} else {
			t.Errorf("Not occurs panic")
		}
	}()

	min()
}

func TestMax2(t *testing.T) {
	var tests = []struct {
		values   []int
		expected int
	}{
		{[]int{1, 2, 30, 40}, 40},
		{[]int{-1, 3}, 3},
	}

	for _, test := range tests {
		if r := max2(test.values[0], test.values[1:]...); r != test.expected {
			t.Errorf("Result = %v, Expected %v", r, test.expected)
		}
	}
}

func TestMin2(t *testing.T) {
	var tests = []struct {
		values   []int
		expected int
	}{
		{[]int{1, 2, 30, 40}, 1},
		{[]int{3, -3}, -3},
	}

	for _, test := range tests {
		if r := min2(test.values[0], test.values[1:]...); r != test.expected {
			t.Errorf("Result = %v, Expected %v", r, test.expected)
		}
	}
}
