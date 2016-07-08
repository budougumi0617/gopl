// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import "fmt"

func Example_length() {
	{
		f := MToF(10.0)                      // "212" m
		fmt.Printf("%v\n", f.String())       // "212" m
		fmt.Printf("%g\n", FToM(f))          // "10" m
		fmt.Printf("%g\n", FToM(0))          // "0" m
		fmt.Printf("%g\n", MToF(0))          // "0" f
		fmt.Printf("%v\n", FToM(3).String()) // "0.9146341463414634" m
	}

	// Output:
	// 32.8f
	// 10
	// 0
	// 0
	// 0.9146341463414634m
}
