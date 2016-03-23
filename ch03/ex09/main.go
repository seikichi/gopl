package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"net/url"
	"strconv"
)

func getFloat64(q url.Values, key string, d float64) float64 {
	f, err := strconv.ParseFloat(q.Get(key), 64)
	if err != nil {
		return d
	}
	return f
}

func main() {
	port := 8080
	flag.IntVar(&port, "p", 8000, "listening port")
	flag.Parse()

	const width, height = 1024, 1024

	handler := func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		x, y, m := getFloat64(query, "x", 0), getFloat64(query, "y", 0), getFloat64(query, "m", 1)
		log.Printf("render mandelbrot (x = %f, y = %f, m = %f)\n", x, y, m)
		render(x, y, m, width, height, w)
	}
	http.HandleFunc("/", handler)
	log.Printf("Serving lissajous on localhost port %d ...\n", port)
	log.Println("Supported queries: x, y, m (defaults are 0, 0, 1)")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil))
}

func render(x, y, m float64, width, height int, w io.Writer) {
	xmin, ymin, xmax, ymax := x-2/m, y-2/m, x+2/m, y+2/m

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img)
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
