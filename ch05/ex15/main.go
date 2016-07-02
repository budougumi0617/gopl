// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

var stdout io.Writer = os.Stdout // modified during testing

func main() {
	x, y := 5, 10
	fmt.Fprintf(stdout, "max(%d, %d) = %d\n", x, y, max(x, y))
}

func max(values ...int) int {
	if len(values) == 0 {
		panic("No input")
	}

	r := math.MinInt64
	for _, v := range values {
		if r < v {
			r = v
		}
	}
	return r
}

func max2(i int, values ...int) int {
	for _, v := range values {
		if i < v {
			i = v
		}
	}
	return i
}

func min(values ...int) int {
	if len(values) == 0 {
		panic("No input")
	}
	r := math.MaxInt64
	for _, v := range values {
		if v < r {
			r = v
		}
	}
	return r
}

func min2(i int, values ...int) int {
	for _, v := range values {
		if v < i {
			i = v
		}
	}
	return i
}
