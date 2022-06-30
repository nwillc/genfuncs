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
	"github.com/nwillc/genfuncs/internal/tests"
	"testing"

	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
)

func TestHeapFunctionExamples(t *testing.T) {
	tests.MaybeRunExamples(t)
	// Heap
	ExampleNewHeap()
}

func ExampleNewHeap() {
	heap := container.NewHeap[int](genfuncs.OrderedLess[int], 3, 1, 4, 2)
	for heap.Len() > 0 {
		fmt.Print(heap.Remove())
	}
	fmt.Println()
	// Output: 1234
}
