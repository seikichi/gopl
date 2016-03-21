package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
)

var dataType = flag.String("type", "complex128", "complex64, complex128, big.Float, or big.Rat")

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	flag.Parse()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	switch *dataType {
	case "complex64":
		renderComplex64(img)
	case "complex128":
		renderComplex128(img)
	case "big.Float":
		renderFloat(img)
	case "big.Rat":
		renderRat(img)
	default:
		flag.Usage()
		os.Exit(-1)
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

const iterations = 200
const contrast = 15

func renderComplex64(img *image.RGBA) {
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrotComplex64(z))
		}
	}
}

func mandelbrotComplex64(z complex64) color.Color {
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		rv, iv := real(v), imag(v)
		if rv*rv+iv*iv > 4 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func renderComplex128(img *image.RGBA) {
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrotComplex128(z))
		}
	}
}

func mandelbrotComplex128(z complex128) color.Color {
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		rv, iv := real(v), imag(v)
		if rv*rv+iv*iv > 4 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func renderFloat(img *image.RGBA) {
	var yminF, ymaxMinF, heightF big.Float
	yminF.SetInt64(ymin)
	ymaxMinF.SetInt64(ymax - ymin)
	heightF.SetInt64(height)

	var xminF, xmaxMinF, widthF big.Float
	xminF.SetInt64(xmin)
	xmaxMinF.SetInt64(xmax - xmin)
	widthF.SetInt64(width)

	var y, x big.Float
	for py := int64(0); py < height; py++ {
		// y := float64(py)/height*(ymax-ymin) + ymin
		y.SetInt64(py)
		y.Quo(&y, &heightF)
		y.Mul(&y, &ymaxMinF)
		y.Add(&y, &yminF)

		for px := int64(0); px < width; px++ {
			// x := float64(px)/width*(xmax-xmin) + xmin
			x.SetInt64(px)
			x.Quo(&x, &widthF)
			x.Mul(&x, &xmaxMinF)
			x.Add(&x, &xminF)

			c := mandelbrotFloat(&x, &y)
			if c == nil {
				c = color.Black
			}
			img.Set(int(px), int(py), c)
		}
	}
}

func mandelbrotFloat(a, b *big.Float) color.Color {
	var x, y, nx, ny, x2, y2, f2, f4, r2, tmp big.Float
	f2.SetInt64(2)
	f4.SetInt64(4)
	x.SetInt64(0)
	y.SetInt64(0)

	defer func() { recover() }()

	for n := uint8(0); n < iterations; n++ {
		// Not update x2 and y2
		// because they are already updated in the previous loop
		nx.Sub(&x2, &y2)
		nx.Add(&nx, a)

		tmp.Mul(&x, &y)
		ny.Mul(&f2, &tmp)
		ny.Add(&ny, b)

		x.Set(&nx)
		y.Set(&ny)

		x2.Mul(&x, &x)
		y2.Mul(&y, &y)
		r2.Add(&x2, &y2)

		if r2.Cmp(&f4) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func renderRat(img *image.RGBA) {
	var yminR, ymaxMinR, heightR big.Rat
	yminR.SetInt64(ymin)
	ymaxMinR.SetInt64(ymax - ymin)
	heightR.SetInt64(height)

	var xminR, xmaxMinR, widthR big.Rat
	xminR.SetInt64(xmin)
	xmaxMinR.SetInt64(xmax - xmin)
	widthR.SetInt64(width)

	var y, x big.Rat
	for py := int64(0); py < height; py++ {
		// y := float64(py)/height*(ymax-ymin) + ymin
		y.SetInt64(py)
		y.Quo(&y, &heightR)
		y.Mul(&y, &ymaxMinR)
		y.Add(&y, &yminR)

		for px := int64(0); px < width; px++ {
			// x := float64(px)/width*(xmax-xmin) + xmin
			x.SetInt64(px)
			x.Quo(&x, &widthR)
			x.Mul(&x, &xmaxMinR)
			x.Add(&x, &xminR)

			c := mandelbrotRat(&x, &y)
			if c == nil {
				c = color.Black
			}
			img.Set(int(px), int(py), c)
		}
	}
}

func mandelbrotRat(a, b *big.Rat) color.Color {
	var x, y, nx, ny, x2, y2, f2, f4, r2, tmp big.Rat
	f2.SetInt64(2)
	f4.SetInt64(4)
	x.SetInt64(0)
	y.SetInt64(0)

	defer func() { recover() }()

	for n := uint8(0); n < iterations; n++ {
		// Not update x2 and y2
		// because they are already updated in the previous loop
		nx.Sub(&x2, &y2)
		nx.Add(&nx, a)

		tmp.Mul(&x, &y)
		ny.Mul(&f2, &tmp)
		ny.Add(&ny, b)

		x.Set(&nx)
		y.Set(&ny)

		x2.Mul(&x, &x)
		y2.Mul(&y, &y)
		r2.Add(&x2, &y2)

		if r2.Cmp(&f4) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
