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

package iterables_test

import (
	"fmt"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/container/iterables"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFold(t *testing.T) {
	sum := 0
	si := container.GSlice[int]{1, 2, 3}
	sum = iterables.Fold[int, int](si, 10, func(r int, i int) int { return r + i })
	assert.Equal(t, 16, sum)

	mi := container.GMap[int, int]{1: 1, 2: 2, 3: 3}
	sum = iterables.Fold[int, int](mi, 10, func(r int, i int) int { return r + i })
	assert.Equal(t, 16, sum)
}

func TestMap(t *testing.T) {
	v := container.GSlice[int]{1, 2, 3}
	want := container.GSlice[string]{"1", "2", "3"}

	got := iterables.Map[int, string](v, func(i int) string { return fmt.Sprint(i) })

	index := 0
	for got.HasNext() {
		assert.Equal(t, want[index], got.Next())
		index++
	}
	assert.Len(t, want, index)
}
