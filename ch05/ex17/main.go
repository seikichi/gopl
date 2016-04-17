package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main url tags...")
	}

	url := os.Args[1]
	tags := os.Args[2:]

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	es := ElementsByTagName(doc, tags...)
	for _, e := range es {
		fmt.Printf("{Data: %+v, Attr: %+v}\n", e.Data, e.Attr)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	set := map[string]bool{}
	for _, n := range name {
		set[n] = true
	}

	var rec func(*html.Node, []*html.Node) []*html.Node

	rec = func(n *html.Node, list []*html.Node) []*html.Node {
		if n.Type == html.ElementNode && set[n.Data] {
			list = append(list, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			list = rec(c, list)
		}
		return list
	}

	return rec(doc, nil)
}
