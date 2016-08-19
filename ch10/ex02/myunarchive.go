// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"fmt"
	"os"

	"github.com/budougumi0617/gotraining/ch10/ex02/unarchive"
	_ "github.com/budougumi0617/gotraining/ch10/ex02/unarchive/tar"
	_ "github.com/budougumi0617/gotraining/ch10/ex02/unarchive/zip"
)

func main() {
	unarchive.List()
	for _, f := range os.Args[1:] {
		err := unarchive.Unarchive(f)
		if err != nil {
			fmt.Printf("%q", err)
		}
	}
}
