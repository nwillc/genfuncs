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

package genfuncs_test

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResult(t *testing.T) {
	var r *genfuncs.Result[int]
	err := fmt.Errorf("ro ruh")
	r = genfuncs.NewError[int](err)
	r.
		OnFailure(func(e error) {
			assert.Equal(t, err, e)
		}).
		OnSuccess(func(_ int) {
			assert.Fail(t, "success on an error")
		})

	assert.Panics(t, func() {
		r.ValueOrPanic()
	})

	assert.Equal(t, 10, r.ValueOr(10))

	r = genfuncs.NewResult(10)
	assert.Equal(t, 10, r.ValueOrPanic())

	r = r.Then(func(i int) *genfuncs.Result[int] {
		return genfuncs.NewResult(i * 10)
	})
	assert.Equal(t, 100, r.ValueOrPanic())
}

func TestResult_Error(t *testing.T) {
	type fields struct {
		value int
		err   error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "no error",
			fields: fields{
				value: 0,
				err:   nil,
			},
		},
		{
			name: "error",
			fields: fields{
				value: 0,
				err:   fmt.Errorf("foo"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r *genfuncs.Result[int]
			if tt.fields.err != nil {
				r = genfuncs.NewError[int](tt.fields.err)
			} else {
				r = genfuncs.NewResult(tt.fields.value)
			}

			if tt.wantErr {
				assert.False(t, r.Ok())
				assert.NotNil(t, r.Error())
			} else {
				assert.True(t, r.Ok())
				assert.Nil(t, r.Error())
			}
		})
	}
}
