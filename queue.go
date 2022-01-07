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

package genfuncs

import "fmt"

var NoSuchElement = fmt.Errorf("no such element")

type Queue[T any] interface {
	Len() int
	Add(t T)
	Remove() T
	Peek() T
}

var _ Queue[bool] = (*Fifo[bool])(nil)

type Fifo[T any] struct {
	slice Slice[T]
}

func NewFifo[T any](t ...T) *Fifo[T] {
	f := &Fifo[T]{}
	f.slice = make([]T, len(t))
	copy(f.slice, t)
	return f
}

func (f *Fifo[T]) Len() int {
	return len(f.slice)
}

func (f *Fifo[T]) Add(t T) {
	f.slice = append(f.slice, t)
}

func (f *Fifo[T]) Peek() T {
	if f.Len() < 1 {
		panic(NoSuchElement)
	}
	return f.slice[0]
}

func (f *Fifo[T]) Remove() T {
	value := f.Peek()
	f.slice = f.slice[1:]
	return value
}
