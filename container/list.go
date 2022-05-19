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

var (
	// List implements Container.
	_ Container[int] = (*List[int])(nil)
)

// ListElement is an element of List.
type ListElement[T any] struct {
	next, prev *ListElement[T]
	list       *List[T]
	Value      T
}

// Next returns the next list element or nil.
func (e *ListElement[T]) Next() (next *ListElement[T]) {
	if e.list == nil || e.next == &e.list.root {
		return next
	}
	next = e.next
	return next
}

// Prev returns the previous list element or nil.
func (e *ListElement[T]) Prev() (prev *ListElement[T]) {
	if e.list == nil || e.prev == &e.list.root {
		return prev
	}
	prev = e.prev
	return prev
}

// Swap the values of two ListElements.
func (e *ListElement[T]) Swap(e2 *ListElement[T]) {
	if e == nil || e2 == nil {
		return
	}
	e.Value, e2.Value = e2.Value, e.Value
}

// List is a doubly linked list, inspired by list.List but reworked to be generic. List implements Container.
type List[T any] struct {
	root ListElement[T]
	len  int
}

// NewList instantiates a new List containing any values provided.
func NewList[T any](values ...T) (l *List[T]) {
	l = new(List[T])
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
	i := 0
	c := len(values)
	for i < c {
		l.Add(values[i])
		i++
	}
}

// AddLeft adds a value to the left of the List.
func (l *List[T]) AddLeft(value T) (e *ListElement[T]) {
	e = l.insertValue(value, &l.root)
	return e
}

// AddRight adds a value to the right of the List.
func (l *List[T]) AddRight(v T) (e *ListElement[T]) {
	e = l.insertValue(v, l.root.prev)
	return e
}

// ForEach invokes the action for each value in the list.
func (l *List[T]) ForEach(action func(value T)) {
	for e := l.PeekLeft(); e != nil; e = e.Next() {
		action(e.Value)
	}
}

// IsSorted returns true if the List is sorted by order.
func (l *List[T]) IsSorted(order genfuncs.BiFunction[T, T, bool]) (ok bool) {
	e1 := l.PeekLeft()
	for i := 1; i < l.Len(); i++ {
		e2 := e1.Next()
		if order(e2.Value, e1.Value) {
			return ok
		}
		e1 = e2
	}
	ok = true
	return ok
}

// Len returns the number of values in the List.
func (l *List[T]) Len() (length int) {
	length = l.len
	return length
}

// PeekLeft returns the leftmost value in the List or nil if empty.
func (l *List[T]) PeekLeft() (e *ListElement[T]) {
	if l.len != 0 {
		e = l.root.next
	}
	return e
}

// PeekRight returns the rightmost value in the List or nil if empty.
func (l *List[T]) PeekRight() (e *ListElement[T]) {
	if l.len != 0 {
		e = l.root.prev
	}
	return e
}

// Remove removes a given value from the List.
func (l *List[T]) Remove(e *ListElement[T]) (t T) {
	if e.list == l {
		l.remove(e)
	}
	t = e.Value
	return t
}

// SortBy sorts the List by the order of the order function. This is not a pure function, the List is sorted, the
// List returned is to allow for fluid call chains. List does not provide efficient indexed access so a Bubble sort is employed.
func (l *List[T]) SortBy(order genfuncs.BiFunction[T, T, bool]) (result *List[T]) {
	result = l
	l.bubbleSort(order)
	return l
}

// Values returns the values in the list as a GSlice.
func (l *List[T]) Values() (values GSlice[T]) {
	values = make(GSlice[T], l.Len())
	i := 0
	l.ForEach(func(value T) {
		values[i] = value
		i++
	})
	return values
}

func (l *List[T]) insertValue(v T, at *ListElement[T]) (le *ListElement[T]) {
	le = l.insert(&ListElement[T]{Value: v}, at)
	return le
}

func (l *List[T]) insert(e, at *ListElement[T]) (le *ListElement[T]) {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	le = e
	return le
}

func (l *List[T]) remove(e *ListElement[T]) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--
}

// bubbleSort implements a Bubble Sort on List. Because List employs next/prev pointers arbitrary indexing in costly and
// this is the fastest sort.
func (l *List[T]) bubbleSort(order genfuncs.BiFunction[T, T, bool]) {
	end := l.Len()
	for end > 1 {
		e1 := l.PeekLeft()
		newEnd := 0
		for i := 1; i < end; i++ {
			e2 := e1.Next()
			if order(e2.Value, e1.Value) {
				e2.Swap(e1)
				newEnd = i
			}
			e1 = e2
		}
		end = newEnd
	}
}
