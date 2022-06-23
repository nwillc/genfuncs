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
	"math/rand"
	"testing"
	"time"
)

var (
	_        fmt.Stringer = (*PersonName)(nil)
	letters               = []string{"t", "e", "s", "t"}
	alphabet              = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "t", "u", "v", "w", "x", "y", "z"}
)

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
			got := tt.args.slice.Any(genfuncs.OrderedEqualTo(tt.args.element))
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
			assert.True(t, result.Equal(tt.want, genfuncs.Ordered[int]))
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
	type args struct {
		slice      container.GSlice[string]
		comparator genfuncs.BiFunction[string, string, bool]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Empty",
			args: args{
				slice:      []string{},
				comparator: genfuncs.OrderedLess[string],
			},
		},
		{
			name: "Single",
			args: args{
				slice:      []string{"a"},
				comparator: genfuncs.OrderedLess[string],
			},
		},
		{
			name: "Double",
			args: args{
				slice:      []string{"a", "b"},
				comparator: genfuncs.OrderedLess[string],
			},
		},
		{
			name: "Double Reverse",
			args: args{
				slice:      []string{"a", "b"},
				comparator: genfuncs.OrderedGreater[string],
			},
		},
		{
			name: "Min Max",
			args: args{
				slice:      letters,
				comparator: genfuncs.OrderedLess[string],
			},
		},
		{
			name: "Max Min",
			args: args{
				slice:      letters,
				comparator: genfuncs.OrderedGreater[string],
			},
		},
		{
			name: "More than 12",
			args: args{
				slice:      alphabet,
				comparator: genfuncs.OrderedLess[string],
			},
		},
		{
			name: "Test duplicates",
			args: args{
				slice:      []string{"d", "z", "d", "a", "d", "a", "d", "a", "d", "a", "a"},
				comparator: genfuncs.OrderedLess[string],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.slice.SortBy(tt.args.comparator)
			assert.True(t, tt.args.slice.IsSorted(tt.args.comparator))
		})
	}
}

func TestGSliceRandomSorts(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	passes := 20

	type args struct {
		count int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Tiny",
			args: args{
				count: 1,
			},
		},
		{
			name: "Small",
			args: args{
				count: 8,
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
			for pass := passes; pass >= 0; pass-- {
				count := tt.args.count + random.Intn(tt.args.count)
				numbers := make(container.GSlice[int], count)
				for i := 0; i < tt.args.count; i++ {
					numbers[i] = random.Int()
				}
				numbers = numbers.SortBy(genfuncs.OrderedLess[int])
				for i := 0; i < count-1; i++ {
					assert.LessOrEqual(t, numbers[i], numbers[i+1])
				}
				for i := 0; i < count; i++ {
					numbers[i] = random.Int()
				}
				numbers = numbers.SortBy(genfuncs.OrderedGreater[int])
				for i := 0; i < count-1; i++ {
					assert.GreaterOrEqual(t, numbers[i], numbers[i+1])
				}
			}
		})
	}
}

func TestRandom(t *testing.T) {
	var s container.GSlice[int] = []int{1, 2, 3}

	for c := 0; c < 2*s.Len(); c++ {
		i := s.Random()
		p := genfuncs.OrderedEqualTo(i)
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
		want     int
		wanPanic bool
	}{
		{
			name: "Mismatched length greater",
			args: args{
				a: []string{"a"},
				b: []string{},
			},
			want: genfuncs.GreaterThan,
		},
		{
			name: "Matched",
			args: args{
				a: []string{"a", "b"},
				b: []string{"a", "b"},
			},
			want: genfuncs.EqualTo,
		},
		{
			name: "Mismatched less",
			args: args{
				a: []string{"a", "b"},
				b: []string{"a", "c"},
			},
			want: genfuncs.LessThan,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, test.args.a.Compare(test.args.b, genfuncs.Ordered[string]))
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
			tt.s.ForEach(func(_ int, i int) {
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
			tt.s.ForEach(func(i, v int) {
				assert.Equal(t, count, i)
				count++
			})
			assert.Equal(t, tt.want, count)
		})
	}
}

func TestGSlice_Values(t *testing.T) {
	s := container.GSlice[int]{1, 2, 3, 4}
	v := s.Values()

	assert.True(t, s.Equal(v, genfuncs.Ordered[int]))
}

func TestGSlice_Equal(t *testing.T) {
	type args struct {
		s2 container.GSlice[int]
	}
	tests := []struct {
		name string
		s1   container.GSlice[int]
		args args
		want bool
	}{
		{
			name: "Equal",
			s1:   container.GSlice[int]{1, 2, 3},
			args: args{
				s2: container.GSlice[int]{1, 2, 3},
			},
			want: true,
		},
		{
			name: "Wrong Ordered",
			s1:   container.GSlice[int]{1, 2, 3},
			args: args{
				s2: container.GSlice[int]{2, 1, 3},
			},
			want: false,
		},
		{
			name: "Different Lengths",
			s1:   container.GSlice[int]{1, 2, 3},
			args: args{
				s2: container.GSlice[int]{1, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.s1.Equal(tt.args.s2, genfuncs.Ordered[int]))
		})
	}
}

func TestSliceIterator_Next(t *testing.T) {
	values := []int{1, 2, 3}
	i := container.NewValuesIterator(values...)

	index := 0
	for i.HasNext() {
		assert.Equal(t, values[index], i.Next())
		index++
	}
	assert.Equal(t, index, len(values))
}
