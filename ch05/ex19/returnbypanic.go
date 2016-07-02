// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout // modified during testing
var stdin io.Reader = os.Stdin   // modified during testing
var stderr io.Writer = os.Stderr // modified during testing

// returnbypanic return non-zero value by panic and recover
func returnbypanic() (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = 1
		}
	}()
	panic("panic")
}

func main() {
	fmt.Fprintf(stdout, "Value: %d\n", returnbypanic())
}
