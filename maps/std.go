package maps

import (
	"iter"
	stdMaps "maps"
)

// Equal reports whether two maps contain the same key/value pairs.
// This is a convenience wrapper for [maps.Equal].
//
// See: https://pkg.go.dev/maps#Equal
func Equal[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool {
	return stdMaps.Equal(m1, m2)
}

// EqualFunc reports whether two maps contain the same key/value pairs,
// using eq to compare values with the same key.
// This is a convenience wrapper for [maps.EqualFunc].
//
// See: https://pkg.go.dev/maps#EqualFunc
func EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool {
	return stdMaps.EqualFunc(m1, m2, eq)
}

// Clone returns a copy of the provided map.
// This is a convenience wrapper for [maps.Clone].
//
// See: https://pkg.go.dev/maps#Clone
func Clone[M ~map[K]V, K comparable, V any](m M) M {
	return stdMaps.Clone(m)
}

// Copy copies all key/value pairs in src into dst, overwriting existing keys.
// This is a convenience wrapper for [maps.Copy].
//
// See: https://pkg.go.dev/maps#Copy
func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {
	stdMaps.Copy(dst, src)
}

// DeleteFunc deletes all entries from the map m for which del returns true.
// This is a convenience wrapper for [maps.DeleteFunc].
//
// See: https://pkg.go.dev/maps#DeleteFunc
func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool) {
	stdMaps.DeleteFunc(m, del)
}

// All returns an iterator over all key/value pairs in the map m.
// This is a convenience wrapper for [maps.All].
//
// See: https://pkg.go.dev/maps#All
func All[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V] {
	return stdMaps.All(m)
}

// Keys returns an iterator over all keys in the map m.
// This is a convenience wrapper for [maps.Keys].
//
// See: https://pkg.go.dev/maps#Keys
func Keys[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[K] {
	return stdMaps.Keys(m)
}

// Values returns an iterator over all values in the map m.
// This is a convenience wrapper for [maps.Values].
//
// See: https://pkg.go.dev/maps#Values
func Values[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[V] {
	return stdMaps.Values(m)
}

// Insert inserts all key/value pairs from seq into the map m.
// This is a convenience wrapper for [maps.Insert].
//
// See: https://pkg.go.dev/maps#Insert
func Insert[Map ~map[K]V, K comparable, V any](m Map, seq iter.Seq2[K, V]) {
	stdMaps.Insert(m, seq)
}

// Collect constructs a map from the key/value pairs produced by seq.
// This is a convenience wrapper for [maps.Collect].
//
// See: https://pkg.go.dev/maps#Collect
func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	return stdMaps.Collect(seq)
}
