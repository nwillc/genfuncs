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

package promises

import (
	"context"
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/results"
)

var (
	PromiseAnyNoPromisesErrorMsg = "no any without promises"
	PromiseNoneFulfilled         = "no promises fulfilled"
)

type promiseResult[T any] struct {
	result *genfuncs.Result[T]
	index  int
}

// All accepts promises and collects their results, returning a container.GSlice of the results in correlating order,
// or if *any* genfuncs.Promise fails then All returns its error and immediately returns.
func All[T any](ctx context.Context, promises ...*genfuncs.Promise[T]) *genfuncs.Promise[container.GSlice[T]] {
	count := len(promises)
	promiseResults := make(container.GSlice[T], count)
	if count == 0 {
		return genfuncs.NewPromiseFromResult(genfuncs.NewResult(promiseResults))
	}
	return genfuncs.NewPromise(
		ctx,
		func(_ context.Context) *genfuncs.Result[container.GSlice[T]] {
			resultChan := make(chan promiseResult[T], count)
			for i := 0; i < count; i++ {
				i := i
				promises[i].
					OnSuccess(func(value T) { resultChan <- promiseResult[T]{result: genfuncs.NewResult(value), index: i} }).
					OnError(func(err error) { resultChan <- promiseResult[T]{result: genfuncs.NewError[T](err), index: i} })
			}
			for i := 0; i < count; i++ {
				select {
				case r := <-resultChan:
					if !r.result.Ok() {
						return results.MapError[T, container.GSlice[T]](r.result)
					}
					promiseResults[r.index] = r.result.OrEmpty()
				}
			}
			return genfuncs.NewResult(promiseResults)
		})
}

// Any returns a Promise that will return the first Promise fulfilled, or an error if none were.
func Any[T any](ctx context.Context, promises ...*genfuncs.Promise[T]) *genfuncs.Promise[T] {
	count := len(promises)
	if count == 0 {
		return genfuncs.NewPromiseFromResult(genfuncs.NewError[T](fmt.Errorf(PromiseAnyNoPromisesErrorMsg)))
	}
	return genfuncs.NewPromise(
		ctx,
		func(_ context.Context) *genfuncs.Result[T] {
			resultChan := make(chan promiseResult[T], count)
			for i := 0; i < count; i++ {
				i := i
				promises[i].
					OnSuccess(func(value T) { resultChan <- promiseResult[T]{result: genfuncs.NewResult(value), index: i} }).
					OnError(func(err error) { resultChan <- promiseResult[T]{result: genfuncs.NewError[T](err), index: i} })
			}
			for i := 0; i < count; i++ {
				select {
				case r := <-resultChan:
					if r.result.Ok() {
						return r.result
					}
				}
			}
			return genfuncs.NewError[T](fmt.Errorf(PromiseNoneFulfilled))
		})
}

// Map will Wait for aPromise and then return a new Promise which then maps its result.
func Map[A, B any](ctx context.Context, aPromise *genfuncs.Promise[A], then genfuncs.Function[A, *genfuncs.Result[B]]) *genfuncs.Promise[B] {
	return genfuncs.NewPromise(ctx, func(_ context.Context) *genfuncs.Result[B] {
		return results.Map[A, B](aPromise.Wait(), then)
	})
}
