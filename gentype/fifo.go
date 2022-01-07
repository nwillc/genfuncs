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

// Fifo implements Queue.
var _ Queue[bool] = (*Fifo[bool])(nil)

// Fifo is a first in last out data structure.
type Fifo[T any] struct {
	slice Slice[T]
}

// NewFifo creates a Fifo containing any provided elements.
func NewFifo[T any](t ...T) *Fifo[T] {
	f := &Fifo[T]{}
	f.slice = make(Slice[T], len(t))
	copy(f.slice, t)
	return f
}

// Add an element to the Fifo.
func (f *Fifo[T]) Add(t T) {
	f.slice = append(f.slice, t)
}

// Len reports the length of the Fifo.
func (f *Fifo[T]) Len() int {
	return len(f.slice)
}

// Peek returns the next element in the Fifo without removing it.
func (f *Fifo[T]) Peek() T {
	if f.Len() < 1 {
		panic(NoSuchElement)
	}
	return f.slice[0]
}

// Remove and return the next element in the Fifo.
func (f *Fifo[T]) Remove() T {
	v := f.Peek()
	f.slice = f.slice[1:]
	return v
}
