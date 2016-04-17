package main

import "testing"

var tests = []struct {
	sep  string
	a    []string
	want string
}{
	{",", []string{}, ""},
	{",", []string{"a"}, "a"},
	{",", []string{"a", "b"}, "a,b"},
}

func TestJoin(t *testing.T) {
	for _, tt := range tests {
		if got := join(tt.sep, tt.a...); got != tt.want {
			t.Errorf("join(%v, %v) = %v; want %v", tt.sep, tt.a, got, tt.want)
		}
	}
}
