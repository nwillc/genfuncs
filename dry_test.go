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
	"strconv"
	"testing"
	"time"
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

func TestMax(t *testing.T) {
	type args struct {
		v []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name:    "No Args",
			args:    args{},
			wantErr: genfuncs.IllegalArguments,
		},
		{
			name: "Greater",
			args: args{
				v: []int{2, 1},
			},
			want: 2,
		},
		{
			name: "Equal",
			args: args{
				v: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "Less",
			args: args{
				v: []int{0, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr != nil {
				assert.Panics(t, func() {
					genfuncs.Max(tt.args.v...)
				})
				return
			}
			v := genfuncs.Max(tt.args.v...)
			assert.Equal(t, v, tt.want, v)
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		v []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name:    "No Args",
			args:    args{},
			wantErr: genfuncs.IllegalArguments,
		},
		{
			name: "Greater",
			args: args{
				v: []int{2, 1},
			},
			want: 1,
		},
		{
			name: "Equal",
			args: args{
				v: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "Less",
			args: args{
				v: []int{0, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr != nil {
				assert.Panics(t, func() {
					genfuncs.Min(tt.args.v...)
				})
				return
			}
			v := genfuncs.Min(tt.args.v...)
			assert.Equal(t, v, tt.want, v)
		})
	}
}

func TestReverse(t *testing.T) {
	reversed := genfuncs.Reverse(genfuncs.LessThanOrdered[int])
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Less",
			args: args{
				a: 1,
				b: 2,
			},
		},
		{
			name: "Equal",
			args: args{
				a: 0,
				b: 0,
			},
		},
		{
			name: "Greater",
			args: args{
				a: 1,
				b: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := genfuncs.LessThanOrdered(tt.args.a, tt.args.b)
			assert.Equal(t, tt.args.a < tt.args.b, v)
			r := reversed(tt.args.a, tt.args.b)
			assert.Equal(t, tt.args.b < tt.args.a, r)
		})
	}
}

func TestNot(t *testing.T) {
	var echo genfuncs.Function[bool, bool] = func(b bool) bool { return b }
	var notEcho = genfuncs.Not(echo)
	assert.Equal(t, echo(true), true)
	assert.Equal(t, notEcho(true), false)
	assert.Equal(t, echo(false), false)
	assert.Equal(t, notEcho(false), true)
}

func TestStringerToString(t *testing.T) {
	now := time.Now()
	ts := genfuncs.StringerToString[time.Time]()

	assert.Equal(t, ts(now), now.String())
}

func TestTransformArgs(t *testing.T) {
	var atoi genfuncs.Function[string, int] = func(s string) int { i, _ := strconv.Atoi(s); return i }
	var adder = func(a, b int) int { return a + b }
	strAdder := genfuncs.TransformArgs(atoi, adder)

	assert.Equal(t, 10, strAdder("5", "5"))
}
