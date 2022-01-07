/*
 *  Copyright (c) 2021,  nwillc@gmail.com
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

package genfuncs

// Heap implements Queue interface
var _ Queue[bool] = (*Heap[bool])(nil)

// Heap implements either a min or max ordered heap of any type.
type Heap[T any] struct {
	slice    Slice[T]
	lessThan LessThan[T]
}

// NewHeap return a heap ordered based on the LessThan.
func NewHeap[T any](lessThan LessThan[T], t ...T) *Heap[T] {
	h := &Heap[T]{lessThan: lessThan}
	h.PushAll(t...)
	return h
}

// Len returns current length of the heap.
func (h *Heap[T]) Len() int { return len(h.slice) }

// Add a value onto the heap.
func (h *Heap[T]) Add(v T) {
	h.slice = append(h.slice, v)
	h.up(h.Len() - 1)
}

// Peek returns the next value without removing it.
func (h *Heap[T]) Peek() T {
	if h.Len() < 1 {
		panic(NoSuchElement)
	}
	n := h.Len() - 1
	return h.slice[n]
}

// PushAll the values onto the Heap.
func (h *Heap[T]) PushAll(values ...T) {
	end := h.Len()
	h.slice = append(h.slice, values...)
	for ; end < h.Len(); end++ {
		h.up(end)
	}
}

// Remove an item off the heap.
func (h *Heap[T]) Remove() T {
	v := h.Peek()
	n := h.Len() - 1
	if n > 0 {
		h.slice.Swap(0, n)
		h.down()
	}
	h.slice = h.slice[0 : h.Len()-1]
	return v
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
