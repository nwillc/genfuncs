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

// Heap implements Queue.
var _ Queue[int] = (*Heap[int])(nil)

// Heap implements either a min or max ordered heap of any type. Heap implements Queue.
type Heap[T any] struct {
	slice    GSlice[T]
	lessThan genfuncs.BiFunction[T, T, bool]
	ordered  bool
}

// NewHeap return a heap ordered based on the LessThan and pushes any values provided.
func NewHeap[T any](lessThan genfuncs.BiFunction[T, T, bool], values ...T) *Heap[T] {
	h := &Heap[T]{lessThan: lessThan}
	h.AddAll(values...)
	return h
}

// Len returns current length of the heap.
func (h *Heap[T]) Len() int { return len(h.slice) }

// Add a value onto the heap.
func (h *Heap[T]) Add(v T) {
	h.slice = append(h.slice, v)
	h.up(h.Len() - 1)
	h.ordered = false
}

// AddAll the values onto the Heap.
func (h *Heap[T]) AddAll(values ...T) {
	end := h.Len()
	h.slice = append(h.slice, values...)
	for ; end < h.Len(); end++ {
		h.up(end)
	}
}

// Peek returns the next element without removing it.
func (h *Heap[T]) Peek() T {
	if h.Len() <= 0 {
		panic(genfuncs.NoSuchElement)
	}
	n := h.Len() - 1
	if n > 0 && !h.ordered {
		h.slice.Swap(0, n)
		h.down()
		h.ordered = true
	}
	v := h.slice[n]
	return v
}

// Remove an item off the heap.
func (h *Heap[T]) Remove() T {
	v := h.Peek()
	h.slice = h.slice[0 : h.Len()-1]
	h.ordered = false
	return v
}

// Values returns a slice of the values in the Heap in no particular order.
func (h *Heap[T]) Values() GSlice[T] {
	return h.slice
}

func (h *Heap[T]) up(jj int) {
	for {
		i := parent(jj)
		if i == jj || !h.lessThan(h.slice[jj], h.slice[i]) {
			break
		}
		h.slice.Swap(i, jj)
		jj = i
	}
}

func (h *Heap[T]) down() {
	n := h.Len() - 1
	i1 := 0
	for {
		j1 := left(i1)
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		j2 := right(i1)
		if j2 < n && h.lessThan(h.slice[j2], h.slice[j1]) {
			j = j2
		}
		if !h.lessThan(h.slice[j], h.slice[i1]) {
			break
		}
		h.slice.Swap(i1, j)
		i1 = j
	}
}

func parent(i int) int { return (i - 1) / 2 }
func left(i int) int   { return (i * 2) + 1 }
func right(i int) int  { return left(i) + 1 }
