// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import "os"

func Example_main() {
	os.Args = []string{"go.out", "unsafe", "sync"}
	main()
	// Output:
	// [internal/race runtime runtime/internal/atomic runtime/internal/sys sync/atomic unsafe]
}
