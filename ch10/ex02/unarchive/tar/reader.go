// Copyright 2016 budougumi0617 All Rights Reserved.

// Package tar is based my untar command.
package tar

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
	unarchive.RegisterFormat("tar", "ustar", 257, nil, nil)
}
