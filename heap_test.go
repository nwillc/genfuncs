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
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/nwillc/genfuncs"
	"github.com/stretchr/testify/assert"
)

var (
	names = []string{"fred", "barney", "pebbles"}
)

func TestNewHeap(t *testing.T) {
	type args struct {
		data     []string
		lessThan genfuncs.LessThan[string]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{
				data:     nil,
				lessThan: strCompare,
			},
			want: nil,
		},
		{
			name: "min a b c",
			args: args{
				data:     []string{"a", "b", "c"},
				lessThan: strCompare,
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "max a b c",
			args: args{
				data:     []string{"a", "b", "c"},
				lessThan: genfuncs.Reverse(strCompare),
			},
			want: []string{"c", "b", "a"},
		},
		{
			name: "min c b a",
			args: args{
				data:     []string{"c", "b", "a"},
				lessThan: strCompare,
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "max c b a",
			args: args{
				data:     []string{"c", "b", "a"},
				lessThan: genfuncs.Reverse(strCompare),
			},
			want: []string{"c", "b", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fifo := genfuncs.NewHeap(tt.args.lessThan, tt.args.data...)
			assert.Equal(t, len(tt.want), fifo.Len())
			for _, e := range tt.want {
				value := fifo.Remove()
				assert.Equal(t, e, value)
			}
		})
	}
}

func TestHeapAddPeekRemove(t *testing.T) {
	numericOrder := genfuncs.OrderedLessThan[int]()
	type args struct {
		data     []int
		lessThan genfuncs.LessThan[int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{
				data:     nil,
				lessThan: numericOrder,
			},
			want: nil,
		},
		{
			name: "min 1 2 3",
			args: args{
				data:     []int{1, 2, 3},
				lessThan: numericOrder,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := genfuncs.NewHeap[int](tt.args.lessThan)
			for _, e := range tt.want {
				heap.Add(e)
			}
			assert.Equal(t, len(tt.want), heap.Len())
			for _, e := range tt.want {
				peek := heap.Peek()
				value := heap.Remove()
				require.Equal(t, peek, value)
				require.Equal(t, e, value)
			}
		})
	}
}
