package sexpr

import (
	"reflect"
	"strings"
	"testing"
)

func TestDecoder(t *testing.T) {
	type Message struct {
		Name, Text string
	}

	const sexprStream = `
((Name "seikichi") (Text "Hello, world!"))
((Name "こんどう") (Text "こんにちは！"))
`

	var got, want Message
	var err error
	dec := NewDecoder(strings.NewReader(sexprStream))

	err = dec.Decode(&got)
	want = Message{"seikichi", "Hello, world!"}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("dec.Decode(&got) = %#v (got becomes %#v); want nil (got becomes %#v)",
			err, got, want)
	}

	err = dec.Decode(&got)
	want = Message{"こんどう", "こんにちは！"}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("dec.Decode(&got) = %#v (got becomes %#v); want nil (got becomes %#v)",
			err, got, want)
	}

	err = dec.Decode(&got)
	if err == nil {
		t.Errorf("dec.Decode(&got) = nil; want !nil")
	}
}
