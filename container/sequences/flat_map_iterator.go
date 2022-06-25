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

package sequences

import (
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
)

var _ container.Iterator[int] = (*flatMapIterator[string, int])(nil)

type flatMapIterator[T, R any] struct {
	outer     container.Iterator[T]
	inner     container.Iterator[R]
	transform func(t T) container.Sequence[R]
	hasData   bool
}

func newFlatMapIterator[T, R any](sequence container.Sequence[T], transform genfuncs.Function[T, container.Sequence[R]]) container.Iterator[R] {
	return &flatMapIterator[T, R]{outer: sequence.Iterator(), transform: transform}
}

func (f *flatMapIterator[T, R]) HasNext() bool {
	if f.hasData && f.inner.HasNext() {
		return true
	}
	f.hasData = false
	for f.outer.HasNext() {
		f.inner = f.transform(f.outer.Next()).Iterator()
		if f.inner.HasNext() {
			f.hasData = true
			return true
		}
	}
	return false
}

func (f *flatMapIterator[T, R]) Next() R {
	if !f.HasNext() {
		panic(genfuncs.NoSuchElement)
	}
	return f.inner.Next()
}
