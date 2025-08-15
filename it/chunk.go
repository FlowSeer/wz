package it

import (
	"iter"
)

func Chunk[T any](it iter.Seq[T], size int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		chunk := make([]T, 0, size)
		for v := range it {
			chunk = append(chunk, v)

			if len(chunk) == size {
				if !yield(chunk) {
					return
				}

				chunk = make([]T, 0, size)
			}
		}
	}
}
