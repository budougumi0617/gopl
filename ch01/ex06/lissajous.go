// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var out io.Writer = os.Stdout // modified during testing
var pallete = []color.Color{color.White, color.Black, color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}}

const (
	whiteIndex = 0 // First palette color
	blackIndex = 1 // Next palette color
	redIndex   = 2 // Next palette color
	greanIndex = 3 // Next palette color
	blueIndex  = 4 // Next palette color
	indexSize  = 5 // Amount of palette color
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(out)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5
		res    = 0.001
		size   = 100
		nframe = 64
		delay  = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframe}
	phase := 0.0
	for i := 0; i < nframe; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(indexSize)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
