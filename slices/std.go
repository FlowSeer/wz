package slices

import (
	"cmp"
	"iter"
	stdSlices "slices"
)

// Equal reports whether two slices are equal.
// This is a convenience wrapper for [slices.Equal].
//
// See: https://pkg.go.dev/slices#Equal
func Equal[S ~[]E, E comparable](s1, s2 S) bool {
	return stdSlices.Equal(s1, s2)
}

// EqualFunc reports whether two slices are equal using a custom equality function.
// This is a convenience wrapper for [slices.EqualFunc].
//
// See: https://pkg.go.dev/slices#EqualFunc
func EqualFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, eq func(E1, E2) bool) bool {
	return stdSlices.EqualFunc(s1, s2, eq)
}

// Compare compares two slices lexicographically.
// This is a convenience wrapper for [slices.Compare].
//
// See: https://pkg.go.dev/slices#Compare
func Compare[S ~[]E, E cmp.Ordered](s1, s2 S) int {
	return stdSlices.Compare(s1, s2)
}

// CompareFunc compares two slices lexicographically using a custom comparison function.
// This is a convenience wrapper for [slices.CompareFunc].
//
// See: https://pkg.go.dev/slices#CompareFunc
func CompareFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, cmp func(E1, E2) int) int {
	return stdSlices.CompareFunc(s1, s2, cmp)
}

// Index returns the index of the first occurrence of v in s, or -1 if not present.
// This is a convenience wrapper for [slices.Index].
//
// See: https://pkg.go.dev/slices#Index
func Index[S ~[]E, E comparable](s S, v E) int {
	return stdSlices.Index(s, v)
}

// IndexFunc returns the first index i satisfying f(s[i]), or -1 if none do.
// This is a convenience wrapper for [slices.IndexFunc].
//
// See: https://pkg.go.dev/slices#IndexFunc
func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	return stdSlices.IndexFunc(s, f)
}

// Contains reports whether v is present in s.
// This is a convenience wrapper for [slices.Contains].
//
// See: https://pkg.go.dev/slices#Contains
func Contains[S ~[]E, E comparable](s S, v E) bool {
	return stdSlices.Contains(s, v)
}

// ContainsFunc reports whether any element of s satisfies f.
// This is a convenience wrapper for [slices.ContainsFunc].
//
// See: https://pkg.go.dev/slices#ContainsFunc
func ContainsFunc[S ~[]E, E any](s S, f func(E) bool) bool {
	return stdSlices.ContainsFunc(s, f)
}

// Insert inserts the values v... into s at index i, returning the modified slice.
// This is a convenience wrapper for [slices.Insert].
//
// See: https://pkg.go.dev/slices#Insert
func Insert[S ~[]E, E any](s S, i int, v ...E) S {
	return stdSlices.Insert(s, i, v...)
}

// Delete removes the elements s[i:j] from s, returning the modified slice.
// This is a convenience wrapper for [slices.Delete].
//
// See: https://pkg.go.dev/slices#Delete
func Delete[S ~[]E, E any](s S, i, j int) S {
	return stdSlices.Delete(s, i, j)
}

// DeleteFunc deletes elements from s for which del returns true.
// This is a convenience wrapper for [slices.DeleteFunc].
//
// See: https://pkg.go.dev/slices#DeleteFunc
func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S {
	return stdSlices.DeleteFunc(s, del)
}

// Replace replaces the elements s[i:j] with v..., returning the modified slice.
// This is a convenience wrapper for [slices.Replace].
//
// See: https://pkg.go.dev/slices#Replace
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S {
	return stdSlices.Replace(s, i, j, v...)
}

// Clone returns a copy of the provided slice.
// This is a convenience wrapper for [slices.Clone].
//
// See: https://pkg.go.dev/slices#Clone
func Clone[S ~[]E, E any](s S) S {
	return stdSlices.Clone(s)
}

// Compact removes consecutive duplicate elements from s.
// This is a convenience wrapper for [slices.Compact].
//
// See: https://pkg.go.dev/slices#Compact
func Compact[S ~[]E, E comparable](s S) S {
	return stdSlices.Compact(s)
}

// CompactFunc removes consecutive elements from s that are equal according to eq.
// This is a convenience wrapper for [slices.CompactFunc].
//
// See: https://pkg.go.dev/slices#CompactFunc
func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S {
	return stdSlices.CompactFunc(s, eq)
}

// Grow increases the capacity of s, returning the updated slice.
// This is a convenience wrapper for [slices.Grow].
//
// See: https://pkg.go.dev/slices#Grow
func Grow[S ~[]E, E any](s S, n int) S {
	return stdSlices.Grow(s, n)
}

// Clip returns a slice with the same elements as s, but with no extra capacity.
// This is a convenience wrapper for [slices.Clip].
//
// See: https://pkg.go.dev/slices#Clip
func Clip[S ~[]E, E any](s S) S {
	return stdSlices.Clip(s)
}

// Backward returns an iterator over the elements of s in reverse order.
// This is a convenience wrapper for [slices.Backward].
//
// See: https://pkg.go.dev/slices#Backward
func Backward[Slice ~[]E, E any](s Slice) iter.Seq2[int, E] {
	return stdSlices.Backward(s)
}

