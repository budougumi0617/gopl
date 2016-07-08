// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html"
)

// visit appends to links each link found in n and returns the result.
func orgvisit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = orgvisit(links, c)
	}
	return links
}

func TestFindLinks(t *testing.T) {

	source, _ := ioutil.ReadFile("./golang.html")
	data := bytes.NewBuffer(source)
	doc, _ := html.Parse(data)
	got := visit(nil, doc)
	// Reset Buffer
	data = bytes.NewBuffer(source)
	doc, _ = html.Parse(data)
	expected := orgvisit(nil, doc)

	for i, s := range got {
		if s != expected[i] {
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
