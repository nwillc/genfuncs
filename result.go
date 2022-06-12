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

import "fmt"

var _ fmt.Stringer = (*Result[int])(nil)

// Result is an implementation of the Maybe pattern. This is mostly for experimentation as it is a poor fit with Go's
// traditional idiomatic error handling.
type Result[T any] struct {
	value T
	err   error
}

// NewResult for a value.
func NewResult[T any](t T) *Result[T] {
	return &Result[T]{value: t}
}

// NewResultFromTuple creates a Result from a value, error tuple.
func NewResultFromTuple[T any](t T, err error) *Result[T] {
	return &Result[T]{value: t, err: err}
}

// NewError for an error.
func NewError[T any](err error) *Result[T] {
	return &Result[T]{err: err}
}

// Error of the Result, nil if Ok().
func (r *Result[T]) Error() error {
	return r.err
}

// Ok returns the status of Result, is it ok, or an error.
func (r *Result[T]) Ok() bool {
	return r.err == nil
}

// OnFailure performs the action is Result is not Ok().
func (r *Result[T]) OnFailure(action func(e error)) *Result[T] {
	if !r.Ok() {
		action(r.err)
	}
	return r
}

// OnSuccess performs action if Result is Ok().
func (r *Result[T]) OnSuccess(action func(t T)) *Result[T] {
	if r.Ok() {
		action(r.value)
	}
	return r
}

// String returns a string representation of Result, either the value or error.
func (r *Result[T]) String() string {
	if r.Ok() {
		return fmt.Sprint(r.value)
	}

	return "error: " + r.err.Error()
}

// Map calls action on Result if it's Ok() and returns its Result, if not Ok() the original result is returned.
func (r *Result[T]) Map(action func(t T) *Result[T]) *Result[T] {
	if r.Ok() {
		return action(r.value)
	}
	return r
}

// OrElse returns the value of the Result if Ok(), or the value v if not.
func (r *Result[T]) OrElse(v T) T {
	if r.Ok() {
		return r.value
	}
	return v
}

// OrEmpty will return the value of the Result or the empty value if Error.
func (r *Result[T]) OrEmpty() T {
	return r.value
}

// MustGet returns the value of the Result if Ok() or if not, panics with the error.
func (r *Result[T]) MustGet() T {
	if r.Ok() {
		return r.value
	}
	panic(r.err)
}
