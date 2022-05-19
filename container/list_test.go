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
	"math/rand"
	"testing"
	"time"
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
	type args struct {
		expect container.GSlice[int]
	}
	tests := []struct {
		name string
		args
	}{
		{
			name: "empty",
			args: args{
				expect: container.GSlice[int]{},
			},
		},
		{
			name: "two",
			args: args{
				expect: container.GSlice[int]{1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := container.NewList[int](tt.args.expect...)
			assert.True(t, tt.args.expect.Equal(l.Values(), genfuncs.Ordered[int]))
		})
	}
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

func TestList_SortBy(t *testing.T) {
	type args struct {
		list  *container.List[int]
		order genfuncs.BiFunction[int, int, bool]
	}
	tests := []struct {
		name string
		args args
		want container.GSlice[int]
	}{
		{
			name: "empty",
			args: args{
				list:  container.NewList[int](),
				order: genfuncs.OrderedLess[int],
			},
			want: container.GSlice[int]{},
		},
		{
			name: "single",
			args: args{
				list:  container.NewList[int](1),
				order: genfuncs.OrderedLess[int],
			},
			want: container.GSlice[int]{1},
		},
		{
			name: "sort ascending",
			args: args{
				list:  container.NewList[int](2, 1, 7, 3),
				order: genfuncs.OrderedLess[int],
			},
			want: container.GSlice[int]{1, 2, 3, 7},
		},
		{
			name: "sort descending",
			args: args{
				list:  container.NewList[int](1, 7, 3, 9),
				order: genfuncs.OrderedGreater[int],
			},
			want: container.GSlice[int]{9, 7, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.list.SortBy(tt.args.order)
			values := tt.args.list.Values()
			assert.True(t, tt.want.Equal(values, genfuncs.Ordered[int]))
		})
	}
}

func TestList_RandomSorts(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	passes := 10

	type args struct {
		count int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "One",
			args: args{
				count: 1,
			},
		},
		{
			name: "Two",
			args: args{
				count: 2,
			},
		},
		{
			name: "Three",
			args: args{
				count: 3,
			},
		},
		{
			name: "Four",
			args: args{
				count: 4,
			},
		},
		{
			name: "Medium",
			args: args{
				count: 16,
			},
		},
		{
			name: "Large",
			args: args{
				count: 64,
			},
		},
		{
			name: "Larger",
			args: args{
				count: 4096,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for pass := passes; pass > 0; pass-- {
				count := tt.args.count + random.Intn(tt.args.count)
				numbers := container.NewList[int]()
				for i := 0; i < count; i++ {
					numbers.Add(random.Int() % 10)
				}
				numbers.SortBy(genfuncs.OrderedLess[int])
				assert.True(t, numbers.IsSorted(genfuncs.OrderedLess[int]))
				for i := 0; i < count; i++ {
					numbers.Add(random.Int())
				}
				numbers = numbers.SortBy(genfuncs.OrderedGreater[int])
				assert.True(t, numbers.IsSorted(genfuncs.OrderedGreater[int]))
			}
		})
	}
}
