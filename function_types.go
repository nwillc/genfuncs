package genfuncs

type Predicate[T any] func(T) bool
type KeySelector[T any, K comparable] func(T) K
type Operation[T, R any] func(R, T) R
