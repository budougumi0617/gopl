package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var stdout io.Writer = os.Stdout // modified during testing
var stdin io.Reader = os.Stdin   // modified during testing
var stderr io.Writer = os.Stderr // modified during testing

func main() {
	//doc, err := html.Parse(stdin)
	doc, err := html.Parse(stdin)
	if err != nil {
		fmt.Fprintf(stderr, "showTextNode: %v\n", err)
		return
	}
	for _, text := range gettext(nil, doc) {
		fmt.Fprintln(stdout, text)
	}
}

// gettext get text nodes
func gettext(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}
	if n.Type == html.TextNode {
		if n.Parent.Data != "script" && n.Parent.Data != "style" {
			//fmt.Fprintf(stderr, "gettext: len(%d), %v\n", len(n.Data), n.Data)
			for _, line := range strings.Split(n.Data, "\n") {
				if len(line) != 0 {
					texts = append(texts, line)
				}
			}

		}
	}
	texts = gettext(texts, n.FirstChild)
	return gettext(texts, n.NextSibling)
}
