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
	"github.com/nwillc/genfuncs/container"
	"math/rand"
	"testing"
	"time"

	"github.com/nwillc/genfuncs"
	"github.com/stretchr/testify/assert"
)

var (
	letters  = []string{"t", "e", "s", "t"}
	alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "t", "u", "v", "w", "x", "y", "z"}
)

func TestSort(t *testing.T) {
	type args struct {
		slice      container.GSlice[string]
		comparator genfuncs.BiFunction[string, string, bool]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty",
			args: args{
				slice:      []string{},
				comparator: genfuncs.SLexicalOrder,
			},
			want: []string{},
		},
		{
			name: "Single",
			args: args{
				slice:      []string{"a"},
				comparator: genfuncs.SLexicalOrder,
			},
			want: []string{"a"},
		},
		{
			name: "Double",
			args: args{
				slice:      []string{"a", "b"},
				comparator: genfuncs.SLexicalOrder,
			},
			want: []string{"a", "b"},
		},
		{
			name: "Double Reverse",
			args: args{
				slice:      []string{"a", "b"},
				comparator: genfuncs.SReverseLexicalOrder,
			},
			want: []string{"b", "a"},
		},
		{
			name: "Min Max",
			args: args{
				slice:      letters,
				comparator: genfuncs.SLexicalOrder,
			},
			want: []string{"e", "s", "t", "t"},
		},
		{
			name: "Max Min",
			args: args{
				slice:      letters,
				comparator: genfuncs.SReverseLexicalOrder,
			},
			want: []string{"t", "t", "s", "e"},
		},
		{
			name: "More than 12",
			args: args{
				slice:      alphabet,
				comparator: genfuncs.SLexicalOrder,
			},
			want: alphabet,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := tt.args.slice.SortBy(tt.args.comparator)
			assert.Equal(t, len(tt.want), len(dst))
			for i := 0; i < len(tt.want); i++ {
				assert.Equal(t, tt.want[i], dst[i])
			}
		})
	}
}

func TestRandomSorts(t *testing.T) {
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
				numbers.Sort(genfuncs.INumericOrder)
				for i := 0; i < count-1; i++ {
					assert.LessOrEqual(t, numbers[i], numbers[i+1])
				}
				for i := 0; i < count; i++ {
					numbers[i] = random.Int()
				}
				numbers.Sort(genfuncs.IReverseNumericOrder)
				for i := 0; i < count-1; i++ {
					assert.GreaterOrEqual(t, numbers[i], numbers[i+1])
				}
			}
		})
	}
}
