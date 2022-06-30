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

package maps

import "github.com/nwillc/genfuncs"

type (
	// Entry can be used to hold onto a key/value.
	Entry[K comparable, V any] struct {
		Key   K
		Value V
	}

	// KeyFor is used for generating keys from types, it accepts any type and returns a comparable key for it.
	KeyFor[T any, K comparable] func(T) *genfuncs.Result[K]

	// KeyValueFor is used to generate a key and value from a type, it accepts any type, and returns a comparable key and
	// any value.
	KeyValueFor[T any, K comparable, V any] func(T) *genfuncs.Result[*Entry[K, V]]

	// ValueFor given a comparable key will return a value for it.
	ValueFor[K comparable, T any] func(K) *genfuncs.Result[T]
)
