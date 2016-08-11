package intset

import (
	"math/rand"
	"testing"
	"time"
)

var seed = time.Now().UTC().UnixNano()
var rng = rand.New(rand.NewSource(seed))

// Benchmarks (Add)

func BenchmarkUint64SliceAdd100000(b *testing.B)    { benchmarkUint64SliceAdd(b, 100000) }
func BenchmarkUint64SliceAdd1000000(b *testing.B)   { benchmarkUint64SliceAdd(b, 1000000) }
func BenchmarkUint64SliceAdd10000000(b *testing.B)  { benchmarkUint64SliceAdd(b, 10000000) }
func BenchmarkUint64SliceAdd100000000(b *testing.B) { benchmarkUint64SliceAdd(b, 100000000) }

func benchmarkUint64SliceAdd(b *testing.B, size int) {
	m := IntSetByUint64Slice{}
	for i := 0; i < b.N; i++ {
		m.Add(rng.Intn(size))
	}
}

func BenchmarkOptimizedIntSetAdd100000(b *testing.B)    { benchmarkOptimizedIntSetAdd(b, 100000) }
func BenchmarkOptimizedIntSetAdd1000000(b *testing.B)   { benchmarkOptimizedIntSetAdd(b, 1000000) }
func BenchmarkOptimizedIntSetAdd10000000(b *testing.B)  { benchmarkOptimizedIntSetAdd(b, 10000000) }
func BenchmarkOptimizedIntSetAdd100000000(b *testing.B) { benchmarkOptimizedIntSetAdd(b, 100000000) }

func benchmarkOptimizedIntSetAdd(b *testing.B, size int) {
	m := OptimizedIntSet{}
	for i := 0; i < b.N; i++ {
		m.Add(rng.Intn(size))
	}
}

func BenchmarkUint32SliceAdd100000(b *testing.B)    { benchmarkUint32SliceAdd(b, 100000) }
func BenchmarkUint32SliceAdd1000000(b *testing.B)   { benchmarkUint32SliceAdd(b, 1000000) }
func BenchmarkUint32SliceAdd10000000(b *testing.B)  { benchmarkUint32SliceAdd(b, 10000000) }
func BenchmarkUint32SliceAdd100000000(b *testing.B) { benchmarkUint32SliceAdd(b, 100000000) }

func benchmarkUint32SliceAdd(b *testing.B, size int) {
	m := IntSetByUint32Slice{}
	for i := 0; i < b.N; i++ {
		m.Add(rng.Intn(size))
	}
}

func BenchmarkMapAdd100000(b *testing.B)    { benchmarkMapAdd(b, 100000) }
func BenchmarkMapAdd1000000(b *testing.B)   { benchmarkMapAdd(b, 1000000) }
func BenchmarkMapAdd10000000(b *testing.B)  { benchmarkMapAdd(b, 10000000) }
func BenchmarkMapAdd100000000(b *testing.B) { benchmarkMapAdd(b, 100000000) }

func benchmarkMapAdd(b *testing.B, size int) {
	m := *NewIntSetByMap()
	for i := 0; i < b.N; i++ {
		m.Add(rng.Intn(size))
	}
}

func BenchmarkBoolMapAdd100000(b *testing.B)    { benchmarkBoolMapAdd(b, 100000) }
func BenchmarkBoolMapAdd1000000(b *testing.B)   { benchmarkBoolMapAdd(b, 1000000) }
func BenchmarkBoolMapAdd10000000(b *testing.B)  { benchmarkBoolMapAdd(b, 10000000) }
func BenchmarkBoolMapAdd100000000(b *testing.B) { benchmarkBoolMapAdd(b, 100000000) }

func benchmarkBoolMapAdd(b *testing.B, size int) {
	m := *NewIntSetByBoolMap()
	for i := 0; i < b.N; i++ {
		m.Add(rng.Intn(size))
	}
}

// Benchmarks (UnionWith)

func BenchmarkUint64SliceUnionWith1000_10(b *testing.B)  { benchmarkUint64SliceUnionWith(b, 1000, 10) }
func BenchmarkUint64SliceUnionWith1000_100(b *testing.B) { benchmarkUint64SliceUnionWith(b, 1000, 100) }

func benchmarkUint64SliceUnionWith(b *testing.B, randSize int, unionSize int) {
	m := IntSetByUint64Slice{}
	for i := 0; i < b.N; i++ {
		t := IntSetByUint64Slice{}
		for j := 0; j < unionSize; j++ {
			t.Add(rng.Intn(randSize))
		}

		m.UnionWith(&t)
	}
}

func BenchmarkOptimizedUnionWith1000_10(b *testing.B)  { benchmarkOptimizedUnionWith(b, 1000, 10) }
func BenchmarkOptimizedUnionWith1000_100(b *testing.B) { benchmarkOptimizedUnionWith(b, 1000, 100) }

func benchmarkOptimizedUnionWith(b *testing.B, randSize int, unionSize int) {
	m := OptimizedIntSet{}
	for i := 0; i < b.N; i++ {
		t := OptimizedIntSet{}
		for j := 0; j < unionSize; j++ {
			t.Add(rng.Intn(randSize))
		}

		m.UnionWith(&t)
	}
}

// Tests

