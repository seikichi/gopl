package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var tests = []struct {
	in            string
	words, images int
}{
	{``, 0, 0},
	{`
<html>
  <head>Hello, world!</head>
  <body>
    <img src="image.png"/>
  </body>
</html>
`, 2, 1},
	{`
<html>
  <head>Hello, world!</head>
  <body>
    <p>I am sleeeeeeeepy!!!!!</p>
    <div>
      <div>
        <img src="image1.png"/>
        <span>Test</span>
      </div>
    </div>
    <img src="image2.png"/>
  </body>
</html>
`, 6, 2},
}

func TestVisit(t *testing.T) {
	for _, tt := range tests {
		doc, _ := html.Parse(strings.NewReader(tt.in))
		words, images := countWordsAndImages(doc)

		if words != tt.words || images != tt.images {
			t.Errorf("countWordsAndImages(Parse(%s)) = %d, %d; want %d, %d",
				tt.in, words, images, tt.words, tt.images)
		}
	}
}
