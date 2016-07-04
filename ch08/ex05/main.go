// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var para int

func main() {
	flag.IntVar(&para, "P", 1, "parallel")
	flag.Parse()

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	done := make(chan struct{})

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < para; i++ {
		go func(i int) {
			for py := 0; py < height; py++ {
				if py%para != i {
					continue
				}

				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					img.Set(px, py, mandelbrot(z))
				}
			}
			done <- struct{}{}
		}(i)
	}

	for i := 0; i < para; i++ {
		<-done
	}

	png.Encode(os.Stdout, img)
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
