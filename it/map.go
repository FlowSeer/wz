package it

import "iter"

func Map[T, R any](i iter.Seq[T], f func(T) R) iter.Seq[R] {
	panic("not implemented")
}

func SkipMap[T, R any](i iter.Seq[T], f func(T) (R, error)) iter.Seq[R] {
	panic("not implemented")
}
