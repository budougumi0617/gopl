// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"fmt"
	"os"
	"strconv"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	n, _ := strconv.ParseUint(os.Args[1], 0, 64)
	fmt.Printf("result %v\n", PopCountByShifting(n))
}

// PopCountByTable returns the population count (number of set bits) of x.
func PopCountByTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountByClear using clear rightmost.
func PopCountByClear(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

// PopCountByShifting using bit shift
func PopCountByShifting(x uint64) int {
	n := 0
	for i := 0; i < 64; i++ {
		if x&1 != 0 {
			n++
		}
		x = x >> 1 // bit shift
	}
	return n
}
