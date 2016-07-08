// Copyright 2016 budougumi0617 All Rights Reserved.
// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

var stdin io.Reader = os.Stdin // modified during testing

func main() {
	dec := xml.NewDecoder(stdin)
	var stack []string   // stack of element names
	var attrs [][]string // stack of element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
			v := []string{}
			for _, a := range tok.Attr {
				v = append(v, a.Value)
			}
			attrs = append(attrs, v) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
			attrs = attrs[:len(attrs)-1] // pop
		case xml.CharData:
			if containsAll(attrs, stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(os.Args[1:], " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(attrs [][]string, x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		} else {
			for _, v := range attrs[0] {
				if v == y[0] {
					y = y[1:]
				}
			}
		}
		x = x[1:]
		attrs = attrs[1:]
	}
	return false
}
