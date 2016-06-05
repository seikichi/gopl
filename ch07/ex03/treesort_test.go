package treesort

import (
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		in   *tree
		want string
	}{
		{&tree{0, nil, nil}, "(0)"},
		{&tree{0, &tree{1, nil, nil}, nil}, "((1) 0)"},
		{&tree{0, &tree{1, nil, nil}, &tree{2, nil, nil}}, "((1) 0 (2))"},
		{
			&tree{0, &tree{1, &tree{5, nil, nil}, &tree{4, nil, nil}}, &tree{2, nil, &tree{3, nil, nil}}},
			"(((5) 1 (4)) 0 (2 (3)))",
		},
	}

	for _, test := range tests {
		if got := test.in.String(); got != test.want {
			t.Errorf("%#v.String() = %q; want %q", test.in, got, test.want)
		}
	}
}
