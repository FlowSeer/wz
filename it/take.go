package it

import (
	"iter"
)

// Take returns an iterator that yields at most n elements from the input iterator.
// If n is negative, Take panics.
// If n is 0, the returned iterator yields no elements.
// If the input iterator has fewer than n elements, all available elements are yielded.
func Take[T any](i iter.Seq[T], n int) []T {
	panic("not implemented")
}
