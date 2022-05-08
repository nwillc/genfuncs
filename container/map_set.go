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

var (
	mapNilEntry          = struct{}{}
	_           Set[int] = (*MapSet[int])(nil)
)

// MapSet is a Set implementation based on a map. MapSet implements Set.
type MapSet[T comparable] struct {
	set GMap[T, struct{}]
}

// NewMapSet returns a new Set containing given values.
func NewMapSet[T comparable](t ...T) Set[T] {
	s := &MapSet[T]{set: make(map[T]struct{})}
	s.AddAll(t...)
	return s
}

// Add element to MapSet.
func (h *MapSet[T]) Add(t T) {
	h.set[t] = mapNilEntry
}

// AddAll elements to MapSet.
func (h *MapSet[T]) AddAll(t ...T) {
	for _, e := range t {
		h.set[e] = mapNilEntry
	}
}

// Contains returns true if MapSet contains element.
func (h *MapSet[T]) Contains(t T) bool {
	_, ok := h.set[t]
	return ok
}

// Len returns the length of the MapSet.
func (h *MapSet[T]) Len() int {
	return h.set.Len()
}

// Remove an element from the MapSet.
func (h *MapSet[T]) Remove(t T) {
	delete(h.set, t)
}

// Values returns the elements in the MapSet as a GSlice.
func (h *MapSet[T]) Values() GSlice[T] {
	return h.set.Keys()
}
