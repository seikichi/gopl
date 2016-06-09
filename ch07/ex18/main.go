package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func NewTree(dec *xml.Decoder) (Node, error) {
	var stack []*Element // stack of elements
	var lastElement *Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elem := stack[len(stack)-1]
			newElem := &Element{tok.Name, tok.Attr, []Node{}}
			elem.Children = append(elem.Children, newElem)
			stack = append(stack, newElem)
		case xml.EndElement:
			lastElement = stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if len(stack) == 0 {
				return nil, errors.New("invalid xml tree")
			}
			elem := stack[len(stack)-1]
			elem.Children = append(elem.Children, CharData(tok))
		}
	}
	return lastElement, nil
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	tree, err := NewTree(dec)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", tree)
}
