package main

import "fmt"

func Example_weight() {
	{
		kg := PToKG(10.0)
		fmt.Printf("%v\n", kg.String())
		fmt.Printf("%g\n", KGToP(kg))
		fmt.Printf("%g\n", KGToP(0))
		fmt.Printf("%g\n", PToKG(0))
		fmt.Printf("%v\n", KGToP(3).String())
	}

	// Output:
	// 4.535970244035199kg
	// 10
	// 0
	// 0
	// 6.6138pond
}
