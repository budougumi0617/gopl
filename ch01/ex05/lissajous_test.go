// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"testing"
)

func TestLissajous(t *testing.T) {
	out = new(bytes.Buffer) // captured output
	main()
	got := out.(*bytes.Buffer).String()
	if size := len(got); size == 0 {
		t.Errorf("Result data size = %d", size)
	}
}
