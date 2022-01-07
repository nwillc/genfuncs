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

package genfuncs_test

import (
	"github.com/nwillc/genfuncs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFifo(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{
				data: nil,
			},
			want: nil,
		},
		{
			name: "a b c",
			args: args{
				data: []string{"a", "b", "c"},
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fifo := genfuncs.NewFifo(tt.args.data...)
			assert.Equal(t, len(tt.want), fifo.Len())
			for _, e := range tt.want {
				value := fifo.Remove()
				assert.Equal(t, e, value)
			}
		})
	}
}

func TestFifoAddPeekRemove(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{
				data: nil,
			},
			want: nil,
		},
		{
			name: "1 2 3",
			args: args{
				data: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fifo := genfuncs.NewFifo[int]()
			for _, e := range tt.want {
				fifo.Add(e)
			}
			assert.Equal(t, len(tt.want), fifo.Len())
			for _, e := range tt.want {
				peek := fifo.Peek()
				value := fifo.Remove()
				assert.Equal(t, peek, value)
				assert.Equal(t, e, value)
			}
		})
	}
}
