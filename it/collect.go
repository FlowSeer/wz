package it

import (
	"iter"
	"slices"
)

// Collect exhausts the iterator i and collects all yielded elements into a slice of type S.
// S must be a slice type of T (e.g., []T or a named slice type).
// The returned slice contains all elements produced by the iterator, in order.
func Collect[T any, S ~[]T](i iter.Seq[T]) S {
	return slices.Collect(i)
}
