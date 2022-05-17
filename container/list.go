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
	"sort"
)

var (
	// List implements Container
	_ Container[int] = (*List[int])(nil)
	// listSorter implements sort.Interface
	_ sort.Interface = (*listSorter[int])(nil)
)

// ListElement is an element of List.
type ListElement[T any] struct {
	next, prev *ListElement[T]
	list       *List[T]
	Value      T
}

// Next returns the next list element or nil.
func (e *ListElement[T]) Next() *ListElement[T] {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil.
func (e *ListElement[T]) Prev() *ListElement[T] {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// List is a doubly linked list, inspired by list.List but reworked to be generic. List implements Container.
type List[T any] struct {
	root ListElement[T]
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

// AddLeft adds a value to the left of the List.
func (l *List[T]) AddLeft(value T) *ListElement[T] {
	return l.insertValue(value, &l.root)
}

// AddRight adds a value to the right of the List.
func (l *List[T]) AddRight(v T) *ListElement[T] {
	return l.insertValue(v, l.root.prev)
}

// ForEach invokes the action for each value in the list.
func (l *List[T]) ForEach(action func(value T)) {
	for e := l.PeekLeft(); e != nil; e = e.Next() {
		action(e.Value)
	}
}

// Get returns the ListElement at index. This traverses all the ListElement from the end nearest the index. If index
// is not within the bounds of the list nil is returned.
func (l *List[T]) Get(index int) *ListElement[T] {
	if index < 0 || index >= l.Len() {
		return nil
	}
	if index < l.Len()/2 {
		// Closer to left end
		i := 0
		for e := l.PeekLeft(); e != nil; e = e.Next() {
			if i == index {
				return e
			}
			i++
		}
	} else {
		// Closer to right end
		i := l.Len() - 1
		for e := l.PeekRight(); e != nil; e = e.Prev() {
			if i == index {
				return e
			}
			i--
		}
	}
	panic("index in bounds but element not found")
}

// Len returns the number of values in the List.
func (l *List[T]) Len() int {
	return l.len
}

// PeekLeft returns the leftmost value in the List or nil if empty.
func (l *List[T]) PeekLeft() *ListElement[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// PeekRight returns the rightmost value in the List or nil if empty.
func (l *List[T]) PeekRight() *ListElement[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// Remove removes a given value from the List.
func (l *List[T]) Remove(e *ListElement[T]) T {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

// SortBy sorts the List by the order of the lessThan function. This is not a pure function, the List is sorted, the
// List returned is to allow for fluid call chains.
func (l *List[T]) SortBy(lessThan genfuncs.BiFunction[T, T, bool]) *List[T] {
	s := listSorter[T]{
		List:  l,
		order: lessThan,
	}
	sort.Sort(s)
	return l
}

// Swap the Value of two elements in the List. If either index is not within the bounds of the List no action is taken.
func (l *List[T]) Swap(i, j int) {
	iElement := l.Get(i)
	jElement := l.Get(j)
	if iElement == nil || jElement == nil {
		return
	}
	iElement.Value, jElement.Value = jElement.Value, iElement.Value
}

// Values returns the values in the list as a GSlice.
func (l *List[T]) Values() GSlice[T] {
	s := make(GSlice[T], l.Len())
	i := 0
	l.ForEach(func(value T) {
		s[i] = value
		i++
	})
	return s
}

func (l *List[T]) insertValue(v T, at *ListElement[T]) *ListElement[T] {
	return l.insert(&ListElement[T]{Value: v}, at)
}

func (l *List[T]) insert(e, at *ListElement[T]) *ListElement[T] {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

func (l *List[T]) remove(e *ListElement[T]) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--
}

type listSorter[T any] struct {
	*List[T]
	order genfuncs.BiFunction[T, T, bool]
}

func (s listSorter[T]) Less(i, j int) bool {
	return s.order(s.Get(i).Value, s.Get(j).Value)
}
