package maps

import (
	"iter"
	stdMaps "maps"
)

func Equal[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool {
	return stdMaps.Equal(m1, m2)
}

func EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool {
	return stdMaps.EqualFunc(m1, m2, eq)
}

func Clone[M ~map[K]V, K comparable, V any](m M) M {
	return stdMaps.Clone(m)
}

func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {
	stdMaps.Copy(dst, src)
}

func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool) {
	stdMaps.DeleteFunc(m, del)
}

func All[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V] {
	return stdMaps.All(m)
}

func Keys[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[K] {
	return stdMaps.Keys(m)
}

func Values[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[V] {
	return stdMaps.Values(m)
}

func Insert[Map ~map[K]V, K comparable, V any](m Map, seq iter.Seq2[K, V]) {
	stdMaps.Insert(m, seq)
}

func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	return stdMaps.Collect(seq)
}
