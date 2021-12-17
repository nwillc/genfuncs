package genfuncs

// KeySelector is used for generating keys from types, it accepts any type and returns a comparable key for it.
type KeySelector[T any, K comparable] func(T) K

// Operation is used to perform operations on its arguments, it accepts two arguments of any type and returns a result
// of the type of the first argument.
type Operation[T, R any] func(R, T) R

// Predicate is used evaluate a value, it accepts any type and returns a bool.
type Predicate[T any] func(T) bool

// Stringer is used to create string representations, it accepts any type and returns a string.
type Stringer[T any] func(T) string

// Transform is used to transform values and types, it accepts an argument of any type and returns any type.
type Transform[T, R any] func(T) R

// TransformKV is used to generate a key and value from a type, it accepts any type, and returns a comparable key and
// any value.
type TransformKV[T any, K comparable, V any] func(T) (K, V)

// ValueSelector is used to select a value for a key, given a comparable key will return a value of any type.
type ValueSelector[K comparable, T any] func(K) T
