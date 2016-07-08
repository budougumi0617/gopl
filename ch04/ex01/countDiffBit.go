// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout // modified during testing

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Fprintf(stdout, "Different bit count is %v\n", countDiffBit(c1, c2))
}

func countDiffBit(sha1, sha2 [32]uint8) int {
	n := 0
	for i := 0; i < len(sha1); i++ {
		s := sha1[i] ^ sha2[i]
		n += popCount(s)
	}
	return n
}

// PopCount using clear rightmost.
func popCount(x uint8) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}
