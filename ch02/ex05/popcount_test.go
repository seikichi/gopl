package popcount_test

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

func TestPopCountByLoop(t *testing.T) {
	for _, tt := range tests {
		if got := PopCountByLoop(tt.in); got != tt.out {
			t.Errorf("PopCountByLoop(0b%b) = %d; want %d", tt.in, got, tt.out)
		}
	}
}

func TestPopCountByShift(t *testing.T) {
	for _, tt := range tests {
		if got := PopCountByShift(tt.in); got != tt.out {
			t.Errorf("PopCountByShift(0b%b) = %d; want %d", tt.in, got, tt.out)
		}
	}
}

func TestPopCountByClear(t *testing.T) {
	for _, tt := range tests {
		if got := PopCountByClear(tt.in); got != tt.out {
			t.Errorf("PopCountByClear(0b%b) = %d; want %d", tt.in, got, tt.out)
		}
	}
}

func TestPopCountByHD(t *testing.T) {
	for _, tt := range tests {
		if got := PopCountByHD(tt.in); got != tt.out {
			t.Errorf("PopCountByHD(0b%b) = %d; want %d", tt.in, got, tt.out)
		}
	}
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountByLoop(x uint64) int {
	count := 0
	for i := uint(0); i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

func PopCountByShift(x uint64) int {
	count := 0
	for i := uint(0); i < 64; i++ {
		count += int(x & 1)
		x = x >> 1
	}
	return count
}

func PopCountByClear(x uint64) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

func PopCountByHD(x uint64) int {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0xFFFFFFFFFFFFFFFF)
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByLoop(0xFFFFFFFFFFFFFFFF)
	}
}

func BenchmarkPopCountByShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShift(0xFFFFFFFFFFFFFFFF)
	}
}

func BenchmarkPopCountByClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClear(0xFFFFFFFFFFFFFFFF)
	}
}

func BenchmarkPopCountByHD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByHD(0xFFFFFFFFFFFFFFFF)
	}
}
