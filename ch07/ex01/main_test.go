package main

import (
	"testing"
)

func TestWordCounter(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"test", 1},
		{"  test ", 1},
		{"Hello, world", 2},
	}

	for _, test := range tests {
		var c WordCounter
		c.Write([]byte(test.in))
		if int(c) != test.want {
			t.Errorf("After c.Write(%q), c becomes %v; want %v", test.in, c, test.want)
		}
	}
}

func TestLineCounter(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"foo", 1},
		{"foo\nbar", 2},
		{"foo\nbar\n", 2},
		{
			`foo
bar
foo`, 3,
		},
	}

	for _, test := range tests {
		var c LineCounter
		c.Write([]byte(test.in))
		if int(c) != test.want {
			t.Errorf("After c.Write(%q), c becomes %v; want %v", test.in, c, test.want)
		}
	}
}
