package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"}, // added

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	corses, err := topoSort(prereqs)
	if err != nil {
		log.Fatal(err)
	}
	for i, course := range corses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	temporary := make(map[string]bool)
	var visitAll func(items []string)

	cycle := false
	visitAll = func(items []string) {
		if cycle {
			return
		}

		for _, item := range items {
			if temporary[item] {
				cycle = true
				return
			}

			if !seen[item] {
				seen[item] = true

				temporary[item] = true
				visitAll(m[item])
				temporary[item] = false

				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)

	if cycle {
		return nil, errors.New("The graph has circular dependency")
	}
	return order, nil
}
