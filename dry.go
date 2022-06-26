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
	LessThan    = -1
	EqualTo     = 0
	GreaterThan = 1

	// Predicates
	IsBlank    = OrderedEqualTo("")
	IsNotBlank = Not(IsBlank)
	F32IsZero  = OrderedEqualTo(float32(0.0))
	F64IsZero  = OrderedEqualTo(0.0)
	IIsZero    = OrderedEqualTo(0)
)

// Empty return an empty value of type T.
func Empty[T any]() (empty T) {
	return empty
}

// OrderedEqual returns true jf a is ordered equal to b.
func OrderedEqual[O constraints.Ordered](a, b O) (orderedEqualTo bool) {
	orderedEqualTo = Ordered(a, b) == EqualTo
	return orderedEqualTo
}

// OrderedEqualTo return a function that returns true if its argument is ordered equal to a.
func OrderedEqualTo[O constraints.Ordered](a O) (fn Function[O, bool]) {
	fn = Curried[O, bool](OrderedEqual[O], a)
	return fn
}

// OrderedGreater returns true if a is ordered greater than b.
func OrderedGreater[O constraints.Ordered](a, b O) (orderedGreaterThan bool) {
	orderedGreaterThan = Ordered(a, b) == GreaterThan
	return orderedGreaterThan
}

// OrderedGreaterThan returns a function that returns true if its argument is ordered greater than a.
func OrderedGreaterThan[O constraints.Ordered](a O) (fn Function[O, bool]) {
	fn = Curried[O, bool](OrderedGreater[O], a)
	return fn
}

// OrderedLess returns true if a is ordered less than b.
func OrderedLess[O constraints.Ordered](a, b O) (orderedLess bool) {
	orderedLess = Ordered(a, b) == LessThan
	return orderedLess
}

// OrderedLessThan returns a function that returns true if its argument is ordered less than a.
func OrderedLessThan[O constraints.Ordered](a O) (fn Function[O, bool]) {
	fn = Curried[O, bool](OrderedLess[O], a)
	return fn
}

// Max returns max value one or more constraints.Ordered values,
func Max[T constraints.Ordered](v ...T) (max T) {
	if len(v) == 0 {
		panic(fmt.Errorf("%w: at leat one value required", IllegalArguments))
	}
	max = v[0]
	for _, val := range v {
		if val > max {
			max = val
		}
	}
	return max
}

// Min returns min value of one or more constraints.Ordered values,
func Min[T constraints.Ordered](v ...T) (min T) {
	if len(v) == 0 {
		panic(fmt.Errorf("%w: at leat one value required", IllegalArguments))
	}
	min = v[0]
	for _, val := range v {
		if val < min {
			min = val
		}
	}
	return min
}

// StringerToString creates a ToString for any type that implements fmt.Stringer.
func StringerToString[T fmt.Stringer]() (fn ToString[T]) {
	fn = func(t T) string { return t.String() }
	return fn
}

// TransformArgs uses the function to transform the arguments to be passed to the operation.
func TransformArgs[T1, T2, R any](transform Function[T1, T2], operation BiFunction[T2, T2, R]) (fn BiFunction[T1, T1, R]) {
	fn = func(a, b T1) R {
		return operation(transform(a), transform(b))
	}
	return fn
}

// Curried takes a BiFunction and one argument, and Curries the function to return a single argument Function.
func Curried[A, R any](operation BiFunction[A, A, R], a A) (fn Function[A, R]) {
	fn = func(b A) R { return operation(b, a) }
	return fn
}

// Not takes a predicate returning and inverts the result.
func Not[T any](predicate Function[T, bool]) (fn Function[T, bool]) {
	fn = func(a T) bool { return !predicate(a) }
	return fn
}

// Ordered performs old school -1/0/1 comparison of constraints.Ordered arguments.
func Ordered[T constraints.Ordered](a, b T) (order int) {
	switch {
	case a < b:
		order = LessThan
	case a > b:
		order = GreaterThan
	default:
		order = EqualTo
	}
	return order
}
