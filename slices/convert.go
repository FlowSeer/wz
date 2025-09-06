package slices

import "golang.org/x/exp/constraints"

// ConvertIntegers converts a slice of integers to a slice of other integers.
func ConvertIntegers[S ~[]E, OS ~[]O, E constraints.Integer, O constraints.Integer](s S) OS {
	res := make([]O, len(s))
	for i, v := range s {
		res[i] = O(v)
	}

	return res
}

// ConvertFloats converts a slice of floats to a slice of other floats.
func ConvertFloats[S ~[]E, OS ~[]O, E constraints.Float, O constraints.Float](s S) OS {
	res := make([]O, len(s))
	for i, v := range s {
		res[i] = O(v)
	}

	return res
}
