package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
	"gopl.io/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		go func() {
		loop:
			for {
				select {
				case link := <-unseenLinks:
					foundLinks := crawl(link)
					go func() { worklist <- foundLinks }()
				case <-done:
					break loop
				}
			}
		}()
	}

	seen := make(map[string]bool)

loop:
	for {
		select {
		case list := <-worklist:
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					select {
					case unseenLinks <- link:
					case <-done:
						break loop
					}
				}
			}
		case <-done:
			break loop
		}
	}
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = done

	resp, err := http.DefaultClient.Do(req)
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

	if cancelled() {
		return nil, nil
	}

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
