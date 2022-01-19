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

package container_test

import (
	"github.com/nwillc/genfuncs/container"
	"testing"

	"github.com/nwillc/genfuncs"
	"github.com/stretchr/testify/assert"
)

var (
	names = []string{"fred", "barney", "pebbles"}
)

func TestHeapNew(t *testing.T) {
	heap := container.NewHeap(genfuncs.SLexicalOrder)
	assert.NotNil(t, heap)
	assert.Equal(t, 0, heap.Len())
}

func TestHeapAddPeekRemove(t *testing.T) {
	type args struct {
		slice    []int
		lessThan genfuncs.BiFunction[int, int, bool]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{
				slice:    nil,
				lessThan: genfuncs.INumericOrder,
			},
			want: nil,
		},
		{
			name: "min 1 2 3",
			args: args{
				slice:    []int{1, 2, 3},
				lessThan: genfuncs.INumericOrder,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "min 3 4 2 1",
			args: args{
				slice:    []int{3, 4, 2, 1},
				lessThan: genfuncs.INumericOrder,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "max 1 2 3",
			args: args{
				slice:    []int{1, 2, 3},
				lessThan: genfuncs.IReverseNumericOrder,
			},
			want: []int{3, 2, 1},
		},
		{
			name: "max 3 2 1",
			args: args{
				slice:    []int{3, 2, 1},
				lessThan: genfuncs.IReverseNumericOrder,
			},
			want: []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := container.NewHeap(tt.args.lessThan, tt.args.slice...)
			assert.Equal(t, len(tt.want), heap.Len())
			for _, ii := range tt.want {
				v1 := heap.Peek()
				v2 := heap.Remove()
				assert.Equal(t, v1, v2)
				assert.Equal(t, ii, v2)
			}
		})
	}
}

func TestHeapInserting(t *testing.T) {
	h := container.NewHeap(genfuncs.INumericOrder, 4, 2, 3, 1)
	assert.Equal(t, 1, h.Remove())
	assert.Equal(t, 2, h.Remove())
	assert.Equal(t, 3, h.Peek())
	h.Add(1)
	assert.Equal(t, 1, h.Remove())
	assert.Equal(t, 3, h.Remove())
	assert.Equal(t, 4, h.Remove())
	assert.PanicsWithError(t, "no such element", func() {
		h.Peek()
	})
}

func TestHeap_Values(t *testing.T) {
	s := container.Slice[int]{1, 2, 3}
	h := container.NewHeap(genfuncs.INumericOrder, s...)
	assert.True(t, s.Compare(h.Values(), genfuncs.EqualComparable[int]))
}
