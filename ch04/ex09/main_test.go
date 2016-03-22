package main

import (
	"reflect"
	"strings"
	"testing"
)

var tests = []struct {
	in   string
	want map[string]int
}{
	{"", map[string]int{}},
	{"世界", map[string]int{"世界": 1}},
	{"AA BB AA", map[string]int{"AA": 2, "BB": 1}},
	{"aa BB AA", map[string]int{"AA": 1, "BB": 1, "aa": 1}},
}

func TestWordFreq(t *testing.T) {
	for _, tt := range tests {
		got := wordfreq(strings.NewReader(tt.in))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("wordfreq(%q) = %v; want %v", tt.in, got, tt.want)
		}
	}
}
