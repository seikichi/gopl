package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %v\n", err)
		os.Exit(1)
	}
	count := mapping(map[string]int{}, doc)
	for k, v := range count {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func mapping(count map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		mapping(count, c)
	}
	return count
}
