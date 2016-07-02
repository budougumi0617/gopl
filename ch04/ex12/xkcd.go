// Copyright 2016 budougumi0617 All Rights Reserved.
// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

var stdout io.Writer = os.Stdout // modified during testing
var stdin io.Reader = os.Stdin   // modified during testing
var dir = "./data"

const templ = `{{.TotalCount}} comics:
{{range .Items}}----------------------------------------
Number: {{.Num}}
URL:    {{.Num | printf "http://xkcd.com/%d/"}}
Link:   {{.Link}}
Title:   {{.SafeTitle}}
Transcript: {{.Transcript | println}}
{{end}}`

var report = template.Must(template.New("comiclist").
	Parse(templ))

// DEBUG for compress test time
var DEBUG = false

func main() {
	if _, err := os.Stat(dir); err != nil {
		lastnum, err := getLastPage()
		if err != nil {
			log.Fatal(err)
		}
		if DEBUG { // compress test time
			lastnum = 10
		}
		getData(lastnum)
	}
	files, _ := filepath.Glob(dir + "/*.json")
	comics := readFiles(files)
	if err := report.Execute(stdout, comics); err != nil {
		log.Fatal(err)
	}
	// main loop
	scanner := bufio.NewScanner(stdin)
	fmt.Fprintln(stdout, "Exit command \\q")
	for {
		fmt.Fprint(stdout, "Input search words:")
		scanner.Scan()
		w := scanner.Text()
		if w == "\\q" {
			fmt.Fprintln(stdout, "Exit xkcd")
			break
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			break
		}
		result := searchComic(w, comics)
		if err := report.Execute(stdout, result); err != nil {
			log.Fatal(err)
		}
	}

}

func searchComic(w string, comics *ComicSearchResult) *ComicSearchResult {
	result := new(ComicSearchResult)
	for _, comic := range comics.Items {

		if strings.Contains(comic.SafeTitle, w) || strings.Contains(comic.Transcript, w) {
			result.TotalCount++
			result.Items = append(result.Items, comic)
		}
	}
	return result
}

func getData(lastnum int) {
	os.Mkdir(dir, 0777)

	for i := 1; i <= lastnum; i++ {
		url := URL + strconv.Itoa(i) + JSONFILE
		file := dir + "/" + strconv.Itoa(i) + ".json"
		fmt.Printf("Try:%v to %v\n", url, file)
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		// We must close resp.Body on all execution paths.
		// (Chapter 5 presents 'defer', which makes this simpler.)
		if resp.StatusCode != http.StatusOK {
			continue
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}
		err = ioutil.WriteFile(file, data, 0644)
		if err != nil {
			continue
		}
	}
}

func readFiles(files []string) *ComicSearchResult {
	fmt.Printf("readFiles : %s\n", files)
	results := new(ComicSearchResult)
	for _, file := range files {
		b, err := ioutil.ReadFile(file)
		result := new(Comic)
		if err = json.Unmarshal(b, result); err != nil {
			continue
		}
		results.TotalCount++
		results.Items = append(results.Items, result)
	}

	return results
}

// Home for http://xkcd.com/info.0.json to get total page number
type Home struct {
	Num int
}

// SearchIssues queries the GitHub issue tracker.
func getLastPage() (int, error) {
	resp, err := http.Get("http://xkcd.com/info.0.json")
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return -1, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result Home
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return -1, err
	}
	return result.Num, nil
}
