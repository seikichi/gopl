package main

import "testing"

var tests = []struct {
	in1, in2 string
	want     bool
}{
	{"", "", true},
	{"aA", "Aa", true},
	{"abc", "cba", true},
	{"Hello, 世界", "世界, Hello", true},
	{"aaA", "aAA", false},
	{"123", "１２３", false},
}

func TestComma(t *testing.T) {
	for _, tt := range tests {
		if got := IsAnagram(tt.in1, tt.in2); got != tt.want {
			t.Errorf("IsAnagram(%q, %q) = %v; want %v", tt.in1, tt.in2, got, tt.want)
		}
	}
}
