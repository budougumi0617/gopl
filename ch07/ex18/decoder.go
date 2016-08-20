// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// Node is CharData or Element.
type Node interface{}

// CharData represent text strings
type CharData string

// Element represent named elements and their attributes
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func decodeXML(dec *xml.Decoder) (Node, error) {
	var stack []*Element // stack of element
	var root Node
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "decoder: %v\n", err)
			return root, err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			nelm := &Element{tok.Name, tok.Attr, []Node{}}
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, nelm)
			}
			stack = append(stack, nelm) // push
			if root == nil {
				root = nelm
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if len(stack) == 0 {
				continue
			}
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, CharData(tok))
		}
	}
	return root, nil
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	node, err := decodeXML(dec)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Result:\n%#v\n", node)
}
