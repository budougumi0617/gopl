// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

var stdout io.Writer = os.Stdout // modified during testing
var stdin io.Reader = os.Stdin   // modified during testing
var stderr io.Writer = os.Stderr // modified during testing

func main() {
	doc, err := html.Parse(stdin)
	if err != nil {
		fmt.Fprintf(stderr, "findlinks1: %v\n", err)
		return
	}

	for _, link := range visit(nil, doc) {
		fmt.Fprintln(stdout, link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	//fmt.Printf("%v\n", *n)
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "link") {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, n.Data+" href:"+a.Val)
			}
		}
	}

	if n.Type == html.ElementNode && (n.Data == "img" || n.Data == "script") {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links = append(links, n.Data+" src:"+a.Val)
			}
		}
	}

	links = visit(links, n.FirstChild)
	return visit(links, n.NextSibling)
}
