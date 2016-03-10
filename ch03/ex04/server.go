package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var stdout io.Writer = os.Stdout // modified during testing
var stderr io.Writer = os.Stderr // modified during testing

var xyscale float64 // pixels per x or y unit
var zscale float64  // pixels per z unit
var width float64
var height float64
var red, green, blue uint64

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func surface(width float64, height float64, r uint8, g uint8, b uint8) string {
	log.Printf("width = %v, height = %v, r = %v, g = %v, b = %v\n", width, height, red, green, blue)
	xyscale = width / 2 / xyrange // pixels per x or y unit
	zscale = height * 0.4         // pixels per z unit
	s := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", int(width), int(height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aok := corner(i+1, j)
			bx, by, bok := corner(i, j)
			cx, cy, cok := corner(i, j+1)
			dx, dy, dok := corner(i+1, j+1)
			if aok && bok && cok && dok {
				s += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=\"rgb(%d,%d,%d)\"/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
			}
		}
	}
	s += "</svg>"
	return s
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	if math.IsNaN(z) {
		fmt.Fprintf(stderr, "f(%d, %d) was NaN.\n", i, j)
		return 0, 0, false
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func parsequery(u *url.URL) (width float64, height float64) {
	width, err := strconv.ParseFloat(u.Query().Get("width"), 64)
	if err != nil {
		width = 600
	}
	height, err = strconv.ParseFloat(u.Query().Get("height"), 64)
	if err != nil {
		height = 320
	}

	red, err = strconv.ParseUint(u.Query().Get("r"), 10, 8)
	if err != nil {
		red = 255
	}
	green, err = strconv.ParseUint(u.Query().Get("g"), 10, 8)
	if err != nil {
		green = 255
	}
	blue, err = strconv.ParseUint(u.Query().Get("b"), 10, 8)
	if err != nil {
		blue = 255
	}
	return
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		width, height = parsequery(r.URL)
		w.Header().Set("Content-Type", "image/svg+xml")
		fmt.Fprintf(w, surface(width, height, uint8(red), uint8(green), uint8(blue)))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
