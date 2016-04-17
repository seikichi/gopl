package main

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var tests = []struct {
	in   string
	want map[string]int
}{
	{
		`<html><head></head><body></body></html>`,
		map[string]int{"html": 1, "body": 1, "head": 1},
	},
	{
		`
<html>
  <head></head>
  <body><a href="http://example.com">Example</a></body>
</html>`,
		map[string]int{"html": 1, "body": 1, "head": 1, "a": 1},
	},
	{
		`\
<html>
  <head></head>
  <body>
  	<a href="http://example.com/1">Example 1</a>
  	<a href="http://example.com/2">Example 2</a>
  	<a href="http://example.com/3">Example 3</a>
  </body>
</html>
	`,
		map[string]int{"html": 1, "body": 1, "head": 1, "a": 3},
	},
	{
		`\
<html>
  <head></head>
  <body>
    <p>hoge<span>fuga</span>poyo</p>
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
  </body>
</html>
	`,
		map[string]int{"html": 1, "body": 1, "head": 1, "a": 3, "div": 4, "p": 1, "span": 1},
	},
}

func TestMapping(t *testing.T) {
	for _, tt := range tests {
		doc, _ := html.Parse(strings.NewReader(tt.in))
		got := mapping(map[string]int{}, doc)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("mapping(nil, Parse(%q)) = %v; want %v", tt.in, got, tt.want)
		}
	}
}
