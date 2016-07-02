// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"io/ioutil"
	"os"
)

func ExampleMain() {
	source, _ := ioutil.ReadFile("./ex17.xml")
	stdin = bytes.NewBuffer(source)
	os.Args = []string{"test", "div", "foo", "url"}
	main()
	source, _ = ioutil.ReadFile("./ex17.xml")
	stdin = bytes.NewBuffer(source)
	os.Args = []string{"test", "div", "div2", "url"}
	main()

	// Output:
	// div foo url: OK!!
	// div div2 url: OK!!
}
