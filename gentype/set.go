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

type Set[T comparable] interface {
	// Add an element to the Queue.
	Add(t T)
	// Contains returns true if the Set contains a given element.
	Contains(t T) bool
	// Len returns length of the Queue.
	Len() int
	// Remove a given element from the Set.
	Remove(t T)
	// Values in the Set as a Slice.
	Values() Slice[T]
}

var _ Set[bool] = (*MapSet[bool])(nil)

// MapSet is a Set implementation based on a map.
type MapSet[T comparable] struct {
	set map[T]struct{}
}

func NewHashSet[T comparable]() *MapSet[T] {
	return &MapSet[T]{set: make(map[T]struct{})}
}

func (h *MapSet[T]) Add(t T) {
	h.set[t] = struct{}{}
}

func (h *MapSet[T]) Contains(t T) bool {
	_, ok := h.set[t]
	return ok
}

func (h *MapSet[T]) Len() int {
	return len(h.set)
}

func (h *MapSet[T]) Remove(t T) {
	delete(h.set, t)
}

func (h *MapSet[T]) Values() Slice[T] {
	return Keys(h.set)
}
