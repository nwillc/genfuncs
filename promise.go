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

package genfuncs

import (
	"context"
	"fmt"
	"sync"
)

// Promise provides asynchronous Result of an action.
type Promise[T any] struct {
	result      *Result[T]
	wg          sync.WaitGroup
	ctx         context.Context
	cancel      context.CancelFunc
	deliverOnce sync.Once
}

var (
	// PromiseNoActionErrorMsg indicates a Promise was created for no action.
	PromiseNoActionErrorMsg = "promise requested with no action"
	// PromisePanicErrorMsg indicates the action of a Promise caused a panic.
	PromisePanicErrorMsg = "promise action panic"
)

// NewPromise creates a Promise for an action.
func NewPromise[T any](ctx context.Context, action func(context.Context) *Result[T]) *Promise[T] {
	if action == nil {
		return NewPromiseFromResult(NewError[T](fmt.Errorf(PromiseNoActionErrorMsg)))
	}
	pctx, cancel := context.WithCancel(ctx)
	p := &Promise[T]{
		ctx:    pctx,
		cancel: cancel,
	}

	p.wg.Add(1)

	go func() {
		done := false
		defer func() {
			if !done {
				recovered := recover()
				var err error
				if validErr, ok := recovered.(error); ok {
					err = fmt.Errorf("%s: %w", PromisePanicErrorMsg, validErr)
				} else {
					err = fmt.Errorf("%s: %+v", PromisePanicErrorMsg, recovered)
				}
				p.deliver(NewError[T](err))
			}
		}()
		result := action(pctx)
		p.deliver(result)
		done = true
	}()

	return p
}

// NewPromiseFromResult returns a completed Promise with the specified result.
func NewPromiseFromResult[T any](result *Result[T]) *Promise[T] {
	return &Promise[T]{
		result: result,
	}
}

// Cancel the Promise which will allow any action that is listening on `<-ctx.Done()` to complete.
func (p *Promise[T]) Cancel() {
	p.cancel()
}

// OnError returns a new Promise with an error handler waiting on the original Promise.
func (p *Promise[T]) OnError(action func(error)) *Promise[T] {
	return NewPromise(p.ctx, func(_ context.Context) *Result[T] {
		return p.Wait().OnError(action)
	})
}

// OnSuccess returns a new Promise with a success handler waiting on the original Promise.
func (p *Promise[T]) OnSuccess(action func(t T)) *Promise[T] {
	return NewPromise(p.ctx, func(_ context.Context) *Result[T] {
		return p.Wait().OnSuccess(action)
	})
}

// Wait on the completion of a Promise.
func (p *Promise[T]) Wait() *Result[T] {
	p.wg.Wait()
	return p.result
}

// deliver on a Promise with a Result.
func (p *Promise[T]) deliver(result *Result[T]) {
	p.deliverOnce.Do(func() {
		p.result = result
		p.wg.Done()
	})
}
