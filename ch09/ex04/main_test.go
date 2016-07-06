package main

import "testing"

func BenchmarkRun10(b *testing.B)    { benchmarkRun(b, 10) }
func BenchmarkRun100(b *testing.B)   { benchmarkRun(b, 100) }
func BenchmarkRun1000(b *testing.B)  { benchmarkRun(b, 1000) }
func BenchmarkRun10000(b *testing.B) { benchmarkRun(b, 10000) }

func benchmarkRun(b *testing.B, n int) {
	in, out := pipeline(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() { in <- struct{}{} }()
		<-out
	}
}
