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
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/internal/tests"
	"testing"
)

var isGreaterThanZero = genfuncs.OrderedGreaterThan(0)
var wordPositions = container.GMap[string, int]{"hello": 1, "world": 2}
var words container.GSlice[string] = []string{"hello", "world"}

func TestFunctionExamples(t *testing.T) {
	tests.MaybeRunExamples(t)
	ExampleGMap_Contains()
	ExampleGMap_Keys()
	ExampleGMap_Values()
	ExampleGSlice_Filter()
	ExampleGSlice_SortBy()
	ExampleGSlice_Swap()
}

func ExampleGMap_Contains() {
	fmt.Println(wordPositions.Contains("hello"))
	fmt.Println(wordPositions.Contains("no"))
	// Output:
	// true
	// false
}

func ExampleGMap_Keys() {
	fmt.Println(wordPositions.Keys().SortBy(genfuncs.OrderedLess[string]))
	// Output: [hello world]
}

func ExampleGMap_Values() {
	wordPositions.Values().ForEach(func(_, i int) { fmt.Println(i) })
	// Unordered Output:
	// 1
	// 2
}

func ExampleGSlice_Filter() {
	var values container.GSlice[int] = []int{1, -2, 2, -3}
	values.Filter(isGreaterThanZero).ForEach(func(_, i int) {
		fmt.Println(i)
	})
	// Unordered Output:
	// 1
	// 2
}

func ExampleGSlice_SortBy() {
	var numbers container.GSlice[int] = []int{1, 0, 9, 6, 0}
	fmt.Println(numbers)
	fmt.Println(numbers.SortBy(genfuncs.OrderedLess[int]))
	// Output:
	// [1 0 9 6 0]
	// [0 0 1 6 9]
}

func ExampleGSlice_Swap() {
	words = words.SortBy(genfuncs.OrderedLess[string])
	words.Swap(0, 1)
	fmt.Println(words)
	// Output: [world hello]
}
