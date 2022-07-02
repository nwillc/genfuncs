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
	"sync"
)

// SyncMap implements Map.
var _ Map[int, int] = (*SyncMap[int, int])(nil)

// SyncMap is a Map implementation employing sync.Map and is therefore GoRoutine safe.
type SyncMap[K any, V any] struct {
	m sync.Map
}

// NewSyncMap creates a new SyncMap instance.
func NewSyncMap[K any, V any]() (syncMap *SyncMap[K, V]) {
	syncMap = &SyncMap[K, V]{}
	return syncMap
}

// Contains returns true if the Map contains the given key.
func (s *SyncMap[K, V]) Contains(key K) (contains bool) {
	_, contains = s.m.Load(key)
	return contains
}

// Delete an entry from the Map.
func (s *SyncMap[K, V]) Delete(key K) {
	s.m.Delete(key)
}

// ForEach traverses the Map applying the given function to all entries. The sync.Map's any types are cast
// to the appropriate types.
func (s *SyncMap[K, V]) ForEach(f func(key K, value V)) {
	s.m.Range(func(k any, v any) bool {
		f(k.(K), v.(V))
		return true
	})
}

// Get the value for the key. The sync.Map any type to cast to the appropriate type. The returned
// ok value will be false if the map is not contained in the Map.
func (s *SyncMap[K, V]) Get(key K) (value V, ok bool) {
	var v any
	v, ok = s.m.Load(key)
	if ok {
		value = v.(V)
	}
	return value, ok
}

// GetAndDelete returns the value from the SyncMap corresponding to the key, returning it, and deletes it.
func (s *SyncMap[K, V]) GetAndDelete(key K) (value V, ok bool) {
	var v any
	v, ok = s.m.LoadAndDelete(key)
	if ok {
		value = v.(V)
	}
	return value, ok
}

// GetOrPut returns the existing value for the key if present. Otherwise, it puts and returns the given value.
// The ok result is true if the value was present, false if put.
func (s *SyncMap[K, V]) GetOrPut(key K, value V) (actual V, ok bool) {
	var v any
	v, ok = s.m.LoadOrStore(key, value)
	if ok {
		actual = v.(V)
	} else {
		actual = value
	}
	return actual, ok
}

func (s *SyncMap[K, V]) Iterator() Iterator[V] {
	return s.Values().Iterator()
}

// Keys returns the keys in the Map by traversing it and casting the sync.Map's any to the appropriate
// type.
func (s *SyncMap[K, V]) Keys() (keys GSlice[K]) {
	s.m.Range(func(k any, _ any) bool {
		keys = append(keys, k.(K))
		return true
	})
	return keys
}

// Len returns the element count. This requires a traversal of the Map.
func (s *SyncMap[K, V]) Len() (length int) {
	length = 0
	s.m.Range(func(_ any, _ any) bool { length++; return true })
	return length
}

// Put a key value pair into the Map.
func (s *SyncMap[K, V]) Put(key K, value V) {
	s.m.Store(key, value)
}

// Values returns the values in the Map, The sync.Map any values is cast to the Map's type.
func (s *SyncMap[K, V]) Values() (values GSlice[V]) {
	s.m.Range(func(_ any, v any) bool {
		values = append(values, v.(V))
		return true
	})
	return values
}
