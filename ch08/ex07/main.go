package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/net/html"
)

func main() {
	worklist := make(chan []string)
	var n int

	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(u string) []string {
	fmt.Println(u)

	tokens <- struct{}{} // acquire a token
	list, err := Extract(u)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	list = selectSameDomain(u, list)

	for _, link := range list {
		save(link)
	}
	return list
}

func selectSameDomain(link string, list []string) []string {
	sourceURL, err := url.Parse(link)
	if err != nil {
		log.Print(err)
		return nil
	}

	filtered := []string{}
	for _, l := range list {
		crawlURL, err := url.Parse(l)
		if err != nil {
			log.Print(err)
			continue
		}
		if crawlURL.Host != sourceURL.Host {
			continue
		}
		filtered = append(filtered, l)
	}
	return filtered
}

func save(link string) {
	resp, err := http.Get(link)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	u, err := url.Parse(link)
	if err != nil {
		log.Println(err)
		return
	}

	dir := filepath.Join("./", u.Host, filepath.Clean(u.Path))
	base := "[content]"

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(filepath.Join(dir, base), bs, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}
}

// The following functions are copied from "gopl.io/ch5/links"

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
					continue // ignore bad URLs
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
