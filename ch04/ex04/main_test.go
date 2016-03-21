package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	s    []int
	d    int
	want []int
}{
	{[]int{}, 42, []int{}},
	{[]int{0}, 42, []int{0}},
	{[]int{1, 2}, 2, []int{1, 2}},
	{[]int{1, 2}, 1, []int{2, 1}},
	{[]int{0, 1, 2, 3, 4, 5}, 2, []int{2, 3, 4, 5, 0, 1}},
	{[]int{0, 1, 2, 3, 4, 5}, -2, []int{4, 5, 0, 1, 2, 3}},
}

func TestRotate(t *testing.T) {
	for _, tt := range tests {
		s := append([]int{}, tt.s...)
		rotate(s, tt.d)
		if !reflect.DeepEqual(s, tt.want) {
			t.Errorf("rotate(%v, %d) changes the argument to %v, but want %v\n", tt.s, tt.d, s, tt.want)
		}
	}
}
