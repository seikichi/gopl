package sexpr

import "testing"

var tests = []struct {
	in   interface{}
	want string
}{
	{true, "t"},
	{false, "nil"},
	{1.0, "1.000000"},
	{1 + 2i, "#C(1.000000 2.000000)"},
	{struct{ v interface{} }{[]int{1, 2, 3}}, `((v ("[]int" (1 2 3))))`},
	{struct{ v interface{} }{nil}, `((v nil))`},
}

func TestMarshal(t *testing.T) {
	for _, tt := range tests {
		b, err := Marshal(tt.in)
		if err != nil || string(b) != tt.want {
			t.Errorf("Marshal(%v) = %q, %v; want %q, nil", tt.in, b, err, tt.want)
		}
	}
}
