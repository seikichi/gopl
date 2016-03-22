package main

import "fmt"

func main() {
	fmt.Println(uniq([]string{"a", "a", "a"}))
}

func uniq(s []string) []string {
	out := s[:0]

	var prev string
	for i, v := range s {
		if i == 0 || prev != v {
			out = append(out, v)
			prev = v
		}
	}
	return out
}
