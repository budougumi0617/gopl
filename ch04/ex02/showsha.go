// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"crypto"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout // modified during testing
var stderr io.Writer = os.Stderr // modified during testing
var msg = "Choose hash. SHA256, SHA384, or SHA512"

func main() {
	showsha(os.Args)
}

func showsha(args []string) {
	types := map[string]crypto.Hash{
		"SHA256": crypto.SHA256,
		"SHA384": crypto.SHA384,
		"SHA512": crypto.SHA512,
	}

	// Parse option
	flags := flag.NewFlagSet("showsha", flag.ContinueOnError)
	flags.SetOutput(stderr)
	var hash string
	flags.StringVar(&hash, "hash", "SHA256", msg)
	flags.Parse(args[1:])

	for _, arg := range flags.Args() {
		switch types[hash] {
		case crypto.SHA256:
			fmt.Fprintf(stdout, "%x\n", sha256.Sum256([]byte(arg)))
		case crypto.SHA384:
			fmt.Fprintf(stdout, "%x\n", sha512.Sum384([]byte(arg)))
		case crypto.SHA512:
			fmt.Fprintf(stdout, "%x\n", sha512.Sum512([]byte(arg)))
		default:
			fmt.Fprintln(stdout, "Invalid Option. "+msg)
			return
		}
	}
}
