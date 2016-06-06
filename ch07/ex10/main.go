package main

import (
	"fmt"
	"os"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len(); i++ {
		j := s.Len() - i - 1
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

type runeSort []rune

func (s runeSort) Len() int           { return len(s) }
func (s runeSort) Less(i, j int) bool { return s[i] < s[j] }
func (s runeSort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	for _, v := range os.Args[1:] {
		if IsPalindrome(runeSort([]rune(v))) {
			fmt.Printf("%q is palindrome!\n", v)
		} else {
			fmt.Printf("%q is not palindrome...\n", v)
		}
	}
}
