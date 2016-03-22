package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	s    []string
	want []string
}{
	{[]string{}, []string{}},
	{[]string{"42"}, []string{"42"}},
	{[]string{"a", "a"}, []string{"a"}},
	{[]string{"a", "b", "a"}, []string{"a", "b", "a"}},
	{[]string{"a", "a", "b", "b", "a"}, []string{"a", "b", "a"}},
	{[]string{"b", "b", "a"}, []string{"b", "a"}},
}

func TestRotate(t *testing.T) {
	for _, tt := range tests {
		s := append([]string{}, tt.s...)
		got := uniq(s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("uniq(%v) = %v, but want %v\n", tt.s, got, tt.want)
		}
	}
}
