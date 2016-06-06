package main

import (
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in   []int
		want bool
	}{
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 1}, true},
		{[]int{1, 2, 1}, true},
		{[]int{1, 2}, false},
		{[]int{2, 1}, false},
	}

	for _, tt := range tests {
		s := sort.IntSlice(tt.in)
		if got := IsPalindrome(s); got != tt.want {
			t.Errorf("IsPalindrome(sort.IntSlice(%v)) = %v; want %v", tt.in, got, tt.want)
		}
	}
}
