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

import "github.com/nwillc/genfuncs"

var _ Queue[int] = (*Deque[int])(nil)

// Deque is a doubly ended implementation of Queue with default behavior of a Fifo but provides left and right access.
// Employs a List for storage.
type Deque[T any] struct {
	list *List[T]
}

// NewDeque creates a Deque containing any provided elements.
func NewDeque[T any](t ...T) *Deque[T] {
	d := &Deque[T]{list: NewList[T]()}
	d.AddAll(t...)
	return d
}

// Add an element to the right of the Deque.
func (d *Deque[T]) Add(t T) {
	d.list.Add(t)
}

// AddAll elements to the right of the Deque.
func (d *Deque[T]) AddAll(t ...T) {
	d.list.AddAll(t...)
}

// AddLeft an element to the left of the Deque.
func (d *Deque[T]) AddLeft(t T) {
	d.list.AddLeft(t)
}

// AddRight an element to the right of the Deque.
func (d *Deque[T]) AddRight(t T) {
	d.list.AddRight(t)
}

// Len reports the length of the Deque.
func (d *Deque[T]) Len() int {
	return d.list.Len()
}

// Peek returns the left most element in the Deque without removing it.
func (d *Deque[T]) Peek() T {
	return d.PeekLeft()
}

// PeekLeft returns the left most element in the Deque without removing it.
func (d *Deque[T]) PeekLeft() T {
	if d.Len() == 0 {
		panic(genfuncs.NoSuchElement)
	}
	return d.list.PeekLeft().Value
}

// PeekRight returns the right most element in the Deque without removing it.
func (d *Deque[T]) PeekRight() T {
	if d.Len() == 0 {
		panic(genfuncs.NoSuchElement)
	}
	return d.list.PeekRight().Value
}

// Remove and return the left most element in the Deque.
func (d *Deque[T]) Remove() T {
	return d.RemoveLeft()
}

// RemoveLeft and return the left most element in the Deque.
func (d *Deque[T]) RemoveLeft() T {
	if d.Len() == 0 {
		panic(genfuncs.NoSuchElement)
	}
	e := d.list.PeekLeft()
	return d.list.Remove(e)
}

// RemoveRight and return the right most element in the Deque.
func (d *Deque[T]) RemoveRight() T {
	if d.Len() == 0 {
		panic(genfuncs.NoSuchElement)
	}
	e := d.list.PeekRight()
	return d.list.Remove(e)
}

// Values in the Deque returned in a new GSlice.
func (d *Deque[T]) Values() GSlice[T] {
	return d.list.Values()
}
