package tempconv

import "fmt"

func Example_one() {
	{
		fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
		boilingF := CToF(BoilingC)
		fmt.Printf("%v\n", boilingF.String())        // "212" °F
		fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
		boilingK := CToK(BoilingC)
		fmt.Printf("%g\n", boilingK-CToK(FreezingC)) // "100" °K
	}

	// Output:
	// 100
	// 212°F
	// 180
	// 100
}

func Example_two() {
	c := FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String

	c = KToC(FreezingK)
	fmt.Println(FreezingK.String()) // "273.15°K"
	fmt.Println(c.String())         // "100°C"
	fmt.Println(c.String())         // "100°C"

	// Output:
	// 100°C
	// 100°C
	// 100°C
	// 100°C
	// 100
	// 100
	// 273.15°K
	// 0°C
	// 0°C
}
