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

package gslices_test

import (
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/container/gslices"
	"strconv"
	"testing"

	"github.com/nwillc/genfuncs"
	"github.com/stretchr/testify/assert"
)

type PersonName struct {
	First string
	Last  string
}

func TestAssociate(t *testing.T) {
	var firstLast genfuncs.MapKeyValueFor[PersonName, string, string] = func(p PersonName) (string, string) { return p.First, p.Last }
	type args struct {
		slice     []PersonName
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
				slice:     []PersonName{},
				transform: firstLast,
			},
			wantSize: 0,
		},
		{
			name: "Two Unique",
			args: args{
				slice: []PersonName{
					{
						First: "fred",
						Last:  "flintstone",
					},
					{
						First: "barney",
						Last:  "rubble",
					},
				},
				transform: firstLast,
			},
			wantSize: 2,
			contains: []string{"fred", "baarney"},
		},
		{
			name: "Duplicate",
			args: args{
				slice: []PersonName{
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
			fNameMap := gslices.Associate(tt.args.slice, tt.args.transform)
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
		slice     []int
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
				slice:     []int{},
				transform: valueSelector,
			},
			wantSize: 0,
		},
		{
			name: "Three Unique",
			args: args{
				slice:     []int{1, 2, 3},
				transform: valueSelector,
			},
			wantSize: 3,
		},
		{
			name: "Duplicate",
			args: args{
				slice:     []int{1, 2, 2},
				transform: valueSelector,
			},
			wantSize: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultMap := gslices.AssociateWith(tt.args.slice, tt.args.transform)
			assert.Equal(t, tt.wantSize, resultMap.Len())
			for _, k := range tt.args.slice {
				v, ok := resultMap[k]
				assert.True(t, ok, "did not find key:", k)
				assert.Equal(t, k*2, v)
			}
		})
	}
}

func TestDistinct(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Empty",
			args: args{
				slice: []int{},
			},
			want: []int{},
		},
		{
			name: "No Duplicates",
			args: args{
				slice: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "Duplicates",
			args: args{
				slice: []int{1, 2, 3, 1, 1, 2, 3, 3, 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distinct := gslices.Distinct(tt.args.slice)
			assert.Equal(t, len(tt.want), distinct.Len())
		})
	}
}

func TestFlatMap(t *testing.T) {
	var trans = func(i int) container.GSlice[string] { return []string{"#", strconv.Itoa(i)} }
	type args struct {
		slice     container.GSlice[int]
		transform func(int) container.GSlice[string]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty",
			args: args{
				slice:     []int{},
				transform: trans,
			},
			want: []string{},
		},
		{
			name: "List",
			args: args{
				slice:     []int{1, 2, 3},
				transform: trans,
			},
			want: []string{"#", "1", "#", "2", "#", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gslices.FlatMap(tt.args.slice, tt.args.transform)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestFold(t *testing.T) {
	si := []int{1, 2, 3}
	sum := gslices.Fold(si, 10, func(r int, i int) int { return r + i })
	assert.Equal(t, 16, sum)
}

func TestGroupBy(t *testing.T) {
	type args struct {
		slice       container.GSlice[int]
		keySelector genfuncs.MapKeyFor[int, string]
	}
	tests := []struct {
		name string
		args args
		want map[string]container.GSlice[int]
	}{
		{
			name: "Odds Evens",
			args: args{
				slice: []int{1, 2, 3, 4},
				keySelector: func(i int) string {
					if i%2 == 0 {
						return "even"
					}
					return "odd"
				},
			},
			want: map[string]container.GSlice[int]{"odd": {1, 3}, "even": {2, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultsMap := gslices.GroupBy(tt.args.slice, tt.args.keySelector)
			assert.Equal(t, len(tt.want), resultsMap.Len())
			for k, v := range tt.want {
				assert.True(t, v.All(func(i int) bool {
					return container.GSlice[int](resultsMap[k]).Any(genfuncs.IsEqualOrdered(i))
				}))
			}
		})
	}
}

func TestMap(t *testing.T) {
	var trans = strconv.Itoa
	type args struct {
		slice     container.GSlice[int]
		transform func(int) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty",
			args: args{
				slice:     []int{},
				transform: trans,
			},
			want: []string{},
		},
		{
			name: "List",
			args: args{
				slice:     []int{1, 2, 3},
				transform: trans,
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gslices.Map(tt.args.slice, tt.args.transform)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestToSet(t *testing.T) {
	s := container.GSlice[string]{"a", "b", "c", "b", "a"}
	set := gslices.ToSet(s)
	assert.Equal(t, 3, set.Len())
	for _, l := range s {
		assert.True(t, set.Contains(l))
	}
}
