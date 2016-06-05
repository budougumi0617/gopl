package main

import "io"

// LimitReader returns a Reader that reads from r
// but stops with EOF after n bytes.
// The underlying implementation is a *LimitedReader.
// https://golang.org/src/io/io.go
func LimitReader(r io.Reader, n int64) io.Reader { return &LimitedReader{r, n} }

// A LimitedReader reads from R but limits the amount of
// data returned to just N bytes. Each call to Read
// updates N to reflect the new amount remaining.
// https://golang.org/src/io/io.go
type LimitedReader struct {
	R io.Reader // underlying reader
	N int64     // max bytes remaining
}

// https://golang.org/src/io/io.go
func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}
