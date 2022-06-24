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

package iterables_test

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/container/iterables"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type PersonName struct {
	First string
	Last  string
}

func TestFold(t *testing.T) {
	sum := 0
	si := container.GSlice[int]{1, 2, 3}
	sum = iterables.Fold[int, int](si, 10, func(r int, i int) int { return r + i })
	assert.Equal(t, 16, sum)

	mi := container.GMap[int, int]{1: 1, 2: 2, 3: 3}
	sum = iterables.Fold[int, int](mi, 10, func(r int, i int) int { return r + i })
	assert.Equal(t, 16, sum)
}

func TestMap(t *testing.T) {
	v := container.GSlice[int]{1, 2, 3}
	want := container.GSlice[string]{"1", "2", "3"}

	got := iterables.Map[int, string](v, func(i int) string { return fmt.Sprint(i) })

	index := 0
	for got.HasNext() {
		assert.Equal(t, want[index], got.Next())
		index++
	}
	assert.Len(t, want, index)
}

func TestAssociate(t *testing.T) {
	var firstLast genfuncs.MapKeyValueFor[PersonName, string, string] = func(p PersonName) (string, string) { return p.First, p.Last }
	type args struct {
		iterable  container.Iterable[PersonName]
		transform genfuncs.MapKeyValueFor[PersonName, string, string]
	}
	tests := []struct {
		name     string
		args     args
		wantSize int
		contains []string
	}{
		{
			name: "Empty",
			args: args{
				iterable:  container.NewList[PersonName](),
				transform: firstLast,
			},
			wantSize: 0,
		},
		{
			name: "Two Unique",
			args: args{
				iterable: container.GMap[string, PersonName]{
					"fred": {
						First: "fred",
						Last:  "flintstone",
					},
					"barney": {
						First: "barney",
						Last:  "rubble",
					},
				},
				transform: firstLast,
			},
			wantSize: 2,
			contains: []string{"fred", "barney"},
		},
		{
			name: "Duplicate",
			args: args{
				iterable: container.GSlice[PersonName]{
					{
						First: "fred",
						Last:  "flintstone",
					},
					{
						First: "fred",
						Last:  "astaire",
					},
				},
				transform: firstLast,
			},
			wantSize: 1,
			contains: []string{"fred"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fNameMap := iterables.Associate(tt.args.iterable, tt.args.transform)
			assert.Equal(t, tt.wantSize, fNameMap.Len())
			for k := range fNameMap {
				_, ok := fNameMap[k]
				assert.True(t, ok)
			}
		})
	}
}

func TestAssociateWith(t *testing.T) {
	var valueSelector genfuncs.MapValueFor[int, int] = func(i int) int { return i * 2 }
	type args struct {
		slice     container.Iterable[int]
		transform genfuncs.MapValueFor[int, int]
	}
	tests := []struct {
		name     string
		args     args
		wantSize int
	}{
		{
			name: "Empty",
			args: args{
				slice:     container.NewList[int](),
				transform: valueSelector,
			},
			wantSize: 0,
		},
		{
			name: "Three Unique",
			args: args{
				slice:     container.GSlice[int]{1, 2, 3},
				transform: valueSelector,
			},
			wantSize: 3,
		},
		{
			name: "Duplicate",
			args: args{
				slice:     container.GMap[string, int]{"1": 1, "2": 2, "two": 2},
				transform: valueSelector,
			},
			wantSize: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultMap := iterables.AssociateWith(tt.args.slice, tt.args.transform)
			assert.Equal(t, tt.wantSize, resultMap.Len())
			for k := range resultMap {
				_, ok := resultMap[k]
				assert.True(t, ok)
			}
		})
	}
}

func TestFlatMap(t *testing.T) {
	var trans = func(i int) container.Iterable[string] { return container.GSlice[string]{"#", strconv.Itoa(i)} }
	type args struct {
		slice     container.Iterable[int]
		transform func(int) container.Iterable[string]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty",
			args: args{
				slice:     container.GSlice[int]{},
				transform: trans,
			},
			want: []string{},
		},
		{
			name: "List",
			args: args{
				slice:     container.NewList[int](1, 2, 3),
				transform: trans,
			},
			want: []string{"#", "1", "#", "2", "#", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := iterables.FlatMap[int, string](tt.args.slice, tt.args.transform)
			index := 0
			for got.HasNext() {
				assert.Equal(t, got.Next(), tt.want[index])
				index++
			}
		})
	}
}
