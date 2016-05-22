package main

import (
	"fmt"
	"io"
	"os"
)

// WriterWithCounter has counter
type WriterWithCounter struct {
	writer io.Writer
	c      int64
}

func (w *WriterWithCounter) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)
	w.c += int64(n)
	return
}

// CounterWriter returns a new `Writer` that wraps the original, and a pointer to an `int64`
func CounterWriter(w io.Writer) (io.Writer, *int64) {
	wwc := new(WriterWithCounter)
	wwc.writer = w
	wwc.c = 0
	return wwc, &(wwc.c)
}

func main() {
	out := os.Stdout
	w, c := CounterWriter(out)
	w.Write([]byte("test input\n"))
	w.Write([]byte("test input\n"))
	fmt.Printf("Byte count %d\n", *c)
}
