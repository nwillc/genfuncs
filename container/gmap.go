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

import "golang.org/x/exp/maps"

// GMap is a generic type corresponding to a standard Go map.
type GMap[K comparable, V any] map[K]V

// Contains returns true if the GMap contains the given key.
func (m GMap[K, V]) Contains(key K) bool {
	_, ok := m[key]
	return ok
}

// Keys return a GSlice containing the keys of the GMap.
func (m GMap[K, V]) Keys() GSlice[K] {
	return maps.Keys(m)
}

// Values returns a GSlice of all the values in the GMap.
func (m GMap[K, V]) Values() GSlice[V] {
	return maps.Values(m)
}
