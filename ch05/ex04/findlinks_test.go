// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html"
)

func TestFindLinks(t *testing.T) {

	source, _ := ioutil.ReadFile("./golang.html")
	data := bytes.NewBuffer(source)
	doc, _ := html.Parse(data)
	got := visit(nil, doc)
	expected := []string{
		"link href:/lib/godoc/style.css",
		"link href:/opensearch.xml",
		"link href:/lib/godoc/jquery.treeview.css",
		"a href:https://golang.org/dl/",
		"img src:../images/img001.gif",
		"img src:../images/img002.gif",
		"script src:https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js",
		"script src:/lib/godoc/jquery.treeview.js",
		"script src:/lib/godoc/jquery.treeview.edit.js",
		"script src:/lib/godoc/playground.js",
		"script src:/lib/godoc/godocs.js",
	}
	for i, s := range expected {
		if s != got[i] {
			t.Errorf("Result:%q, Expected:%q", got[i], expected[i])
		}
	}

}

func TestMain(t *testing.T) {
	source, _ := ioutil.ReadFile("./golang.html")
	stdin = bytes.NewBuffer(source)
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if len(got) == 0 {
		t.Errorf("Fail result")
	}
}
