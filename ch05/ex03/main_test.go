package main

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var tests = []struct {
	in   string
	want []string
}{
	{
		`<html><head></head><body>こんにちは</body></html>`,
		[]string{"こんにちは"},
	},
	{
		`<html><head></head><script>こんにちは</script><style>こんにちは</style></html>`,
		[]string{},
	},
	{
		"<html>" +
			"<head><title>タイトル</title></head>" +
			"<body>" +
			"<script>console.log('Hello, world!');</script>" +
			"<p>こんにちは</p>" +
			"<div><p>A</p><div><div><p>B</p></div><p>C</p></div></div>" +
			"</body></html>",
		[]string{"タイトル", "こんにちは", "A", "B", "C"},
	},
}

func TestMapping(t *testing.T) {
	for _, tt := range tests {
		doc, _ := html.Parse(strings.NewReader(tt.in))
		got := findtexts([]string{}, doc)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("mapping(nil, Parse(%q)) = %q; want %q", tt.in, got, tt.want)
		}
	}
}
