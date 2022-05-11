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
	"reflect"
	"sync"
)

var _ Map[int, int] = (*SyncMap[int, int])(nil)

// SyncMap is a Map implementation employing sync,Map and is therefore GoRoutine safe,
type SyncMap[K any, V any] struct {
	m sync.Map
}

// NewSyncMap creates a new SyncMap instance.
func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{}
}

// Contains returns true if the Map contains the given key.
func (s *SyncMap[K, V]) Contains(key K) bool {
	_, ok := s.m.Load(key)
	return ok
}

// Len returns the element count. This requires a traversal of the Map.
func (s *SyncMap[K, V]) Len() int {
	count := 0
	s.m.Range(func(_ any, _ any) bool { count++; return true })
	return count
}

// Values returns the values in the Map, reflection is used to cast the sync.Map any values to the Map's type.
func (s *SyncMap[K, V]) Values() GSlice[V] {
	var gs GSlice[V]
	s.m.Range(func(_ any, v any) bool {
		vv := reflect.ValueOf(v).Interface().(V)
		gs = append(gs, vv)
		return true
	})
	return gs
}

// Delete an entry from the Map.
func (s *SyncMap[K, V]) Delete(key K) {
	s.m.Delete(key)
}

// Get the value for the key. Reflection is used to cast the sync.Map any type to the appropriate type. The returned
// ok value will be false if the map is not contained in the Map.
func (s *SyncMap[K, V]) Get(key K) (value V, ok bool) {
	var vv V
	v, ok := s.m.Load(key)
	if ok {
		vv = reflect.ValueOf(v).Interface().(V)
	}
	return vv, ok
}

// Put a key value pair into the Map.
func (s *SyncMap[K, V]) Put(key K, value V) {
	s.m.Store(key, value)
}

// ForEach traverses the Map applying the given function to all entries. Reflection is used to cast sync.Map's any types
// to the appropriate types.
func (s *SyncMap[K, V]) ForEach(f func(key K, value V)) {
	s.m.Range(func(k any, v any) bool {
		kk := reflect.ValueOf(k).Interface().(K)
		vv := reflect.ValueOf(v).Interface().(V)
		f(kk, vv)
		return true
	})
}

// Keys returns the keys in the Map by traversing is and using reflection to cast the sync.Map's any to the appropriate
// types.
func (s *SyncMap[K, V]) Keys() GSlice[K] {
	var gs GSlice[K]
	s.m.Range(func(k any, _ any) bool {
		kk := reflect.ValueOf(k).Interface().(K)
		gs = append(gs, kk)
		return true
	})
	return gs
}
