package main

import (
	"flag"
	"fmt"
	"math"
	"os"
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

var function = flag.String("function", "eggbox", "the function to visualize (eggbox, monguls or saddle)")

var f func(x, y float64) float64

func init() {
	flag.Parse()

	switch *function {
	case "monguls":
		f = func(x, y float64) float64 {
			return math.Sin(-x) * math.Pow(1.5, -math.Hypot(x, y))
		}
	case "eggbox":
		f = func(x, y float64) float64 {
			return math.Pow(2, math.Sin(y)) * math.Pow(2, math.Sin(x)) / 12
		}
	case "saddle":
		f = func(x, y float64) float64 {
			return math.Sin(x*y/10) / 10
		}
	default:
		flag.Usage()
		os.Exit(-1)
	}
}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	min, max := calcMinMaxHeight()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ab := corner(i+1, j)
			bx, by, bb := corner(i, j)
			cx, cy, cb := corner(i, j+1)
			dx, dy, db := corner(i+1, j+1)
			if !ab || !bb || !cb || !db {
				continue
			}
			fill := computeColorCode(i, j, min, max)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, fill)
		}
	}
	fmt.Println("</svg>")
}

func calcMinMaxHeight() (float64, float64) {
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

func computeColorCode(i, j int, min, max float64) string {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	coeff := (z - min) / (max - min)
	red := uint8(255 * coeff)
	blue := uint8(255 * (1 - coeff))

	return fmt.Sprintf("#%02x00%02x", red, blue)
}

func corner(i, j int) (float64, float64, bool) {
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
