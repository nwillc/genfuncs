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
	"time"
)

func TestFutureAwait(t *testing.T) {
	type args struct {
		action func() *genfuncs.Result[int]
	}
	tests := []struct {
		name         string
		args         args
		want         int
		wantOk       bool
		wantErrorMsg string
	}{
		{
			name:         "nil action",
			args:         args{},
			wantOk:       false,
			wantErrorMsg: genfuncs.PromiseNoActionErrorMsg,
		},
		{
			name: "deliver action",
			args: args{
				action: func() *genfuncs.Result[int] {
					return genfuncs.NewResult(1)
				},
			},
			wantOk: true,
			want:   1,
		},
		{
			name: "deliver slow action",
			args: args{
				action: func() *genfuncs.Result[int] {
					time.Sleep(400 * time.Millisecond)
					return genfuncs.NewResult(2)
				},
			},
			wantOk: true,
			want:   2,
		},
		{
			name: "deliver error",
			args: args{
				action: func() *genfuncs.Result[int] {
					return genfuncs.NewError[int](genfuncs.NoSuchElement)
				},
			},
			wantOk:       false,
			wantErrorMsg: genfuncs.NoSuchElement.Error(),
		},
		{
			name: "action panic",
			args: args{
				action: func() *genfuncs.Result[int] {
					panic("sky is falling")
				},
			},
			wantOk:       false,
			wantErrorMsg: genfuncs.PromisePanicErrorMsg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			future := genfuncs.NewPromise[int](tt.args.action)
			result := future.Await()
			assert.Equal(t, tt.wantOk, result.Ok())
			if !tt.wantOk {
				assert.Contains(t, result.Error().Error(), tt.wantErrorMsg)
				return
			}
			assert.Equal(t, tt.want, result.OrEmpty())
		})
	}
}

func TestPromise_Catch(t *testing.T) {
	var errorCount int
	errorAdd := func(err error) { errorCount++ }

	type args struct {
		f1 *genfuncs.Result[int]
	}
	tests := []struct {
		name       string
		args       args
		wantOk     bool
		errorCount int
	}{
		{
			name: "no error",
			args: args{
				f1: genfuncs.NewResult(1),
			},
			wantOk:     true,
			errorCount: 0,
		},
		{
			name: "error",
			args: args{
				f1: genfuncs.NewError[int](genfuncs.NoSuchElement),
			},
			wantOk:     false,
			errorCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorCount = 0
			p1 := genfuncs.NewPromise(func() *genfuncs.Result[int] {
				return tt.args.f1
			})
			p2 := p1.Catch(errorAdd)
			result := p2.Await()
			assert.Equal(t, tt.wantOk, result.Ok())
			assert.Equal(t, tt.errorCount, errorCount)
		})
	}
}
