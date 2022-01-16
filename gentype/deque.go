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

// Deque implements Queue.
var _ Queue[bool] = (*Deque[bool])(nil)

// Deque is a doubly ended queue with default behavior of a Fifo.
type Deque[T any] struct {
	slice Slice[T]
}

// NewDeque creates a Deque containing any provided elements.
func NewDeque[T any](t ...T) *Deque[T] {
	f := &Deque[T]{}
	f.slice = make(Slice[T], len(t))
	copy(f.slice, t)
	return f
}

// Add an element to the Deque.
func (d *Deque[T]) Add(t T) {
	d.AddRight(t)
}

// AddAll elements to the Deque.
func (d *Deque[T]) AddAll(t ...T) {
	for _, e := range t {
		d.AddRight(e)
	}
}

// AddLeft an element to the Deque.
func (d *Deque[T]) AddLeft(t T) {
	d.slice = append(d.slice[:1], d.slice[0:]...)
	d.slice[0] = t
}

// AddRight an element to the Deque.
func (d *Deque[T]) AddRight(t T) {
	d.slice = append(d.slice, t)
}

// Len reports the length of the Deque.
func (d *Deque[T]) Len() int {
	return len(d.slice)
}

// Peek returns the next element in the Deque without removing it.
func (d *Deque[T]) Peek() T {
	return d.PeekLeft()
}

// PeekLeft returns the next element in the Deque without removing it.
func (d *Deque[T]) PeekLeft() T {
	if d.Len() < 1 {
		panic(NoSuchElement)
	}
	return d.slice[0]
}

// PeekRight returns the next element in the Deque without removing it.
func (d *Deque[T]) PeekRight() T {
	if d.Len() < 1 {
		panic(NoSuchElement)
	}
	return d.slice[d.Len()-1]
}

// Remove and return the next element in the Deque.
func (d *Deque[T]) Remove() T {
	return d.RemoveLeft()
}

// RemoveLeft and return the next element in the Deque.
func (d *Deque[T]) RemoveLeft() T {
	v := d.PeekLeft()
	d.slice = d.slice[1:]
	return v
}

// RemoveRight and return the next element in the Deque.
func (d *Deque[T]) RemoveRight() T {
	v := d.PeekRight()
	d.slice = d.slice[:d.Len()-1]
	return v
}
