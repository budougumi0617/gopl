package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestCountWords(t *testing.T) {
	source, _ := ioutil.ReadFile("./ex05.html")
	data := bytes.NewBuffer(source)
	doc, _ := html.Parse(data)
	w := countWords(0, doc)

	if w != 3 {
		t.Errorf("Words expects 3, but actual was %d\n", w)
	}

}

func TestCountWordsAndImages(t *testing.T) {
	source, _ := ioutil.ReadFile("./ex05.html")
	data := bytes.NewBuffer(source)
	doc, _ := html.Parse(data)
	i := countImages(0, doc)
	if i != 2 {
		t.Errorf("Images expects 2, but actual was %d\n", i)
	}

}

func TestMain(t *testing.T) {
	os.Args = []string{"", "https://github.com/budougumi0617/"}
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if len(got) == 0 {
		t.Errorf("No output")
	}
}
