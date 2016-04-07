package main

import (
	"bytes"
	"testing"
)

func TestReturnByPanic(t *testing.T) {
	if r := returnbypanic(); r == 0 {
		t.Errorf("Result:%d, Expected:0", r)
	}
}

func TestMain(t *testing.T) {
	expected := "Value: 1\n"
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if got != expected {
		t.Errorf("Result:%v, Expected:%v", got, expected)
	}
}
