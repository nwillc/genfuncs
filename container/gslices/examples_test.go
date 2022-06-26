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

package gslices_test

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/container/gslices"
	"github.com/nwillc/genfuncs/internal/tests"
	"strings"
	"testing"
)

var words container.GSlice[string] = []string{"hello", "world"}

func TestFunctionExamples(t *testing.T) {
	tests.MaybeRunExamples(t)
	ExampleDistinct()
	ExampleFlatMap()
	ExampleGroupBy()
	ExampleMap()
}

func ExampleDistinct() {
	values := []int{1, 2, 2, 3, 1, 3}
	gslices.Distinct(values).ForEach(func(_, i int) {
		fmt.Println(i)
	})
	// Unordered Output:
	// 1
	// 2
	// 3
}

func ExampleFlatMap() {
	slicer := func(s string) container.GSlice[string] { return strings.Split(s, "") }
	fmt.Println(gslices.FlatMap(words.SortBy(genfuncs.OrderedLess[string]), slicer))
	// Output: [h e l l o w o r l d]
}

func ExampleGroupBy() {
	oddEven := func(i int) string {
		if i%2 == 0 {
			return "EVEN"
		}
		return "ODD"
	}
	numbers := []int{1, 2, 3, 4}
	grouped := gslices.GroupBy(numbers, oddEven)
	fmt.Println(grouped["ODD"])
	// Output: [1 3]
}

func ExampleMap() {
	numbers := []int{69, 88, 65, 77, 80, 76, 69}
	toString := func(i int) string { return string(rune(i)) }
	fmt.Println(gslices.Map(numbers, toString))
	// Output: [E X A M P L E]
}
