// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTopoSort(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("----------------------------\nTry count %d\n", i+1)
		corses := topoSort(prereqs)
		for i, course := range corses {
			fmt.Printf("%d:\t%s\n", i+1, course)
		}
		fmt.Printf("\nCheck answer\n")
		for n, corse := range corses {
			studied := make(map[string]bool)
			for _, c := range corses[:n] {
				studied[c] = true
			}
			requests := prereqs[corse]
			for request := range requests {
				fmt.Printf("%q requests %q\n", corse, request)
				if _, ok := studied[request]; ok != true {
					t.Errorf("%s is requested to alredy study %s", corse, request)
				}

			}
		}
	}

}

func TestMain(t *testing.T) {
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if len(got) == 0 {
		t.Errorf("No output")
	}
}
