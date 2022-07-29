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
	"context"
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPromise_Wait(t *testing.T) {
	type args struct {
		action func(context.Context) *genfuncs.Result[int]
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
				action: func(_ context.Context) *genfuncs.Result[int] {
					return genfuncs.NewResult(1)
				},
			},
			wantOk: true,
			want:   1,
		},
		{
			name: "deliver slow action",
			args: args{
				action: func(_ context.Context) *genfuncs.Result[int] {
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
				action: func(_ context.Context) *genfuncs.Result[int] {
					return genfuncs.NewError[int](genfuncs.NoSuchElement)
				},
			},
			wantOk:       false,
			wantErrorMsg: genfuncs.NoSuchElement.Error(),
		},
		{
			name: "action panic message",
			args: args{
				action: func(_ context.Context) *genfuncs.Result[int] {
					panic("sky is falling")
				},
			},
			wantOk:       false,
			wantErrorMsg: genfuncs.PromisePanicErrorMsg,
		},
		{
			name: "action panic error",
			args: args{
				action: func(_ context.Context) *genfuncs.Result[int] {
					panic(genfuncs.NoSuchElement)
				},
			},
			wantOk:       false,
			wantErrorMsg: genfuncs.NoSuchElement.Error(),
		},
		{
			name: "action panic nil",
			args: args{
				action: func(_ context.Context) *genfuncs.Result[int] {
					panic(nil)
				},
			},
			wantOk:       false,
			wantErrorMsg: genfuncs.PromisePanicErrorMsg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			future := genfuncs.NewPromise[int](context.Background(), tt.args.action)
			result := future.Wait()
			assert.Equal(t, tt.wantOk, result.Ok())
			if !tt.wantOk {
				assert.Contains(t, result.Error().Error(), tt.wantErrorMsg)
				return
			}
			assert.Equal(t, tt.want, result.OrEmpty())
		})
	}
}

func TestPromise_OnError(t *testing.T) {
	var errorCount int
	errorAdd := func(err error) { errorCount++ }

	type args struct {
		aPanic bool
		f1     *genfuncs.Result[int]
	}
	tests := []struct {
		name       string
		args       args
		wantOk     bool
		errorCount int
		errorMsg   string
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
			errorMsg:   genfuncs.NoSuchElement.Error(),
		},
		{
			name: "panic",
			args: args{
				aPanic: true,
			},
			wantOk:     false,
			errorCount: 1,
			errorMsg:   genfuncs.PromisePanicErrorMsg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorCount = 0
			p1 := genfuncs.NewPromise(context.Background(), func(_ context.Context) *genfuncs.Result[int] {
				if tt.args.aPanic {
					panic(tt.name)
				}
				return tt.args.f1
			})
			p2 := p1.OnError(errorAdd)
			result := p2.Wait()
			assert.Equal(t, tt.wantOk, result.Ok())
			assert.Equal(t, tt.errorCount, errorCount)
			if !tt.wantOk {
				assert.Contains(t, result.Error().Error(), tt.errorMsg)
			}
		})
	}
}

func TestPromise_OnSuccess_OnError(t *testing.T) {
	count := 0
	p := genfuncs.NewPromise[bool](
		context.Background(),
		func(_ context.Context) *genfuncs.Result[bool] {
			return genfuncs.NewResult(true)
		}).
		OnSuccess(func(_ bool) {
			count++
		}).
		OnError(func(e error) {
			count++
		})
	r := p.Wait()
	assert.True(t, r.Ok())
	assert.True(t, r.OrEmpty())
	assert.Equal(t, 1, count)
}

func TestPromise_MultiWait(t *testing.T) {
	count := 0
	p := genfuncs.NewPromise[bool](
		context.Background(),
		func(_ context.Context) *genfuncs.Result[bool] {
			return genfuncs.NewResult(true)
		}).
		OnSuccess(func(_ bool) {
			count++
		})
	r := p.Wait()
	assert.True(t, r.Ok())
	assert.True(t, r.OrEmpty())
	assert.Equal(t, 1, count)

	// Safe to call wait again.
	r = p.Wait()
	assert.True(t, r.Ok())
	assert.True(t, r.OrEmpty())
	assert.Equal(t, 1, count)
}

func TestPromise_Cancel(t *testing.T) {
	waitForCancel := 500 * time.Millisecond
	cancelAction := func(ctx context.Context) *genfuncs.Result[bool] {
		select {
		case <-ctx.Done():
			return genfuncs.NewError[bool](fmt.Errorf("cancelled"))
		case <-time.After(waitForCancel):
			return genfuncs.NewResult(true)
		}
	}

	ctx := context.Background()
	p := genfuncs.NewPromise(ctx, cancelAction)
	r := p.Wait()
	assert.True(t, r.Ok())

	ctx = context.Background()
	p = genfuncs.NewPromise(ctx, cancelAction)
	go func() {
		time.Sleep(waitForCancel / 2)
		p.Cancel()
	}()
	r = p.Wait()
	assert.False(t, r.Ok())
	assert.Contains(t, r.Error().Error(), "cancel")
}
