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

// List implements Container
var _ Container[int] = (*List[int])(nil)

// Element is an element of List.
type Element[T any] struct {
	next, prev *Element[T]
	list       *List[T]
	Value      T
}

// Next returns the next list element or nil.
func (e *Element[T]) Next() *Element[T] {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil.
func (e *Element[T]) Prev() *Element[T] {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// List represents a doubly linked list, based on list.List but made type aware with generics.
type List[T any] struct {
	root Element[T]
	len  int
}

// NewList instantiates a new List containing any values provided.
func NewList[T any](values ...T) *List[T] {
	l := new(List[T])
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	l.AddAll(values...)
	return l
}

// Add a value to the right of the List.
func (l *List[T]) Add(value T) {
	l.AddRight(value)
}

// AddAll values to the right of the List.
func (l *List[T]) AddAll(values ...T) {
	for _, tt := range values {
		l.Add(tt)
	}
}

// Len returns the number of values in the List.
func (l *List[T]) Len() int {
	return l.len
}

// PeekLeft returns the leftmost value in the List or nil if empty.
func (l *List[T]) PeekLeft() *Element[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// PeekRight returns the rightmost value in the List or nil if empty.
func (l *List[T]) PeekRight() *Element[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// AddLeft adds a value to the left of the List.
func (l *List[T]) AddLeft(value T) *Element[T] {
	return l.insertValue(value, &l.root)
}

// AddRight adds a value to the right of the List.
func (l *List[T]) AddRight(v T) *Element[T] {
	return l.insertValue(v, l.root.prev)
}

// Remove removes a given value from the List.
func (l *List[T]) Remove(e *Element[T]) T {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

// Values returns the values in the list as a GSlice.
func (l *List[T]) Values() GSlice[T] {
	s := make(GSlice[T], l.Len())
	if l.Len() > 0 {
		i := 0
		for e := l.PeekLeft(); e != nil; e = e.Next() {
			s[i] = e.Value
			i++
		}
	}
	return s
}

func (l *List[T]) insertValue(v T, at *Element[T]) *Element[T] {
	return l.insert(&Element[T]{Value: v}, at)
}

func (l *List[T]) insert(e, at *Element[T]) *Element[T] {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

func (l *List[T]) remove(e *Element[T]) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--
}
