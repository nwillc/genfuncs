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
	"github.com/nwillc/genfuncs/container"
	"testing"
	"time"

	"github.com/nwillc/genfuncs"
	"github.com/stretchr/testify/assert"
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
		predicate genfuncs.Function[string, bool]
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
			got := container.GSlice[string](tt.args.slice).All(tt.args.predicate)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestAny(t *testing.T) {
	type args struct {
		slice     container.GSlice[string]
		predicate genfuncs.Function[string, bool]
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

func TestContains(t *testing.T) {
	type args struct {
		slice   container.GSlice[string]
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
			got := tt.args.slice.Any(genfuncs.IsEqualComparable(tt.args.element))
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		slice     container.GSlice[int]
		predicate genfuncs.Function[int, bool]
	}
	tests := []struct {
		name string
		args args
		want container.GSlice[int]
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
			assert.True(t, result.Compare(tt.want, genfuncs.EqualComparable[int]))
		})
	}
}

func TestFind(t *testing.T) {
	type args struct {
		slice     container.GSlice[float32]
		predicate genfuncs.Function[float32, bool]
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
		slice     container.GSlice[float32]
		predicate genfuncs.Function[float32, bool]
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

func TestJoinToString(t *testing.T) {
	personStringer := genfuncs.StringerToString[PersonName]()
	type args struct {
		slice     container.GSlice[PersonName]
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

func TestSortBy(t *testing.T) {
	timeComparator := genfuncs.TransformArgs(
		func(t time.Time) int64 { return t.Unix() },
		genfuncs.I64NumericOrder,
	)
	type args struct {
		slice      container.GSlice[time.Time]
		comparator genfuncs.BiFunction[time.Time, time.Time, bool]
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
				comparator: genfuncs.Reverse(timeComparator),
			},
			want: []time.Time{now.Add(time.Second), now, now.Add(-time.Second)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorted := tt.args.slice.SortBy(tt.args.comparator)
			assert.Equal(t, len(tt.want), len(sorted))
			assert.True(t, sorted.Compare(tt.want, genfuncs.EqualComparable[time.Time]))
		})
	}
}

func TestRandom(t *testing.T) {
	var s container.GSlice[int] = []int{1, 2, 3}

	for c := 0; c < 2*len(s); c++ {
		i := s.Random()
		p := genfuncs.IsEqualComparable(i)
		assert.True(t, s.Any(p))
	}
}

func TestCompare(t *testing.T) {
	type args struct {
		a container.GSlice[string]
		b container.GSlice[string]
	}
	tests := []struct {
		name     string
		args     args
		want     bool
		wanPanic bool
	}{
		{
			name: "Mismatched length",
			args: args{
				a: []string{"a"},
				b: []string{},
			},
			want: false,
		},
		{
			name: "Matched",
			args: args{
				a: []string{"a", "b"},
				b: []string{"a", "b"},
			},
			want: true,
		},
		{
			name: "Mismatched",
			args: args{
				a: []string{"a", "b"},
				b: []string{"a", "c"},
			},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.args.a.Compare(test.args.b, genfuncs.EqualComparable[string])
		})
	}
}

func TestSlice_ForEach(t *testing.T) {
	tests := []struct {
		name string
		s    container.GSlice[int]
		want int
	}{
		{
			name: "Nil",
			want: 0,
		},
		{
			name: "Empty",
			s:    container.GSlice[int]{},
			want: 0,
		},
		{
			name: "Two",
			s:    container.GSlice[int]{1, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := 0
			tt.s.ForEach(func(i int) {
				count++
			})
			assert.Equal(t, tt.want, count)
		})
	}
}

func TestSlice_ForEachI(t *testing.T) {
	tests := []struct {
		name string
		s    container.GSlice[int]
		want int
	}{
		{
			name: "Nil",
			want: 0,
		},
		{
			name: "Empty",
			s:    container.GSlice[int]{},
			want: 0,
		},
		{
			name: "Two",
			s:    container.GSlice[int]{0, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := 0
			tt.s.ForEachI(func(i, v int) {
				assert.Equal(t, count, i)
				count++
			})
			assert.Equal(t, tt.want, count)
		})
	}
}