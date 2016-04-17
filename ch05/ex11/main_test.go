package main

import "testing"

func TestTopoSort(t *testing.T) {
	in := map[string][]string{
		"a": {"b"},
		"b": {"c"},
		"c": {"d"},
		"d": {"a"},
	}
	_, err := topoSort(in)
	if err == nil {
		t.Errorf("topoSort(%v) = _, nil, but want error", in)
	}
}
