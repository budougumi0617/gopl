package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var stdout io.Writer = os.Stdout // modified during testing
var stderr io.Writer = os.Stderr // modified during testing

func main() {
	fetch()
}

func fetch() {
	fmt.Printf("%v", os.Args[1:])
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(stdout, "resp.status %s\n", resp.Status)
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Fprintf(stdout, "%s", b)
	}
}
