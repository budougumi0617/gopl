package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/budougumi0617/GoTraining/ch02/ex01"
)

var stdout io.Writer = os.Stdout // modified during testing
var stderr io.Writer = os.Stderr // modified during testing
var stdin io.Reader = os.Stdin   // modified during testing

func man() {
	fmt.Fprintln(stdout, "\"unitconv NUMBER...\"Convert temperature, length, or weight.")
}

func showTemp(n float64) {
	f := tempconv.Fahrenheit(n)
	c := tempconv.Celsius(n)
	fmt.Fprintf(stdout, "%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func showLength(n float64) {
	f := Feet(n)
	m := Meter(n)
	fmt.Fprintf(stdout, "%s = %s, %s = %s\n",
		f, FToM(f), m, MToF(m))
}

func showWeight(n float64) {
	p := Pond(n)
	kg := Kilogram(n)
	fmt.Fprintf(stdout, "%s = %s, %s = %s\n",
		p, PToKG(p), kg, KGToP(kg))
}

func manual() {
	scanner := bufio.NewScanner(stdin)
	fmt.Fprintln(stdout, "Ctrl+C Terminate this command.\nInput NUMBER")
	for scanner.Scan() {
		convert(strings.Split(scanner.Text(), " "))
		fmt.Fprintln(stdout, "Input NUMBER")
	}
}

func convert(nums []string) {
	fmt.Printf("Enter convert\n")
	for _, arg := range nums {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(stderr, "unitconv: %v\n", err)
			man()
			os.Exit(1)
		}
		showTemp(t)
		showWeight(t)
		showLength(t)
		fmt.Fprintln(stdout, "---------------------------")
	}
}

func main() {
	if len(os.Args) == 1 {
		man()
		manual()
		return
	}
	convert(os.Args[1:])
}
