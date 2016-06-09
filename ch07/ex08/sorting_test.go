package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()

	var tracks = []*Track{
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	stdout = new(bytes.Buffer) // captured output
	printTracks(tracks)
	expected := stdout.(*bytes.Buffer).String()
	if got != expected {
		t.Errorf("Result = \n%q\nExpected = \n%q", got, expected)
	}
}

func TestCompareSortResult(t *testing.T) {
	stdout = new(bytes.Buffer) // captured output
	compareSortResult()
	got := stdout.(*bytes.Buffer).String()
	stdout = new(bytes.Buffer) // captured output
	fmt.Fprintln(stdout, "Sort by Title, Year, and Length")
	var tracks = []*Track{
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	printTracks(tracks)
	fmt.Fprintln(stdout, "Sort by Title, Year, and Length using sort.Stable")
	tracks = []*Track{
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	}
	printTracks(tracks)
	expected := stdout.(*bytes.Buffer).String()
	if got != expected {
		t.Errorf("Result = \n%q\nExpected = \n%q", got, expected)
	}
}
