package main

import (
	"crypto/sha256"
	"testing"
	"testing/quick"
)

func naiveCount(b1, b2 []byte) int {
	s1 := sha256.Sum256(b1)
	s2 := sha256.Sum256(b2)

	ret := 0
	for i := 0; i < len(s1); i++ {
		v1, v2 := s1[i], s2[i]
		for j := uint8(0); (v1>>j) != 0 || (v2>>j) != 0; j++ {
			if (v1>>j)&1 != (v2>>j)&1 {
				ret++
			}
		}
	}
	return ret
}

func TestDifferentBitCountInSha256(t *testing.T) {
	err := quick.CheckEqual(DifferentBitCountInSha256, naiveCount, nil)
	if err != nil {
		t.Error(err)
	}
}
