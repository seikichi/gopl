package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
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
		t.Errorf("cli.Run(%q) = %q; want %q", input, got, want)
	}
}
