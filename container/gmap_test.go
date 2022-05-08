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
	"strconv"
	"testing"
)

func TestGMapContains(t *testing.T) {
	var m container.GMap[string, bool] = map[string]bool{"a": true}
	assert.Equal(t, true, m.Contains("a"))
	delete(m, "a")
	assert.Equal(t, false, m.Contains("a"))
}

func TestKeys(t *testing.T) {
	type args struct {
		m container.GMap[string, string]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty",
			args: args{
				m: nil,
			},
			want: nil,
		},
		{
			name: "One",
			args: args{
				m: map[string]string{"one": "one"},
			},
			want: []string{"one"},
		},
		{
			name: "Two",
			args: args{
				m: map[string]string{"one": "one", "two": "two"},
			},
			want: []string{"one", "two"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys := tt.args.m.Keys()
			assert.Equal(t, len(tt.want), len(keys))
			for _, k := range keys {
				_, ok := tt.args.m[k]
				assert.True(t, ok)
			}
		})
	}
}

func TestValues(t *testing.T) {
	type args struct {
		m container.GMap[string, int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Empty",
			args: args{
				m: nil,
			},
			want: nil,
		},
		{
			name: "One",
			args: args{
				m: map[string]int{"1": 1},
			},
			want: []int{1},
		},
		{
			name: "Two",
			args: args{
				m: map[string]int{"1": 1, "5": 5},
			},
			want: []int{1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := tt.args.m.Values()
			assert.Equal(t, len(tt.want), len(values))
			for _, v := range values {
				k := strconv.Itoa(v)
				_, ok := tt.args.m[k]
				assert.True(t, ok)
			}
		})
	}
}

func TestGMap_Filter(t *testing.T) {
	m := container.GMap[string, string]{"a": "A", "b": "B", "c": "C"}
	type args struct {
		greaterThan string
	}
	tests := []struct {
		name string
		args args
		want container.GSlice[string]
	}{
		{
			name: "none",
			args: args{
				greaterThan: "Z",
			},
			want: container.GSlice[string]{},
		},
		{
			name: "greater than a",
			args: args{
				greaterThan: "A",
			},
			want: container.GSlice[string]{"B", "C"},
		},
		{
			name: "greater than b",
			args: args{
				greaterThan: "B",
			},
			want: container.GSlice[string]{"C"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filtered := m.Filter(genfuncs.IsGreaterOrdered(tt.args.greaterThan)).Values().SortBy(genfuncs.LessOrdered[string])
			assert.True(t, filtered.Compare(tt.want, genfuncs.Order[string]) == 0)
		})
	}
}

func TestGMap_FilterKeys(t *testing.T) {
	m := container.GMap[string, string]{"a": "A", "b": "B", "c": "C"}
	type args struct {
		greaterThan string
	}
	tests := []struct {
		name string
		args args
		want container.GSlice[string]
	}{
		{
			name: "none",
			args: args{
				greaterThan: "z",
			},
			want: container.GSlice[string]{},
		},
		{
			name: "greater than a",
			args: args{
				greaterThan: "a",
			},
			want: container.GSlice[string]{"B", "C"},
		},
		{
			name: "greater than b",
			args: args{
				greaterThan: "b",
			},
			want: container.GSlice[string]{"C"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filtered := m.FilterKeys(genfuncs.IsGreaterOrdered(tt.args.greaterThan)).Values().SortBy(genfuncs.LessOrdered[string])
			assert.True(t, filtered.Compare(tt.want, genfuncs.Order[string]) == 0)
		})
	}
}

func TestGMap_All(t *testing.T) {
	m := container.GMap[string, string]{"a": "a", "b": "b", "c": "c"}
	type args struct {
		predicate genfuncs.Function[string, bool]
		m         container.GMap[string, string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				predicate: genfuncs.IsNotBlank,
				m:         make(container.GMap[string, string]),
			},
			want: true,
		},
		{
			name: "all",
			args: args{
				predicate: genfuncs.IsNotBlank,
				m:         m,
			},
			want: true,
		},
		{
			name: "some",
			args: args{
				predicate: genfuncs.IsEqualOrdered("a"),
				m:         m,
			},
			want: false,
		},
		{
			name: "none",
			args: args{
				predicate: genfuncs.IsEqualOrdered("z"),
				m:         m,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.args.m.All(tt.args.predicate))
		})
	}
}

func TestGMap_Any(t *testing.T) {
	m := container.GMap[string, string]{"a": "a", "b": "b", "c": "c"}
	type args struct {
		predicate genfuncs.Function[string, bool]
		m         container.GMap[string, string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				predicate: genfuncs.IsNotBlank,
				m:         make(container.GMap[string, string]),
			},
			want: false,
		},
		{
			name: "all",
			args: args{
				predicate: genfuncs.IsNotBlank,
				m:         m,
			},
			want: true,
		},
		{
			name: "some",
			args: args{
				predicate: genfuncs.IsEqualOrdered("a"),
				m:         m,
			},
			want: true,
		},
		{
			name: "none",
			args: args{
				predicate: genfuncs.IsEqualOrdered("z"),
				m:         m,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.args.m.Any(tt.args.predicate))
		})
	}
}

func TestGMap_ForEach(t *testing.T) {

	tests := []struct {
		name string
		m    container.GMap[int, int]
		want int
	}{
		{
			name: "empty",
			m:    container.GMap[int, int]{},
			want: 0,
		},
		{
			name: "with values",
			m:    container.GMap[int, int]{1: 1, 2: 2},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum := 0
			action := func(k, v int) { sum = sum + k + v }
			tt.m.ForEach(action)
			assert.Equal(t, tt.want, sum)
		})
	}
}

func TestGMap_GetOrElse(t *testing.T) {
	m := container.GMap[string, int]{"0": 0, "1": 1, "2": 2}
	def := func() int { return -1 }
	assert.Equal(t, 0, m.GetOrElse("0", def))
	assert.Equal(t, -1, m.GetOrElse("1000", def))
}
