// Copyright 2016 budougumi0617 All Rights Reserved.
package tempconv

import "fmt"

// Celsius For C
type Celsius float64

// Fahrenheit For F
type Fahrenheit float64

// Kelvin For K
type Kelvin float64

// Temperature type
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
	FreezingK     Kelvin  = 273.15
	BoilingK      Kelvin  = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Fahrenheit.
func CToK(c Celsius) Kelvin { return Kelvin(c + AbsoluteZeroC) }

// KToC converts a Celsius temperature to Fahrenheit.
func KToC(k Kelvin) Celsius { return Celsius(k - FreezingK) }
