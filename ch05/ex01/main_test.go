package main

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var tests = []struct {
	in   string
	want []string
}{
	{
		``,
		[]string{},
	},
	{
		`<a href="http://example.com">Example</a>`,
		[]string{"http://example.com"},
	},
	{
		`\
<a href="http://example.com/1">Example 1</a>
<a href="http://example.com/2">Example 2</a>
<a href="http://example.com/3">Example 3</a>
`,
		[]string{"http://example.com/1", "http://example.com/2", "http://example.com/3"},
	},
	{
		`\
<div>
  <div>
    <div>
      <a href="http://example.com/1">Example 1</a>
    </div>
  </div>
  <a href="http://example.com/2">Example 2</a>
  <div>
    <a href="http://example.com/3">Example 3</a>
  </div>
</div>
`,
		[]string{"http://example.com/1", "http://example.com/2", "http://example.com/3"},
	},
}

func TestVisit(t *testing.T) {
	for _, tt := range tests {
		doc, _ := html.Parse(strings.NewReader(tt.in))
		got := visit(nil, doc)
		want := append([]string(nil), tt.want...)

		sort.Strings(got)
		sort.Strings(want)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("visit(nil, Parse(%q)) = %q; want %q", tt.in, got, want)
		}
	}
}