var addTests = []struct {
	init  []int
	add   int
	tests []int
}{
	{[]int{}, 1, []int{1, 2, 3}},
	{[]int{1}, 1, []int{1, 2, 3}},
	{[]int{1}, 2, []int{1, 2, 3}},
	{[]int{}, 1, []int{1, 2, 3}},
	{[]int{1}, 10000, []int{1, 2, 3, 10000, 20000, 30000}},
	{[]int{10000}, 1, []int{1, 2, 3, 10000, 20000, 30000}},
	{[]int{10000}, 10000, []int{1, 2, 3, 10000, 20000, 30000}},
}

var unionTests = []struct {
	init  []int
	union []int
	tests []int
}{
	{[]int{}, []int{}, []int{1, 2, 3}},
	{[]int{1}, []int{}, []int{1, 2, 3}},
	{[]int{}, []int{1}, []int{1, 2, 3}},
	{[]int{1}, []int{1}, []int{1, 2, 3}},
	{[]int{1}, []int{2}, []int{1, 2, 3}},
	{[]int{10000}, []int{}, []int{1, 2, 3, 10000, 20000, 30000}},
	{[]int{}, []int{10000}, []int{1, 2, 3, 10000, 20000, 30000}},
	{[]int{10000}, []int{10000}, []int{1, 2, 3, 10000, 20000, 30000}},
	{[]int{10000}, []int{1}, []int{1, 2, 3, 10000, 20000, 30000}},
	{[]int{1}, []int{10000}, []int{1, 2, 3, 10000, 20000, 30000}},
}

func TestIntSetByUint64SliceAdd(t *testing.T) {
	for _, test := range addTests {
		s := &IntSetByUint64Slice{}
		m := NewIntSetByMap()
		for _, x := range test.init {
			s.Add(x)
			m.Add(x)
		}

		s.Add(test.add)
		m.Add(test.add)

		for _, x := range test.tests {
			if got, want := s.Has(x), m.Has(x); got != want {
				t.Errorf("After add %d to %q, set.Has(%d) = %v, want %v",
					test.add, test.init, x, got, want)
			}
		}
	}
}

func TestIntSetByUint64SliceUnionWith(t *testing.T) {
	for _, test := range unionTests {
		s1, s2 := &IntSetByUint64Slice{}, &IntSetByUint64Slice{}
		m1, m2 := NewIntSetByMap(), NewIntSetByMap()
		for _, x := range test.init {
			s1.Add(x)
			m1.Add(x)
		}

		for _, x := range test.union {
			s2.Add(x)
			m2.Add(x)
		}

		s1.UnionWith(s2)
		m1.UnionWith(m2)

		for _, x := range test.tests {
			if got, want := s1.Has(x), m1.Has(x); got != want {
				t.Errorf("After union %q to %q, set.Has(%d) = %v, want %v",
					test.union, test.init, x, got, want)
			}
		}
	}
}

func TestOptimizedIntSetAdd(t *testing.T) {
	for _, test := range addTests {
		s := &OptimizedIntSet{}
		m := NewIntSetByMap()
		for _, x := range test.init {
			s.Add(x)
			m.Add(x)
		}

		s.Add(test.add)
		m.Add(test.add)

		for _, x := range test.tests {
			if got, want := s.Has(x), m.Has(x); got != want {
				t.Errorf("After add %d to %q, set.Has(%d) = %v, want %v",
					test.add, test.init, x, got, want)
			}
		}
	}
}

func TestOptimizedIntSetUnionWith(t *testing.T) {
	for _, test := range unionTests {
		s1, s2 := &OptimizedIntSet{}, &OptimizedIntSet{}
		m1, m2 := NewIntSetByMap(), NewIntSetByMap()
		for _, x := range test.init {
			s1.Add(x)
			m1.Add(x)
		}

		for _, x := range test.union {
			s2.Add(x)
			m2.Add(x)
		}

		s1.UnionWith(s2)
		m1.UnionWith(m2)

		for _, x := range test.tests {
			if got, want := s1.Has(x), m1.Has(x); got != want {
				t.Errorf("After union %q to %q, set.Has(%d) = %v, want %v",
					test.union, test.init, x, got, want)
			}
		}
	}
}

func TestIntSetByUint32SliceAdd(t *testing.T) {
	for _, test := range addTests {
		s := &IntSetByUint32Slice{}
		m := NewIntSetByMap()
		for _, x := range test.init {
			s.Add(x)
			m.Add(x)
		}

		s.Add(test.add)
		m.Add(test.add)

		for _, x := range test.tests {
			if got, want := s.Has(x), m.Has(x); got != want {
				t.Errorf("After add %d to %q, set.Has(%d) = %v, want %v",
					test.add, test.init, x, got, want)
			}
		}
	}
}

func TestIntSetByUint32SliceUnionWith(t *testing.T) {
	for _, test := range unionTests {
		s1, s2 := &IntSetByUint32Slice{}, &IntSetByUint32Slice{}
		m1, m2 := NewIntSetByMap(), NewIntSetByMap()
		for _, x := range test.init {
			s1.Add(x)
			m1.Add(x)
		}

		for _, x := range test.union {
			s2.Add(x)
			m2.Add(x)
		}

		s1.UnionWith(s2)
		m1.UnionWith(m2)

		for _, x := range test.tests {
			if got, want := s1.Has(x), m1.Has(x); got != want {
				t.Errorf("After union %q to %q, set.Has(%d) = %v, want %v",
					test.union, test.init, x, got, want)
			}
		}
	}
}
