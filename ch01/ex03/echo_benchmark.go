package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout // modified during testing

func echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Fprintln(out, s)
}

func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(out, s)
}

func echo3(args []string) {
	fmt.Fprintln(out, strings.Join(args[1:], " "))
}

func echo4(args []string) {
	fmt.Fprintln(out, args[1:])
}

func main() {
	echo1(os.Args)
	echo2(os.Args)
	echo3(os.Args)
	echo4(os.Args)
}
