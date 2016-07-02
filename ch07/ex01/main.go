// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// WordsCounter count of words
type WordsCounter int

func (c *WordsCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewBuffer(p))
	sc.Split(bufio.ScanWords)
	var lc int
	for sc.Scan() {
		lc++
	}
	*c += WordsCounter(lc)
	return lc, nil
}

// LinesCounter count of lines
type LinesCounter int

func (c *LinesCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewBuffer(p))
	var lc int
	for sc.Scan() {
		lc++
	}
	*c += LinesCounter(lc)
	return lc, nil
}

func main() {
	var wc WordsCounter
	wc.Write([]byte("test foo bar"))
	fmt.Printf("Count %d\n", wc)
	var lc LinesCounter
	b := []byte("teaste\nesateat\ntest test\ntet")
	lc.Write(b)

	fmt.Printf("Count %d\n", lc)

}
