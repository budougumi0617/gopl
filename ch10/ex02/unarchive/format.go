// Copyright 2016 budougumi0617 All Rights Reserved.

// Package unarchive is based my unarchive command.
package unarchive

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// ErrFormat indicates that decoding encountered an unknown format.
var ErrFormat = errors.New("unarchive: unknown format")

// A format holds an image format's name, magic header and how to unarchive it.
type format struct {
	name, magic string
	// Location of the magic number deffers from each archive format.
	offset    int
	unpacking func(string) error
}

// Formats is the list of registered formats.
var formats []format

// RegisterFormat registers an archive format
func RegisterFormat(name, magic string, offset int,
	unpacking func(string) error) {
	formats = append(formats, format{name, magic, offset, unpacking})
}

// Match reports whether magic matches b. Magic may contain "?" wildcards.
func match(magic string, b []byte) bool {
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c && magic[i] != '?' {
			return false
		}
	}
	return true
}

// Sniff determines the format of reader's data.
func sniff(reader io.Reader) format {
	r := bufio.NewReader(reader)
	for _, f := range formats {
		b, err := r.Peek(f.offset + len(f.magic))
		if err == nil && match(f.magic, b[f.offset:]) {
			return f
		}
	}
	return format{}
}

// Unarchive extracts arhive file.
func Unarchive(filename string) error {
	file, _ := os.Open(filename)
	defer file.Close()
	f := sniff(file)
	if f.unpacking == nil {
		return ErrFormat
	}
	fmt.Printf("Use %s\n", f.name)
	if err := f.unpacking(filename); err != nil {
		log.Println("yshimizu in format")
		return err
	}
	return nil
}

// List shows all format in formats
func List() {
	fmt.Println("Support format is below:")
	for _, f := range formats {
		fmt.Printf("%s\n", f.name)
	}
}
