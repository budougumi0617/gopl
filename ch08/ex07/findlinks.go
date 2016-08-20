// Copyright 2016 budougumi0617 All Rights Reserved.

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

var tokens = make(chan struct{}, 20)
var maxdepth int
var stdout io.Writer = os.Stdout // modified during testing
var stderr io.Writer = os.Stderr // modified during testing
var domains []string

type Item struct {
	url   string
	depth int
}

func crawl(item Item) []Item {
	var urls []Item
	if item.depth < maxdepth {
		depth := item.depth + 1
		tokens <- struct{}{} // acquire a token
		list, err := Extract(item.url)
		<-tokens // release the token
		if err != nil {
			log.Print(err)
		}
		for _, url := range list {
			local, err := makefile(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
				continue
			}
			fmt.Fprintf(os.Stderr, "%s => %s.\n", url, local)
		}

		fmt.Printf("----------------Crawled depth %d\n", depth)
	}
	return urls
}

// based on ch05ex13
func makefile(link string) (name string, err error) {

	lurl, _ := url.Parse(link)
	dir, fname := path.Split(lurl.Path)
	if len(fname) == 0 {
		fname = "index.html"
		lurl, _ = url.Parse(link + "index.html")
	}

	if err = os.MkdirAll("./local/"+lurl.Host+dir, 0755); err != nil {
		fmt.Fprintln(stderr, err)
		return
	}
	response, err := http.Get(lurl.String())
	if err != nil {
		fmt.Fprintln(stderr, err)
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return
	}
	// TODO Need to replace URL in file.
	file, err := os.OpenFile("./local/"+lurl.Host+dir+fname, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Fprintln(stderr, err)
		return
	}

	file.Write(body)
	return file.Name(), file.Close()
}

func main() {
	flag.IntVar(&maxdepth, "depth", 3, "max crawl depth")
	flag.Parse()
	// Save original domains
	for _, u := range os.Args[1:] {
		u, _ := url.Parse(u)
		d := strings.Replace(u.Host, "www.", "", 1)
		domains = append(domains, d)
	}
	worklist := make(chan []Item)
	var n int // number of pending sends to worklist

	n++
	go func() {
		var urls []Item
		for _, url := range flag.Args() {
			urls = append(urls, Item{url, 0})
		}
		worklist <- urls
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, item := range list {
			if !seen[item.url] {
				seen[item.url] = true
				n++
				go func(item Item) {
					worklist <- crawl(item)
				}(item)
			}
		}
	}
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				for _, d := range domains {
					// Ignore other domain.
					if strings.Contains(link.Host, d) {
						links = append(links, link.String())
					}
				}

			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
