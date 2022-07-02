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

// Heap implements an ordered heap of any type which can be min heap or max heap depending on the compare provided.
// Heap implements Queue.
type Heap[T any] struct {
	slice   GSlice[T]
	compare genfuncs.BiFunction[T, T, bool]
	ordered bool
}

// NewHeap return a heap ordered based on the compare and adds any values provided.
func NewHeap[T any](compare genfuncs.BiFunction[T, T, bool], values ...T) (heap *Heap[T]) {
	heap = &Heap[T]{
		compare: compare,
		slice:   make(GSlice[T], 0, len(values)),
	}
	heap.AddAll(values...)
	return heap
}

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

// Len returns current length of the heap.
func (h *Heap[T]) Len() (length int) { length = h.slice.Len(); return length }

// Peek returns the next element without removing it.
func (h *Heap[T]) Peek() (value T) {
	if h.Len() <= 0 {
		panic(genfuncs.NoSuchElement)
	}
	n := h.Len() - 1
	if n > 0 && !h.ordered {
		h.slice.Swap(0, n)
		h.down(0)
		h.ordered = true
	}
	value = h.slice[n]
	return value
}

// Remove an item off the heap.
func (h *Heap[T]) Remove() (value T) {
	value = h.Peek()
	h.slice = h.slice[0 : h.Len()-1]
	h.ordered = false
	return value
}

// Values returns a slice of the values in the Heap in no particular order.
func (h *Heap[T]) Values() (values GSlice[T]) {
	values = h.slice
	return values
}

func (h *Heap[T]) up(i int) {
	for {
		iParent := parent(i)
		if i < 1 || iParent == i || !h.compare(h.slice[i], h.slice[iParent]) {
			break
		}
		h.slice.Swap(iParent, i)
		i = iParent
	}
}

func (h *Heap[T]) down(i int) {
	length := h.Len() - 1
	for {
		l := left(i)
		if l < 0 || l >= length {
			break
		}
		j := l
		r := right(i)
		if r < length && h.compare(h.slice[r], h.slice[l]) {
			j = r
		}
		if !h.compare(h.slice[j], h.slice[i]) {
			break
		}
		h.slice.Swap(i, j)
		i = j
	}
}

func parent(i int) (p int) { p = (i - 1) / 2; return p }
func left(i int) (l int)   { l = (i * 2) + 1; return l }
func right(i int) (r int)  { r = left(i) + 1; return r }
