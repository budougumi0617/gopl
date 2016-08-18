// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"fmt"

	"github.com/budougumi0617/gotraining/ch10/ex02/unarchive"
	_ "github.com/budougumi0617/gotraining/ch10/ex02/unarchive/tar"
	_ "github.com/budougumi0617/gotraining/ch10/ex02/unarchive/zip"
)

func main() {
	fmt.Printf("%s\n", "test")
	unarchive.List()
}
