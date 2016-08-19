// Copyright 2016 budougumi0617 All Rights Reserved.

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 5000, 5000
	)
	for num := 1; num <= runtime.NumCPU(); num++ {
		img := image.NewRGBA(image.Rect(0, 0, width, height))

		runtime.GOMAXPROCS(num)
		wg := sync.WaitGroup{}
		start := time.Now()

		yranges := make([][]int, num)
		for i := range yranges {
			yranges[i] = make([]int, 0)
		}
		for y := 0; y < height; y++ {
			i := y % num
			yranges[i] = append(yranges[i], y)
		}
		for i := 0; i < num; i++ {
			wg.Add(1)
			go func(yrange []int) {
				for _, py := range yrange {
					y := float64(py)/height*(ymax-ymin) + ymin
					for px := 0; px < width; px++ {
						x := float64(px)/width*(xmax-xmin) + xmin
						z := complex(x, y)
						// Image point (px, py) represents complex value z.
						img.Set(px, py, mandelbrot(z))
					}
				}
				wg.Done()
			}(yranges[i])
		}
		wg.Wait()
		png.Encode(os.Stdout, img) // NOTE: ignoring errors
		log.Printf("\nMAX CPU NUM %d\nUSE CPU NUM %d\nCalculated time : %v\n", runtime.NumCPU(), num, time.Since(start))
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
