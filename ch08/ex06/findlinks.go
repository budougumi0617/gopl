// Copyright 2016 budougumi0617 All Rights Reserved.

package main

import (
	"flag"
	"fmt"
	"log"
)

var tokens = make(chan struct{}, 20)
var maxdepth int

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
			urls = append(urls, Item{url, depth})
			fmt.Printf("depth %d, url %s\n", item.depth, url)
		}
		fmt.Printf("----------------Crawled depth %d\n", depth)
	}
	return urls
}

func main() {
	flag.IntVar(&maxdepth, "depth", 3, "max crawl depth")
	flag.Parse()
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
