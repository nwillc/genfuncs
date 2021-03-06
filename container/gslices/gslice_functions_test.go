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
	"github.com/nwillc/genfuncs/container/maps"
	"github.com/nwillc/genfuncs/container/sequences"
	"strconv"
	"testing"

	"github.com/nwillc/genfuncs"
	"github.com/stretchr/testify/assert"
)

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

func TestGroupBy(t *testing.T) {
	type args struct {
		slice       container.GSlice[int]
		keySelector maps.KeyFor[int, string]
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
				keySelector: func(i int) *genfuncs.Result[string] {
					if i%2 == 0 {
						return genfuncs.NewResult("even")
					}
					return genfuncs.NewResult("odd")
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
				assert.Equal(t, genfuncs.EqualTo, sequences.Compare[int](v, resultsMap[k], genfuncs.Ordered[int]))
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
