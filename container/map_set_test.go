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
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/container/sequences"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	set := container.NewMapSet[string]()
	assert.Equal(t, 0, set.Len())
	assert.False(t, set.Contains("foo"))
	set.Add("foo")
	assert.True(t, set.Contains("foo"))
	s := set.Values()
	assert.Equal(t, "foo", s[0])
}

func TestRemove(t *testing.T) {
	set := container.NewMapSet("a")
	assert.Equal(t, true, set.Contains("a"))
	set.Remove("a")
	assert.Equal(t, false, set.Contains("a"))
}

func TestMapSet_Iterator(t *testing.T) {
	m := container.NewMapSet[int]()
	s := sequences.NewSequence(1, 2, 3, 5)

	sequences.Collect[int](s, m)
	mapIterator := m.Iterator()
	count := 0
	for ; mapIterator.HasNext(); count++ {
		v := mapIterator.Next()
		assert.True(t, sequences.Any[int](s, genfuncs.OrderedEqualTo(v)))
	}
	assert.Equal(t, count, m.Len())
}
