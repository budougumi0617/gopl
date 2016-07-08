// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import "fmt"

// Feet For F
type Feet float64

// Meter For M
type Meter float64

func (f Feet) String() string  { return fmt.Sprintf("%gf", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

// FToM converts Feet to Meter.
func FToM(f Feet) Meter {
	if f == 0.0 {
		return Meter(0.0)
	}
	return Meter(f / 3.28)
}

// MToF converts Meter to Feet.
func MToF(m Meter) Feet {
	if m == 0.0 {
		return Feet(0.0)
	}
	return Feet(m * 3.28)
}