// Values returns an iterator over the values of s.
// This is a convenience wrapper for [slices.Values].
//
// See: https://pkg.go.dev/slices#Values
func Values[Slice ~[]E, E any](s Slice) iter.Seq[E] {
	return stdSlices.Values(s)
}

// AppendSeq appends the values from seq to s, returning the updated slice.
// This is a convenience wrapper for [slices.AppendSeq].
//
// See: https://pkg.go.dev/slices#AppendSeq
func AppendSeq[Slice ~[]E, E any](s Slice, seq iter.Seq[E]) Slice {
	return stdSlices.AppendSeq(s, seq)
}

// Collect collects all values from seq into a slice.
// This is a convenience wrapper for [slices.Collect].
//
// See: https://pkg.go.dev/slices#Collect
func Collect[E any](seq iter.Seq[E]) []E {
	return stdSlices.Collect(seq)
}

// Sorted returns a sorted slice containing the values from seq.
// This is a convenience wrapper for [slices.Sorted].
//
// See: https://pkg.go.dev/slices#Sorted
func Sorted[E cmp.Ordered](seq iter.Seq[E]) []E {
	return stdSlices.Sorted(seq)
}

// SortedFunc returns a sorted slice containing the values from seq, using cmp for ordering.
// This is a convenience wrapper for [slices.SortedFunc].
//
// See: https://pkg.go.dev/slices#SortedFunc
func SortedFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E {
	return stdSlices.SortedFunc(seq, cmp)
}

// SortedStableFunc returns a stably sorted slice containing the values from seq, using cmp for ordering.
// This is a convenience wrapper for [slices.SortedStableFunc].
//
// See: https://pkg.go.dev/slices#SortedStableFunc
func SortedStableFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E {
	return stdSlices.SortedStableFunc(seq, cmp)
}

// Chunk splits s into chunks of size n, returning an iterator over the chunks.
// This is a convenience wrapper for [slices.Chunk].
//
// See: https://pkg.go.dev/slices#Chunk
func Chunk[Slice ~[]E, E any](s Slice, n int) iter.Seq[Slice] {
	return stdSlices.Chunk(s, n)
}

// Sort sorts the slice x in increasing order.
// This is a convenience wrapper for [slices.Sort].
//
// See: https://pkg.go.dev/slices#Sort
func Sort[S ~[]E, E cmp.Ordered](x S) {
	stdSlices.Sort(x)
}

// SortFunc sorts the slice x using a custom comparison function.
// This is a convenience wrapper for [slices.SortFunc].
//
// See: https://pkg.go.dev/slices#SortFunc
func SortFunc[S ~[]E, E any](x S, cmp func(a, b E) int) {
	stdSlices.SortFunc(x, cmp)
}

// SortStableFunc stably sorts the slice x using a custom comparison function.
// This is a convenience wrapper for [slices.SortStableFunc].
//
// See: https://pkg.go.dev/slices#SortStableFunc
func SortStableFunc[S ~[]E, E any](x S, cmp func(a, b E) int) {
	stdSlices.SortStableFunc(x, cmp)
}

// IsSorted reports whether x is sorted in increasing order.
// This is a convenience wrapper for [slices.IsSorted].
//
// See: https://pkg.go.dev/slices#IsSorted
func IsSorted[S ~[]E, E cmp.Ordered](x S) bool {
	return stdSlices.IsSorted(x)
}

// IsSortedFunc reports whether x is sorted according to cmp.
// This is a convenience wrapper for [slices.IsSortedFunc].
//
// See: https://pkg.go.dev/slices#IsSortedFunc
func IsSortedFunc[S ~[]E, E any](x S, cmp func(a, b E) int) bool {
	return stdSlices.IsSortedFunc(x, cmp)
}

// Min returns the minimum element in x.
// This is a convenience wrapper for [slices.Min].
//
// See: https://pkg.go.dev/slices#Min
func Min[S ~[]E, E cmp.Ordered](x S) E {
	return stdSlices.Min(x)
}

// MinFunc returns the minimum element in x according to cmp.
// This is a convenience wrapper for [slices.MinFunc].
//
// See: https://pkg.go.dev/slices#MinFunc
func MinFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E {
	return stdSlices.MinFunc(x, cmp)
}

// Max returns the maximum element in x.
// This is a convenience wrapper for [slices.Max].
//
// See: https://pkg.go.dev/slices#Max
func Max[S ~[]E, E cmp.Ordered](x S) E {
	return stdSlices.Max(x)
}

// MaxFunc returns the maximum element in x according to cmp.
// This is a convenience wrapper for [slices.MaxFunc].
//
// See: https://pkg.go.dev/slices#MaxFunc
func MaxFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E {
	return stdSlices.MaxFunc(x, cmp)
}

// BinarySearch searches for target in a sorted slice x.
// This is a convenience wrapper for [slices.BinarySearch].
//
// See: https://pkg.go.dev/slices#BinarySearch
func BinarySearch[S ~[]E, E cmp.Ordered](x S, target E) (int, bool) {
	return stdSlices.BinarySearch(x, target)
}

// BinarySearchFunc searches for target in a sorted slice x using cmp.
// This is a convenience wrapper for [slices.BinarySearchFunc].
//
// See: https://pkg.go.dev/slices#BinarySearchFunc
func BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool) {
	return stdSlices.BinarySearchFunc(x, target, cmp)
}
