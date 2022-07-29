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

package promises_test

import (
	"context"
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/container/gslices"
	"github.com/nwillc/genfuncs/promises"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		f1 func(ctx context.Context) *genfuncs.Result[int]
		f2 func(int) *genfuncs.Result[string]
	}
	tests := []struct {
		name   string
		args   args
		wantOk bool
		want   string
	}{
		{
			name: "both pass",
			args: args{
				f1: func(_ context.Context) *genfuncs.Result[int] { return genfuncs.NewResult(1) },
				f2: func(i int) *genfuncs.Result[string] { return genfuncs.NewResult(fmt.Sprintf("%d", i)) },
			},
			wantOk: true,
			want:   "1",
		},
		{
			name: "first fails",
			args: args{
				f1: func(_ context.Context) *genfuncs.Result[int] { return genfuncs.NewError[int](fmt.Errorf("first")) },
				f2: func(i int) *genfuncs.Result[string] { return genfuncs.NewResult(fmt.Sprintf("%d", i)) },
			},
			wantOk: false,
			want:   "first",
		},
		{
			name: "second fails",
			args: args{
				f1: func(_ context.Context) *genfuncs.Result[int] { return genfuncs.NewResult(1) },
				f2: func(i int) *genfuncs.Result[string] { return genfuncs.NewError[string](fmt.Errorf("second")) },
			},
			wantOk: false,
			want:   "second",
		},
		{
			name: "both fail",
			args: args{
				f1: func(_ context.Context) *genfuncs.Result[int] { return genfuncs.NewError[int](fmt.Errorf("first")) },
				f2: func(i int) *genfuncs.Result[string] { return genfuncs.NewError[string](fmt.Errorf("second")) },
			},
			wantOk: false,
			want:   "first",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			p1 := genfuncs.NewPromise[int](ctx, tt.args.f1)
			p2 := promises.Map(ctx, p1, tt.args.f2)
			result := p2.Wait()
			assert.Equal(t, tt.wantOk, result.Ok())
			if !result.Ok() {
				assert.Contains(t, tt.want, result.Error().Error())
				return
			}
			assert.Equal(t, tt.want, result.OrEmpty())
		})
	}
}

func TestAll(t *testing.T) {
	type args struct {
		actions []func(ctx context.Context) *genfuncs.Result[int]
	}
	tests := []struct {
		name   string
		args   args
		want   container.GSlice[int]
		wantOk bool
	}{
		{
			name: "empty",
			args: args{
				actions: []func(context.Context) *genfuncs.Result[int]{},
			},
			want:   container.GSlice[int]{},
			wantOk: true,
		},
		{
			name: "one success",
			args: args{
				actions: []func(context.Context) *genfuncs.Result[int]{
					func(_ context.Context) *genfuncs.Result[int] { return genfuncs.NewResult(1) },
				},
			},
			want:   container.GSlice[int]{1},
			wantOk: true,
		},
		{
			name: "two success",
			args: args{
				actions: []func(ctx context.Context) *genfuncs.Result[int]{
					func(_ context.Context) *genfuncs.Result[int] { return genfuncs.NewResult(1) },
					func(_ context.Context) *genfuncs.Result[int] { return genfuncs.NewResult(2) },
				},
			},
			want:   container.GSlice[int]{1, 2},
			wantOk: true,
		},
		{
			name: "two promises one error",
			args: args{
				actions: []func(ctx context.Context) *genfuncs.Result[int]{
					func(_ context.Context) *genfuncs.Result[int] { return genfuncs.NewResult(1) },
					func(_ context.Context) *genfuncs.Result[int] { return genfuncs.NewError[int](genfuncs.NoSuchElement) },
				},
			},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			var p []*genfuncs.Promise[int] = gslices.Map(
				tt.args.actions,
				func(f func(context.Context) *genfuncs.Result[int]) *genfuncs.Promise[int] {
					return genfuncs.NewPromise[int](context.Background(), f)
				})
			all := promises.All(ctx, p...)
			result := all.Wait()
			assert.Equal(t, tt.wantOk, result.Ok())
			if tt.wantOk {
				assert.Equal(t, tt.want, result.OrEmpty())
			}
		})
	}
}

func TestAny(t *testing.T) {
	type args struct {
		actions []func(context.Context) *genfuncs.Result[string]
	}
	tests := []struct {
		name   string
		args   args
		want   string
		wantOk bool
	}{
		{
			name: "empty",
			args: args{
				actions: []func(context.Context) *genfuncs.Result[string]{},
			},
			want:   promises.PromiseAnyNoPromisesErrorMsg,
			wantOk: false,
		},
		{
			name: "single success",
			args: args{
				actions: []func(context.Context) *genfuncs.Result[string]{
					func(_ context.Context) *genfuncs.Result[string] { return genfuncs.NewResult("one") },
				},
			},
			want:   "one",
			wantOk: true,
		},
		{
			name: "all error",
			args: args{
				actions: []func(context.Context) *genfuncs.Result[string]{
					func(_ context.Context) *genfuncs.Result[string] {
						return genfuncs.NewError[string](genfuncs.NoSuchElement)
					},
				},
			},
			want:   promises.PromiseNoneFulfilled,
			wantOk: false,
		},
		{
			name: "second success",
			args: args{
				actions: []func(context.Context) *genfuncs.Result[string]{
					func(_ context.Context) *genfuncs.Result[string] {
						return genfuncs.NewError[string](genfuncs.NoSuchElement)
					},
					func(_ context.Context) *genfuncs.Result[string] { return genfuncs.NewResult("second") },
				},
			},
			want:   "second",
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			var p []*genfuncs.Promise[string] = gslices.Map(
				tt.args.actions,
				func(f func(context.Context) *genfuncs.Result[string]) *genfuncs.Promise[string] {
					return genfuncs.NewPromise[string](ctx, f)
				})
			one := promises.Any(ctx, p...)
			result := one.Wait()
			assert.Equal(t, tt.wantOk, result.Ok())
			if !tt.wantOk {
				assert.Contains(t, tt.want, result.Error().Error())
				return
			}
			assert.Equal(t, tt.want, result.OrEmpty())
		})
	}
}
