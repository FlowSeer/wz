package slices

import "fmt"

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
