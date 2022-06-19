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

package promise

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/result"
)

var PromiseAnyNoPromisesErrorMsg = "no any of none"

// Map will Await aPromise and then return a new Promise which then maps its result.
func Map[A, B any](aPromise *genfuncs.Promise[A], then genfuncs.Function[A, *genfuncs.Result[B]]) *genfuncs.Promise[B] {
	return genfuncs.NewPromise(func() *genfuncs.Result[B] {
		result1 := aPromise.Await()
		if !result1.Ok() {
			return result.MapError[A, B](result1)
		}
		return then(result1.OrEmpty())
	})
}

type promiseResult[T any] struct {
	result *genfuncs.Result[T]
	index  int
}

// All accepts promises and collects their results, returning a container.GSlice of the results in correlating order,
// or if *any* genfuncs.Promise fails then All returns its error and immediately returns.
func All[T any](promises ...*genfuncs.Promise[T]) *genfuncs.Promise[container.GSlice[T]] {
	count := len(promises)
	results := make(container.GSlice[T], count)
	if len(promises) == 0 {
		return genfuncs.NewPromiseFromResult(genfuncs.NewResult(results))
	}
	return genfuncs.NewPromise(func() *genfuncs.Result[container.GSlice[T]] {
		resultChan := make(chan promiseResult[T], count)

		for i := 0; i < count; i++ {
			index := i
			_ = Map[T, T](promises[i], func(value T) *genfuncs.Result[T] {
				r := genfuncs.NewResult(value)
				resultChan <- promiseResult[T]{result: r, index: index}
				return r
			})
			_ = promises[i].Catch(func(err error) {
				r := genfuncs.NewError[T](err)
				resultChan <- promiseResult[T]{result: r, index: index}
			})
		}

		for i := 0; i < count; i++ {
			select {
			case r := <-resultChan:
				if !r.result.Ok() {
					return genfuncs.NewError[container.GSlice[T]](r.result.Error())
				}
				results[r.index] = r.result.OrEmpty()
			}
		}
		return genfuncs.NewResult(results)
	})
}

func Any[T any](promises ...*genfuncs.Promise[T]) *genfuncs.Promise[T] {
	count := len(promises)
	if count == 0 {
		return genfuncs.NewPromiseFromResult(genfuncs.NewError[T](fmt.Errorf(PromiseAnyNoPromisesErrorMsg)))
	}
	return genfuncs.NewPromise(func() *genfuncs.Result[T] {
		resultChan := make(chan promiseResult[T], count)

		for i := 0; i < count; i++ {
			index := i
			_ = Map[T, T](promises[i], func(value T) *genfuncs.Result[T] {
				r := genfuncs.NewResult(value)
				resultChan <- promiseResult[T]{result: r, index: index}
				return r
			})
			_ = promises[i].Catch(func(err error) {
				r := genfuncs.NewError[T](err)
				resultChan <- promiseResult[T]{result: r, index: index}
			})
		}

		for i := 0; i < count; i++ {
			select {
			case r := <-resultChan:
				if r.result.Ok() {
					return r.result
				}
			}
		}
		return genfuncs.NewError[T](fmt.Errorf("none"))
	})
}
