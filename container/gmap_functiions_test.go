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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatMerge(t *testing.T) {
	m1 := make(container.GMap[string, container.GSlice[string]])
	m2 := make(container.GMap[string, container.GSlice[string]])

	m1["a"] = container.GSlice[string]{"1"}
	m2["a"] = container.GSlice[string]{"2"}
	m2["b"] = container.GSlice[string]{"1"}

	m3 := container.MapMerge(m1, m2)
	assert.Equal(t, 2, len(m3))
	assert.Equal(t, 2, len(m3["a"]))
	assert.Equal(t, 1, len(m3["b"]))
}
