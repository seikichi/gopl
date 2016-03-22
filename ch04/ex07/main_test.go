package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	s    []byte
	want []byte
}{
	{[]byte(""), []byte("")},
	{[]byte("あ!!"), []byte("!!あ")},
	{[]byte("!あ!"), []byte("!あ!")},
	{[]byte("Hello"), []byte("olleH")},
	{[]byte("こんにちは!!"), []byte("!!はちにんこ")},
	{[]byte("こんにちは世界"), []byte("界世はちにんこ")},
}

func TestRotate(t *testing.T) {
	for _, tt := range tests {
		got := reverse(append([]byte{}, tt.s...))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("uniq(%q) = %q, but want %q\n", tt.s, got, tt.want)
		}
	}
}
