// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var stdout io.Writer = os.Stdout // modified during testing

func main() {
	url := os.Args[1]
	//tags := os.Args[2:]
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	doc, _ := html.Parse(resp.Body)
	images := ElementByTagName(doc, "img")

	resp, _ = http.Get(url)
	doc, _ = html.Parse(resp.Body)
	headings := ElementByTagName(doc, "h1", "h2", "h3", "h4")
	fmt.Fprintln(stdout, "img")
	for i, image := range images {
		fmt.Fprintf(stdout, "No.%d is %s\n", i, image.Data)
		for _, a := range image.Attr {
			fmt.Fprintf(stdout, "\t\"%s\" element, value is \"%s\"\n", a.Key, a.Val)
		}
	}
	fmt.Fprintln(stdout, "headings")
	for i, head := range headings {
		fmt.Fprintf(stdout, "No.%d is %s\n", i, head.Data)
		for _, a := range head.Attr {
			fmt.Fprintf(stdout, "\t\"%s\" element, value is \"%s\"\n", a.Key, a.Val)
		}
	}
}

// ElementByTagName that given an HTML node tree and zero or moe names, returns all the elements that match one of those names
func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	if len(name) == 0 {
		return nil
	}
	if doc.Type == html.ElementNode {
		for _, tag := range name {
			if doc.Data == tag {
				nodes = append(nodes, doc)
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementByTagName(c, name...)...)
	}
	return nodes
}
