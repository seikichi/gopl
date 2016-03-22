package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	s    []byte
	want []byte
}{
	{[]byte("Hello"), []byte("Hello")},
	{[]byte("こんにちは　 　世界"), []byte("こんにちは 世界")},
	{[]byte("　 Ｇ　　ｏ 　"), []byte(" Ｇ ｏ ")},
}

func TestRotate(t *testing.T) {
	for _, tt := range tests {
		s := append([]byte{}, tt.s...)
		got := compress(s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("uniq(%q) = %q, but want %q\n", tt.s, got, tt.want)
		}
	}
}
