package sl

import "slices"

// Chunk splits the slice s into chunks of size size.
// If the length of s is not a multiple of size, the last chunk will be smaller.
// If size is less than or equal to 0, Chunk panics.
func Chunk[T any, S ~[]T](s S, size int) []S {
	return slices.Collect(slices.Chunk(s, size))
}
