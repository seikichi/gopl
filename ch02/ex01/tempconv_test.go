package tempconv

import (
	"math"
	"testing"
)

const eps = 1e-10

var tests = []struct {
	k Kelvin
	c Celsius
	f Fahrenheit
}{
	{0, -273.15, -459.67},
	{1, -272.15, -457.87},
}

func TestKToC(t *testing.T) {
	for _, tt := range tests {
		if got := KToC(tt.k); math.Abs(float64(got-tt.c)) > eps {
			t.Errorf("KToC(%q) = %q; want %q", tt.k, got, tt.c)
		}
	}
}

func TestKToF(t *testing.T) {
	for _, tt := range tests {
		if got := KToF(tt.k); math.Abs(float64(got-tt.f)) > eps {
			t.Errorf("KToF(%q) = %q; want %q", tt.k, got, tt.f)
		}
	}
}

func TestCToK(t *testing.T) {
	for _, tt := range tests {
		if got := CToK(tt.c); math.Abs(float64(got-tt.k)) > eps {
			t.Errorf("CToK(%q) = %q; want %q", tt.c, got, tt.k)
		}
	}
}

func TestFToK(t *testing.T) {
	for _, tt := range tests {
		if got := FToK(tt.f); math.Abs(float64(got-tt.k)) > eps {
			t.Errorf("FToK(%q) = %q; want %q", tt.f, got, tt.k)
		}
	}
}
