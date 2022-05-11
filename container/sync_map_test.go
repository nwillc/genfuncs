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

package container_test

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSyncMap_ForEach(t *testing.T) {
	m := container.NewSyncMap[int, string]()
	m.Put(1, "1")
	m.Put(2, "2")
	m.Put(3, "3")

	all := container.GSlice[string]{}
	m.ForEach(func(k int, v string) {
		all = append(all, fmt.Sprintf("%d.%s", k, v))
	})
	all = all.SortBy(genfuncs.LessOrdered[string])
	assert.True(t, container.GSlice[string]{"1.1", "2.2", "3.3"}.Equal(all, genfuncs.Order[string]))
}

func TestSyncMap_Len(t *testing.T) {
	m := container.NewSyncMap[int, string]()
	m.Put(1, "1")
	m.Put(2, "2")
	m.Put(3, "3")
	assert.Equal(t, 3, m.Len())
}

func TestSyncMap_Values(t *testing.T) {
	m := container.NewSyncMap[int, string]()
	m.Put(1, "1")
	m.Put(2, "2")
	m.Put(3, "3")
	values := m.Values().SortBy(genfuncs.LessOrdered[string])
	assert.True(t, container.GSlice[string]{"1", "2", "3"}.Equal(values, genfuncs.Order[string]))
}

func TestSyncMap_Keys(t *testing.T) {
	m := container.NewSyncMap[int, string]()
	m.Put(1, "1")
	m.Put(2, "2")
	m.Put(3, "3")
	values := m.Keys().SortBy(genfuncs.LessOrdered[int])
	assert.True(t, container.GSlice[int]{1, 2, 3}.Equal(values, genfuncs.Order[int]))
}

func TestSyncMap_Get(t *testing.T) {
	m := container.NewSyncMap[int, string]()
	m.Put(1, "1")
	m.Put(2, "2")
	m.Put(3, "3")

	var v string
	var ok bool
	v, ok = m.Get(4)
	assert.Equal(t, "", v)
	assert.False(t, ok)
	v, ok = m.Get(2)
	assert.Equal(t, "2", v)
	assert.True(t, ok)
}

func TestSyncMap_Delete(t *testing.T) {
	m := container.NewSyncMap[int, string]()
	m.Put(1, "1")
	m.Put(2, "2")
	m.Put(3, "3")

	var v string
	var ok bool
	v, ok = m.Get(2)
	assert.Equal(t, "2", v)
	assert.True(t, ok)
	m.Delete(2)
	v, ok = m.Get(2)
	assert.Equal(t, "", v)
	assert.False(t, ok)
}

func TestSyncMap_Contains(t *testing.T) {
	m := container.NewSyncMap[int, string]()
	m.Put(1, "1")
	m.Put(2, "2")
	m.Put(3, "3")

	var ok bool
	ok = m.Contains(2)
	assert.True(t, ok)
	ok = m.Contains(5)
	assert.False(t, ok)
}
