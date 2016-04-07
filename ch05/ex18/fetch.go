package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

var stdin io.Reader = os.Stdin   // modified during testing
var stdout io.Writer = os.Stdout // modified during testing
var stderr io.Writer = os.Stderr // modified during testing

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

// Fetch downloads the URL and returns the
// name and length of the local file.
func newfetch(url string) (filename string, n int64, err error) {
	var f *(os.File)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		resp.Body.Close()
		// Close file, but prefer error from Copy, if any.
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err = os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	return local, n, err

}

func main() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Fprintf(stdout, "%s => %s (%d bytes).\n", url, local, n)
	}
}
