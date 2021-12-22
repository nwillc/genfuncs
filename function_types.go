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

// ComparedOrder is the type returned by a Comparator.
type ComparedOrder int

var (
	LessThan    ComparedOrder = -1
	EqualTo     ComparedOrder = 0
	GreaterThan ComparedOrder = 1
)

// Comparator compares a to b and returns LessThan, EqualTo or GreaterThan based on a relative to b.
type Comparator[T any] func(a, b T) ComparedOrder

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

func ReverseComparator[T any](comparator Comparator[T]) Comparator[T] {
	return func(a, b T) ComparedOrder { return comparator(b, a) }
}
