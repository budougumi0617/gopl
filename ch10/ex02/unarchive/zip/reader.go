// Copyright 2016 budougumi0617 All Rights Reserved.

// Package zip is based my unzip command.
package zip

import (
	"io"

	"github.com/budougumi0617/gotraining/ch10/ex02/unarchive"
)

func read(string) (io.Reader, error) {
	return nil, nil
}
func unpacking(io.Reader) error {
	return nil
}

func init() {
	unarchive.RegisterFormat("zip", "", 0, read, unpacking)
}
