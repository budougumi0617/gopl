package main

import "fmt"

// Pond For P
type Pond float64

// Kilogram For Kg
type Kilogram float64

func (p Pond) String() string      { return fmt.Sprintf("%gpond", p) }
func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }

// PToKG converts Pond to Kilogram.
func PToKG(p Pond) Kilogram {
	if p == 0.0 {
		return Kilogram(0.0)
	}
	return Kilogram(p / 2.2046)
}

// KGToP converts Kilogram to Pond.
func KGToP(kg Kilogram) Pond {
	if kg == 0.0 {
		return Pond(0.0)
	}
	return Pond(kg * 2.2046)
}
