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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewList(t *testing.T) {
	l := container.NewList[int]()
	assert.NotNil(t, l)
	assert.Equal(t, 0, l.Len())
	assert.Nil(t, l.PeekLeft())
	assert.Nil(t, l.PeekRight())
}

func TestList_AddRight(t *testing.T) {
	l := container.NewList[string]("1")
	assert.Equal(t, 1, l.Len())
	assert.Equal(t, "1", l.PeekRight().Value)
	assert.Equal(t, "1", l.PeekLeft().Value)
	l.AddRight("2")
	assert.Equal(t, 2, l.Len())
	assert.Equal(t, "1", l.PeekLeft().Value)
	assert.Equal(t, "2", l.PeekRight().Value)

}

func TestList_AddLeft(t *testing.T) {
	l := container.NewList[string]("1")
	assert.Equal(t, 1, l.Len())
	assert.Equal(t, "1", l.PeekRight().Value)
	assert.Equal(t, "1", l.PeekLeft().Value)
	l.AddLeft("2")
	assert.Equal(t, 2, l.Len())
	assert.Equal(t, "1", l.PeekRight().Value)
	assert.Equal(t, "2", l.PeekLeft().Value)
}

func TestList_Remove(t *testing.T) {
	l := container.NewList[int](1, 2)
	e := l.PeekLeft()
	assert.Equal(t, 1, e.Value)
	v := l.Remove(e)
	assert.Equal(t, 1, v)
	e = l.PeekLeft()
	assert.Equal(t, 2, e.Value)
}

func TestList_Values(t *testing.T) {
	expect := container.GSlice[int]{1, 2}
	l := container.NewList[int](expect...)
	assert.True(t, expect.Equal(l.Values(), genfuncs.Order[int]))
}

func TestElement_NextPrev(t *testing.T) {
	l := container.NewList[int](1, 2)
	left := l.PeekLeft()
	right := l.PeekRight()

	assert.Equal(t, left.Next(), right)
	assert.Nil(t, left.Prev())

	assert.Equal(t, right.Prev(), left)
	assert.Nil(t, right.Next())
}
