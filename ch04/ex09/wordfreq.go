package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout // modified during testing

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(stdout, "Execute with the name of text file")
		return
	}
	fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(stdout, "Failed file open")
		return
	}
	words := make(map[string]int) // a map of strings
	input := bufio.NewScanner(fp)
	input.Split(bufio.ScanWords) // split by space
	for input.Scan() {
		word := input.Text()
		words[word]++
	}
	for k, v := range words {
		fmt.Fprintf(stdout, "\"%s\" was found %d times\n", k, v)
	}
}
