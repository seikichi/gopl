package main

import "testing"

var maxTests = []struct {
	in   []int
	want int
}{
	{[]int{1}, 1},
	{[]int{1, 3, 2}, 3},
}

var minTests = []struct {
	in   []int
	want int
}{
	{[]int{1}, 1},
	{[]int{1, 3, 2}, 1},
}

func TestMax(t *testing.T) {
	for _, tt := range maxTests {
		got, _ := max(tt.in...)
		if got != tt.want {
			t.Errorf("max(%v) = %v; want %v", tt.in, got, tt.want)
		}
	}
}

func TestMax2(t *testing.T) {
	for _, tt := range maxTests {
		if got := max2(tt.in[0], tt.in[1:]...); got != tt.want {
			t.Errorf("max2(%v) = %v; want %v", tt.in, got, tt.want)
		}
	}
}

func TestMin(t *testing.T) {
	for _, tt := range minTests {
		got, _ := min(tt.in...)
		if got != tt.want {
			t.Errorf("min(%v) = %v; want %v", tt.in, got, tt.want)
		}
	}
}

func TestMin2(t *testing.T) {
	for _, tt := range minTests {
		if got := min2(tt.in[0], tt.in[1:]...); got != tt.want {
			t.Errorf("min2(%v) = %v; want %v", tt.in, got, tt.want)
		}
	}
}
