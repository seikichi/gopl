package main

import "testing"

func TestTopoSort(t *testing.T) {
	seen := map[string]bool{}
	for _, course := range topoSort(prereqs) {
		for dep := range prereqs[course] {
			if !seen[dep] {
				t.Errorf("%s should appear before %s, but not", dep, course)
			}
		}
		seen[course] = true
	}
}
