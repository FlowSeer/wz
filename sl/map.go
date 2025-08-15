package sl

import "fmt"

// Map applies the function f to each element of the slice s and returns a new slice
// containing the results. The output is of the same length as the input.
func Map[T any, S ~[]T, R any, RS ~[]R](s S, f func(T) R) RS {
	res := make(RS, len(s))

	for i, v := range s {
		res[i] = f(v)
	}

	return res
}

// MapUntil applies the function f to each element of the slice s and returns a new slice
// containing the mapped results. The mapping function f returns a value and a boolean flag.
// Mapping stops at the first element for which f returns false, and the resulting slice
// contains only the mapped values up to (but not including) that point. The output slice
// may be shorter than the input.
func MapUntil[T any, S ~[]T, R any, RS ~[]R](s S, f func(T) (R, bool)) RS {
	res := make(RS, 0, len(s))

	for _, v := range s {
		r, ok := f(v)
		if !ok {
			break
		}
		res = append(res, r)
	}

	return res
}

// TryMap applies the function f to each element of the slice s and returns a new slice
// containing the results. If f returns an error for any element, TryMap returns nil and
// the error, wrapping the index of the failed element. Otherwise, it returns the mapped slice.
func TryMap[T any, S ~[]T, R any, RS ~[]R](s S, f func(T) (R, error)) (RS, error) {
	res := make(RS, len(s))

	for i, v := range s {
		r, err := f(v)
		if err != nil {
			return nil, fmt.Errorf("failed to map element %d: %w", i, err)
		}
		res[i] = r
	}

	return res, nil
}

// SkipMap applies the function f to each element of the slice s and returns a new slice
// containing only the successfully mapped results. If f returns an error for an element,
// that element is skipped. The output slice may be shorter than the input.
func SkipMap[T any, S ~[]T, R any, RS ~[]R](s S, f func(T) (R, error)) RS {
	// We assume the optimistic case that the mapping function will not return an error.
	res := make(RS, 0, len(s))

	for _, v := range s {
		r, err := f(v)
		if err != nil {
			continue
		}

		res = append(res, r)
	}

	return res
}
