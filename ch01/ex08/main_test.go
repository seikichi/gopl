package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
})

func TestFetch(t *testing.T) {
	ts := httptest.NewServer(handler)
	defer ts.Close()

	outStream := new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: ioutil.Discard}

	input := []string{"fetch", strings.Replace(ts.URL, "http://", "", 1)}
	cli.Run(input)
	if got, want := outStream.String(), "Hello, world!\n"; got != want {
		t.Errorf("cli.Run(%q) = %q; want %q", input, got, want)
	}
}
