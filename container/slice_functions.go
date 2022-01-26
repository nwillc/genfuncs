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

package container

import (
	"github.com/nwillc/genfuncs"
)

// Associate returns a map containing key/values created by applying a function to elements of the slice.
func Associate[T, V any, K comparable](slice Slice[T], keyValueFor genfuncs.MapKeyValueFor[T, K, V]) map[K]V {
	m := make(map[K]V)
	for _, e := range slice {
		k, v := keyValueFor(e)
		m[k] = v
	}
	return m
}

// AssociateWith returns a Map where keys are elements from the given sequence and values are produced by the
// valueSelector function applied to each element.
func AssociateWith[K comparable, V any](slice Slice[K], valueFor genfuncs.MapValueFor[K, V]) map[K]V {
	m := make(map[K]V)
	for _, k := range slice {
		v := valueFor(k)
		m[k] = v
	}
	return m
}

// Distinct returns a slice containing only distinct elements from the given slice.
func Distinct[T comparable](slice Slice[T]) Slice[T] {
	return NewMapSet(slice...).Values()
}

// FlatMap returns a slice of all elements from results of transform function being invoked on each element of
// original slice, and those resultant slices concatenated.
func FlatMap[T, R any](slice Slice[T], function genfuncs.Function[T, Slice[R]]) Slice[R] {
	var results []R
	for _, e := range slice {
		results = append(results, function(e)...)
	}
	return results
}

// Fold accumulates a value starting with initial value and applying operation from left to right to current
// accumulated value and each element.
func Fold[T, R any](slice Slice[T], initial R, biFunction genfuncs.BiFunction[R, T, R]) R {
	r := initial
	for _, t := range slice {
		r = biFunction(r, t)
	}
	return r
}

// GroupBy groups elements of the slice by the key returned by the given keySelector function applied to
// each element and returns a map where each group key is associated with a slice of corresponding elements.
func GroupBy[T any, K comparable](slice Slice[T], keyFor genfuncs.MapKeyFor[T, K]) map[K]Slice[T] {
	m := make(map[K]Slice[T])
	for _, e := range slice {
		k := keyFor(e)
		m[k] = append(m[k], e)
	}
	return m
}

// Map returns a slice containing the results of applying the given transform function to each element in the original slice.
func Map[T, R any](slice Slice[T], function genfuncs.Function[T, R]) Slice[R] {
	var results = make([]R, len(slice))
	for i, e := range slice {
		results[i] = function(e)
	}
	return results
}

func ToSet[T comparable](slice Slice[T]) *MapSet[T] {
	return NewMapSet(slice...)
}
