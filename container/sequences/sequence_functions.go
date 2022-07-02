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
	"github.com/nwillc/genfuncs/container/maps"
	"github.com/nwillc/genfuncs/results"
	"strings"
)

// All returns true if all elements in the sequence match the predicate.
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
func Associate[T any, K comparable, V any](sequence container.Sequence[T], keyValueFor maps.KeyValueFor[T, K, V]) (result *genfuncs.Result[container.GMap[K, V]]) {
	iterator := sequence.Iterator()
	m := make(container.GMap[K, V])
	for iterator.HasNext() {
		keyValueFor(iterator.Next()).
			OnSuccess(func(kv *maps.Entry[K, V]) {
				m[kv.Key] = kv.Value
			})

	}
	return genfuncs.NewResult(m)
}

// AssociateWith returns a Map where keys are elements from the given sequence and values are produced by the
// valueSelector function applied to each element.
func AssociateWith[K comparable, V any](sequence container.Sequence[K], valueFor maps.ValueFor[K, V]) (result *genfuncs.Result[container.GMap[K, V]]) {
	iterator := sequence.Iterator()
	m := make(container.GMap[K, V])
	var t K
	for iterator.HasNext() {
		t = iterator.Next()
		value := valueFor(t)
		if !value.Ok() {
			return results.MapError[V, container.GMap[K, V]](value)
		}
		m[t] = value.OrEmpty()
	}
	return genfuncs.NewResult(m)
}

func Collect[T any](s container.Sequence[T], c container.Container[T]) {
	iterator := s.Iterator()
	for iterator.HasNext() {
		c.Add(iterator.Next())
	}
}

// Compare two sequences with a comparator returning less/equal/greater (-1/0/1) and return comparison of the two.
func Compare[T any](s1, s2 container.Sequence[T], comparator func(t1, t2 T) int) int {
	i1 := s1.Iterator()
	i2 := s2.Iterator()
	for i1.HasNext() && i2.HasNext() {
		cmp := comparator(i1.Next(), i2.Next())
		if cmp != genfuncs.EqualTo {
			return cmp
		}
	}
	if i2.HasNext() {
		return genfuncs.LessThan
	}
	if i1.HasNext() {
		return genfuncs.GreaterThan
	}
	return genfuncs.EqualTo
}

// Find returns the first element matching the given predicate, or Result error of NoSuchElement if not found.
func Find[T any](sequence container.Sequence[T], predicate genfuncs.Function[T, bool]) *genfuncs.Result[T] {
	iterator := sequence.Iterator()
	var result T
	for iterator.HasNext() {
		result = iterator.Next()
		if predicate(result) {
			return genfuncs.NewResult(result)
		}
	}
	return genfuncs.NewError[T](genfuncs.NoSuchElement)
}

// FindLast returns the last element matching the given predicate, or Result error of NoSuchElement if not found.
func FindLast[T any](sequence container.Sequence[T], predicate genfuncs.Function[T, bool]) *genfuncs.Result[T] {
	iterator := sequence.Iterator()
	result := genfuncs.NewError[T](genfuncs.NoSuchElement)
	var t T
	for iterator.HasNext() {
		t = iterator.Next()
		if predicate(t) {
			result = genfuncs.NewResult(t)
		}
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
func Fold[T, R any](sequence container.Sequence[T], initial R, operation genfuncs.BiFunction[R, T, R]) (result R) {
	iterator := sequence.Iterator()
	result = initial
	for iterator.HasNext() {
		result = operation(result, iterator.Next())
	}
	return result
}

// ForEach calls action for each element of a Sequence.
func ForEach[T any](sequence container.Sequence[T], action func(t T)) {
	iterator := sequence.Iterator()
	for iterator.HasNext() {
		action(iterator.Next())
	}
}

// IsSorted returns true if the GSlice is sorted by order.
func IsSorted[T any](sequence container.Sequence[T], order genfuncs.BiFunction[T, T, bool]) (ok bool) {
	iterator := sequence.Iterator()
	var current, last T
	ok = true
	first := true
	for iterator.HasNext() {
		current = iterator.Next()
		if first {
			first = false
		} else {
			if order(current, last) {
				ok = false
				break
			}
		}
		last = current
	}
	return ok
}

// JoinToString creates a string from all the elements of a Sequence using the stringer on each, separating them using separator, and
// using the given prefix and postfix.
func JoinToString[T any](
	sequence container.Sequence[T],
	stringer genfuncs.ToString[T],
	separator string,
	prefix string,
	postfix string,
) string {
	var sb strings.Builder
	iterator := sequence.Iterator()
	sb.WriteString(prefix)
	first := true
	for iterator.HasNext() {
		if first {
			first = false
		} else {
			sb.WriteString(separator)
		}
		sb.WriteString(stringer(iterator.Next()))
	}
	sb.WriteString(postfix)
	return sb.String()
}

// Map elements in a Sequence to a new Sequence having applied the transform to them.
func Map[T, R any](sequence container.Sequence[T], transform genfuncs.Function[T, R]) container.Sequence[R] {
	return container.NewIteratorSequence[R](transformIterator[T, R]{iterator: sequence.Iterator(), transform: transform})
}

// NewSequence creates a sequence from the provided values.
func NewSequence[T any](values ...T) (sequence container.Sequence[T]) {
	var slice container.GSlice[T] = values
	return slice
}
