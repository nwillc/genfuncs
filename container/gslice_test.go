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
	"github.com/nwillc/genfuncs/container/sequences"
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
			got := sequences.Any[string](tt.args.slice, genfuncs.OrderedEqualTo(tt.args.element))
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
			assert.Equal(t, genfuncs.EqualTo, sequences.Compare[int](result, tt.want, genfuncs.Ordered[int]))
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
			assert.True(t, sequences.IsSorted[string](tt.args.slice, tt.args.comparator))
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
		p := genfuncs.OrderedEqualTo(s.Random())
		assert.True(t, sequences.Any[int](s, p))
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

	assert.Equal(t, genfuncs.EqualTo, sequences.Compare[int](s, v, genfuncs.Ordered[int]))
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
			assert.Equal(t, tt.want, sequences.Compare[int](tt.s1, tt.args.s2, genfuncs.Ordered[int]) == genfuncs.EqualTo)
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

func Test_sliceIterator_NoHasNext(t *testing.T) {
	s := container.GSlice[int]{}
	iterator := s.Iterator()
	assert.False(t, iterator.HasNext())
	assert.Panics(t, func() {
		_ = iterator.Next()
	})
}
