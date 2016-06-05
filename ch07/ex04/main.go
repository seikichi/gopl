package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type reader struct {
	s string
	i int64
}

func (r *reader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n := copy(b, r.s[r.i:])
	r.i += int64(n)
	return n, nil
}

func newReader(s string) io.Reader {
	return &reader{s: s, i: 0}
}

func main() {
	var doc *html.Node
	var err error

	if len(os.Args) == 2 {
		doc, err = html.Parse(newReader(os.Args[1]))
	} else {
		doc, err = html.Parse(os.Stdin)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
