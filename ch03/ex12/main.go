package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s s1 s2", os.Args[0])
		os.Exit(-1)
	}
	s1, s2 := os.Args[1], os.Args[2]
	if IsAnagram(s1, s2) {
		fmt.Printf("%v and %v are anagrams!\n", s1, s2)
	} else {
		fmt.Printf("%v and %v are not anagrams...\n", s1, s2)
	}
}

func IsAnagram(s1, s2 string) bool {
	c1 := countRunes(s1)
	c2 := countRunes(s2)
	return reflect.DeepEqual(c1, c2)
}

func countRunes(s string) map[rune]int {
	c := map[rune]int{}
	for _, r := range s {
		c[r]++
	}
	return c
}
