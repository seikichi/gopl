package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

	input := []string{"fetch", ts.URL}
	cli.Run(input)
	want := `HTTP/1.1 200 OK

Hello, world!
`

	if got := outStream.String(); got != want {
		t.Errorf("cli.Run(%q) = %q; want %q", input, got, want)
	}
}
