package genfuncs

type Predicate[T any] func(T) bool
type KeySelector[T any, K comparable] func(T) K
type ValueSelector[K comparable, T any] func(K) T
type Operation[T, R any] func(R, T) R
type Stringer[T any] func(T) string
type TransformKV[T any, K comparable, V any] func(T) (K,V)
type Transform[T, R any]  func(T) R
