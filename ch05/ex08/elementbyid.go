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
	id := os.Args[2]
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	doc, _ := html.Parse(resp.Body)
	n := ElementByID(doc, id)
	if n == nil {
		fmt.Fprintf(stdout, "Not found id element\n")
		return
	}
	for _, a := range n.Attr {
		fmt.Fprintf(stdout, "\"%s\" has \"%s\" element, value is \"%s\"\n",
			n.Data, a.Key, a.Val)
	}
}

// ElementByID finds the first HTML element with specified id attribute
func ElementByID(doc *html.Node, id string) *html.Node {
	isTarget := func(n *html.Node) bool { // find ID in current node
		if e != nil {
			return true
		}
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				e = n
				return true
			}
		}
		return false
	}
	found := func(n *html.Node) bool { // ID does not find yet
		return e != nil
	}
	forEachNode(doc, isTarget, found)
	return e
}

var e *html.Node // result node

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil && pre(n) {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil && post(n) {
		return
	}
}
