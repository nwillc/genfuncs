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
	"os"
	"testing"
)

var greaterThanZero = genfuncs.IsGreaterOrdered(0)
var wordPositions = container.GMap[string, int]{"hello": 1, "world": 2}
var words container.GSlice[string] = []string{"hello", "world"}

func TestFunctionExamples(t *testing.T) {
	if _, ok := os.LookupEnv("RUN_EXAMPLES"); !ok {
		t.Skip("skipping: RUN_EXAMPLES not set")
	}
	ExampleGMap_Contains()
	ExampleGMap_Keys()
	ExampleGMap_Values()
	ExampleGSlice_All()
	ExampleGSlice_Any()
	ExampleGSlice_Filter()
	ExampleGSlice_Find()
	ExampleGSlice_FindLast()
	ExampleGSlice_JoinToString()
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
	fmt.Println(wordPositions.Keys().SortBy(genfuncs.LessOrdered[string]))
	// Output: [hello world]
}

func ExampleGMap_Values() {
	wordPositions.Values().ForEach(func(_, i int) { fmt.Println(i) })
	// Unordered Output:
	// 1
	// 2
}

func ExampleGSlice_All() {
	var numbers container.GSlice[int] = []int{1, 2, 3, 4}
	fmt.Println(numbers.All(greaterThanZero))
	// Output: true
}

func ExampleGSlice_Any() {
	var fruits container.GSlice[string] = []string{"apple", "banana", "grape"}
	isPear := genfuncs.IsEqualOrdered("pear")
	fmt.Println(fruits.Any(isPear))
	fmt.Println(fruits.Any(genfuncs.Not(isPear)))
	// Output:
	// false
	// true
}

func ExampleGSlice_Filter() {
	var values container.GSlice[int] = []int{1, -2, 2, -3}
	values.Filter(greaterThanZero).ForEach(func(_, i int) {
		fmt.Println(i)
	})
	// Unordered Output:
	// 1
	// 2
}

func ExampleGSlice_Find() {
	var values container.GSlice[int] = []int{-1, -2, 2, -3}
	fmt.Println(values.Find(greaterThanZero))
	// Output: 2 true
}

func ExampleGSlice_FindLast() {
	var values container.GSlice[int] = []int{-1, -2, 2, 3}
	fmt.Println(values.FindLast(greaterThanZero))
	// Output: 3 true
}

func ExampleGSlice_JoinToString() {
	var toString genfuncs.ToString[string] = func(s string) string { return s }
	fmt.Println(words.SortBy(genfuncs.LessOrdered[string]).JoinToString(
		toString,
		" ",
		"> ",
		" <",
	))
	// Output: > hello world <
}

func ExampleGSlice_SortBy() {
	var numbers container.GSlice[int] = []int{1, 0, 9, 6, 0}
	fmt.Println(numbers)
	fmt.Println(numbers.SortBy(genfuncs.LessOrdered[int]))
	// Output:
	// [1 0 9 6 0]
	// [0 0 1 6 9]
}

func ExampleGSlice_Swap() {
	words = words.SortBy(genfuncs.LessOrdered[string])
	words.Swap(0, 1)
	fmt.Println(words)
	// Output: [world hello]
}
