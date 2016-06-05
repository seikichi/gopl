package main

import (
	"testing"
)

func TestNewReader(t *testing.T) {
	tests := []struct {
		in       string
		byteSize int
		wantN    int
	}{
		{"foo", 2, 2},
		{"foo", 3, 3},
		{"foo", 4, 3},
	}

	for _, test := range tests {
		r := newReader(test.in)
		b := make([]byte, test.byteSize)

		n, err := r.Read(b)
		if err != nil || n != test.wantN {
			t.Errorf("newReader(%q).Read(make([]byte, %d)) = %d, %v ; want %d, nil",
				test.in, test.byteSize, n, err, test.wantN)
		}

		want := string([]byte(test.in)[:test.byteSize])
		if string(b) != want {
			t.Errorf("After newReader(%q).Read(b = make([]byte, %d)), b becomes %v; want %v",
				test.in, test.byteSize, string(b), want)
		}
	}
}
