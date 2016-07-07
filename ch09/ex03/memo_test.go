package memo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func httpGetBody(url string, cancel Cancel) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, cancel Cancel) (interface{}, error)
}

func Sequential(t *testing.T, m M, c Cancel) {
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, c)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

func Concurrent(t *testing.T, m M, c Cancel) {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url, c)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}

func Test(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()

	c := make(Cancel)
	Sequential(t, m, c)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()

	c := make(Cancel)
	Concurrent(t, m, c)
}

func waitForCancel(_ string, cancel Cancel) (interface{}, error) {
	<-cancel
	return nil, errors.New("cancelled")
}

func TestCancel(t *testing.T) {
	m := New(waitForCancel)
	defer m.Close()

	c := make(chan struct{})
	results := make(chan result)
	go func() {
		v, err := m.Get("", c)
		results <- result{v, err}
	}()

	<-time.After(1 * time.Second)
	close(c)
	r := <-results
	if r.value != nil || r.err == nil {
		t.Errorf("New(waitForCancel).Get(...) = %#v, %#v; want nil, error (!= nil)",
			r.value, r.err)
	}
}
