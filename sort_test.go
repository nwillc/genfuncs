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

package genfuncs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	letters                       = []string{"t", "e", "s", "t"}
	strCompare Comparator[string] = func(a, b string) ComparedOrder {
		if a == b {
			return EqualTo
		}
		if a < b {
			return LessThan
		}
		return GreaterThan
	}
)

func TestInsertionSort(t *testing.T) {
	type args struct {
		slice      []string
		comparator Comparator[string]
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
				comparator: strCompare,
			},
			want: []string{},
		},
		{
			name: "Sort Min Max",
			args: args{
				slice:      letters,
				comparator: strCompare,
			},
			want: []string{"e", "s", "t", "t"},
		},
		{
			name: "Sort Max Min",
			args: args{
				slice:      letters,
				comparator: ReverseComparator(strCompare),
			},
			want: []string{"t", "t", "s", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := make([]string, len(tt.args.slice))
			copy(dst, tt.args.slice)
			InsertionSort(dst, tt.args.comparator)
			assert.Equal(t, len(tt.want), len(dst))
			for i := 0; i < len(tt.want); i++ {
				assert.Equal(t, tt.want[i], dst[i])
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	type args struct {
		slice      []string
		comparator Comparator[string]
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
				comparator: strCompare,
			},
			want: []string{},
		},
		{
			name: "Sort Min Max",
			args: args{
				slice:      letters,
				comparator: strCompare,
			},
			want: []string{"e", "s", "t", "t"},
		},
		{
			name: "Sort Max Min",
			args: args{
				slice:      letters,
				comparator: ReverseComparator(strCompare),
			},
			want: []string{"t", "t", "s", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := make([]string, len(tt.args.slice))
			copy(dst, tt.args.slice)
			HeapSort(dst, tt.args.comparator)
			assert.Equal(t, len(tt.want), len(dst))
			for i := 0; i < len(tt.want); i++ {
				assert.Equal(t, tt.want[i], dst[i])
			}
		})
	}
}
