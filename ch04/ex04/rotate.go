// Rev reverses a slice.
package main

import (
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout // modified during testing

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Fprintln(stdout, a) // "[0 1 2 3 4 5]"
	rotate(a, 3)
	fmt.Fprintln(stdout, a) // "[2 3 4 5 0 1]"
}

// rotate slice
func rotate(s []int, order int) {
	if len(s) < order || order < 0 {
		return
	}
	for i := 0; i < order; i++ {
		for j := 0; j < len(s)-1; j++ {
			s[j], s[j+1] = s[j+1], s[j]
		}
	}
}
