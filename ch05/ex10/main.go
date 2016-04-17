package main

import "fmt"

func set(ss ...string) map[string]bool {
	ret := map[string]bool{}
	for _, s := range ss {
		ret[s] = true
	}
	return ret
}

var prereqs = map[string]map[string]bool{
	"algorithms": set("data structures"),
	"calculus":   set("linear algebra"),

	"compilers": set(
		"data structures",
		"formal languages",
		"computer organization",
	),

	"data structures":       set("discrete math"),
	"databases":             set("data structures"),
	"discrete math":         set("intro to programming"),
	"formal languages":      set("discrete math"),
	"networks":              set("operating systems"),
	"operating systems":     set("data structures", "computer organization"),
	"programming languages": set("data structures", "computer organization"),
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(item map[string]bool)

	visitAll = func(items map[string]bool) {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	for item := range m {
		visitAll(map[string]bool{item: true})
	}
	return order
}
