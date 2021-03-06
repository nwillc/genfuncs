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
	"golang.org/x/exp/slices"
	"math/rand"
	"time"
)

var (
	_      HasValues[int] = (*GSlice[int])(nil)
	_      Sequence[int]  = (*GSlice[int])(nil)
	_      Iterator[int]  = (*sliceIterator[int])(nil)
	random                = rand.New(rand.NewSource(time.Now().Unix()))
)

// GSlice is a generic type corresponding to a standard Go slice that implements HasValues.
type (
	GSlice[T any]        []T
	sliceIterator[T any] struct {
		slice []T
		index int
	}
)

// Filter returns a slice containing only elements matching the given predicate.
func (s GSlice[T]) Filter(predicate genfuncs.Function[T, bool]) GSlice[T] {
	length := len(s)
	var results []T
	var t T
	for i := 0; i < length; i++ {
		t = s[i]
		if predicate(t) {
			results = append(results, t)
		}
	}
	return results
}

// ForEach element of the GSlice invoke given function with the element. Syntactic sugar for a range that intends to
// traverse all the elements, i.e. no exiting midway through.
func (s GSlice[T]) ForEach(action func(i int, t T)) {
	length := s.Len()
	for i := 0; i < length; i++ {
		action(i, s[i])
	}
}

// Iterator returns an Iterator that will iterate over the GSlice.
func (s GSlice[T]) Iterator() Iterator[T] {
	return NewSliceIterator[T](s)
}

// Len is the number of elements in the GSlice.
func (s GSlice[T]) Len() int {
	return len(s)
}

// Random returns a random element of the GSlice.
func (s GSlice[T]) Random() (t T) {
	t = s[random.Intn(s.Len())]
	return t
}

// SortBy copies a slice, sorts the copy applying the Ordered and returns it. This is not a pure function, the GSlice
// is sorted in place, the returned slice is to allow for fluid calls in chains.
func (s GSlice[T]) SortBy(order genfuncs.BiFunction[T, T, bool]) (sorted GSlice[T]) {
	slices.SortStableFunc(s, order)
	sorted = s
	return sorted
}

// Swap two values in the slice.
func (s GSlice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Values is the GSlice itself.
func (s GSlice[T]) Values() (values GSlice[T]) {
	values = s
	return values
}

func NewSliceIterator[T any](slice []T) Iterator[T] {
	return &sliceIterator[T]{slice: slice}
}

func NewValuesIterator[T any](values ...T) Iterator[T] {
	return NewSliceIterator(values)
}

func (s *sliceIterator[T]) HasNext() bool {
	return s.index < len(s.slice)
}

func (s *sliceIterator[T]) Next() (value T) {
	if !s.HasNext() {
		panic(genfuncs.NoSuchElement)
	}
	value = s.slice[s.index]
	s.index++
	return value
}
