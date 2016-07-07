package popcount

import (
	"strconv"
	"strings"
	"testing"
)

func fromBin(s string) uint64 {
	v, _ := strconv.ParseUint(s, 2, 64)
	return v
}

var tests = []struct {
	in  uint64
	out int
}{
	{fromBin("0"), 0},
	{fromBin("1"), 1},
	{fromBin("111"), 3},
	{fromBin("1001"), 2},
	{fromBin("100000000"), 1},
	{fromBin("1" + strings.Repeat("0", 63)), 1},
	{fromBin(strings.Repeat("10000000", 8)), 8},
}

func TestPopCount(t *testing.T) {
	for _, tt := range tests {
		if got := PopCount(tt.in); got != tt.out {
			t.Errorf("PopCount(0b%b) = %d; want %d", tt.in, got, tt.out)
		}
	}
}
