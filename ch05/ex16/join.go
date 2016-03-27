package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var stdin io.Reader = os.Stdin   // modified during testing
var stdout io.Writer = os.Stdout // modified during testing
var stderr io.Writer = os.Stderr // modified during testing

func main() {
	args := []string{"one", "two", "three", "four"}
	sep := " "
	fmt.Fprintf(stdout, "strings.Join:\n\"%v\"\n", strings.Join(args, sep))
	fmt.Fprintf(stdout, "ex16.join:\n\"%v\"\n",
		join(sep, args[0], args[1], args[2], args[3]))
	emp := []string{}
	fmt.Fprintf(stdout, "strings.Join:\n\"%v\"\n", strings.Join(emp, sep))
}

func join(sep string, a ...string) string {
	if a == nil {
		return ""
	}
	return strings.Join(a, sep)
}
