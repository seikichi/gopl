package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}

	buf := bytes.Buffer{}
	n := len(s)
	buf.Grow(n + n/3)

	r := n - 3*((n-1)/3)
	buf.WriteString(s[:r])
	s = s[r:]

	for len(s) > 0 {
		buf.WriteString(",")
		buf.WriteString(s[:3])
		s = s[3:]
	}
	return buf.String()
}
