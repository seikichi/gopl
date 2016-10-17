package sexpr

import (
	"reflect"
	"strings"
	"testing"
)

func TestToken(t *testing.T) {
	const sexprStream = `
((Name "seikichi") (Text "Hello, world!"))
((1 2) (3 (4 5)))
`

	dec := NewDecoder(strings.NewReader(sexprStream))

	wants := []Token{
		StartList{},
		StartList{}, Symbol{"Name"}, String{"seikichi"}, EndList{},
		StartList{}, Symbol{"Text"}, String{"Hello, world!"}, EndList{},
		EndList{},

		StartList{},
		StartList{}, Int{1}, Int{2}, EndList{},
		StartList{}, Int{3}, StartList{}, Int{4}, Int{5}, EndList{}, EndList{},
		EndList{},
	}

	for _, want := range wants {
		got, err := dec.Token()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("dec.Token() = (%v, %v); want (%v, nil)",
				got, err, want)
		}
	}
}
