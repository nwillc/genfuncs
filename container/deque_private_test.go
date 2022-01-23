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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeque_prev(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Prev at 0",
			args: args{
				i: 0,
			},
			want: 15,
		},
		{
			name: "Prev at 15",
			args: args{
				i: 15,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDeque[int](1)
			assert.Equal(t, tt.want, d.prev(tt.args.i))
		})
	}
}

func TestExpandContract(t *testing.T) {
	s := make(Slice[int], minimumCapacity)
	for i := 0; i < minimumCapacity; i++ {
		s[i] = i
	}
	d := NewDeque[int]()
	// expand to hold 3*minimumCapacity
	d.AddAll(s...)
	d.AddAll(s...)
	d.AddAll(s...)
	assert.Equal(t, 3*minimumCapacity, d.Len())
	assert.Equal(t, 4*minimumCapacity, d.Cap())
	// remove 2*minimumCapacity
	for i := 0; i < 2*minimumCapacity; i++ {
		d.Remove()
	}
	assert.Equal(t, minimumCapacity, d.Len())
	assert.Equal(t, 2*minimumCapacity, d.Cap())
}
