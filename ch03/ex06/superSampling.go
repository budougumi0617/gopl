// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"os"
)

var stdout io.Writer = os.Stdout // modified during testing
const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	qwidth, qheight        = width * 4, height * 4
)

func main() {
	img := quadrupleImage()
	img = superSampling(img)
	png.Encode(stdout, img) // NOTE: ignoring errors
}

func quadrupleImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, qwidth, qheight))
	for py := 0; py < qheight; py++ {
		y := float64(py)/qheight*(ymax-ymin) + ymin
		for px := 0; px < qwidth; px++ {
			x := float64(px)/qwidth*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := 255 - contrast*n
			g := contrast * n
			b := g
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.Black
}

func superSampling(s image.Image) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			// Image point (px, py) represents complex value z.
			img.Set(px, py, averaging(px*4, py*4, s))
		}
	}
	return img
}

func averaging(x, y int, s image.Image) color.Color {
	// weight coefficient
	//   1 2 1
	//   2 4 2
	//   1 2 1
	var filter = []struct {
		fx, fy int
		fw     uint32
	}{
		{-1, -1, 1},
		{0, -1, 2},
		{1, -1, 1},
		{-1, 0, 2},
		{0, 0, 4},
		{1, -0, 2},
		{-1, 1, 1},
		{0, 1, 2},
		{1, 1, 1},
	}
	var r, g, b uint32
	for _, f := range filter {
		cr, cg, cb, _ := s.At(x+f.fx, y+f.fy).RGBA()
		r += cr * f.fw
		g += cg * f.fw
		b += cb * f.fw
	}
	return color.RGBA{uint8(r / 9 >> 8), uint8(g / 9 >> 8), uint8(b / 9 >> 8), 255}
}
