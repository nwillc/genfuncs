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

package sequences_test

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/container/maps"
	"github.com/nwillc/genfuncs/container/sequences"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
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
		sequence  container.Sequence[string]
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
				sequence:  sequences.NewSequence[string](),
				predicate: func(s string) bool { return s == "a" },
			},
			want: true,
		},
		{
			name: "Some Not All",
			args: args{
				sequence:  sequences.NewSequence("b", "c"),
				predicate: func(s string) bool { return s == "b" },
			},
			want: false,
		},
		{
			name: "All",
			args: args{
				sequence:  sequences.NewSequence("b", "a", "c"),
				predicate: func(s string) bool { return len(s) == 1 },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sequences.All(tt.args.sequence, tt.args.predicate)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestAny(t *testing.T) {
	type args struct {
		sequence  container.Sequence[string]
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
				sequence:  sequences.NewSequence[string](),
				predicate: func(s string) bool { return s == "a" },
			},
			want: false,
		},
		{
			name: "Not Found",
			args: args{
				sequence:  sequences.NewSequence("b", "c"),
				predicate: func(s string) bool { return s == "a" },
			},
			want: false,
		},
		{
			name: "Found",
			args: args{
				sequence:  sequences.NewSequence("b", "a", "c"),
				predicate: func(s string) bool { return s == "a" },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sequences.Any(tt.args.sequence, tt.args.predicate)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestCompare(t *testing.T) {
	type args[T any] struct {
		s1 container.Sequence[T]
		s2 container.Sequence[T]
	}
	tests := []struct {
		name string
		args args[int]
		want int
	}{
		{
			name: "empty",
			args: args[int]{
				s1: sequences.NewSequence[int](),
				s2: sequences.NewSequence[int](),
			},
			want: genfuncs.EqualTo,
		},
		{
			name: "simple equal",
			args: args[int]{
				s1: sequences.NewSequence[int](2, 1),
				s2: sequences.NewSequence[int](2, 1),
			},
			want: genfuncs.EqualTo,
		},
		{
			name: "simple less",
			args: args[int]{
				s1: sequences.NewSequence[int](1),
				s2: sequences.NewSequence[int](2),
			},
			want: genfuncs.LessThan,
		},
		{
			name: "simple greater",
			args: args[int]{
				s1: sequences.NewSequence[int](1, 2),
				s2: sequences.NewSequence[int](1, 1),
			},
			want: genfuncs.GreaterThan,
		},
		{
			name: "shorter less",
			args: args[int]{
				s1: sequences.NewSequence[int](1),
				s2: sequences.NewSequence[int](1, 2),
			},
			want: genfuncs.LessThan,
		},
		{
			name: "shorter less",
			args: args[int]{
				s1: sequences.NewSequence[int](1),
				s2: sequences.NewSequence[int](1, 2),
			},
			want: genfuncs.LessThan,
		},
		{
			name: "longer greater",
			args: args[int]{
				s1: sequences.NewSequence[int](1, 2),
				s2: sequences.NewSequence[int](1),
			},
			want: genfuncs.GreaterThan,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sequences.Compare(tt.args.s1, tt.args.s2, genfuncs.Ordered[int])
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFind(t *testing.T) {
	type args struct {
		sequence  container.Sequence[float32]
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
				sequence:  sequences.NewSequence[float32](),
				predicate: func(f float32) bool { return false },
			},
			wantFound: false,
		},
		{
			name: "Not Found",
			args: args{
				sequence:  sequences.NewSequence[float32](1.0, 2.0, 3.0),
				predicate: func(f float32) bool { return false },
			},
			wantFound: false,
		},
		{
			name: "Found",
			args: args{
				sequence:  sequences.NewSequence[float32](1.0, 2.0, 3.0),
				predicate: func(f float32) bool { return f > 1.0 },
			},
			want:      2.0,
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sequences.Find(tt.args.sequence, tt.args.predicate)
			if !tt.wantFound {
				assert.False(t, got.Ok())
				return
			}
			assert.True(t, got.Ok())
			assert.Equal(t, tt.want, got.MustGet())
		})
	}
}

func TestFindLast(t *testing.T) {
	type args struct {
		sequence  container.Sequence[float32]
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
				sequence:  sequences.NewSequence[float32](),
				predicate: func(f float32) bool { return false },
			},
			wantFound: false,
		},
		{
			name: "Not Found",
			args: args{
				sequence:  sequences.NewSequence[float32](1.0, 2.0, 3.0),
				predicate: func(f float32) bool { return f < 0.0 },
			},
			wantFound: false,
		},
		{
			name: "Found",
			args: args{
				sequence:  sequences.NewSequence[float32](1.0, 2.0, 3.0),
				predicate: func(f float32) bool { return f > 1.0 },
			},
			want:      3.0,
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sequences.FindLast(tt.args.sequence, tt.args.predicate)
			if !tt.wantFound {
				assert.False(t, got.Ok())
				return
			}
			assert.True(t, got.Ok())
			assert.Equal(t, tt.want, got.MustGet())
		})
	}
}

func TestFold(t *testing.T) {
	sum := 0
	si := container.GSlice[int]{1, 2, 3}
	sum = sequences.Fold[int, int](si, 10, func(r int, i int) int { return r + i })
	assert.Equal(t, 16, sum)

	mi := container.GMap[int, int]{1: 1, 2: 2, 3: 3}
	sum = sequences.Fold[int, int](mi, 10, func(r int, i int) int { return r + i })
	assert.Equal(t, 16, sum)
}

func TestForEach(t *testing.T) {
	tests := []struct {
		name     string
		sequence container.Sequence[int]
		want     int
	}{
		{
			name:     "Empty",
			sequence: sequences.NewSequence[int](),
			want:     0,
		},
		{
			name:     "Two",
			sequence: sequences.NewSequence[int](1, 1),
			want:     2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := 0
			sequences.ForEach(tt.sequence, func(i int) {
				count++
			})
			assert.Equal(t, tt.want, count)
		})
	}
}

