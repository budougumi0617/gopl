// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var stdout io.Writer = os.Stdout // modified during testing

// Poster JSON structure for http://www.omdbapi.com/
type Poster struct {
	Title  string
	Poster string
}

func main() {
	result, err := SearchPoster(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	SavePoster(*result)
}

// SearchPoster get movie data
func SearchPoster(keywords []string) (*Poster, error) {
	q := url.QueryEscape(strings.Join(keywords, "+"))
	resp, err := http.Get("http://www.omdbapi.com/?t=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Poster
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SavePoster save image
func SavePoster(p Poster) {
	if p.Poster == "N/A" {
		fmt.Fprintln(stdout, "Not exist image")
		return
	}

	resp, err := http.Get(p.Poster)
	if err != nil {
		fmt.Printf("Error\n")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error StatusCode %v\n", resp.StatusCode)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error read\n")
	}
	ioutil.WriteFile("./"+p.Title+".jpg", data, 0644)
}
