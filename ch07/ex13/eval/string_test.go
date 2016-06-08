package eval

import (
	"reflect"
	"testing"
)

func TestExprString(t *testing.T) {
	tests := []struct{ s string }{
		{"1 + 1"},
		{"(x + y) + z"},
		{"-sin(x + y)"},
	}
	for _, tt := range tests {
		e, _ := Parse(tt.s)

		got, err := Parse(e.String())

		if err != nil {
			t.Errorf("Parse(%v.String()) = _, %s; want _, nil", e, err)
			continue
		}

		if !reflect.DeepEqual(e, got) {
			t.Errorf("%#v != Parse(%#v.String()); got %#v", e, e, got)
		}
	}
}
