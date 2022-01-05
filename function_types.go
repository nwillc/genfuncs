/*
 *  Copyright (c) 2021,  nwillc@gmail.com
 *
 *  Permission to use, copy, modify, and/or distribute this software for any
 *  purpose with or without fee is hereby granted, provided that the above
 *  copyright notice and this permission notice appear in all copies.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 *  WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 *  MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 *  ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 *  WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 *  ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 *  OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package genfuncs

import (
	"constraints"
	"fmt"
)

// BiFunction accepts two arguments and produces a result.
type BiFunction[T, U, R any] func(T, U) R

// LessThan compares two arguments of the same type and returns true if the first is less than the second.
type LessThan[T any] BiFunction[T, T, bool]

// Function accepts one argument and produces a result.
type Function[T, R any] func(T) R

// KeyFor is used for generating keys from types, it accepts any type and returns a comparable key for it.
type KeyFor[T any, K comparable] Function[T, K]

// KeyValueFor is used to generate a key and value from a type, it accepts any type, and returns a comparable key and
// any value.
type KeyValueFor[T any, K comparable, V any] func(T) (K, V)

// Predicate is used evaluate a value, it accepts any type and returns a bool.
type Predicate[T any] func(T) bool

func (p Predicate[T]) Not() Predicate[T] { return func(a T) bool { return !p(a) } }

// IsLessThan creates a Predicate that tests if its argument is less than a given value.
func IsLessThan[T constraints.Ordered](a T) Predicate[T] { return func(b T) bool { return b < a } }

// IsGreaterThan creates a Predicate that tests if its argument is greater than a given value.
func IsGreaterThan[T constraints.Ordered](a T) Predicate[T] { return func(b T) bool { return b > a } }

// Stringer is used to create string representations, it accepts any type and returns a string.
type Stringer[T any] func(T) string

// ValueFor given a comparable key will return a value for it.
type ValueFor[K comparable, T any] Function[K, T]

// OrderedLessThan will create a LessThan from any type included in the constraints.Ordered constraint.
func OrderedLessThan[T constraints.Ordered]() LessThan[T] {
	return func(a, b T) bool {
		return a < b
	}
}

// Reverse reverses a LessThan to facilitate reverse sort ordering.
func Reverse[T any](lessThan LessThan[T]) LessThan[T] {
	return func(a, b T) bool { return lessThan(b, a) }
}

// StringerStringer creates a Stringer for any type that implements fmt.Stringer.
func StringerStringer[T fmt.Stringer]() Stringer[T] {
	return func(t T) string { return t.String() }
}

// TransformLessThan composites an existing LessThan[R] and transform Function[T,R] into a new LessThan[T]. The
// transform is used to convert the arguments before they are passed to the lessThan.
func TransformLessThan[T, R any](transform Function[T, R], lessThan LessThan[R]) LessThan[T] {
	return func(a, b T) bool {
		return lessThan(transform(a), transform(b))
	}
}
