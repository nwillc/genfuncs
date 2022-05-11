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

package gmaps_test

import (
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/container/gmaps"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatMerge(t *testing.T) {
	var m1 container.Map[string, container.GSlice[string]] = container.GMap[string, container.GSlice[string]]{}
	var m2 container.Map[string, container.GSlice[string]] = container.GMap[string, container.GSlice[string]]{}

	m1.Put("a", container.GSlice[string]{"1"})
	m2.Put("a", container.GSlice[string]{"2"})
	m2.Put("b", container.GSlice[string]{"1"})

	m3 := gmaps.MapMerge(m1, m2)
	assert.Equal(t, 2, m3.Len())
	v, ok := m3.Get("a")
	assert.True(t, ok)
	assert.Equal(t, 2, v.Len())
	v, ok = m3.Get("b")
	assert.True(t, ok)
	assert.Equal(t, 1, v.Len())
}

func TestMap(t *testing.T) {
	var m container.Map[int, int] = container.GMap[int, int]{1: 2, 3: 4, 5: 6}

	sums := gmaps.Map(m, func(k int, v int) int { return k + v })
	want := container.GSlice[int]{3, 7, 11}
	assert.ElementsMatch(t, want, sums)
}
