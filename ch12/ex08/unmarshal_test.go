package sexpr

import (
	"reflect"
	"testing"
)

func TestUnmarshalBool(t *testing.T) {
	type s struct {
		Flag bool
	}

	tests := []struct {
		in   []byte
		want s
	}{
		{[]byte("((Flag t))"), s{true}},
		{[]byte("((Flag nil))"), s{false}},
	}

	for _, test := range tests {
		var got s
		err := Unmarshal(test.in, &got)
		if err != nil || !reflect.DeepEqual(test.want, got) {
			t.Errorf("Unmarshal([]byte(\"%s\"), &got) = %v, (got becomes %v); want nil (got becomes %v)",
				test.in, err, got, test.want)
		}
	}
}

func TestUnmarshalComplex(t *testing.T) {
	type s struct {
		Val complex128
	}

	tests := []struct {
		in   []byte
		want s
	}{
		{[]byte("((Val #C(1.0 0.0)))"), s{1}},
		{[]byte("((Val #C(0.0 1.0)))"), s{1i}},
		{[]byte("((Val #C(1.0 2.0)))"), s{1 + 2i}},
	}

	for _, test := range tests {
		var got s
		err := Unmarshal(test.in, &got)
		if err != nil || !reflect.DeepEqual(test.want, got) {
			t.Errorf("Unmarshal([]byte(\"%s\"), &got) = %v, (got becomes %v); want nil (got becomes %v)",
				test.in, err, got, test.want)
		}
	}
}

func TestUnmarshalInterface(t *testing.T) {
	type s struct {
		Val interface{}
	}

	AddType("[]int", reflect.TypeOf([]int{}))

	tests := []struct {
		in   []byte
		want s
	}{
		{[]byte(`((Val ("[]int" (1 2 3))))`), s{[]int{1, 2, 3}}},
	}

	for _, test := range tests {
		var got s
		err := Unmarshal(test.in, &got)
		if err != nil || !reflect.DeepEqual(test.want, got) {
			t.Errorf("Unmarshal([]byte(\"%s\"), &got) = %v, (got becomes %v); want nil (got becomes %v)",
				test.in, err, got, test.want)
		}
	}
}
