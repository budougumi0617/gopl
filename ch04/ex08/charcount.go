// Copyright 2016 budougumi0617 All Rights Reserved.
// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

var stdout io.Writer = os.Stdout // modified during testing
var stdin io.Reader = os.Stdin   // modified during testing

var countsByCategory = make(map[string]int) // counts of Unicode category

// RuneSlice for sorting []rune
type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	counts := make(map[rune]int) // counts of Unicode characters

	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		cntChars(r)

	}
	fmt.Fprintf(stdout, "rune\tcount\n")
	var runes []rune
	for rune := range counts {
		runes = append(runes, rune)
	}
	sort.Sort(RuneSlice(runes))
	for _, r := range runes {
		fmt.Fprintf(stdout, "%q\t%d\n", r, counts[r])
	}
	fmt.Fprint(stdout, "\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Fprintf(stdout, "%d\t%d\n", i, n)
		}
	}
	var categories []string
	for category := range countsByCategory {
		categories = append(categories, category)
	}
	sort.Strings(categories)
	fmt.Fprint(stdout, "\nCategory\tcount\n")
	for _, c := range categories {
		fmt.Fprintf(stdout, "%s\t%d\n", c, countsByCategory[c])
	}
	if invalid > 0 {
		fmt.Fprintf(stdout, "\n%d invalid UTF-8 characters\n", invalid)
	}
}

// cntChars counts the letters, digits, and punctuation marks in a text.
func cntChars(c rune) {
	// Refer to example in https://golang.org/pkg/unicode/#pkg-overview
	if unicode.IsControl(c) {
		countsByCategory["Control"]++
	}
	if unicode.IsDigit(c) {
		countsByCategory["Digit"]++
	}
	if unicode.IsGraphic(c) {
		countsByCategory["Graphic"]++
	}
	if unicode.IsLetter(c) {
		countsByCategory["Letter"]++
	}
	if unicode.IsLower(c) {
		countsByCategory["Lower"]++
	}
	if unicode.IsMark(c) {
		countsByCategory["Mark"]++
	}
	if unicode.IsNumber(c) {
		countsByCategory["Number"]++
	}
	if unicode.IsPrint(c) {
		countsByCategory["Printable"]++
	}
	if !unicode.IsPrint(c) {
		countsByCategory["NotPrintable"]++
	}
	if unicode.IsPunct(c) {
		countsByCategory["Punct"]++
	}
	if unicode.IsSpace(c) {
		countsByCategory["Space"]++
	}
	if unicode.IsSymbol(c) {
		countsByCategory["Symbol"]++
	}
	if unicode.IsTitle(c) {
		countsByCategory["Title"]++
	}
	if unicode.IsUpper(c) {
		countsByCategory["Upper"]++
	}
}
