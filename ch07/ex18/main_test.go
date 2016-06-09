package main

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func TestNewTree(t *testing.T) {
	tests := []struct {
		in   string
		want Node
	}{
		{
			"<foo>bar</foo>",
			&Element{
				Type:     xml.Name{Space: "", Local: "foo"},
				Attr:     []xml.Attr{},
				Children: []Node{CharData("bar")},
			},
		},
		{
			"<foo>A<bar>B</bar>C<bar><bar>D</bar></bar></foo>",
			&Element{
				Type: xml.Name{Space: "", Local: "foo"},
				Attr: []xml.Attr{},
				Children: []Node{
					CharData("A"),
					&Element{
						Type:     xml.Name{Space: "", Local: "bar"},
						Attr:     []xml.Attr{},
						Children: []Node{CharData("B")},
					},
					CharData("C"),
					&Element{
						Type: xml.Name{Space: "", Local: "bar"},
						Attr: []xml.Attr{},
						Children: []Node{
							&Element{
								Type:     xml.Name{Space: "", Local: "bar"},
								Attr:     []xml.Attr{},
								Children: []Node{CharData("D")},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		dec := xml.NewDecoder(strings.NewReader(tt.in))
		got, err := NewTree(dec)
		if err != nil {
			t.Errorf("NewTree(reader from %q) = _, %s; want _, nil", tt.in, err)
		} else if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("NewTree(reader from %q) = %s, nil; want %s, nil", tt.in, got, tt.want)
		}
	}
}
