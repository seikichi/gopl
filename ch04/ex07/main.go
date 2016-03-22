package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(string(reverse([]byte("こんにちは!!"))))
}

type result struct {
	s    []byte
	i, j int
}

// func (res result) pushFront(r rune) {
// 	copy(s[ai:], []byte(string(rl)))
// 	i += utf8.RuneLen(r)
// }

// func (res result) pushBack(r rune) {
// }

func reverse(s []byte) []byte {
	// Note: ai <= bi < bj <= aj
	ai, aj := 0, len(s)
	bi, bj := 0, len(s)
	que := []rune{}
	for bi < bj {
		rl, dl := utf8.DecodeLastRune(s[bi:bj])
		bj -= dl

		for bi < bj && (bi-ai) < dl {
			rh, dh := utf8.DecodeRune(s[bi:bj])
			bi += dh
			que = append(que, rh)
		}
		copy(s[ai:], []byte(string(rl)))
		ai += dl

		for len(que) > 0 {
			rh := que[0]
			dh := utf8.RuneLen(rh)
			if bj >= aj-dh {
				break
			}
			copy(s[aj-dh:], []byte(string(rh)))
			aj -= dh
			que = que[1:]
		}
	}
	for len(que) > 0 && ai < aj {
		rh := que[0]
		copy(s[ai:], []byte(string(rh)))
		ai += utf8.RuneLen(rh)
	}
	return s
}
