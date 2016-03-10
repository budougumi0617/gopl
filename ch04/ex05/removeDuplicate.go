package main

import (
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout // modified during testing

func main() {
	s := []string{"test1", "sss", "sss", "sss", "a", "ttt", "ttt", "test2"}
	fmt.Fprintln(stdout, s)
	s = removeDuplicate(s)
	fmt.Fprintln(stdout, s)
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]

}

func removeDuplicate(slice []string) []string {
	for i := 0; i < len(slice)-1; {
		if slice[i] == slice[i+1] {
			slice = remove(slice, i+1)
		} else {
			i++
		}
	}
	return slice
}
