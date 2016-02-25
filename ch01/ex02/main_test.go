package main

import (
	"bytes"
	"testing"
)

var tests = []struct {
	in  []string
	out string
}{
	{[]string{""}, "\n"},
	{[]string{"echo"}, "\n"},
	{[]string{"echo", "Hello", "world!"}, "1\tHello\n2\tworld!\n"},
}

func TestEcho(t *testing.T) {
	for _, tt := range tests {
		outStream := new(bytes.Buffer)
		cli := &CLI{outStream: outStream}
		cli.Run(tt.in)
		if got := outStream.String(); got != tt.out {
			t.Errorf("cli.Run(%q) outputs %q; want %q", tt.in, got, tt.out)
		}
	}
}
