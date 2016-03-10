// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var stdout io.Writer = os.Stdout // modified during testing

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		x, y, s := parsequery(r.URL)
		w.Header().Set("Content-Type", "image/png")
		png.Encode(w, createImage(x, y, s))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func parsequery(u *url.URL) (x, y, s float64) {
	x, err := strconv.ParseFloat(u.Query().Get("x"), 64)
	if err != nil {
		x = 0
	}
	y, err = strconv.ParseFloat(u.Query().Get("y"), 64)
	if err != nil {
		y = 0
	}
	s, err = strconv.ParseFloat(u.Query().Get("scale"), 64)
	if err != nil {
		s = 2
	}
	return
}
func createImage(ix, iy, s float64) *image.RGBA {
	const (
		width, height = 1024, 1024
	)
	var xmin, ymin, xmax, ymax float64 = -s, -s, +s, +s
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x+ix, y+iy)
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
