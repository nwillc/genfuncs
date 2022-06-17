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

package result_test

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/result"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	transform := func(i int) string { return fmt.Sprintf("%d", i) }
	type args struct {
		t *genfuncs.Result[int]
	}
	tests := []struct {
		name  string
		args  args
		wantR *genfuncs.Result[string]
	}{
		{
			name: "10",
			args: args{
				t: genfuncs.NewResult(10),
			},
			wantR: genfuncs.NewResult("10"),
		},
		{
			name: "error",
			args: args{
				t: genfuncs.NewError[int](fmt.Errorf("foo")),
			},
			wantR: genfuncs.NewError[string](fmt.Errorf("foo")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR := result.Map(tt.args.t, transform)
			assert.Equal(t, tt.wantR.Ok(), gotR.Ok())
			assert.Equal(t, tt.wantR.Error(), gotR.Error())
			assert.Equal(t, tt.wantR.OrEmpty(), gotR.OrEmpty())
		})
	}
}

func TestMapError(t *testing.T) {
	iErr := genfuncs.NewError[int](fmt.Errorf("foo"))
	sErr := result.MapError[int, string](iErr)
	assert.Equal(t, sErr.Error().Error(), "foo")
}
