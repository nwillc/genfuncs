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
	"math/rand"
	"strings"
	"time"
)

var (
	random = rand.New(rand.NewSource(time.Now().Unix()))
)

// Slice is a generic type corresponding to a standard Go slice.
type Slice[T any] []T

// All returns true if all elements of slice match the predicate.
func (s Slice[T]) All(predicate genfuncs.Function[T, bool]) bool {
	for _, e := range s {
		if !predicate(e) {
			return false
		}
	}
	return true
}

// Any returns true if any element of the slice matches the predicate.
func (s Slice[T]) Any(predicate genfuncs.Function[T, bool]) bool {
	for _, e := range s {
		if predicate(e) {
			return true
		}
	}
	return false
}

// Compare one Slice to another, applying a comparison to each pair of corresponding entries. Compare returns true
// if all the pair's comparison return true. While the comparison might test equality it could have any behavior.
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
func (s Slice[T]) Filter(predicate genfuncs.Function[T, bool]) Slice[T] {
	var results []T
	s.ForEach(func(t T) {
		if predicate(t) {
			results = append(results, t)
		}
	})
	return results
}

// Find returns the first element matching the given predicate and true, or false when no such element was found.
func (s Slice[T]) Find(predicate genfuncs.Function[T, bool]) (T, bool) {
	for _, t := range s {
		if predicate(t) {
			return t, true
		}
	}
	var t T
	return t, false
}

// FindLast returns the last element matching the given predicate and true, or false when no such element was found.
func (s Slice[T]) FindLast(predicate genfuncs.Function[T, bool]) (T, bool) {
	var last T
	var found = false
	s.ForEach(func(t T) {
		if predicate(t) {
			found = true
			last = t
		}
	})
	return last, found
}

// ForEach element of the Slice invoke given function with the element. Syntactic sugar for a range that intends to
// traverse all the elements, i.e. no exiting midway through.
func (s Slice[T]) ForEach(fn func(t T)) {
	for _, t := range s {
		fn(t)
	}
}

// ForEachI element of the Slice invoke given function with the element and its index in the Slice.
// Syntactic sugar for a range that intends to traverse all the elements, i.e. no exiting midway through.
func (s Slice[T]) ForEachI(fn func(i int, t T)) {
	for i, t := range s {
		fn(i, t)
	}
}

// JoinToString creates a string from all the elements using the stringer on each, separating them using separator, and
// using the given prefix and postfix.
func (s Slice[T]) JoinToString(stringer genfuncs.ToString[T], separator string, prefix string, postfix string) string {
	var sb strings.Builder
	sb.WriteString(prefix)
	last := len(s) - 1
	s.ForEachI(func(i int, t T) {
		sb.WriteString(stringer(t))
		if i != last {
			sb.WriteString(separator)
		}
	})
	sb.WriteString(postfix)
	return sb.String()
}

func (s Slice[T]) Random() T {
	return s[random.Intn(len(s))]
}

// SortBy copies a slice, sorts the copy applying the Comparator and returns it.
func (s Slice[T]) SortBy(lessThan genfuncs.BiFunction[T, T, bool]) Slice[T] {
	dst := make([]T, len(s))
	copy(dst, s)
	Slice[T](dst).Sort(lessThan)
	return dst
}

// Swap two values in the slice.
func (s Slice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// inBounds panics if given index out of Slice's bounds.
func (s Slice[T]) inBounds(i int) {
	if i < 0 || i > len(s)-1 {
		panic(genfuncs.NoSuchElement)
	}
}
