// Copyright 2016 budougumi0617 All Rights Reserved.
//Package ex01 Customize echo
package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout // modified during testing

//Prints its command-line argments
func main() {
	var s, sep string
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(out, s)
}
