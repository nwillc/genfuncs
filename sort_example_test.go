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
	"fmt"
	"github.com/nwillc/genfuncs"
	"strings"
	"testing"
)

var order genfuncs.Comparator[string] = func(a, b string) genfuncs.ComparedOrder {
	switch {
	case a < b:
		return genfuncs.LessThan
	case a > b:
		return genfuncs.GreaterThan
	default:
		return genfuncs.EqualTo
	}
}

func TestSwapExamples(t *testing.T) {
	ExampleInsertionSort()
	ExampleHeapSort()
}

func ExampleInsertionSort() {
	letters := strings.Split("example", "")

	genfuncs.InsertionSort(letters, order)
	fmt.Println(letters) // [a e e l m p x]
	genfuncs.InsertionSort(letters, genfuncs.ReverseComparator(order))
	fmt.Println(letters) // [x p m l e e a]
}

func ExampleHeapSort() {
	letters := strings.Split("example", "")

	genfuncs.HeapSort(letters, order)
	fmt.Println(letters) // [a e e l m p x]
	genfuncs.HeapSort(letters, genfuncs.ReverseComparator(order))
	fmt.Println(letters) // [x p m l e e a]
}