func TestIsSorted(t *testing.T) {
	type args struct {
		sequence container.Sequence[int]
		order    genfuncs.BiFunction[int, int, bool]
	}
	tests := []struct {
		args args
		name string
		want bool
	}{
		{
			args: args{
				sequence: sequences.NewSequence[int](),
				order:    genfuncs.OrderedLess[int],
			},
			name: "empty asc",
			want: true,
		},
		{
			args: args{
				sequence: sequences.NewSequence[int](),
				order:    genfuncs.OrderedGreater[int],
			},
			name: "empty desc",
			want: true,
		},
		{
			args: args{
				sequence: sequences.NewSequence[int](1, 2),
				order:    genfuncs.OrderedLess[int],
			},
			name: "two asc",
			want: true,
		},
		{
			args: args{
				sequence: sequences.NewSequence[int](8, 2),
				order:    genfuncs.OrderedGreater[int],
			},
			name: "two desc",
			want: true,
		},
		{
			args: args{
				sequence: sequences.NewSequence[int](1, 2, 9),
				order:    genfuncs.OrderedLess[int],
			},
			name: "three asc",
			want: true,
		},
		{
			args: args{
				sequence: sequences.NewSequence[int](8, 2, 0),
				order:    genfuncs.OrderedGreater[int],
			},
			name: "three desc",
			want: true,
		},
		{
			args: args{
				sequence: sequences.NewSequence[int](2, 2, 9),
				order:    genfuncs.OrderedLess[int],
			},
			name: "duplicate asc",
			want: true,
		},
		{
			args: args{
				sequence: sequences.NewSequence[int](8, 2, 2, 0),
				order:    genfuncs.OrderedGreater[int],
			},
			name: "duplicate desc",
			want: true,
		},
		{
			args: args{
				sequence: sequences.NewSequence[int](2, 10, 9),
				order:    genfuncs.OrderedLess[int],
			},
			name: "out of asc",
			want: false,
		},
		{
			args: args{
				sequence: sequences.NewSequence[int](8, 2, 2, 3),
				order:    genfuncs.OrderedGreater[int],
			},
			name: "out of desc",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sequences.IsSorted(tt.args.sequence, tt.args.order)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestJoinToString(t *testing.T) {
	personStringer := genfuncs.StringerToString[PersonName]()
	type args struct {
		sequence  container.Sequence[PersonName]
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
				sequence:  sequences.NewSequence[PersonName](),
				separator: "",
			},
			want: "",
		},
		{
			name: "One",
			args: args{
				sequence: sequences.NewSequence[PersonName](
					PersonName{
						First: "fred",
						Last:  "flintstone",
					}),
				separator: ", ",
			},
			want: "fred flintstone",
		},
		{
			name: "Two",
			args: args{
				sequence: sequences.NewSequence[PersonName](
					PersonName{
						First: "fred",
						Last:  "flintstone",
					},
					PersonName{
						First: "barney",
						Last:  "rubble",
					}),
				separator: ", ",
			},
			want: "fred flintstone, barney rubble",
		},
		{
			name: "Two With Prefix Postfix",
			args: args{
				sequence: sequences.NewSequence[PersonName](
					PersonName{
						First: "fred",
						Last:  "flintstone",
					},
					PersonName{
						First: "barney",
						Last:  "rubble",
					}),
				separator: ", ",
				prefix:    "[",
				postfix:   "]",
			},
			want: "[fred flintstone, barney rubble]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sequences.JoinToString(tt.args.sequence, personStringer, tt.args.separator, tt.args.prefix, tt.args.postfix)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMap(t *testing.T) {
	v := container.GSlice[int]{1, 2, 3}
	want := container.GSlice[string]{"1", "2", "3"}

	got := sequences.Map[int, string](v, func(i int) string { return fmt.Sprint(i) }).Iterator()

	index := 0
	for got.HasNext() {
		assert.Equal(t, want[index], got.Next())
		index++
	}
	assert.Len(t, want, index)
}

func TestAssociate(t *testing.T) {
	var firstLast maps.KeyValueFor[PersonName, string, string] = func(p PersonName) *genfuncs.Result[*maps.Entry[string, string]] {
		return genfuncs.NewResult(maps.NewEntry(p.First, p.Last))
	}
	type args struct {
		sequence  container.Sequence[PersonName]
		transform maps.KeyValueFor[PersonName, string, string]
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
				sequence:  sequences.NewSequence[PersonName](),
				transform: firstLast,
			},
			wantSize: 0,
		},
		{
			name: "Two Unique",
			args: args{
				sequence: container.GMap[string, PersonName]{
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
				sequence: container.GSlice[PersonName]{
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
		{
			name: "list one",
			args: args{
				sequence:  container.NewList(PersonName{First: "Donald", Last: "Duck"}),
				transform: firstLast,
			},
			wantSize: 1,
			contains: []string{"Donald"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sequences.Associate(tt.args.sequence, tt.args.transform).
				OnFailure(func(e error) {
					assert.Fail(t, "failed associate")
				}).
				OnSuccess(func(fNameMap container.GMap[string, string]) {
					assert.Equal(t, tt.wantSize, fNameMap.Len())
					for k := range fNameMap {
						_, ok := fNameMap[k]
						assert.True(t, ok)
					}
				})
		})
	}
}

func TestAssociateWith(t *testing.T) {
	var valueSelector maps.ValueFor[int, int] = func(i int) *genfuncs.Result[int] { return genfuncs.NewResult(i * 2) }
	type args struct {
		sequence  container.Sequence[int]
		transform maps.ValueFor[int, int]
	}
	tests := []struct {
		name     string
		args     args
		wantSize int
	}{
		{
			name: "Empty",
			args: args{
				sequence:  container.NewList[int](),
				transform: valueSelector,
			},
			wantSize: 0,
		},
		{
			name: "Three Unique",
			args: args{
				sequence:  container.GSlice[int]{1, 2, 3},
				transform: valueSelector,
			},
			wantSize: 3,
		},
		{
			name: "Duplicate",
			args: args{
				sequence:  container.GMap[string, int]{"1": 1, "2": 2, "two": 2},
				transform: valueSelector,
			},
			wantSize: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultMap := sequences.AssociateWith(tt.args.sequence, tt.args.transform)
			assert.True(t, resultMap.Ok())
			m := resultMap.OrEmpty()
			assert.Equal(t, tt.wantSize, m.Len())
			sequences.ForEach(tt.args.sequence, func(k int) {
				assert.True(t, m.Contains(k))
				assert.Equal(t, k*2, m[k])
			})
		})
	}
}

func TestFlatMap(t *testing.T) {
	var trans = func(i int) container.Sequence[string] { return sequences.NewSequence("#", strconv.Itoa(i)) }
	type args struct {
		sequence  container.Sequence[int]
		transform func(int) container.Sequence[string]
	}
	tests := []struct {
		name string
		args args
		want container.Sequence[string]
	}{
		{
			name: "Empty",
			args: args{
				sequence:  container.GSlice[int]{},
				transform: trans,
			},
			want: sequences.NewSequence[string](),
		},
		{
			name: "List",
			args: args{
				sequence:  container.NewList[int](1, 2, 3),
				transform: trans,
			},
			want: sequences.NewSequence("#", "1", "#", "2", "#", "3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sequences.FlatMap[int, string](tt.args.sequence, tt.args.transform)
			assert.Equal(t, genfuncs.EqualTo, sequences.Compare[string](got, tt.want, genfuncs.Ordered[string]))
		})
	}
}

func TestNewSequence(t *testing.T) {
	values := []int{1, 2, 3}
	iterator := sequences.NewSequence(values...).Iterator()
	for _, v := range values {
		assert.True(t, iterator.HasNext())
		assert.Equal(t, v, iterator.Next())
	}
}
