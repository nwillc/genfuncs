/*
 *  Copyright (c) 2021,  nwillc@gmail.com
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

// BiFunction accepts two arguments and produces a result.
type BiFunction[T, U, R any] func(T, U) R

// Function is a single argument function.
type Function[T, R any] func(T) R

// MapKeyFor is used for generating keys from types, it accepts any type and returns a comparable key for it.
type MapKeyFor[T any, K comparable] func(T) K

// MapKeyValueFor is used to generate a key and value from a type, it accepts any type, and returns a comparable key and
// any value.
type MapKeyValueFor[T any, K comparable, V any] func(T) (K, V)

// MapValueFor given a comparable key will return a value for it.
type MapValueFor[K comparable, T any] func(K) T

// ToString is used to create string representations, it accepts any type and returns a string.
type ToString[T any] func(T) string
