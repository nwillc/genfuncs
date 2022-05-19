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
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestHeapNew(t *testing.T) {
	heap := container.NewHeap[string](genfuncs.OrderedLess[string])
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
				lessThan: genfuncs.OrderedLess[int],
			},
			want: nil,
		},
		{
			name: "min 1 2 3",
			args: args{
				slice:    []int{1, 2, 3},
				lessThan: genfuncs.OrderedLess[int],
			},
			want: []int{1, 2, 3},
		},
		{
			name: "min 3 4 2 1",
			args: args{
				slice:    []int{3, 4, 2, 1},
				lessThan: genfuncs.OrderedLess[int],
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "min 3 4 1 2 1",
			args: args{
				slice:    []int{3, 4, 1, 2, 1},
				lessThan: genfuncs.OrderedLess[int],
			},
			want: []int{1, 1, 2, 3, 4},
		},
		{
			name: "max 1 2 3",
			args: args{
				slice:    []int{1, 2, 3},
				lessThan: genfuncs.OrderedGreater[int],
			},
			want: []int{3, 2, 1},
		},
		{
			name: "max 3 1 2",
			args: args{
				slice:    []int{3, 1, 2},
				lessThan: genfuncs.OrderedGreater[int],
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

func TestRandomHeaps(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	passes := 20

	type args struct {
		count int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Tiny",
			args: args{
				count: 2,
			},
		},
		{
			name: "Small",
			args: args{
				count: 5,
			},
		},
		{
			name: "Medium",
			args: args{
				count: 15,
			},
		},
		{
			name: "Large",
			args: args{
				count: 70,
			},
		},
		{
			name: "Larger",
			args: args{
				count: 3000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for pass := passes; pass >= 0; pass-- {
				heap := container.NewHeap[int](genfuncs.OrderedLess[int])
				for i := 0; i < tt.args.count; i++ {
					heap.Add(random.Int())
				}
				start := heap.Peek()
				for i := 0; i < tt.args.count; i++ {
					next := heap.Remove()
					assert.LessOrEqual(t, start, next)
					start = next
				}
				heap = container.NewHeap[int](genfuncs.OrderedGreater[int])
				for i := 0; i < tt.args.count; i++ {
					heap.Add(random.Int())
				}
				start = heap.Peek()
				for i := 0; i < tt.args.count; i++ {
					next := heap.Remove()
					assert.GreaterOrEqual(t, start, next)
					start = next
				}
			}
		})
	}

}

func TestHeapInserting(t *testing.T) {
	h := container.NewHeap[int](genfuncs.OrderedLess[int], 4, 2, 3, 1)
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
	s := container.GSlice[int]{1, 2, 3}
	h := container.NewHeap[int](genfuncs.OrderedLess[int], s...)
	assert.Equal(t, genfuncs.EqualTo, s.Compare(h.Values(), genfuncs.Ordered[int]))
}
