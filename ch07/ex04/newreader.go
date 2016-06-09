package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type myreader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

func (mr *myreader) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	if mr.i >= int64(len(mr.s)) {
		return 0, io.EOF
	}
	mr.prevRune = -1
	n = copy(b, mr.s[mr.i:])
	mr.i += int64(n)
	return
}

// NewReader returns Reader
func NewReader(s string) io.Reader {
	return &myreader{s: string(`<!DOCTYPE html>
  <html>
  <head>
    <title>One</title>
  </head>
  <body>
  <a href="https://golang.org/dl/" id="start">
  test
  </a>

  <a href="https://google.com/" id="start">
  test2
  </a>

  </body>
  </html>
  `)}
}

func main() {

	doc, err := html.Parse(NewReader("OK?"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
