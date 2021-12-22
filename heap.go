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

// Heap implements either a min or max ordered heap of any type.
type Heap[T any] struct {
	slice      []T
	comparator Comparator[T]
}

// NewMinHeap return a heap ordered min to max value.
func NewHeap[T any](comparator Comparator[T]) *Heap[T] {
	return &Heap[T]{comparator: comparator}
}

// Len returns current length of the heap.
func (h *Heap[T]) Len() int { return len(h.slice) }

// Push an item onto the heap.
func (h *Heap[T]) Push(v T) {
	h.slice = append(h.slice, v)
	h.up(h.Len() - 1)
}

// Pop an item off the heap.
func (h *Heap[T]) Pop() T {
	n := h.Len() - 1
	if n > 0 {
		swap(h.slice, 0, n)
		h.down()
	}
	v := h.slice[n]
	h.slice = h.slice[0:n]
	return v
}

func (h *Heap[T]) up(jj int) {
	for {
		i := parent(jj)
		if i == jj || h.comparator(h.slice[jj], h.slice[i]) >= EqualTo {
			break
		}
		swap(h.slice, i, jj)
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
		if j2 < n && h.comparator(h.slice[j2], h.slice[j1]) == LessThan {
			j = j2
		}
		if h.comparator(h.slice[j], h.slice[i1]) >= EqualTo {
			break
		}
		swap(h.slice, i1, j)
		i1 = j
	}
}

func swap[T any](slice []T, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func parent(i int) int { return (i - 1) / 2 }
func left(i int) int   { return (i * 2) + 1 }
func right(i int) int  { return left(i) + 1 }
