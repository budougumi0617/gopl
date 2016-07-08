// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var stdout io.Writer = os.Stdout // modified during testing
var stdin io.Reader = os.Stdin   // modified during testing
var stderr io.Writer = os.Stderr // modified during testing

func main() {
	for _, url := range os.Args[1:] {
		w, i, err := CountWordsAndImages(url)
		if err == nil {
			fmt.Fprintf(stdout, "words %d, images, %d\n", w, i)
		}
	}
}

// CountWordsAndImages does an HTTP GET request for the HTML
// documet url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	words = countWords(0, n)
	images = countImages(0, n)
	return
}

func countWords(words int, n *html.Node) int {
	if n == nil {
		return words
	}
	if n.Type == html.TextNode {
		if n.Parent.Data != "script" && n.Parent.Data != "style" {
			//fmt.Fprintf(stderr, "gettext: len(%d), %v\n", len(n.Data), n.Data)
			wordlist := strings.Fields(n.Data)
			words += len(wordlist)
		}
	}
	words += countWords(0, n.FirstChild)
	return countWords(words, n.NextSibling)

}

func countImages(images int, n *html.Node) int {
	if n == nil {
		return images
	}
	if n.Type == html.ElementNode && (n.Data == "img") {
		for _, a := range n.Attr {
			if a.Key == "src" {
				images++
			}
		}
	}

	images += countImages(0, n.FirstChild)
	return countImages(images, n.NextSibling)
}
