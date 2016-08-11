// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"fmt"
	"sync"
)

func popCountShiftMask(x uint64) int {
	count := 0
	mask := uint64(1)
	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}
		mask <<= 1
	}
	return count
}

func popCountShiftValue(x uint64) int {
	count := 0
	mask := uint64(1)
	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}
		x >>= 1
	}
	return count
}

func popCountClearRightmost(x uint64) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

// pc[i] is the population count of i.
var loadTableOnce sync.Once
var pc [256]byte

func table() [256]byte {
	loadTableOnce.Do(func() {
		for i := range pc {
			pc[i] = pc[i/2] + byte(i&1)
		}
	})
	return pc
}

// PopCountTable returns the population count (number of set bits) of x.
func popCountTable(x uint64) int {
	// Calling Table here makes PopCountTable 4 times slower, from 8ns -> 34ns
	// per call.
	pc := table()
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	fmt.Println("Exercise is in test file.")
}
