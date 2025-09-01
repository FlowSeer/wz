package slices

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/constraints"
)

// ToString returns a new slice of strings by calling the String method
// of each element in the input slice s. The input slice must contain
// elements that implement the fmt.Stringer interface.
//
// Example:
//
//	type person struct {
//	    name string
//	}
//	func (p person) String() string { return p.name }
//
//	people := []person{{"Alice"}, {"Bob"}}
//	names := slices.ToString(people) // []string{"Alice", "Bob"}
func ToString[S ~[]E, E fmt.Stringer](s S) []string {
	res := make([]string, len(s))

	for i, v := range s {
		res[i] = v.String()
	}

	return res
}

// ToIntString returns a new slice of strings by converting each integer element
// in the input slice s to its string representation using strconv.Itoa.
// The input slice must contain elements of an integer type.
//
// Example:
//
//	ints := []int{1, 2, 3}
//	strs := slices.ToIntString(ints) // []string{"1", "2", "3"}
func ToIntString[S ~[]E, E constraints.Integer](s S) []string {
	res := make([]string, len(s))

	for i, v := range s {
		res[i] = strconv.Itoa(int(v))
	}

	return res
}

// ToFloatString returns a new slice of strings by converting each float element
// in the input slice s to its string representation using strconv.FormatFloat.
// The input slice must contain elements of a float type.
//
// Example:
//
//	floats := []float64{1.23, 4.56}
//	strs := slices.ToFloatString(floats) // []string{"1.23", "4.56"}
func ToFloatString[S ~[]E, E constraints.Float](s S) []string {
	res := make([]string, len(s))

	for i, v := range s {
		res[i] = strconv.FormatFloat(float64(v), 'f', -1, 64)
	}

	return res
}
