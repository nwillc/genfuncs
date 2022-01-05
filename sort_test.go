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
	"github.com/nwillc/genfuncs"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	letters    = []string{"t", "e", "s", "t"}
	strCompare = genfuncs.OrderedLessThan[string]()
)

func TestSort(t *testing.T) {
	type args struct {
		slice      genfuncs.Slice[string]
		comparator genfuncs.LessThan[string]
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
			name: "Single",
			args: args{
				slice:      []string{"a"},
				comparator: strCompare,
			},
			want: []string{"a"},
		},
		{
			name: "Double",
			args: args{
				slice:      []string{"a", "b"},
				comparator: strCompare,
			},
			want: []string{"a", "b"},
		},
		{
			name: "Double Reverse",
			args: args{
				slice:      []string{"a", "b"},
				comparator: genfuncs.Reverse(strCompare),
			},
			want: []string{"b", "a"},
		},
		{
			name: "Min Max",
			args: args{
				slice:      letters,
				comparator: strCompare,
			},
			want: []string{"e", "s", "t", "t"},
		},
		{
			name: "Max Min",
			args: args{
				slice:      letters,
				comparator: genfuncs.Reverse(strCompare),
			},
			want: []string{"t", "t", "s", "e"},
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
