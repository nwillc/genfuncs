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

package gentype_test

import (
	"testing"

	"github.com/nwillc/genfuncs/gentype"
	"github.com/stretchr/testify/assert"
)

func TestDeque_New(t *testing.T) {
	deque := gentype.NewDeque[int]()
	assert.NotNil(t, deque)
	assert.Equal(t, 0, deque.Len())
}

func TestDeque_Bounds(t *testing.T) {
	deque := gentype.NewDeque[bool]()
	assert.PanicsWithError(t, gentype.NoSuchElement.Error(), func() {
		_ = deque.Remove()
	})
	assert.PanicsWithError(t, gentype.NoSuchElement.Error(), func() {
		_ = deque.RemoveRight()
	})
	assert.PanicsWithError(t, gentype.NoSuchElement.Error(), func() {
		_ = deque.RemoveLeft()
	})
	assert.PanicsWithError(t, gentype.NoSuchElement.Error(), func() {
		_ = deque.Peek()
	})
	assert.PanicsWithError(t, gentype.NoSuchElement.Error(), func() {
		_ = deque.PeekRight()
	})
	assert.PanicsWithError(t, gentype.NoSuchElement.Error(), func() {
		_ = deque.PeekLeft()
	})
}

func TestDeque_Inserting(t *testing.T) {
	deque := gentype.NewDeque(1, 3, 2, 4)
	assert.Equal(t, 1, deque.Remove())
	assert.Equal(t, 3, deque.Remove())
	assert.Equal(t, 2, deque.Peek())
	deque.Add(6)
	assert.Equal(t, 2, deque.Remove())
	assert.Equal(t, 4, deque.Remove())
	assert.Equal(t, 6, deque.Remove())
}

func TestDeque_AddAll(t *testing.T) {
	type args struct {
		slice gentype.Slice[string]
	}
	tests := []struct {
		name string
		want gentype.Slice[string]
		args args
	}{
		{
			name: "A B C",
			want: []string{"A", "B", "C"},
			args: args{
				slice: []string{"A", "B", "C"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := gentype.NewDeque[string]()
			d.AddAll(tt.args.slice...)
			for _, e := range tt.want {
				assert.Equal(t, e, d.Remove())
			}
		})
	}
}

func TestDequeFifo_AddPeekRemove(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{
				slice: nil,
			},
			want: nil,
		},
		{
			name: "1 2 3",
			args: args{
				slice: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "3 4 2 1",
			args: args{
				slice: []int{3, 4, 2, 1},
			},
			want: []int{3, 4, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := gentype.NewDeque(tt.args.slice...)
			assert.Equal(t, len(tt.want), d.Len())
			for _, ii := range tt.want {
				v1 := d.Peek()
				v2 := d.Remove()
				assert.Equal(t, v1, v2)
				assert.Equal(t, ii, v2)
			}
		})
	}
}

func TestDequeLifo_AddPeekRemove(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{
				slice: nil,
			},
			want: nil,
		},
		{
			name: "1 2 3",
			args: args{
				slice: []int{1, 2, 3},
			},
			want: []int{3, 2, 1},
		},
		{
			name: "3 4 2 1",
			args: args{
				slice: []int{3, 4, 2, 1},
			},
			want: []int{1, 2, 4, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := gentype.NewDeque(tt.args.slice...)
			assert.Equal(t, len(tt.want), d.Len())
			for _, ii := range tt.want {
				v1 := d.PeekRight()
				v2 := d.RemoveRight()
				assert.Equal(t, v1, v2)
				assert.Equal(t, ii, v2)
			}
		})
	}
}
