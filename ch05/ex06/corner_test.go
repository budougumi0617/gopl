// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"testing"
)

func orgcorner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
func TestCorner(t *testing.T) {
	var tests = []struct {
		x, y int
	}{
		{10, 20},
		{20, 10},
		{0, 0},
	}

	for _, test := range tests {
		sx, sy := corner(test.x, test.y)
		osx, osy := orgcorner(test.x, test.y)
		if sx != osx || sy != osy {
			t.Errorf("Expected:\n sx = %v, sy = %v\nActual:\n sx = %v, sy = %v",
				osx, osy, sx, sy)
		}
	}
}

func TestMain(t *testing.T) {
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if len(got) == 0 {
		t.Errorf("No output")
	}
}
