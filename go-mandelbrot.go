// Draws the Mandelbrot set https://en.wikipedia.org/wiki/Mandelbrot_set
// Author: Christophe Pallier <christophe@pallier.org>
// LICENSE: GPL-3
package main

import (
	"flag"
	"image/color"
	"strconv"

	"github.com/crazy3lf/colorconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

// good values for xmin, xmax=-2.0, 0.47   ymin=-1.12 ymax=1.12

var (
	w = 1024
	h = 1024

	xmin = -2.00
	xmax = 0.47
	ymin = -1.12
	ymax = 1.12

	maxIterations = 512
	palette       []color.Color
)

func createBWPalette(max int) {
	palette = make([]color.Color, max)
	for i := 0; i < max; i++ {
		palette[i] = color.RGBA{uint8(i % 256),
			uint8(i % 256),
			uint8(i % 256),
			0xff}
	}
}

func createColorPalette(max int) {
	var h, s, v float64
	var r, g, b uint8
	var err error

	palette = make([]color.Color, max)

	for i := 0; i < max; i++ {
		h = 360.0 * float64(i) / float64(max)
		s = 1.0
		if i == max-1 {
			v = 0.0
		} else {
			v = 1.0
		}

		r, g, b, err = colorconv.HSVToRGB(h, s, v)
		if err != nil {
			panic("hsv2rgb")
		}

		palette[i] = color.RGBA{r, g, b, 0xff}
	}
}

// translate from window coordinate to mandelspace coordinates
func mapToMandelSpace(px, py, w, h int) complex128 {
	x := xmin + (float64(px)/float64(w))*(xmax-xmin)
	y := ymin + (float64(py)/float64(h))*(ymax-ymin)
	return complex(x, y)
}

func niter(z0 complex128, MaxIterations int) int {
	var i int
	var z complex128
	for i = 0; i < maxIterations-1; i++ {
		z = (z * z) + z0
		if float64(real(z)*real(z)+imag(z)*imag(z)) > 4 {
			break
		}
	}
	return i
}

func getPixelColor(px, py, w, h int) color.Color {
	z := mapToMandelSpace(px, py, w, h)
	i := niter(z, maxIterations)
	return palette[i]
}

func main() {
	var err error

	flag.Parse()
	coords := flag.Args()
	if len(coords) == 4 {
		xmin, err = strconv.ParseFloat(coords[0], 64)
		if err != nil {
			panic("coords[0] should be a float")
		}
		xmax, err = strconv.ParseFloat(coords[1], 64)
		if err != nil {
			panic("coords[1] should be a float")
		}
		ymin, err = strconv.ParseFloat(coords[2], 64)
		if err != nil {
			panic("coords[2] should be a float")
		}
		ymax, err = strconv.ParseFloat(coords[3], 64)
		if err != nil {
			panic("coords[3] should be a float")
		}

	}

	a := app.New()
	win := a.NewWindow("Mandelbrot")

	//win.SetFullScreen(true)

	createColorPalette(maxIterations)

	raster := canvas.NewRasterWithPixels(getPixelColor)
	win.SetContent(raster)

	win.Resize(fyne.NewSize(float32(w), float32(h)))

	win.ShowAndRun()
}
