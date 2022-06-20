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

package results

import (
	"github.com/nwillc/genfuncs"
)

// Map one Result type to another employing a transform.
func Map[T, R any](t *genfuncs.Result[T], transform genfuncs.Function[T, *genfuncs.Result[R]]) (r *genfuncs.Result[R]) {
	if t.Ok() {
		r = transform(t.OrEmpty())
	} else {
		r = genfuncs.NewError[R](t.Error())
	}
	return r
}

// MapError Result of one type to another.
func MapError[T, R any](t *genfuncs.Result[T]) (r *genfuncs.Result[R]) {
	r = genfuncs.NewError[R](t.Error())
	return r
}
