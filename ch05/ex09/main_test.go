package main

import "testing"

var tests = []struct {
	in   string
	f    func(string) string
	want string
}{
	{"", func(s string) string { return s }, ""},
	{"$hoge $hoge", func(s string) string { return s }, "hoge hoge"},
	{"foo $bar foo", func(s string) string { return s + s }, "foo barbar foo"},
}

func TestVisit(t *testing.T) {
	for _, tt := range tests {
		got := expand(tt.in, tt.f)
		if got != tt.want {
			t.Errorf("expand(%q, %v) = %q; want %q", tt.in, tt.f, got, tt.want)
		}
	}
}
