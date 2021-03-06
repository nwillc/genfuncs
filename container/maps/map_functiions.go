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

import (
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
)

// Map returns a GSlice containing the results of applying the given transform function to each element in the GMap.
func Map[K comparable, V any, R any](m container.Map[K, V], transform genfuncs.BiFunction[K, V, R]) (result container.GSlice[R]) {
	result = make(container.GSlice[R], m.Len())
	i := 0
	m.ForEach(func(k K, v V) {
		result[i] = transform(k, v)
		i++
	})
	return result
}

// Merge merges maps of container.GSlice's together into a new map appending the container.GSlice's when collisions occur.
func Merge[K comparable, V any](mv ...container.Map[K, container.GSlice[V]]) (result container.GMap[K, container.GSlice[V]]) {
	result = make(container.GMap[K, container.GSlice[V]])
	for _, m := range mv {
		m.ForEach(func(k K, v container.GSlice[V]) {
			v1 := result[k]
			result[k] = append(v1, v...)
		})
	}
	return result
}

func NewEntry[K comparable, V any](k K, v V) *Entry[K, V] {
	return &Entry[K, V]{Key: k, Value: v}
}
