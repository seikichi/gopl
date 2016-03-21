package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		visualize(w, func(x, y float64) float64 {
			r := math.Hypot(x, y)
			return math.Sin(r) / r
		})
	}
	http.HandleFunc("/", handler)
	fmt.Fprintf(os.Stderr, "Serving svg visualizer on localhost port 8000 ...\n")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func visualize(w io.Writer, f func(float64, float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	min, max := calcMinMaxHeight(f)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ab := corner(i+1, j, f)
			bx, by, bb := corner(i, j, f)
			cx, cy, cb := corner(i, j+1, f)
			dx, dy, db := corner(i+1, j+1, f)
			if !ab || !bb || !cb || !db {
				continue
			}
			fill := computeColorCode(i, j, min, max, f)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, fill)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func calcMinMaxHeight(f func(float64, float64) float64) (float64, float64) {
	min := math.Inf(+1)
	max := math.Inf(-1)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			z := f(x, y)

			if math.IsNaN(z) || math.IsInf(z, 0) {
				continue
			}

			min = math.Min(min, z)
			max = math.Max(max, z)
		}
	}
	return min, max
}

func computeColorCode(i, j int, min, max float64, f func(float64, float64) float64) string {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	coeff := (z - min) / (max - min)
	red := uint8(255 * coeff)
	blue := uint8(255 * (1 - coeff))

	return fmt.Sprintf("#%02x00%02x", red, blue)
}

func corner(i, j int, f func(float64, float64) float64) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	finite := true
	z := f(x, y)

	if math.IsNaN(z) || math.IsInf(z, 0) {
		finite = false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, finite
}
