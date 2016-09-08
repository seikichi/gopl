package sexpr

import "testing"

var tests = []struct {
	in   interface{}
	want string
}{
	{true, "true"},
	{false, "false"},
	{1.0, "1.000000"},
	{[]int{1, 2, 3}, "[1, 2, 3]"},
	{map[string]int{"x": 10, "y": 20}, `{"x": 10, "y": 20}`},
	{map[string][]int{"x": []int{10}, "y": []int{1, 2}}, `{"x": [10], "y": [1, 2]}`},
	{struct {
		x string
		y []int
	}{"A", []int{1, 2, 3}}, `{"x": "A", "y": [1, 2, 3]}`},
}

func TestMarshal(t *testing.T) {
	for _, tt := range tests {
		b, err := Marshal(tt.in)
		if err != nil || string(b) != tt.want {
			t.Errorf("Marshal(%v) = %q, %v; want %q, nil", tt.in, b, err, tt.want)
		}
	}
}
