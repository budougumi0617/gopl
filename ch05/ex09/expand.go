// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var stdout io.Writer = os.Stdout // modified during testing

func expand(s string, f func(string) string) string {
	rep := regexp.MustCompile(`[$][^[:space:]]*[[:space:]]?\z?`)
	for w := rep.FindString(s); w != ""; w = rep.FindString(s) {
		r := regexp.MustCompile("[$]" + w[1:])
		s = r.ReplaceAllString(s, f(w[1:]))
	}
	return s
}

func main() {
	str := "$The $Go Programming $Language"
	fmt.Fprintln(stdout, str)
	str = expand(str, strings.ToUpper)
	fmt.Fprintln(stdout, str)
}
