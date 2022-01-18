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

func TestEqualComparable(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Equal",
			args: args{
				a: "a",
				b: "a",
			},
			want: true,
		},
		{
			name: "Not Equal",
			args: args{
				a: "a",
				b: "b",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := genfuncs.EqualComparable(tt.args.a, tt.args.b)
			assert.Equal(t, v, tt.want, v)
		})
	}
}

func TestIsEqualComparable(t *testing.T) {
	type args struct {
		a genfuncs.Function[string, bool]
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Equal",
			args: args{
				a: genfuncs.IsEqualComparable("a"),
				b: "a",
			},
			want: true,
		},
		{
			name: "Not Equal",
			args: args{
				a: genfuncs.IsEqualComparable("a"),
				b: "b",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.args.a(tt.args.b)
			assert.Equal(t, v, tt.want, v)
		})
	}
}

func TestGreaterThanOrdered(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Greater",
			args: args{
				a: 2,
				b: 1,
			},
			want: true,
		},
		{
			name: "Equal",
			args: args{
				a: 1,
				b: 1,
			},
			want: false,
		},
		{
			name: "Less",
			args: args{
				a: 0,
				b: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := genfuncs.GreaterThanOrdered(tt.args.a, tt.args.b)
			assert.Equal(t, v, tt.want, v)
		})
	}
}

func TestIsGreaterThanOrdered(t *testing.T) {
	type args struct {
		a genfuncs.Function[int, bool]
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Greater",
			args: args{
				a: genfuncs.IsGreaterThanOrdered(2),
				b: 1,
			},
			want: true,
		},
		{
			name: "Equal",
			args: args{
				a: genfuncs.IsGreaterThanOrdered(1),
				b: 1,
			},
			want: false,
		},
		{
			name: "Less",
			args: args{
				a: genfuncs.IsGreaterThanOrdered(0),
				b: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.args.a(tt.args.b)
			assert.Equal(t, v, tt.want, v)
		})
	}
}

func TestLessThanOrdered(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Greater",
			args: args{
				a: 2,
				b: 1,
			},
			want: false,
		},
		{
			name: "Equal",
			args: args{
				a: 1,
				b: 1,
			},
			want: false,
		},
		{
			name: "Less",
			args: args{
				a: 0,
				b: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := genfuncs.LessThanOrdered(tt.args.a, tt.args.b)
			assert.Equal(t, v, tt.want, v)
		})
	}
}

func TestIsLessThanOrdered(t *testing.T) {
	type args struct {
		a genfuncs.Function[int, bool]
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Greater",
			args: args{
				a: genfuncs.IsLessThanOrdered(2),
				b: 1,
			},
			want: false,
		},
		{
			name: "Equal",
			args: args{
				a: genfuncs.IsLessThanOrdered(1),
				b: 1,
			},
			want: false,
		},
		{
			name: "Less",
			args: args{
				a: genfuncs.IsLessThanOrdered(0),
				b: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.args.a(tt.args.b)
			assert.Equal(t, v, tt.want, v)
		})
	}
}
