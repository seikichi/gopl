package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s s1 s2\n", os.Args[0])
	}
	fmt.Printf("Different bit count in shar256: %d\n",
		DifferentBitCountInSha256([]byte(os.Args[1]), []byte(os.Args[2])))
}

func DifferentBitCountInSha256(b1, b2 []byte) int {
	s1 := sha256.Sum256(b1)
	s2 := sha256.Sum256(b2)

	sum := 0
	for i := 0; i < len(s1); i++ {
		sum += popCount(uint64(s1[i] ^ s2[i]))
	}
	return sum
}

func popCount(x uint64) int {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}
