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

package gentype

import (
	"github.com/nwillc/genfuncs"
	"strings"
)

// Slice is a generic type corresponding to a standard Go slice.
type Slice[T any] []T

// All returns true if all elements of slice match the predicate.
func (s Slice[T]) All(predicate genfuncs.Predicate[T]) bool {
	for _, e := range s {
		if !predicate(e) {
			return false
		}
	}
	return true
}

// Any returns true if any element of the slice matches the predicate.
func (s Slice[T]) Any(predicate genfuncs.Predicate[T]) bool {
	for _, e := range s {
		if predicate(e) {
			return true
		}
	}
	return false
}

// Compare one Slice to another, applying a comparison to each pair of corresponding entries. Compare returns true
// if all the pairs comparison return true. While the comparison might test equality it could have any behavior.
func (s Slice[T]) Compare(s2 Slice[T], comparison genfuncs.BiFunction[T, T, bool]) bool {
	if len(s) != len(s2) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if !comparison(s[i], s2[i]) {
			return false
		}
	}
	return true
}

// Filter returns a slice containing only elements matching the given predicate.
func (s Slice[T]) Filter(predicate genfuncs.Predicate[T]) Slice[T] {
	var results []T
	for _, t := range s {
		if predicate(t) {
			results = append(results, t)
		}
	}
	return results
}

// Find returns the first element matching the given predicate and true, or false when no such element was found.
func (s Slice[T]) Find(predicate genfuncs.Predicate[T]) (T, bool) {
	for _, t := range s {
		if predicate(t) {
			return t, true
		}
	}
	var t T
	return t, false
}

// FindLast returns the last element matching the given predicate and true, or false when no such element was found.
func (s Slice[T]) FindLast(predicate genfuncs.Predicate[T]) (T, bool) {
	var last T
	var found = false
	for _, t := range s {
		if predicate(t) {
			found = true
			last = t
		}
	}
	return last, found
}

// JoinToString creates a string from all the elements using the stringer on each, separating them using separator, and
// using the given prefix and postfix.
func (s Slice[T]) JoinToString(stringer genfuncs.Stringer[T], separator string, prefix string, postfix string) string {
	var sb strings.Builder
	sb.WriteString(prefix)
	last := len(s) - 1
	for i, e := range s {
		sb.WriteString(stringer(e))
		if i == last {
			continue
		}
		sb.WriteString(separator)
	}
	sb.WriteString(postfix)
	return sb.String()
}

// SortBy copies a slice, sorts the copy applying the Comparator and returns it.
func (s Slice[T]) SortBy(lessThan genfuncs.LessThan[T]) Slice[T] {
	dst := make([]T, len(s))
	copy(dst, s)
	Slice[T](dst).Sort(lessThan)
	return dst
}

// Swap two values in the slice.
func (s Slice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
