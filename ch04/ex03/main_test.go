package main

import "testing"

var tests = []struct {
	in, want [size]int
}{
	{
		[size]int{1, 1, 1, 1, 1, 1, 1, 1},
		[size]int{1, 1, 1, 1, 1, 1, 1, 1},
	},
	{
		[size]int{1, 2, 3, 4, 5, 6, 7, 8},
		[size]int{8, 7, 6, 5, 4, 3, 2, 1},
	},
}

func TestReverse(t *testing.T) {
	for _, tt := range tests {
		orig := tt.in
		reverse(&tt.in)
		if tt.in != tt.want {
			t.Errorf("reverse(%v) changes the argument to %v, but want %v", orig, tt.in, tt.want)
		}
	}
}
