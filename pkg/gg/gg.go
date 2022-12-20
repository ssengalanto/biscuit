package gg

import "golang.org/x/exp/constraints"

// Prepend prepends an element in the slice.
func Prepend[T any](x []T, y T) []T {
	x = append([]T{y}, x...)
	return x
}

// FindIndexOf searches an element in a slice based on a predicate and returns the index.
// It returns -1 if the element is not found.
func FindIndexOf[T any](collection []T, predicate func(item T) bool) int {
	for i, item := range collection {
		if predicate(item) {
			return i
		}
	}

	return -1
}

// Itob converts integer value to boolean value.
func Itob[T constraints.Integer](n T) bool {
	return n > 0
}
