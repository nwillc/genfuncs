/*
 *  Copyright (c) 2021,  nwillc@gmail.com
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

package genfuncs_test

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

var _ fmt.Stringer = (*PersonName)(nil)

type PersonName struct {
	First string
	Last  string
}

func (p PersonName) String() string {
	return p.First + " " + p.Last
}

func TestAll(t *testing.T) {
	type args struct {
		slice     []string
		predicate genfuncs.Predicate[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty",
			args: args{
				slice:     []string{},
				predicate: func(s string) bool { return s == "a" },
			},
			want: true,
		},
		{
			name: "Some Not All",
			args: args{
				slice:     []string{"b", "c"},
				predicate: func(s string) bool { return s == "b" },
			},
			want: false,
		},
		{
			name: "All",
			args: args{
				slice:     []string{"b", "a", "c"},
				predicate: func(s string) bool { return len(s) == 1 },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := genfuncs.Slice[string](tt.args.slice).All(tt.args.predicate)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestAny(t *testing.T) {
	type args struct {
		slice     genfuncs.Slice[string]
		predicate genfuncs.Predicate[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty",
			args: args{
				slice:     []string{},
				predicate: func(s string) bool { return s == "a" },
			},
			want: false,
		},
		{
			name: "Not Found",
			args: args{
				slice:     []string{"b", "c"},
				predicate: func(s string) bool { return s == "a" },
			},
			want: false,
		},
		{
			name: "Found",
			args: args{
				slice:     []string{"b", "a", "c"},
				predicate: func(s string) bool { return s == "a" },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.slice.Any(tt.args.predicate)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestAssociate(t *testing.T) {
	var firstLast genfuncs.KeyValueFor[PersonName, string, string] = func(p PersonName) (string, string) { return p.First, p.Last }
	type args struct {
		slice     []PersonName
		transform genfuncs.KeyValueFor[PersonName, string, string]
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
			fNameMap := genfuncs.Associate(tt.args.slice, tt.args.transform)
			assert.Equal(t, tt.wantSize, len(fNameMap))
			for k, _ := range fNameMap {
				_, ok := fNameMap[k]
				assert.True(t, ok)
			}
		})
	}
}

func TestAssociateWith(t *testing.T) {
	var valueSelector genfuncs.ValueFor[int, int] = func(i int) int { return i * 2 }
	type args struct {
		slice     []int
		transform genfuncs.ValueFor[int, int]
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
			resultMap := genfuncs.AssociateWith(tt.args.slice, tt.args.transform)
			assert.Equal(t, tt.wantSize, len(resultMap))
			for _, k := range tt.args.slice {
				v, ok := resultMap[k]
				assert.True(t, ok, "did not find key:", k)
				assert.Equal(t, k*2, v)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		slice   genfuncs.Slice[string]
		element string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty",
			args: args{
				slice:   []string{},
				element: "foo",
			},
			want: false,
		},
		{
			name: "Not Found",
			args: args{
				slice:   []string{"b", "c"},
				element: "a",
			},
			want: false,
		},
		{
			name: "Found",
			args: args{
				slice:   []string{"b", "a", "c"},
				element: "a",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.slice.Contains(tt.args.element, genfuncs.OrderedComparator[string]())
			assert.Equal(t, got, tt.want)
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
			distinct := genfuncs.Distinct(tt.args.slice)
			assert.Equal(t, len(tt.want), len(distinct))
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		slice     genfuncs.Slice[int]
		predicate genfuncs.Predicate[int]
	}
	tests := []struct {
		name string
		args args
		want genfuncs.Slice[int]
	}{
		{
			name: "Empty",
			args: args{
				slice:     []int{},
				predicate: func(i int) bool { return true },
			},
			want: nil,
		},
		{
			name: "Evens",
			args: args{
				slice:     []int{1, 2, 3, 4},
				predicate: func(i int) bool { return i%2 == 0 },
			},
			want: []int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.slice.Filter(tt.args.predicate)
			assert.Equal(t, len(tt.want), len(result))
			for _, v := range result {
				assert.True(t, result.Any(func(i int) bool { return i == v }))
			}
		})
	}
}

func TestFind(t *testing.T) {
	type args struct {
		slice     genfuncs.Slice[float32]
		predicate genfuncs.Predicate[float32]
	}
	tests := []struct {
		name      string
		args      args
		want      float32
		wantFound bool
	}{
		{
			name: "Empty",
			args: args{
				slice:     []float32{},
				predicate: func(f float32) bool { return false },
			},
			wantFound: false,
		},
		{
			name: "Not Found",
			args: args{
				slice:     []float32{1.0, 2.0, 3.0},
				predicate: func(f float32) bool { return false },
			},
			wantFound: false,
		},
		{
			name: "Found",
			args: args{
				slice:     []float32{1.0, 2.0, 3.0},
				predicate: func(f float32) bool { return f > 1.0 },
			},
			want:      2.0,
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.args.slice.Find(tt.args.predicate)
			if !tt.wantFound {
				assert.False(t, ok)
				return
			}
			assert.True(t, ok)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFindLast(t *testing.T) {
	type args struct {
		slice     genfuncs.Slice[float32]
		predicate genfuncs.Predicate[float32]
	}
	tests := []struct {
		name      string
		args      args
		want      float32
		wantFound bool
	}{
		{
			name: "Empty",
			args: args{
				slice:     []float32{},
				predicate: func(f float32) bool { return false },
			},
			wantFound: false,
		},
		{
			name: "Not Found",
			args: args{
				slice:     []float32{1.0, 2.0, 3.0},
				predicate: func(f float32) bool { return f < 0.0 },
			},
			wantFound: false,
		},
		{
			name: "Found",
			args: args{
				slice:     []float32{1.0, 2.0, 3.0},
				predicate: func(f float32) bool { return f > 1.0 },
			},
			want:      3.0,
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.args.slice.FindLast(tt.args.predicate)
			if !tt.wantFound {
				assert.False(t, ok)
				return
			}
			assert.True(t, ok)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFlatMap(t *testing.T) {
	var trans genfuncs.Function[int, []string] = func(i int) []string { return []string{"#", strconv.Itoa(i)} }
	type args struct {
		slice     genfuncs.Slice[int]
		transform genfuncs.Function[int, []string]
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
			got := genfuncs.FlatMap(tt.args.slice, tt.args.transform)
			assert.Equal(t, len(tt.want), len(got))
			for _, s := range tt.want {
				assert.True(t, genfuncs.Slice[string](got).Contains(s, genfuncs.OrderedComparator[string]()))
			}
		})
	}
}

func TestFold(t *testing.T) {
	si := []int{1, 2, 3}
	sum := genfuncs.Fold(si, 10, func(r int, i int) int { return r + i })
	assert.Equal(t, 16, sum)
}

func TestGroupBy(t *testing.T) {
	type args struct {
		slice       genfuncs.Slice[int]
		keySelector genfuncs.KeyFor[int, string]
	}
	tests := []struct {
		name string
		args args
		want map[string]genfuncs.Slice[int]
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
			want: map[string]genfuncs.Slice[int]{"odd": {1, 3}, "even": {2, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultsMap := genfuncs.GroupBy(tt.args.slice, tt.args.keySelector)
			assert.Equal(t, len(tt.want), len(resultsMap))
			for k, v := range tt.want {
				assert.True(t, v.All(func(i int) bool {
					return genfuncs.Slice[int](resultsMap[k]).Contains(i, genfuncs.OrderedComparator[int]())
				}))
			}
		})
	}
}

func TestJoinToString(t *testing.T) {
	personStringer := genfuncs.StringerStringer[PersonName]()
	type args struct {
		slice     genfuncs.Slice[PersonName]
		separator string
		prefix    string
		postfix   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty",
			args: args{
				slice:     []PersonName{},
				separator: "",
			},
			want: "",
		},
		{
			name: "One",
			args: args{
				slice: []PersonName{
					{
						First: "fred",
						Last:  "flintstone",
					},
				},
				separator: ", ",
			},
			want: "fred flintstone",
		},
		{
			name: "Two",
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
				separator: ", ",
			},
			want: "fred flintstone, barney rubble",
		},
		{
			name: "Two With Prefix Postfix",
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
				separator: ", ",
				prefix:    "[",
				postfix:   "]",
			},
			want: "[fred flintstone, barney rubble]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.slice.JoinToString(personStringer, tt.args.separator, tt.args.prefix, tt.args.postfix)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMap(t *testing.T) {
	var trans genfuncs.Function[int, string] = strconv.Itoa
	type args struct {
		slice     genfuncs.Slice[int]
		transform genfuncs.Function[int, string]
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
			got := genfuncs.Map(tt.args.slice, tt.args.transform)
			assert.Equal(t, len(tt.want), len(got))
			for _, s := range tt.want {
				assert.True(t, genfuncs.Slice[string](got).Contains(s, genfuncs.OrderedComparator[string]()))
			}
		})
	}
}

func TestSortBy(t *testing.T) {
	timeComparator := genfuncs.FunctionComparator[time.Time, int64](
		func(t time.Time) int64 { return t.Unix() },
		genfuncs.OrderedComparator[int64](),
	)
	type args struct {
		slice      genfuncs.Slice[time.Time]
		comparator genfuncs.Comparator[time.Time]
	}
	now := time.Now()
	tests := []struct {
		name string
		args args
		want []time.Time
	}{
		{
			name: "Empty",
			args: args{
				slice:      []time.Time{},
				comparator: timeComparator,
			},
			want: []time.Time{},
		},
		{
			name: "Min Max",
			args: args{
				slice:      []time.Time{now.Add(time.Second), now, now.Add(-time.Second)},
				comparator: timeComparator,
			},
			want: []time.Time{now.Add(-time.Second), now, now.Add(time.Second)},
		},
		{
			name: "Max Min",
			args: args{
				slice:      []time.Time{now.Add(time.Second), now.Add(-time.Second), now},
				comparator: genfuncs.ReverseComparator(timeComparator),
			},
			want: []time.Time{now.Add(time.Second), now, now.Add(-time.Second)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorted := tt.args.slice.SortBy(tt.args.comparator)
			assert.Equal(t, len(tt.want), len(sorted))
			for i, tm := range tt.want {
				assert.Equal(t, tm, sorted[i])
			}
		})
	}
}
