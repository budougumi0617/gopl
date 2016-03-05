package main

import (
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var stdout io.Writer = os.Stdout // modified during testing

func main() {
	s := "Hello　世界"
	fmt.Fprintln(stdout, s)
	s = string(replaceSpace([]byte(s)))
	fmt.Fprintln(stdout, s)
}

func replaceSpace(s []byte) []byte {
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRune(s[i:])
		if unicode.IsSpace(r) {
			copy(s[i+1:], s[i+size:])
			s[i] = ' '
			s = s[:len(s)-size+1] // len(slice)-len(unicode space)+len(ASCII space)
			i++
		} else {
			i += size
		}
	}
	return s
}
