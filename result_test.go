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
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResult(t *testing.T) {
	var r *genfuncs.Result[int]
	err := fmt.Errorf("ro ruh")
	r = genfuncs.NewError[int](err)
	r.
		OnFailure(func(e error) {
			assert.Equal(t, err, e)
		}).
		OnSuccess(func(_ int) {
			assert.Fail(t, "success on an error")
		})

	assert.Panics(t, func() {
		r.ValueOrPanic()
	})

	assert.Equal(t, 10, r.ValueOr(10))

	r = genfuncs.NewResult(10)
	assert.Equal(t, 10, r.ValueOrPanic())

	r = r.Then(func(i int) *genfuncs.Result[int] {
		return genfuncs.NewResult(i * 10)
	})
	assert.Equal(t, 100, r.ValueOrPanic())
}

func TestResult_Error(t *testing.T) {
	type args struct {
		result *genfuncs.Result[int]
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no error",
			args: args{
				result: genfuncs.NewResult(1),
			},
		},
		{
			name: "error",
			args: args{
				result: genfuncs.NewError[int](fmt.Errorf("foo")),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				assert.False(t, tt.args.result.Ok())
				assert.NotNil(t, tt.args.result.Error())
			} else {
				assert.True(t, tt.args.result.Ok())
				assert.Nil(t, tt.args.result.Error())
			}
		})
	}
}

func TestResult_OnSuccess(t *testing.T) {
	flag := -1
	action := func(i int) { flag = i }
	type args struct {
		result *genfuncs.Result[int]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "error",
			args: args{
				result: genfuncs.NewError[int](fmt.Errorf("")),
			},
		},
		{
			name: "10",
			args: args{
				result: genfuncs.NewResult(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			action(0)
			assert.Equal(t, 0, flag)
			tt.args.result.OnSuccess(action)
			if tt.args.result.Ok() {
				assert.Equal(t, flag, tt.args.result.ValueOrPanic())
			} else {
				assert.Equal(t, flag, 0)
			}
		})
	}
}

func TestResult_String(t *testing.T) {
	type args struct {
		result *genfuncs.Result[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "value",
			args: args{
				result: genfuncs.NewResult(10),
			},
			want: "10",
		},
		{
			name: "error",
			args: args{
				result: genfuncs.NewError[int](fmt.Errorf("no value")),
			},
			want: "error: no value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.args.result.String())
		})
	}
}

func TestResult_Then(t *testing.T) {
	type args struct {
		value *genfuncs.Result[string]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{
				value: genfuncs.NewResult("foo"),
			},
			want: "foosimple",
		},
		{
			name: "error",
			args: args{
				value: genfuncs.NewError[string](fmt.Errorf("foo")),
			},
			want: "error: foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.args.value.Then(func(s string) *genfuncs.Result[string] { return genfuncs.NewResult(s + tt.name) })
			assert.Equal(t, tt.want, v.String())
		})
	}
}

func TestResult_ValueOr(t *testing.T) {
	type args struct {
		result *genfuncs.Result[int]
		value  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "10",
			args: args{
				result: genfuncs.NewResult(10),
				value:  100,
			},
			want: 10,
		},
		{
			name: "error",
			args: args{
				result: genfuncs.NewError[int](fmt.Errorf("foo")),
				value:  100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.args.result.ValueOr(tt.args.value))
		})
	}
}
