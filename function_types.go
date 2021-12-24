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

// ComparedOrder is the type returned by a Comparator.
type ComparedOrder int

var (
	LessThan    ComparedOrder = -1
	EqualTo     ComparedOrder = 0
	GreaterThan ComparedOrder = 1
)

// BiFunction accepts two arguments and produces a result.
type BiFunction[T, U, R any] func(T, U) R

// Comparator compares two arguments of the same type and returns LessThan, EqualTo or GreaterThan based relative order.
type Comparator[T any] BiFunction[T, T, ComparedOrder]

// Function accepts one argument and produces a result.
type Function[T, R any] func(T) R

// KeyFor is used for generating keys from types, it accepts any type and returns a comparable key for it.
type KeyFor[T any, K comparable] Function[T, K]

// KeyValueFor is used to generate a key and value from a type, it accepts any type, and returns a comparable key and
// any value.
type KeyValueFor[T any, K comparable, V any] func(T) (K, V)

// Predicate is used evaluate a value, it accepts any type and returns a bool.
type Predicate[T any] func(T) bool

// Stringer is used to create string representations, it accepts any type and returns a string.
type Stringer[T any] func(T) string

// ValueFor given a comparable key will return a value for it.
type ValueFor[K comparable, T any] Function[K, T]

// OrderedComparator will create a Comparator from any type included in the constraints.Ordered constraint.
func OrderedComparator[T constraints.Ordered]() Comparator[T] {
	return func(a, b T) ComparedOrder {
		switch {
		case a < b:
			return LessThan
		case a > b:
			return GreaterThan
		default:
			return EqualTo
		}
	}
}

// ReverseComparator reverses a Comparator to facilitate switching sort orderings.
func ReverseComparator[T any](comparator Comparator[T]) Comparator[T] {
	return func(a, b T) ComparedOrder { return comparator(b, a) }
}

// StringerStringer creates a Stringer for any type that implements fmt.Stringer.
func StringerStringer[T fmt.Stringer]() Stringer[T] {
	return func(t T) string { return t.String() }
}

// FunctionComparator composites an existing Comparator[R] and Function[T,R] into a new Comparator[T].
func FunctionComparator[T, R any](transform Function[T, R], comparator Comparator[R]) Comparator[T] {
	return func(a, b T) ComparedOrder {
		return comparator(transform(a), transform(b))
	}
}
