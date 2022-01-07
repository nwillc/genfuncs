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
	"github.com/nwillc/genfuncs/gentype"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFifoNew(t *testing.T) {
	fifo := gentype.NewFifo[int]()
	assert.NotNil(t, fifo)
	assert.Equal(t, 0, fifo.Len())
}

func TestFifoInserting(t *testing.T) {
	fifo := gentype.NewFifo(4, 2, 3, 1)
	assert.Equal(t, 4, fifo.Remove())
	assert.Equal(t, 2, fifo.Remove())
	assert.Equal(t, 3, fifo.Peek())
	fifo.Add(6)
	assert.Equal(t, 3, fifo.Remove())
	assert.Equal(t, 1, fifo.Remove())
	assert.Equal(t, 6, fifo.Remove())
}

func TestFifoAddPeekRemove(t *testing.T) {
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
			fifo := gentype.NewFifo(tt.args.slice...)
			assert.Equal(t, len(tt.want), fifo.Len())
			for _, ii := range tt.want {
				v1 := fifo.Peek()
				v2 := fifo.Remove()
				assert.Equal(t, v1, v2)
				assert.Equal(t, ii, v2)
			}
		})
	}
}
