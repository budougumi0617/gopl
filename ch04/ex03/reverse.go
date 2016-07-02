// Copyright 2016 budougumi0617 All Rights Reserved.
// Rev reverses a slice.
package main

import (
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout // modified during testing

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Fprintln(stdout, a) // "[5 4 3 2 1 0]"
	reverse6(&a)
	fmt.Fprintln(stdout, a) // "[5 4 3 2 1 0]"
}

// reverse reverses array of ints in place.
func reverse6(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
