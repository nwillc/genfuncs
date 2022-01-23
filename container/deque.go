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

const minimumCapacity = 16

// Deque implements Queue.
var _ Queue[bool] = (*Deque[bool])(nil)

// Deque is a doubly ended Queue with default behavior of a Fifo but provides left and right access.
type Deque[T any] struct {
	slice Slice[T]
	head  int
	tail  int
	count int
}

// NewDeque creates a Deque containing any provided elements.
func NewDeque[T any](t ...T) *Deque[T] {
	d := &Deque[T]{}
	d.AddAll(t...)
	return d
}

// Add an element to the right of the Deque.
func (d *Deque[T]) Add(t T) {
	d.AddRight(t)
}

// AddAll elements to the right of the Deque.
func (d *Deque[T]) AddAll(t ...T) {
	for _, e := range t {
		d.AddRight(e)
	}
}

// AddLeft an element to the left of the Deque.
func (d *Deque[T]) AddLeft(t T) {
	d.expand()
	d.head = d.prev(d.head)
	d.slice[d.head] = t
	d.count++
}

// AddRight an element to the right of the Deque.
func (d *Deque[T]) AddRight(t T) {
	d.expand()
	d.slice[d.tail] = t
	d.tail = d.next(d.tail)
	d.count++
}

// Len reports the length of the Deque.
func (d *Deque[T]) Len() int {
	return d.count
}

// Peek returns the left most element in the Deque without removing it.
func (d *Deque[T]) Peek() T {
	return d.PeekLeft()
}

// PeekLeft returns the left most element in the Deque without removing it.
func (d *Deque[T]) PeekLeft() T {
	d.slice.inBounds(d.head)
	return d.slice[d.head]
}

// PeekRight returns the right most element in the Deque without removing it.
func (d *Deque[T]) PeekRight() T {
	p := d.prev(d.tail)
	d.slice.inBounds(p)
	return d.slice[p]
}

// Remove and return the left most element in the Deque.
func (d *Deque[T]) Remove() T {
	return d.RemoveLeft()
}

// RemoveLeft and return the left most element in the Deque.
func (d *Deque[T]) RemoveLeft() T {
	v := d.PeekLeft()
	d.head = d.next(d.head)
	d.count--
	d.contract()
	return v
}

// RemoveRight and return the right most element in the Deque.
func (d *Deque[T]) RemoveRight() T {
	v := d.PeekRight()
	d.tail = d.prev(d.tail)
	d.count--
	d.contract()
	return v
}

// Values in the Deque returned in a new Slice.
func (d *Deque[T]) Values() Slice[T] {
	newSlice := make(Slice[T], d.Len())
	d.copy(newSlice)
	return newSlice
}

// Cap returns the capacity of the Deque.
func (d *Deque[T]) Cap() int {
	return len(d.slice)
}

// expand the Deque capacity if needed.
func (d *Deque[T]) expand() {
	if d.Len() != d.Cap() {
		return
	}
	if d.Cap() == 0 {
		d.slice = make(Slice[T], minimumCapacity)
		return
	}
	d.resize()
}

// contract Deque capacity if only 1/4 full.
func (d *Deque[T]) contract() {
	if d.Cap() > minimumCapacity && (d.Len()<<2) == d.Cap() {
		d.resize()
	}
}

// resize the Deque to fit exactly twice its current contents.  This is
// used to grow the queue when it is full, and also to shrink it when it is
// only a quarter full.
func (d *Deque[T]) resize() {
	newSlice := make([]T, d.count<<1)
	d.copy(newSlice)
	d.head = 0
	d.tail = d.count
	d.slice = newSlice
}

// copy the values, in order, from the Deque to a Slice.
func (d *Deque[T]) copy(slice Slice[T]) {
	if d.tail > d.head {
		copy(slice, d.slice[d.head:d.tail])
	} else {
		n := copy(slice, d.slice[d.head:])
		copy(slice[n:], d.slice[:d.tail])
	}
}

// prev returns the previous buffer position wrapping around buffer.
func (d *Deque[T]) prev(i int) int {
	return (i - 1) & (d.Cap() - 1) // bitwise modulus
}

// next returns the next buffer position wrapping around buffer.
func (d *Deque[T]) next(i int) int {
	return (i + 1) & (d.Cap() - 1) // bitwise modulus
}
