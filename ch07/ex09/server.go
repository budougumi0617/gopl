package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

var templ = template.Must(template.New("tracklist").Parse(`
  <!DOCTYPE html>
<html>
<head>
<title>ch07ex09</title>
<style>
table {
	border-collapse: collapse;
}
td, th {
	border: solid 1px;
	padding: 0.5em;
  text-align: right;
}
</style>
</head>
<body>
  <table>
  <tr>
	  <th><a href="./?by=title">Title</a></th>
	  <th><a href="./?by=artist">Artist</a></th>
	  <th><a href="./?by=album">Album</a></th>
	  <th><a href="./?by=year">Year</a></th>
	  <th><a href="./?by=length">Length</a></th>
	</tr>
    {{range .}}
    <tr>
      <td>{{.Title}}</td>
			<td>{{.Artist}}</td>
			<td>{{.Album}}</td>
			<td>{{.Year}}</td>
			<td>{{.Length}}</td>
    </tr>
    {{end}}
  </table>
</body>
</html>`))

type database []*Track

var preLess lessFunc

func main() {
	db := database(getTracks())
	preLess = nil
	http.HandleFunc("/", db.sort)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getLess(s string) (less lessFunc) {
	switch s {
	case "title":
		less = byTitle
	case "artist":
		less = byArtist
	case "album":
		less = byAlbum
	case "year":
		less = byYear
	case "length":
		less = byLength
	default:
		less = byTitle
	}
	return
}

func (db database) sort(w http.ResponseWriter, req *http.Request) {
	s := req.URL.Query().Get("by")
	if s != "" {
		less := getLess(s)
		if preLess != nil {
			sort.Sort(getMultiSort([]*Track(db), less, preLess))
		} else {
			sort.Sort(getMultiSort([]*Track(db), less))
		}
		preLess = less
	}

	if err := templ.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

// Track for record
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type lessFunc func(p1, p2 *Track) bool

func byTitle(t1, t2 *Track) bool  { return t1.Title < t2.Title }
func byArtist(t1, t2 *Track) bool { return t1.Artist < t2.Artist }
func byAlbum(t1, t2 *Track) bool  { return t1.Album < t2.Album }
func byYear(t1, t2 *Track) bool   { return t1.Year < t2.Year }
func byLength(t1, t2 *Track) bool { return t1.Length < t2.Length }

type multiSort struct {
	tracks []*Track
	less   []lessFunc
}

func (x multiSort) Len() int { return len(x.tracks) }
func (x multiSort) Less(i, j int) bool {
	p, q := x.tracks[i], x.tracks[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(x.less)-1; k++ {
		less := x.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return x.less[k](p, q)
}
func (x multiSort) Swap(i, j int) { x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i] }

func getMultiSort(t []*Track, less ...lessFunc) *multiSort {
	return &multiSort{tracks: t, less: less}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func getTracks() []*Track {
	return []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
}
