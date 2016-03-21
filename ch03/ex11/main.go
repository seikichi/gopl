package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	sp := strings.Split(s, ".")
	integer := sp[0]
	fractional := ""
	if len(sp) > 1 {
		fractional = sp[1]
	}

	buf := bytes.Buffer{}

	if len(integer) > 0 && (integer[0] == '+' || integer[0] == '-') {
		buf.WriteRune(rune(integer[0]))
		integer = integer[1:]
	}

	n := len(integer)
	r := n % 3
	if r == 0 && n >= 3 {
		r = 3
	}
	buf.WriteString(integer[:r])
	integer = integer[r:]

	for len(integer) > 0 {
		buf.WriteString(",")
		buf.WriteString(integer[:3])
		integer = integer[3:]
	}

	if len(fractional) > 0 {
		buf.WriteString(".")
		buf.WriteString(fractional)
	}

	return buf.String()
}
