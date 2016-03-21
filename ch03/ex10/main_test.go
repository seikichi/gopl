package main

import "testing"

var tests = []struct{ in, want string }{
	{"", ""},
	{"1", "1"},
	{"12", "12"},
	{"123", "123"},
	{"1234", "1,234"},
	{"123456", "123,456"},
	{"1234567", "1,234,567"},
}

func TestComma(t *testing.T) {
	for _, tt := range tests {
		if got := comma(tt.in); got != tt.want {
			t.Errorf("comma(%q) = %q; want %q", tt.in, got, tt.want)
		}
	}
}
