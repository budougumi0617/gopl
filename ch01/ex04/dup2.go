package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout // modified during testing
var in io.Reader = os.Stdin   // modified during testing
var err = os.Stderr           // modified during testing

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Fprintln(out, "CTRL+C Finish input")
		countLines("stdout", in, counts)
	} else {
		for _, arg := range files {
			f, error := os.Open(arg)
			if error != nil {
				fmt.Fprintf(err, "dup2: %v\n", error)
				continue
			}
			countLines(arg, f, counts)
			f.Close()
		}
	}
	fmt.Fprintf(out, "\t%v\n", counts)
	if len(counts) == 0 {
		return
	}
	for line, e := range counts {
		fmt.Fprintf(out, "%s\t%d\n", line, len(e))
		for _, file := range e {
			fmt.Fprintf(out, "\t%s\n", file)
		}
	}
}

func countLines(arg string, f io.Reader, counts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {

		if _, ok := counts[input.Text()]; ok != true {

			counts[input.Text()] = make([]string, 0)
			counts[input.Text()] = append(counts[input.Text()], arg)
		} else {

			counts[input.Text()] = append(counts[input.Text()], arg)

		}
	}

}
