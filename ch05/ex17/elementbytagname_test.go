// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestElementByTagName(t *testing.T) {
	source, _ := ioutil.ReadFile("../ex08/ex08.html")
	data := bytes.NewBuffer(source)
	doc, _ := html.Parse(data)
	nodes := ElementByTagName(doc, "a")
	if len(nodes) != 3 {
		t.Errorf("Not found element: %q", nodes)
	}
	for _, n := range nodes {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val != "sameid" {
				t.Errorf("Element id: %s", a.Val)
			}
		}
	}

}

func TestMain(t *testing.T) {
	os.Args = []string{"", "https://github.com/budougumi0617/"}
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if len(got) == 0 {
		t.Errorf("Cannot find")
	}
}
