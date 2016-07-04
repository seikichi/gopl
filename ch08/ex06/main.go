package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

type work struct {
	links []string
	depth int
}

func main() {
	depth := 10
	flag.IntVar(&depth, "depth", depth, "depth")
	flag.Parse()

	worklist := make(chan *work)
	n := 0

	n++
	go func() { worklist <- &work{links: flag.Args(), depth: 1} }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		w := <-worklist
		if w.depth > depth {
			continue
		}
		for _, link := range w.links {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string, depth int) {
					worklist <- &work{links: crawl(link), depth: depth + 1}
				}(link, w.depth)
			}
		}
	}
}

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
