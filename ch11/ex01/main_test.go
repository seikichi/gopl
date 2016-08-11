package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharcount(t *testing.T) {
	var tests = []struct {
		input   string
		counts  map[rune]int
		utflen  [utf8.UTFMax + 1]int
		invalid int
	}{
		{"Hello", map[rune]int{'H': 1, 'e': 1, 'l': 2, 'o': 1}, utflenFromMap(map[int]int{1: 5}), 0},
		{"あ", map[rune]int{'あ': 1}, utflenFromMap(map[int]int{utf8.RuneLen('あ'): 1}), 0},
		{
			"あiueお",
			map[rune]int{'あ': 1, 'i': 1, 'u': 1, 'e': 1, 'お': 1},
			utflenFromMap(map[int]int{utf8.RuneLen('あ'): 2, 1: 3}),
			0,
		},
	}

	for _, test := range tests {
		in := bufio.NewReader(strings.NewReader(test.input))
		counts, utflen, invalid, err := charcount(in)
		if !reflect.DeepEqual(counts, test.counts) || utflen != test.utflen || invalid != test.invalid || err != nil {
			t.Errorf("charcount(%q) = %v, %v, %v, %v ; want %v, %v, %v, nil", test.input,
				counts, utflen, invalid, err,
				test.counts, test.utflen, test.invalid)
		}
	}
}

func utflenFromMap(m map[int]int) [utf8.UTFMax + 1]int {
	ret := [utf8.UTFMax + 1]int{}
	for k, v := range m {
		ret[k] = v
	}
	return ret
}
