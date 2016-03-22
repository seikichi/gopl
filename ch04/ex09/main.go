package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	f := wordfreq(os.Stdin)

	words := []string{}
	for w := range f {
		words = append(words, w)
	}
	sort.Strings(words)

	fmt.Println("word\tcount")
	for _, w := range words {
		fmt.Printf("%s\t%d\n", w, f[w])
	}
}

func wordfreq(r io.Reader) map[string]int {
	f := map[string]int{}

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		f[s.Text()]++
	}

	return f
}
