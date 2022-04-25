/*
 *  Copyright (c) 2022,  nwillc@gmail.com
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
	"fmt"
	"golang.org/x/exp/constraints"
)

var (
	// Orderings
	OrderedLess    = -1
	OrderedEqual   = 0
	OrderedGreater = 1

	// Predicates
	IsBlank    = IsEqualOrdered("")
	IsNotBlank = Not(IsBlank)
	F32IsZero  = IsEqualOrdered(float32(0.0))
	F64IsZero  = IsEqualOrdered(0.0)
	IIsZero    = IsEqualOrdered(0)
)

// EqualOrder tests if constraints.Ordered a equal to b.
func EqualOrder[O constraints.Ordered](a, b O) bool {
	return Order(a, b) == OrderedEqual
}

// IsEqualOrdered return a EqualOrder for a.
func IsEqualOrdered[O constraints.Ordered](a O) Function[O, bool] {
	return Curried[O, bool](EqualOrder[O], a)
}

// GreaterOrdered tests if constraints.Ordered a is greater than b.
func GreaterOrdered[O constraints.Ordered](a, b O) bool {
	return Order(a, b) == OrderedGreater
}

// IsGreaterOrdered returns a function that returns true if its argument is greater than a.
func IsGreaterOrdered[O constraints.Ordered](a O) Function[O, bool] {
	return Curried[O, bool](GreaterOrdered[O], a)
}

// LessOrdered tests if constraints.Ordered a is less than b.
func LessOrdered[O constraints.Ordered](a, b O) bool {
	return Order(a, b) == OrderedLess
}

// IsLessOrdered returns a function that returns true if its argument is less than a.
func IsLessOrdered[O constraints.Ordered](a O) Function[O, bool] {
	return Curried[O, bool](LessOrdered[O], a)
}

// Max returns max value one or more constraints.Ordered values,
func Max[T constraints.Ordered](v ...T) T {
	if len(v) == 0 {
		panic(fmt.Errorf("%w: at leat one value required", IllegalArguments))
	}
	max := v[0]
	for _, val := range v {
		if val > max {
			max = val
		}
	}
	return max
}

// Min returns min value of one or more constraints.Ordered values,
func Min[T constraints.Ordered](v ...T) T {
	if len(v) == 0 {
		panic(fmt.Errorf("%w: at leat one value required", IllegalArguments))
	}
	min := v[0]
	for _, val := range v {
		if val < min {
			min = val
		}
	}
	return min
}

// StringerToString creates a ToString for any type that implements fmt.Stringer.
func StringerToString[T fmt.Stringer]() ToString[T] {
	return func(t T) string { return t.String() }
}

// TransformArgs uses the function to transform the arguments to be passed to the operation.
func TransformArgs[T1, T2, R any](transform Function[T1, T2], operation BiFunction[T2, T2, R]) BiFunction[T1, T1, R] {
	return func(a, b T1) R {
		return operation(transform(a), transform(b))
	}
}

// Curried takes a BiFunction and one argument, and Curries the function to return a single argument Function.
func Curried[A, R any](operation BiFunction[A, A, R], a A) Function[A, R] {
	return func(b A) R { return operation(b, a) }
}

// Not takes a predicate returning and inverts the result.
func Not[T any](predicate Function[T, bool]) Function[T, bool] {
	return func(a T) bool { return !predicate(a) }
}

// Order old school -1/0/1 order of constraints.Ordered.
func Order[T constraints.Ordered](a, b T) int {
	switch {
	case a < b:
		return OrderedLess
	case a > b:
		return OrderedGreater
	default:
		return OrderedEqual
	}
}
