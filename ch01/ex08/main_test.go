package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetchWithoutHttp(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	}))
	defer ts.Close()

	outStream := new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: ioutil.Discard}

	input := []string{"fetch", strings.Replace(ts.URL, "http://", "", 1)}
	cli.Run(input)
	if got, want := outStream.String(), "Hello, world!\n"; got != want {
		t.Errorf("cli.Run(%q) = %q; want %q", input, got, want)
	}
}

// ... copied from `github.com/seikichi/gopl/ch01/ex07`

func TestFetch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	}))
	defer ts.Close()

	outStream := new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: ioutil.Discard}

	input := []string{"fetch", ts.URL}
	exit := cli.Run(input)
	if exit != ExitCodeOK {
		t.Errorf("cli.Run(%q) = %q; want %q", input, exit, ExitCodeOK)
	}
	if got, want := outStream.String(), "Hello, world!\n"; got != want {
		t.Errorf("cli.Run(%q) outputs %q; want %q", input, got, want)
	}
}

func TestFetchWithHTTPError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(nil)
	}))
	defer ts.Close()

	outStream := new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: ioutil.Discard}

	input := []string{"fetch", ts.URL}
	exit := cli.Run(input)
	if exit != ExitCodeHTTPError {
		t.Errorf("cli.Run(%q) = %q; want %q", input, exit, ExitCodeHTTPError)
	}
	if got, want := outStream.String(), ""; got != want {
		t.Errorf("cli.Run(%q) outputs %q; want %q", input, got, want)
	}
}

type DummyWriter struct{}

func (w *DummyWriter) Write(p []byte) (int, error) {
	return 0, errors.New("")
}

func TestFetchWithCopyError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	}))
	defer ts.Close()

	cli := &CLI{outStream: &DummyWriter{}, errStream: ioutil.Discard}

	input := []string{"fetch", ts.URL}
	exit := cli.Run(input)
	if exit != ExitCodeCopyError {
		t.Errorf("cli.Run(%q) = %q; want %q", input, exit, ExitCodeCopyError)
	}
}
