// Copyright 2016 budougumi0617 All Rights Reserved.
// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

var stdout io.Writer = os.Stdout // modified during testing

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

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func main() {
	fmt.Println("Sort by Title, Year, and Length")
	sort.Sort(getMultiSort(tracks, byTitle, byYear, byLength))
	printTracks(tracks)
}

func compareSortResult() {
	fmt.Fprintln(stdout, "Sort by Title, Year, and Length")
	sort.Sort(getMultiSort(tracks, byTitle, byYear, byLength))
	printTracks(tracks)

	fmt.Fprintln(stdout, "Sort by Title, Year, and Length using sort.Stable")
	tracks = getTracks()
	sort.Stable(getMultiSort(tracks, byTitle))
	sort.Stable(getMultiSort(tracks, byYear))
	sort.Stable(getMultiSort(tracks, byLength))
	printTracks(tracks)
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var tracks = getTracks()

func getTracks() []*Track {
	return []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
}
