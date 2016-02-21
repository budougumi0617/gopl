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
	counts := make(map[string]map[string]int)
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
	//fmt.Fprintf(out, "\t%v\n", counts)
	if len(counts) == 0 {
		return
	}

	for line, e := range counts {
		result := 0
		for _, c := range e {
			result += c
		}
		if result < 2 {
			continue
		}
		fmt.Fprintf(out, "%s\t%d\n", line, result)
		for file := range e {
			fmt.Fprintf(out, "\t%s\n", file)
		}
	}
}

func countLines(arg string, f io.Reader, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {

		if _, ok := counts[input.Text()]; ok != true {
			counts[input.Text()] = make(map[string]int, 0)
		}
		counts[input.Text()][arg]++
	}

}
