package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var supersampling = flag.Bool("supersampling", true, "enables supersampling")

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	flag.Parse()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	if *supersampling {
		renderWithSupersampling(img)
	} else {
		render(img)
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func render(img *image.RGBA) {
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
}

func renderWithSupersampling(img *image.RGBA) {
	const samplings = 4
	dy := [samplings]float64{+0.25, +0.25, -0.25, -0.25}
	dx := [samplings]float64{+0.25, -0.25, +0.25, -0.25}

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			r, g, b, a := 0, 0, 0, 0
			for i := 0; i < samplings; i++ {
				sy := (dy[i]+float64(py))/height*(ymax-ymin) + ymin
				sx := (dx[i]+float64(px))/width*(xmax-xmin) + xmin
				sc := mandelbrot(complex(sx, sy))
				mr, mg, mb, ma := sc.RGBA()
				r += int(mr)
				g += int(mg)
				b += int(mb)
				a += int(ma)
			}
			c := color.RGBA{
				uint8((r >> 8) / samplings),
				uint8((g >> 8) / samplings),
				uint8((b >> 8) / samplings),
				uint8((a >> 8) / samplings),
			}
			img.Set(px, py, c)
		}
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			if n%2 == 0 {
				return color.RGBA{0xff, 0x80, 0x00, 0xff}
			}
			return color.RGBA{0x00, 0x80, 0xff, 0xff}
		}
	}
	return color.Black
}
