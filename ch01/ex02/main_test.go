package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	outStream := new(bytes.Buffer)
	cli := &CLI{outStream: outStream}

	input := []string{"echo", "Hello", "world!"}
	cli.Run(input)
	if got, want := outStream.String(), "1\tHello\n2\tworld!\n"; got != want {
		t.Errorf("cli.Run(%q) = %q; want %q", input, got, want)
	}
}
