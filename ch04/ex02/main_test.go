package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"testing"
)

var tests = []struct {
	in   []byte
	args []string
	out  string
}{
	{
		[]byte("Hello, world!"),
		[]string{"main"},
		fmt.Sprintf("%x\n", sha256.Sum256([]byte("Hello, world!"))),
	},
	{
		[]byte("Hello, world!"),
		[]string{"main", "-type", "sha256"},
		fmt.Sprintf("%x\n", sha256.Sum256([]byte("Hello, world!"))),
	},
	{
		[]byte("Hello, world!"),
		[]string{"main", "-type", "sha384"},
		fmt.Sprintf("%x\n", sha512.Sum384([]byte("Hello, world!"))),
	},
	{
		[]byte("Hello, world!"),
		[]string{"main", "-type", "sha512"},
		fmt.Sprintf("%x\n", sha512.Sum512([]byte("Hello, world!"))),
	},
}

func TestMain(t *testing.T) {
	for _, tt := range tests {
		outStream := new(bytes.Buffer)
		cli := &CLI{outStream: outStream, inStream: bytes.NewBuffer(tt.in)}
		cli.Run(tt.args)
		if got := outStream.String(); got != tt.out {
			t.Errorf("main(%q) outputs %q; want %q", tt.in, got, tt.out)
		}
	}
}
