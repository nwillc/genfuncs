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
	"fmt"
	"sync"
)

// Promise provides asynchronous Result of an action.
type Promise[T any] struct {
	result  *Result[T]
	pending bool
	mutex   sync.Mutex
	wg      sync.WaitGroup
}

var (
	// PromiseNoActionErrorMsg indicates a Promise was created for no action.
	PromiseNoActionErrorMsg = "promise requested with no action"
	// PromisePanicErrorMsg indicates the action of a Promise caused a panic.
	PromisePanicErrorMsg = "promise action panic"
)

// NewPromise creates a Promise for an action.
func NewPromise[T any](action func() *Result[T]) *Promise[T] {
	if action == nil {
		return NewPromiseFromResult(NewError[T](fmt.Errorf(PromiseNoActionErrorMsg)))
	}

	p := &Promise[T]{
		pending: true,
	}

	p.wg.Add(1)

	go func() {
		defer p.deliverErrorOnPanic()
		p.deliver(action())
	}()

	return p
}

// NewPromiseFromResult returns a completed Promise with the specified result.
func NewPromiseFromResult[T any](result *Result[T]) *Promise[T] {
	return &Promise[T]{
		result:  result,
		pending: false,
	}
}

// Await the completion of a Promise.
func (p *Promise[T]) Await() *Result[T] {
	p.wg.Wait()
	return p.result
}

// Catch returns a new Promise that adds an error handler to a Promise.
func (p *Promise[T]) Catch(err func(error)) *Promise[T] {
	return NewPromise[T](func() *Result[T] {
		result := p.Await()
		if !result.Ok() {
			err(result.Error())
		}
		return result
	})
}

// deliver on a Promise with a Result.
func (p *Promise[T]) deliver(value *Result[T]) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if !p.pending {
		return
	}
	p.result = value
	p.pending = false
	p.wg.Done()
}

// deliverErrorOnPanic converts an action panic to an error Result.
func (p *Promise[T]) deliverErrorOnPanic() {
	var err error
	recovered := recover()
	if validErr, ok := recovered.(error); ok {
		err = fmt.Errorf("%s: %w", PromisePanicErrorMsg, validErr)
	} else {
		err = fmt.Errorf("%s: %+v", PromisePanicErrorMsg, recovered)
	}
	p.deliver(NewError[T](err))
}
