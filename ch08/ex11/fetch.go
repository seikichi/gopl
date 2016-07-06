package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string, done <-chan struct{}) (filename string, n int64, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", 0, err
	}
	req.Cancel = done

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

type result struct {
	filename string
	n        int64
}

func main() {
	done := make(chan struct{})
	results := make(chan result, len(os.Args[1:]))

	for _, url := range os.Args[1:] {
		go func(url string) {
			local, n, err := fetch(url, done)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
				return
			}
			results <- result{local, n}
		}(url)
	}

	r := <-results
	close(done)
	fmt.Printf("%s (%d bytes).\n", r.filename, r.n)
}
