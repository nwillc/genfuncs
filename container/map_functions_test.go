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
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	type args struct {
		m map[string]string
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
			keys := container.Keys(tt.args.m)
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
		m map[string]int
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
			values := container.Values(tt.args.m)
			assert.Equal(t, len(tt.want), len(values))
			for _, v := range values {
				k := strconv.Itoa(v)
				_, ok := tt.args.m[k]
				assert.True(t, ok)
			}
		})
	}
}

func TestMapContains(t *testing.T) {
	m := map[string]bool{"a": true}
	assert.Equal(t, true, container.Contains(m, "a"))
	delete(m, "a")
	assert.Equal(t, false, container.Contains(m, "a"))
}
