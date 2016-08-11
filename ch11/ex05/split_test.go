package split

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s    string
		sep  string
		want []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"a,b", ",", []string{"a", "b"}},
		{"a", ",", []string{"a"}},
	}

	for _, test := range tests {
		if got := strings.Split(test.s, test.sep); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Split(%q, %q) returned %q, want %q",
				test.s, test.sep, got, test.want)
		}
	}
}
