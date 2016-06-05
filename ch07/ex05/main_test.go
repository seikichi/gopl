package main

import (
	"strings"
	"testing"
)

func TestLimitedReader(t *testing.T) {
	tests := []struct {
		in string
		n  int
	}{
		{"foo", 1},
	}

	for _, test := range tests {
		r := LimitReader(strings.NewReader(test.in), int64(test.n))
		b := make([]byte, len(test.in))
		n, err := r.Read(b)
		if err != nil || n != test.n {
			t.Errorf("LimitReader(...).Read(...) = %d, %v ; want %d, nil", n, err, test.n)
		}

		want := string([]byte(test.in)[:test.n])
		got := string(b[:n])
		if got != want {
			t.Errorf("After LimitReader(...).Read(b), b becomes %q; want %q", got, want)
		}
	}
}
