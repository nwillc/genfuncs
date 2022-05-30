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

package gslices

import (
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
)

// Associate returns a map containing key/values created by applying a function to elements of the slice.
func Associate[T, V any, K comparable](slice container.GSlice[T], keyValueFor genfuncs.MapKeyValueFor[T, K, V]) (result container.GMap[K, V]) {
	length := len(slice)
	result = make(container.GMap[K, V], length)
	for i := 0; i < length; i++ {
		k, v := keyValueFor(slice[i])
		result[k] = v
	}
	return result
}

// AssociateWith returns a Map where keys are elements from the given sequence and values are produced by the
// valueSelector function applied to each element.
func AssociateWith[T comparable, V any](slice container.GSlice[T], valueFor genfuncs.MapValueFor[T, V]) (result container.GMap[T, V]) {
	length := len(slice)
	result = make(container.GMap[T, V], length)
	var t T
	for i := 0; i < length; i++ {
		t = slice[i]
		result[t] = valueFor(t)
	}
	return result
}

// Distinct returns a slice containing only distinct elements from the given slice.
func Distinct[T comparable](slice container.GSlice[T]) (distinct container.GSlice[T]) {
	distinct = ToSet(slice).Values()
	return distinct
}

// FlatMap returns a slice of all elements from results of transform being invoked on each element of
// original slice, and those resultant slices concatenated.
func FlatMap[T, R any](slice container.GSlice[T], transform genfuncs.Function[T, container.GSlice[R]]) (result container.GSlice[R]) {
	length := len(slice)
	for i := 0; i < length; i++ {
		result = append(result, transform(slice[i])...)
	}
	return result
}

// Fold accumulates a value starting with initial value and applying operation from left to right to current
// accumulated value and each element.
func Fold[T, R any](slice container.GSlice[T], initial R, operation genfuncs.BiFunction[R, T, R]) (result R) {
	length := len(slice)
	result = initial
	for i := 0; i < length; i++ {
		result = operation(result, slice[i])
	}
	return result
}

// GroupBy groups elements of the slice by the key returned by the given keySelector function applied to
// each element and returns a map where each group key is associated with a slice of corresponding elements.
func GroupBy[T any, K comparable](slice container.GSlice[T], keyFor genfuncs.MapKeyFor[T, K]) (result container.GMap[K, container.GSlice[T]]) {
	length := len(slice)
	result = make(container.GMap[K, container.GSlice[T]], length)
	var t T
	var k K
	var v container.GSlice[T]
	for i := 0; i < length; i++ {
		t = slice[i]
		k = keyFor(t)
		v = result[k]
		result[k] = append(v, t)
	}
	return result
}

// Map returns a new container.GSlice containing the results of applying the given transform function to each element in the original slice.
func Map[T, R any](slice container.GSlice[T], transform genfuncs.Function[T, R]) container.GSlice[R] {
	length := len(slice)
	var results = make(container.GSlice[R], length)
	for i := 0; i < length; i++ {
		results[i] = transform(slice[i])
	}
	return results
}

// ToSet creates a Set from the elements of the GSlice.
func ToSet[T comparable](slice container.GSlice[T]) (set container.Set[T]) {
	set = container.NewMapSet(slice...)
	return set
}
