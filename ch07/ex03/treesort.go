package treesort

import (
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

// String returns the sequence of values in the tree.
func (t *tree) String() string {
	values := []string{}
	if t.left != nil {
		values = append(values, t.left.String())
	}
	values = append(values, fmt.Sprintf("%d", t.value))
	if t.right != nil {
		values = append(values, t.right.String())
	}
	return fmt.Sprintf("(%s)", strings.Join(values, " "))
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
