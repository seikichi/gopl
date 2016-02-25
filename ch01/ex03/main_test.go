package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

var tests = []struct {
	in  []string
	out string
}{
	{[]string{""}, "\n"},
	{[]string{"echo"}, "\n"},
	{[]string{"echo", "Hello", "world!"}, "Hello world!\n"},
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

func TestEchoInefficiently(t *testing.T) {
	for _, tt := range tests {
		outStream := new(bytes.Buffer)
		cli := &CLI{outStream: outStream}
		cli.RunInefficiently(tt.in)
		if got := outStream.String(); got != tt.out {
			t.Errorf("cli.Run(%q) outputs %q; want %q", tt.in, got, tt.out)
		}
	}
}

func BenchmarkRun10(b *testing.B)   { benchmarkRun(b, 10) }
func BenchmarkRun100(b *testing.B)  { benchmarkRun(b, 100) }
func BenchmarkRun1000(b *testing.B) { benchmarkRun(b, 1000) }

func BenchmarkRunInefficiently10(b *testing.B)   { benchmarkRunInefficiently(b, 10) }
func BenchmarkRunInefficiently100(b *testing.B)  { benchmarkRunInefficiently(b, 100) }
func BenchmarkRunInefficiently1000(b *testing.B) { benchmarkRunInefficiently(b, 1000) }

func createTestCase(n int) []string {
	input := []string{}
	for i := 0; i < n; i++ {
		input = append(input, "a")
	}
	return input
}

func benchmarkRun(b *testing.B, n int) {
	input := createTestCase(n)
	cli := &CLI{outStream: ioutil.Discard}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cli.Run(input)
	}
}

func benchmarkRunInefficiently(b *testing.B, n int) {
	input := createTestCase(n)
	cli := &CLI{outStream: ioutil.Discard}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cli.RunInefficiently(input)
	}
}
