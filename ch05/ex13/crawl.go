package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

var stdout io.Writer = os.Stdout // modified during testing
var stderr io.Writer = os.Stderr // modified during testing

var roots []string

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	roots = make([]string, len(worklist))
	copy(roots, worklist)
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(u string) []string {
	fmt.Fprintln(stdout, u)
	list, err := Extract(u)
	if err != nil {
		log.Print(err)
	}
	for _, link := range list {
		lurl, _ := url.Parse(link)
		for _, r := range roots {
			ourl, _ := url.Parse(r)
			if ourl.Host == lurl.Host {
				makefile(lurl)
			}
		}
	}

	return list
}

func makefile(lurl *url.URL) {
	dir, fname := path.Split(lurl.Path)
	if len(fname) == 0 { // direcotry path
		return
	}

	if err := os.MkdirAll("./local/"+lurl.Host+dir, 0777); err != nil {
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
	file, err := os.OpenFile("./local/"+lurl.Host+dir+fname, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Fprintln(stderr, err)
		return
	}

	file.Write(body)
	file.Close()
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}
