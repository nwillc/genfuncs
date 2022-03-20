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

	F32NumericOrder        = LessThanOrdered[float32]
	F32ReverseNumericOrder = Reverse(F32NumericOrder)
	INumericOrder          = LessThanOrdered[int]
	IReverseNumericOrder   = Reverse(INumericOrder)
	I64NumericOrder        = LessThanOrdered[int64]
	I64ReverseNumericOrder = Reverse(I64NumericOrder)
	SLexicalOrder          = LessThanOrdered[string]
	SReverseLexicalOrder   = Reverse(SLexicalOrder)

	// Predicates

	IsBlank    = IsEqualComparable("")
	IsNotBlank = Not(IsBlank)

	F32IsZero = IsEqualComparable(float32(0.0))
	F64IsZero = IsEqualComparable(0.0)
	IIsZero   = IsEqualComparable(0)
)

// EqualComparable tests equality of two given comparable values.
func EqualComparable[C comparable](a, b C) bool {
	return a == b
}

// IsEqualComparable creates a Predicate that tests equality with a given comparable value.
func IsEqualComparable[C comparable](a C) Function[C, bool] {
	return Curried(EqualComparable[C], a)
}

// GreaterThanOrdered tests if constraints.Ordered a is greater than b.
func GreaterThanOrdered[O constraints.Ordered](a, b O) bool {
	return a > b
}

// IsGreaterThanOrdered return a GreaterThanOrdered for a.
func IsGreaterThanOrdered[O constraints.Ordered](a O) Function[O, bool] {
	return Curried(GreaterThanOrdered[O], a)
}

// LessThanOrdered tests if constraints.Ordered a is less than b.
func LessThanOrdered[O constraints.Ordered](a, b O) bool {
	return a < b
}

// IsLessThanOrdered returns a LessThanOrdered for a.
func IsLessThanOrdered[O constraints.Ordered](a O) Function[O, bool] {
	return Curried(LessThanOrdered[O], a)
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

// Reverse reverses a LessThan to facilitate reverse sort ordering.
func Reverse[T any](lessThan BiFunction[T, T, bool]) BiFunction[T, T, bool] {
	return func(a, b T) bool { return lessThan(b, a) }
}

// StringerToString creates a ToString for any type that implements fmt.Stringer.
func StringerToString[T fmt.Stringer]() ToString[T] {
	return func(t T) string { return t.String() }
}

// TransformArgs uses the function to transform the arguments to be passed to the BiFunction.
func TransformArgs[T1, T2, R any](function Function[T1, T2], biFunction BiFunction[T2, T2, R]) BiFunction[T1, T1, R] {
	return func(a, b T1) R {
		return biFunction(function(a), function(b))
	}
}

// Curried takes a BiFunction and one argument, and Curries the function to return a single argument Function.
func Curried[A, B, R any](biFunction BiFunction[A, B, R], a A) Function[B, R] {
	return func(b B) R { return biFunction(a, b) }
}

// Not takes a Function returning a bool and returns a Function that inverts the result.
func Not[T any](function Function[T, bool]) Function[T, bool] {
	return func(a T) bool { return !function(a) }
}
