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

package container

import (
	"github.com/nwillc/genfuncs"
	"golang.org/x/exp/maps"
)

var _ HasValues[int] = (GMap[int, int])(nil)

// GMap is a generic type corresponding to a standard Go map and implements HasValues.
type GMap[K comparable, V any] map[K]V

// All returns true if all values in GMap satisfy the predicate.
func (m GMap[K, V]) All(predicate genfuncs.Function[V, bool]) bool {
	for _, v := range m {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// Any returns true if any values in GMap satisfy the predicate.
func (m GMap[K, V]) Any(predicate genfuncs.Function[V, bool]) bool {
	for _, v := range m {
		if predicate(v) {
			return true
		}
	}
	return false
}

// Contains returns true if the GMap contains the given key.
func (m GMap[K, V]) Contains(key K) bool {
	_, ok := m[key]
	return ok
}

// Filter a GMap by a predicate, returning a new GMap that contains only values that satisfy the predicate.
func (m GMap[K, V]) Filter(predicate genfuncs.Function[V, bool]) GMap[K, V] {
	result := make(GMap[K, V])
	for k, v := range m {
		if !predicate(v) {
			continue
		}
		result[k] = v
	}
	return result
}

// FilterKeys returns a new GMap that contains only values whose key satisfy the predicate.
func (m GMap[K, V]) FilterKeys(predicate genfuncs.Function[K, bool]) GMap[K, V] {
	result := make(GMap[K, V])
	for k, v := range m {
		if !predicate(k) {
			continue
		}
		result[k] = v
	}
	return result
}

// ForEach performs the given action on each entry in the GMap.
func (m GMap[K, V]) ForEach(action func(k K, v V)) {
	for k, v := range m {
		action(k, v)
	}
}

// Keys return a GSlice containing the keys of the GMap.
func (m GMap[K, V]) Keys() GSlice[K] {
	return maps.Keys(m)
}

// Len is the number of elements in the GMap.
func (m GMap[K, V]) Len() int {
	return len(m)
}

// Values returns a GSlice of all the values in the GMap.
func (m GMap[K, V]) Values() GSlice[V] {
	return maps.Values(m)
}
