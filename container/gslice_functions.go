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
func Associate[T, V any, K comparable](slice GSlice[T], keyValueFor genfuncs.MapKeyValueFor[T, K, V]) GMap[K, V] {
	m := make(GMap[K, V])
	slice.ForEach(func(t T) {
		k, v := keyValueFor(t)
		m[k] = v
	})
	return m
}

// AssociateWith returns a Map where keys are elements from the given sequence and values are produced by the
// valueSelector function applied to each element.
func AssociateWith[T comparable, V any](slice GSlice[T], valueFor genfuncs.MapValueFor[T, V]) GMap[T, V] {
	m := make(GMap[T, V])
	slice.ForEach(func(t T) {
		v := valueFor(t)
		m[t] = v
	})
	return m
}

// Distinct returns a slice containing only distinct elements from the given slice.
func Distinct[T comparable](slice GSlice[T]) GSlice[T] {
	return NewMapSet(slice...).Values()
}

// FlatMap returns a slice of all elements from results of transform function being invoked on each element of
// original slice, and those resultant slices concatenated.
func FlatMap[T, R any](slice GSlice[T], function genfuncs.Function[T, GSlice[R]]) GSlice[R] {
	var results GSlice[R]
	slice.ForEach(func(t T) {
		results = append(results, function(t)...)
	})
	return results
}

// Fold accumulates a value starting with initial value and applying operation from left to right to current
// accumulated value and each element.
func Fold[T, R any](slice GSlice[T], initial R, biFunction genfuncs.BiFunction[R, T, R]) R {
	r := initial
	slice.ForEach(func(t T) {
		r = biFunction(r, t)
	})
	return r
}

// GroupBy groups elements of the slice by the key returned by the given keySelector function applied to
// each element and returns a map where each group key is associated with a slice of corresponding elements.
func GroupBy[T any, K comparable](slice GSlice[T], keyFor genfuncs.MapKeyFor[T, K]) GMap[K, GSlice[T]] {
	m := make(GMap[K, GSlice[T]])
	slice.ForEach(func(t T) {
		k := keyFor(t)
		group := m[k]
		m[k] = append(group, t)
	})
	return m
}

// Map returns a slice containing the results of applying the given transform function to each element in the original slice.
func Map[T, R any](slice GSlice[T], function genfuncs.Function[T, R]) GSlice[R] {
	var results = make(GSlice[R], len(slice))
	slice.ForEachI(func(i int, t T) {
		results[i] = function(t)
	})
	return results
}

func ToSet[T comparable](slice GSlice[T]) *MapSet[T] {
	return NewMapSet(slice...)
}
