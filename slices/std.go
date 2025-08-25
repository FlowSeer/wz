package slices

import (
	"cmp"
	"iter"
	stdSlices "slices"
)

func Equal[S ~[]E, E comparable](s1, s2 S) bool {
	return stdSlices.Equal(s1, s2)
}

func EqualFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, eq func(E1, E2) bool) bool {
	return stdSlices.EqualFunc(s1, s2, eq)
}

func Compare[S ~[]E, E cmp.Ordered](s1, s2 S) int {
	return stdSlices.Compare(s1, s2)
}

func CompareFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, cmp func(E1, E2) int) int {
	return stdSlices.CompareFunc(s1, s2, cmp)
}

func Index[S ~[]E, E comparable](s S, v E) int {
	return stdSlices.Index(s, v)
}

func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	return stdSlices.IndexFunc(s, f)
}

func Contains[S ~[]E, E comparable](s S, v E) bool {
	return stdSlices.Contains(s, v)
}

func ContainsFunc[S ~[]E, E any](s S, f func(E) bool) bool {
	return stdSlices.ContainsFunc(s, f)
}

func Insert[S ~[]E, E any](s S, i int, v ...E) S {
	return stdSlices.Insert(s, i, v...)
}

func Delete[S ~[]E, E any](s S, i, j int) S {
	return stdSlices.Delete(s, i, j)
}

func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S {
	return stdSlices.DeleteFunc(s, del)
}

func Replace[S ~[]E, E any](s S, i, j int, v ...E) S {
	return stdSlices.Replace(s, i, j, v...)
}

func Clone[S ~[]E, E any](s S) S {
	return stdSlices.Clone(s)
}

func Compact[S ~[]E, E comparable](s S) S {
	return stdSlices.Compact(s)
}

func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S {
	return stdSlices.CompactFunc(s, eq)
}

func Grow[S ~[]E, E any](s S, n int) S {
	return stdSlices.Grow(s, n)
}

func Clip[S ~[]E, E any](s S) S {
	return stdSlices.Clip(s)
}

func Backward[Slice ~[]E, E any](s Slice) iter.Seq2[int, E] {
	return stdSlices.Backward(s)
}

func Values[Slice ~[]E, E any](s Slice) iter.Seq[E] {
	return stdSlices.Values(s)
}

func AppendSeq[Slice ~[]E, E any](s Slice, seq iter.Seq[E]) Slice {
	return stdSlices.AppendSeq(s, seq)
}

func Collect[E any](seq iter.Seq[E]) []E {
	return stdSlices.Collect(seq)
}

func Sorted[E cmp.Ordered](seq iter.Seq[E]) []E {
	return stdSlices.Sorted(seq)
}

func SortedFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E {
	return stdSlices.SortedFunc(seq, cmp)
}

func SortedStableFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E {
	return stdSlices.SortedStableFunc(seq, cmp)
}

func Chunk[Slice ~[]E, E any](s Slice, n int) iter.Seq[Slice] {
	return stdSlices.Chunk(s, n)
}

func Sort[S ~[]E, E cmp.Ordered](x S) {
	stdSlices.Sort(x)
}

func SortFunc[S ~[]E, E any](x S, cmp func(a, b E) int) {
	stdSlices.SortFunc(x, cmp)
}

func SortStableFunc[S ~[]E, E any](x S, cmp func(a, b E) int) {
	stdSlices.SortStableFunc(x, cmp)
}

func IsSorted[S ~[]E, E cmp.Ordered](x S) bool {
	return stdSlices.IsSorted(x)
}

func IsSortedFunc[S ~[]E, E any](x S, cmp func(a, b E) int) bool {
	return stdSlices.IsSortedFunc(x, cmp)
}

func Min[S ~[]E, E cmp.Ordered](x S) E {
	return stdSlices.Min(x)
}

func MinFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E {
	return stdSlices.MinFunc(x, cmp)
}

func Max[S ~[]E, E cmp.Ordered](x S) E {
	return stdSlices.Max(x)
}

func MaxFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E {
	return stdSlices.MaxFunc(x, cmp)
}

func BinarySearch[S ~[]E, E cmp.Ordered](x S, target E) (int, bool) {
	return stdSlices.BinarySearch(x, target)
}

func BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool) {
	return stdSlices.BinarySearchFunc(x, target, cmp)
}
