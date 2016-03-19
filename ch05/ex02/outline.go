// Outline prints the outline of an HTML document tree.
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

var counts = make(map[string]int)

func main() {
	doc, err := html.Parse(stdin)
	if err != nil {
		fmt.Fprintf(stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
	for k, v := range counts {
		fmt.Fprintf(stdout, "Element name:%s\tcount:%d\n", k, v)
	}
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		counts[n.Data]++
		fmt.Fprintln(stdout, stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
