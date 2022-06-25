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

package sequences

import (
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
)

func All[T any](sequence container.Sequence[T], predicate genfuncs.Function[T, bool]) (result bool) {
	result = true
	iterator := sequence.Iterator()
	for iterator.HasNext() {
		if !predicate(iterator.Next()) {
			result = false
			break
		}
	}
	return result
}

// Any returns true if any element of the sequence matches the predicate.
func Any[T any](sequence container.Sequence[T], predicate genfuncs.Function[T, bool]) (result bool) {
	result = false
	iterator := sequence.Iterator()
	for iterator.HasNext() {
		if predicate(iterator.Next()) {
			result = true
			break
		}
	}
	return result
}

// Associate returns a map containing key/values created by applying a function to each value of the container.Iterator
// returned by the container.Sequence.
func Associate[T, V any, K comparable](iterable container.Sequence[T], keyValueFor genfuncs.MapKeyValueFor[T, K, V]) (result container.GMap[K, V]) {
	iterator := iterable.Iterator()
	result = make(container.GMap[K, V])
	for iterator.HasNext() {
		k, v := keyValueFor(iterator.Next())
		result[k] = v
	}
	return result
}

// AssociateWith returns a Map where keys are elements from the given sequence and values are produced by the
// valueSelector function applied to each element.
func AssociateWith[K comparable, V any](iterable container.Sequence[K], valueFor genfuncs.MapValueFor[K, V]) (result container.GMap[K, V]) {
	iterator := iterable.Iterator()
	result = make(container.GMap[K, V])
	var t K
	for iterator.HasNext() {
		t = iterator.Next()
		result[t] = valueFor(t)
	}
	return result
}

// FlatMap returns a sequence of all elements from results of transform being invoked on each element of
// original sequence, and those resultant slices concatenated.
func FlatMap[T, R any](sequence container.Sequence[T], transform genfuncs.Function[T, container.Sequence[R]]) (result container.Sequence[R]) {
	return container.NewIteratorSequence(newFlatMapIterator(sequence, transform))
}

// Fold accumulates a value starting with an initial value and applying operation to each value of the container.Iterator
// returned by the container.Sequence.
func Fold[T, R any](iterable container.Sequence[T], initial R, operation genfuncs.BiFunction[R, T, R]) (result R) {
	iterator := iterable.Iterator()
	result = initial
	for iterator.HasNext() {
		result = operation(result, iterator.Next())
	}
	return result
}

func Map[T, R any](iterable container.Sequence[T], transform genfuncs.Function[T, R]) container.Sequence[R] {
	return container.NewIteratorSequence[R](transformIterator[T, R]{iterator: iterable.Iterator(), transform: transform})
}

func SequenceOf[T any](values ...T) (sequence container.Sequence[T]) {
	var slice container.GSlice[T] = values
	return slice
}
