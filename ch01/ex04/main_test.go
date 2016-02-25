package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestDup(t *testing.T) {
	outStream := new(bytes.Buffer)
	cli := &CLI{
		inStream:  new(bytes.Buffer),
		outStream: outStream,
		errStream: new(bytes.Buffer),
	}

	tmp1, _ := ioutil.TempFile("", "1")
	defer os.Remove(tmp1.Name())
	tmp1.Write([]byte("hello\nworld\n"))

	tmp2, _ := ioutil.TempFile("", "2")
	defer os.Remove(tmp2.Name())
	tmp2.Write([]byte("hello"))

	input := []string{"main", tmp1.Name(), tmp2.Name()}
	cli.Run(input)
	want := fmt.Sprintf("2\t%s,%s\thello\n", tmp1.Name(), tmp2.Name())
	if got := outStream.String(); got != want {
		t.Errorf("cli.Run(%q) outputs %q; want %q", input, got, want)
	}
}

func TestDupFromStdin(t *testing.T) {
	input := `hoge
fuga
hoge
fuga
fuga
`
	outStream := new(bytes.Buffer)
	cli := &CLI{
		inStream:  strings.NewReader(input),
		outStream: outStream,
		errStream: new(bytes.Buffer),
	}
	cli.Run([]string{"main"})
	want := `3	-,-,-	fuga
2	-,-	hoge
`
	if got := outStream.String(); got != want {
		t.Errorf("cli.Run(%q) outputs %q; want %q", input, got, want)
	}
}
