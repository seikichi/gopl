package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("%s url id\n", os.Args[0])
	}

	url, id := os.Args[1], os.Args[2]

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("parsing HTML: %s", err)
	}

	elem := ElementByID(doc, id)
	if elem == nil {
		log.Fatalf("id = %s not found in %s\n", id, url)
	}

	fmt.Println(elem)
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var result *html.Node

	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}

		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				result = n
				return false
			}
		}
		return true
	}
	post := func(n *html.Node) bool { return true }

	forEachNode(doc, pre, post)
	return result
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) (cont bool)) bool {
	if pre != nil {
		if !pre(n) {
			return false
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}

	if post != nil {
		if !post(n) {
			return false
		}
	}

	return true
}
