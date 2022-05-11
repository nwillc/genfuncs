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

type SyncMap[K any, V any] struct {
	m sync.Map
}

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{}
}

func (s *SyncMap[K, V]) Len() int {
	count := 0
	s.m.Range(func(_ any, _ any) bool { count++; return true })
	return count
}

func (s *SyncMap[K, V]) Values() GSlice[V] {
	var gs GSlice[V]
	s.m.Range(func(_ any, v any) bool {
		vv := reflect.ValueOf(v).Interface().(V)
		gs = append(gs, vv)
		return true
	})
	return gs
}

func (s *SyncMap[K, V]) Delete(key K) {
	s.m.Delete(key)
}

func (s *SyncMap[K, V]) Get(key K) (value V, ok bool) {
	var vv V
	v, ok := s.m.Load(key)
	if ok {
		vv = reflect.ValueOf(v).Interface().(V)
	}
	return vv, ok
}

func (s *SyncMap[K, V]) Put(key K, value V) {
	s.m.Store(key, value)
}

func (s *SyncMap[K, V]) ForEach(f func(key K, value V)) {
	s.m.Range(func(k any, v any) bool {
		kk := reflect.ValueOf(k).Interface().(K)
		vv := reflect.ValueOf(v).Interface().(V)
		f(kk, vv)
		return true
	})
}

func (s *SyncMap[K, V]) Keys() GSlice[K] {
	var gs GSlice[K]
	s.m.Range(func(k any, _ any) bool {
		kk := reflect.ValueOf(k).Interface().(K)
		gs = append(gs, kk)
		return true
	})
	return gs
}
