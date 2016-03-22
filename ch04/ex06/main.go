package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "こんにちは　 　世界 (<- 全角スペースで半角スペースを囲んでいます)"
	fmt.Println(string(compress([]byte(s))))
}

func compress(s []byte) []byte {
	out := s[:0]
	i := 0
	var prev rune
	for i < len(s) {
		r, d := utf8.DecodeRune(s[i:])
		if !unicode.IsSpace(r) {
			out = append(out, s[i:i+d]...)
		} else if !unicode.IsSpace(prev) {
			out = append(out, []byte(" ")...)
		}
		prev = r
		i += d
	}
	return out
}
