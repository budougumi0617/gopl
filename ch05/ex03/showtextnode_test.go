package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestMain(t *testing.T) {
	expected := "The Go Programming Language\n...\nThe Go Programming Language\n"
	source, _ := ioutil.ReadFile("./test.html")
	stdin = bytes.NewBuffer(source)
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if got != expected {
		t.Errorf("Expected:\n%v\n\nActual:\n%v", expected, got)
	}
}
