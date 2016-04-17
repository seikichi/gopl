package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var tests = []struct {
	in    string
	id    string
	exist bool
	data  string
}{
	{``, "test", false, ""},
	{`<div id="test" data-test="test">`, "test", true, "test"},
	{`
<html>
  <head><title>Title</title></heat>
  <body>
    <div>
     <p id="tes"></p>
     <p id="est"></p>
     <p id="test" data-test="FooBar"></p>
    </div>
  </body>
</html>
`, "test", true, "FooBar"},
}

func TestVisit(t *testing.T) {
	for _, tt := range tests {
		doc, _ := html.Parse(strings.NewReader(tt.in))
		elem := ElementByID(doc, tt.id)

		if elem == nil && tt.exist {
			t.Errorf("ElementByID(Parse(%v), %s) = %+v; want non nil result", tt.in, tt.id, elem)
		}

		if elem != nil && !tt.exist {
			t.Errorf("ElementByID(Parse(%v), %s) = %+v; want nil", tt.in, tt.id, elem)
		}

		if elem != nil && tt.exist {
			found, val := false, ""
			for _, a := range elem.Attr {
				if a.Key == "data-test" {
					found = true
					val = a.Val
				}
			}
			if !found || val != tt.data {
				t.Errorf("ElementByID(Parse(%v), %s) = %+v; want node with data-test=%v",
					tt.in, tt.id, elem, tt.data)
			}
		}
	}
}
