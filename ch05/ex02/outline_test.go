package main

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	expects := []string{"Element name:div\tcount:33\n", "Element name:script\tcount:9\n"}
	source, err := ioutil.ReadFile("../ex01/golang.html")
	if err != nil {
		t.Errorf("Failed file open, %v\n", err)
	}
	stdin = bytes.NewBuffer(source)
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	for _, expected := range expects {
		if !strings.Contains(got, expected) {
			t.Errorf("Not exitst expect string %s\ngot =:%v", expected, got)
		}
	}
}
