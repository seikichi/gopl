package intset

import (
	"reflect"
	"testing"
)

func toIntSet(xs ...int) *IntSet {
	s := &IntSet{}
	for _, x := range xs {
		s.Add(x)
	}
	return s
}

func eq(lhs, rhs *IntSet) bool {
	for i := 0; i < len(lhs.words); i++ {
		if i < len(rhs.words) && lhs.words[i] != rhs.words[i] {
			return false
		}
		if i >= len(rhs.words) && lhs.words[i] != 0 {
			return false
		}
	}
	for i := len(lhs.words); i < len(rhs.words); i++ {
		if rhs.words[i] != 0 {
			return false
		}
	}
	return true
}

var lenTests = []struct {
	in   *IntSet
	want int
}{
	{in: toIntSet(), want: 0},
	{in: toIntSet(1), want: 1},
	{in: toIntSet(1, 2, 3), want: 3},
	{in: toIntSet(1, 9, 42, 144), want: 4},
}

func TestLen(t *testing.T) {
	for _, tt := range lenTests {
		if got := tt.in.Len(); got != tt.want {
			t.Errorf("%s.Len() = %d; want %d", tt.in, got, tt.want)
		}
	}
}

var removeTests = []struct {
	elems []int
	x     int
	want  *IntSet
}{
	{elems: []int{}, x: 42, want: toIntSet()},
	{elems: []int{1}, x: 1, want: toIntSet()},
	{elems: []int{1, 2, 3}, x: 3, want: toIntSet(1, 2)},
	{elems: []int{1, 9, 42, 144}, x: 42, want: toIntSet(1, 9, 144)},
	{elems: []int{1, 9, 42, 144}, x: 10, want: toIntSet(1, 9, 42, 144)},
}

func TestRemove(t *testing.T) {
	for _, tt := range removeTests {
		in := toIntSet(tt.elems...)
		in.Remove(tt.x)
		if !eq(in, tt.want) {
			t.Errorf("%s.Remove(%d) changes the set to %s; want %s",
				toIntSet(tt.elems...), tt.x, in, tt.want)
		}
	}
}

var clearTests = []struct {
	elems []int
}{
	{elems: []int{}},
	{elems: []int{1}},
	{elems: []int{1, 2, 3}},
	{elems: []int{1, 9, 42, 144}},
}

func TestClear(t *testing.T) {
	for _, tt := range clearTests {
		in := toIntSet(tt.elems...)
		in.Clear()
		if !eq(in, &IntSet{}) {
			t.Errorf("%s.Clear() changes the set to %s; want {}",
				toIntSet(tt.elems...), in)
		}
	}
}

var copyTests = []struct {
	elems []int
}{
	{elems: []int{}},
	{elems: []int{1}},
	{elems: []int{1, 2, 3}},
	{elems: []int{1, 9, 42, 144}},
}

func TestCopy(t *testing.T) {
	for _, tt := range clearTests {
		in := toIntSet(tt.elems...)
		got := in.Copy()
		if !eq(in, got) {
			t.Errorf("%s.Copy() = %s; want %s", in, got, in)
		}
	}
}

var addAllTests = []struct {
	elems []int
	args  []int
	want  *IntSet
}{
	{elems: []int{}, args: []int{}, want: toIntSet()},
	{elems: []int{}, args: []int{1}, want: toIntSet(1)},
	{elems: []int{1}, args: []int{1}, want: toIntSet(1)},
	{elems: []int{1, 2}, args: []int{42, 144}, want: toIntSet(1, 2, 42, 144)},
}

func TestAddAll(t *testing.T) {
	for _, tt := range addAllTests {
		in := toIntSet(tt.elems...)
		in.AddAll(tt.args...)
		if !eq(in, tt.want) {
			t.Errorf("%s.AddAll(%v...) changes the set to %s; want %s",
				toIntSet(tt.elems...), tt.args, in, tt.want)
		}
	}
}

var intersectWithTests = []struct {
	in, arg, want *IntSet
}{
	{in: toIntSet(), arg: toIntSet(), want: toIntSet()},
	{in: toIntSet(1), arg: toIntSet(), want: toIntSet()},
	{in: toIntSet(), arg: toIntSet(1), want: toIntSet()},
	{in: toIntSet(1), arg: toIntSet(1), want: toIntSet(1)},
	{
		in:   toIntSet(1, 9, 44, 144),
		arg:  toIntSet(2, 9, 144, 200),
		want: toIntSet(9, 144),
	},
}

func TestIntersectWith(t *testing.T) {
	for _, tt := range intersectWithTests {
		original := tt.in.Copy()
		tt.in.IntersectWith(tt.arg)
		if !eq(tt.in, tt.want) {
			t.Errorf("%s.IntersectWith(%v) changes the set to %s; want %s",
				original, tt.arg, tt.in, tt.want)
		}
	}
}

var differenceWithTests = []struct {
	in, arg, want *IntSet
}{
	{in: toIntSet(), arg: toIntSet(), want: toIntSet()},
	{in: toIntSet(1), arg: toIntSet(), want: toIntSet(1)},
	{in: toIntSet(), arg: toIntSet(1), want: toIntSet()},
	{in: toIntSet(1), arg: toIntSet(1), want: toIntSet()},
	{
		in:   toIntSet(1, 9, 44, 144),
		arg:  toIntSet(2, 9, 144, 200),
		want: toIntSet(1, 44),
	},
}

func TestDifferenceWith(t *testing.T) {
	for _, tt := range differenceWithTests {
		original := tt.in.Copy()
		tt.in.DifferenceWith(tt.arg)
		if !eq(tt.in, tt.want) {
			t.Errorf("%s.DifferenceWith(%v) changes the set to %s; want %s",
				original, tt.arg, tt.in, tt.want)
		}
	}
}

var symmetricDifferenceWithTests = []struct {
	in, arg, want *IntSet
}{
	{in: toIntSet(), arg: toIntSet(), want: toIntSet()},
	{in: toIntSet(1), arg: toIntSet(), want: toIntSet(1)},
	{in: toIntSet(), arg: toIntSet(1), want: toIntSet(1)},
	{in: toIntSet(1), arg: toIntSet(1), want: toIntSet()},
	{
		in:   toIntSet(1, 9, 44, 144),
		arg:  toIntSet(2, 9, 144, 200),
		want: toIntSet(1, 2, 44, 200),
	},
}

func TestSymmetricDifferenceWith(t *testing.T) {
	for _, tt := range symmetricDifferenceWithTests {
		original := tt.in.Copy()
		tt.in.SymmetricDifferenceWith(tt.arg)
		if !eq(tt.in, tt.want) {
			t.Errorf("%s.SymmetricDifferenceWith(%v) changes the set to %s; want %s",
				original, tt.arg, tt.in, tt.want)
		}
	}
}

var elemsTest = []struct {
	in   *IntSet
	want []uint64
}{
	{in: toIntSet(), want: []uint64{}},
	{in: toIntSet(1), want: []uint64{1}},
	{in: toIntSet(1, 100), want: []uint64{1, 100}},
	{in: toIntSet(1, 9, 16, 100), want: []uint64{1, 9, 16, 100}},
}

func TestElems(t *testing.T) {
	for _, tt := range elemsTest {
		got := tt.in.Elems()
		if got == nil {
			got = []uint64{}
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s.Elems() = %v; want %v", tt.in, got, tt.want)
		}
	}
}
