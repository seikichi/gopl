package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var id = flag.String("id", "", "search id")
var class = flag.String("class", "", "search class")

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // stack of element names
	var elem xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elem = tok
			stack = append(stack, tok.Name.Local) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if !containsAll(stack, os.Args[1:]) {
				continue
			}

			if *id != "" && !hasAttr(elem.Attr, "id", *id) {
				continue
			}

			if *class != "" && !hasAttr(elem.Attr, "class", *class) {
				continue
			}

			fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
		}
	}
}

func hasAttr(attrs []xml.Attr, name, value string) bool {
	for _, a := range attrs {
		if a.Name.Local == name && a.Value == value {
			return true
		}
	}
	return false
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
