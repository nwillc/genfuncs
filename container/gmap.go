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

// GMap implements the Map interface.
var (
	_ Map[int, int]    = (GMap[int, int])(nil)
	_ Iterable[string] = (GMap[int, string])(nil)
	_ Iterator[string] = (*gmapIterator[int, string])(nil)
)

// GMap is a generic type employing the standard Go map and implementation Map.
type (
	GMap[K comparable, V any]         map[K]V
	gmapIterator[K comparable, V any] struct {
		gmap     GMap[K, V]
		iterator Iterator[K]
	}
)

// All returns true if all values in GMap satisfy the predicate.
func (m GMap[K, V]) All(predicate genfuncs.Function[V, bool]) (ok bool) {
	for _, v := range m {
		if !predicate(v) {
			return ok
		}
	}
	ok = true
	return ok
}

// Any returns true if any values in GMap satisfy the predicate.
func (m GMap[K, V]) Any(predicate genfuncs.Function[V, bool]) (ok bool) {
	for _, v := range m {
		if predicate(v) {
			ok = true
			return ok
		}
	}
	return ok
}

// Contains returns true if the GMap contains the given key.
func (m GMap[K, V]) Contains(key K) (isTrue bool) {
	_, isTrue = m[key]
	return isTrue
}

// Delete an entry in the GMap.
func (m GMap[K, V]) Delete(key K) {
	delete(m, key)
}

// Filter a GMap by a predicate, returning a new GMap that contains only values that satisfy the predicate.
func (m GMap[K, V]) Filter(predicate genfuncs.Function[V, bool]) (result GMap[K, V]) {
	result = make(GMap[K, V])
	for k, v := range m {
		if !predicate(v) {
			continue
		}
		result[k] = v
	}
	return result
}

// FilterKeys returns a new GMap that contains only values whose key satisfy the predicate.
func (m GMap[K, V]) FilterKeys(predicate genfuncs.Function[K, bool]) (result GMap[K, V]) {
	result = make(GMap[K, V])
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

// Get returns an entry from the Map. The returned bool indicates if the key is in the Map.
func (m GMap[K, V]) Get(key K) (v V, ok bool) {
	v, ok = m[key]
	return v, ok
}

// GetOrElse returns the value at the given key if it exists or returns the result of defaultValue.
func (m GMap[K, V]) GetOrElse(k K, defaultValue func() V) (value V) {
	ok := false
	value, ok = m[k]
	if ok {
		return value
	}
	value = defaultValue()
	return value
}

func (m GMap[K, V]) Iterator() Iterator[V] {
	return &gmapIterator[K, V]{gmap: m, iterator: m.Keys().Iterator()}
}

// Keys return a GSlice containing the keys of the GMap.
func (m GMap[K, V]) Keys() (keys GSlice[K]) {
	keys = maps.Keys(m)
	return keys
}

// Len is the number of elements in the GMap.
func (m GMap[K, V]) Len() (length int) {
	length = len(m)
	return length
}

// Put a key and value in the Map.
func (m GMap[K, V]) Put(key K, value V) {
	m[key] = value
}

// Values returns a GSlice of all the values in the GMap.
func (m GMap[K, V]) Values() (values GSlice[V]) {
	values = maps.Values(m)
	return values
}

func (g gmapIterator[K, V]) HasNext() bool {
	return g.iterator.HasNext()
}

func (g gmapIterator[K, V]) Next() V {
	if !g.HasNext() {
		panic(genfuncs.NoSuchElement)
	}
	k := g.iterator.Next()
	return g.gmap[k]
}
