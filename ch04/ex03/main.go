package main

import "fmt"

const size = 8

func main() {
	a := [size]int{0, 1, 2, 3, 4, 5, 6, 7}
	reverse(&a)
	fmt.Println(a)
}

func reverse(s *[size]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
