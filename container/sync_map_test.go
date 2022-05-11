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

var testSyncMap container.Map[int, string]

func init() {
	testSyncMap = container.NewSyncMap[int, string]()
	testSyncMap.Put(1, "1")
	testSyncMap.Put(2, "2")
	testSyncMap.Put(3, "3")
}

func TestSyncMap_ForEach(t *testing.T) {
	all := container.GSlice[string]{}
	testSyncMap.ForEach(func(k int, v string) {
		all = append(all, fmt.Sprintf("%d.%s", k, v))
	})
	all = all.SortBy(genfuncs.LessOrdered[string])
	assert.True(t, container.GSlice[string]{"1.1", "2.2", "3.3"}.Equal(all, genfuncs.Order[string]))
}

func TestSyncMap_Len(t *testing.T) {
	assert.Equal(t, 3, testSyncMap.Len())
}

func TestSyncMap_Values(t *testing.T) {
	values := testSyncMap.Values().SortBy(genfuncs.LessOrdered[string])
	assert.True(t, container.GSlice[string]{"1", "2", "3"}.Equal(values, genfuncs.Order[string]))
}

func TestSyncMap_Keys(t *testing.T) {
	values := testSyncMap.Keys().SortBy(genfuncs.LessOrdered[int])
	assert.True(t, container.GSlice[int]{1, 2, 3}.Equal(values, genfuncs.Order[int]))
}

func TestSyncMap_Get(t *testing.T) {
	var v string
	var ok bool
	v, ok = testSyncMap.Get(4)
	assert.Equal(t, "", v)
	assert.False(t, ok)
	v, ok = testSyncMap.Get(2)
	assert.Equal(t, "2", v)
	assert.True(t, ok)
}

func TestSyncMap_Delete(t *testing.T) {
	var v string
	var ok bool
	testSyncMap.Put(20, "20")
	v, ok = testSyncMap.Get(20)
	assert.Equal(t, "20", v)
	assert.True(t, ok)
	testSyncMap.Delete(20)
	v, ok = testSyncMap.Get(20)
	assert.Equal(t, "", v)
	assert.False(t, ok)
}

func TestSyncMap_Contains(t *testing.T) {
	var ok bool
	ok = testSyncMap.Contains(2)
	assert.True(t, ok)
	ok = testSyncMap.Contains(5)
	assert.False(t, ok)
}
