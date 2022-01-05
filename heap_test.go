/*
 *  Copyright (c) 2021,  nwillc@gmail.com
 *
 *  Permission to use, copy, modify, and/or distribute this software for any
 *  purpose with or without fee is hereby granted, provided that the above
 *  copyright notice and this permission notice appear in all copies.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 *  WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF¬
 *  MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 *  ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 *  WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 *  ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 *  OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package genfuncs_test

import (
	"testing"

	"github.com/nwillc/genfuncs"
	"github.com/stretchr/testify/assert"
)

var (
	names = []string{"fred", "barney", "pebbles"}
)

func Test_New(t *testing.T) {
	var heap *genfuncs.Heap[string]

	heap = genfuncs.NewHeap(strCompare)
	assert.NotNil(t, heap)
	assert.Equal(t, 0, heap.Len())
}

func Test_MinHeap(t *testing.T) {
	heap := genfuncs.NewHeap(strCompare)
	heap.PushAll(names...)
	assert.Equal(t, 3, heap.Len())
	assert.Equal(t, "barney", heap.Pop())
	assert.Equal(t, "fred", heap.Pop())
	assert.Equal(t, "pebbles", heap.Pop())
	assert.Equal(t, 0, heap.Len())
}

func Test_MaxHeap(t *testing.T) {
	heap := genfuncs.NewHeap(genfuncs.Reverse(strCompare))
	heap.PushAll(names...)
	assert.Equal(t, 3, heap.Len())
	assert.Equal(t, "pebbles", heap.Pop())
	assert.Equal(t, "fred", heap.Pop())
	assert.Equal(t, "barney", heap.Pop())
	assert.Equal(t, 0, heap.Len())
}
