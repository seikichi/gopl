package main

import (
	"io/ioutil"
	"testing"
)

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
