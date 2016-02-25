package main

import (
	"bytes"
	"testing"
)

func TestPrintValues(t *testing.T) {
	outStream := new(bytes.Buffer)
	printValues("0", outStream)
	want := `0.00K = -273.15°C = -459.67°F
0.00°F = 255.37K = -17.78°C
0.00°C = 273.15K = 32.00°F
0.00m = 0.00ft
0.00ft = 0.00m
0.00kg = 0.00lb
0.00lb = 0.00kg

`
	if got := outStream.String(); got != want {
		t.Errorf("printValues(%q, ...) outputs %q; want %q", 0, got, want)
	}
}
