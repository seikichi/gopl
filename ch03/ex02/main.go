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
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ab := corner(i+1, j, f)
			bx, by, bb := corner(i, j, f)
			cx, cy, cb := corner(i, j+1, f)
			dx, dy, db := corner(i+1, j+1, f)
			if ab && bb && cb && db {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, f func(x, y float64) float64) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	finite := true
	// Compute surface height z.
	z := f(x, y)

	if math.IsNaN(z) || math.IsInf(z, 0) {
		finite = false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, finite
}
