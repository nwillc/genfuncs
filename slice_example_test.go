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
	"strconv"
	"strings"
	"testing"

	"github.com/nwillc/genfuncs"
)

func TestSliceExamples(t *testing.T) {
	ExampleAll()
	ExampleAny()
	ExampleAssociate()
	ExampleAssociateWith()
	ExampleContains()
	ExampleDistinct()
	ExampleFilter()
	ExampleFind()
	ExampleFindLast()
	ExampleFlatMap()
	ExampleFold()
	ExampleGroupBy()
	ExampleJoinToString()
	ExampleMap()
	ExampleSortBy()
	ExampleSwap()
}

func ExampleAll() {
	// numbers := []float32{1, 2.2, 3.0, 4}
	// positive := func(i float32) bool { return i > 0 }
	// fmt.Println(genfuncs.All(numbers, positive)) // true
}

func ExampleAny() {
	// fruits := []string{"apple", "banana", "grape"}
	// isApple := func(fruit string) bool { return fruit == "apple" }
	// isPear := func(fruit string) bool { return fruit == "pear" }
	// fmt.Println(genfuncs.Any(fruits, isApple)) // true
	// fmt.Println(genfuncs.Any(fruits, isPear))  // false
}

func ExampleAssociate() {
	byLastName := func(n string) (string, string) {
		parts := strings.Split(n, " ")
		return parts[1], n
	}
	names := []string{"fred flintstone", "barney rubble"}
	nameMap := genfuncs.Associate(names, byLastName)
	fmt.Println(nameMap["rubble"]) // barney rubble
}

func ExampleAssociateWith() {
	oddEven := func(i int) string {
		if i%2 == 0 {
			return "EVEN"
		}
		return "ODD"
	}
	numbers := []int{1, 2, 3, 4}
	odsEvensMap := genfuncs.AssociateWith(numbers, oddEven)
	fmt.Println(odsEvensMap[2]) // EVEN
	fmt.Println(odsEvensMap[3]) // ODD
}

func ExampleContains() {
	var values genfuncs.Slice[float32] = []float32{1.0, .5, 42}
	fmt.Println(values.Contains(.5, genfuncs.OrderedComparator[float32]()))    // true
	fmt.Println(values.Contains(3.142, genfuncs.OrderedComparator[float32]())) // false
}

func ExampleDistinct() {
	values := []int{1, 2, 2, 3, 1, 3}
	fmt.Println(genfuncs.Distinct(values)) // [1 2 3]
}

func ExampleFilter() {
	var values genfuncs.Slice[int] = []int{1, -2, 2, -3}
	isPositive := func(i int) bool { return i > 0 }
	fmt.Println(values.Filter(isPositive)) // [1 2]
}

func ExampleFind() {
	var values genfuncs.Slice[int] = []int{-1, -2, 2, -3}
	isPositive := func(i int) bool { return i > 0 }
	fmt.Println(values.Find(isPositive)) // 2 true
}

func ExampleFindLast() {
	var values genfuncs.Slice[int] = []int{-1, -2, 2, 3}
	isPositive := func(i int) bool { return i > 0 }
	fmt.Println(values.FindLast(isPositive)) // 3 true
}

func ExampleFlatMap() {
	words := []string{"hello", " ", "world"}
	slicer := func(s string) []string { return strings.Split(s, "") }
	fmt.Println(genfuncs.FlatMap(words, slicer)) // [h e l l o   w o r l d]
}

func ExampleFold() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := func(a int, b int) int { return a + b }
	fmt.Println(genfuncs.Fold(numbers, 0, sum)) // 15
}

func ExampleGroupBy() {
	oddEven := func(i int) string {
		if i%2 == 0 {
			return "EVEN"
		}
		return "ODD"
	}
	numbers := []int{1, 2, 3, 4}
	grouped := genfuncs.GroupBy(numbers, oddEven)
	fmt.Println(grouped["ODD"]) // [1 3]
}

func ExampleJoinToString() {
	var values genfuncs.Slice[bool] = []bool{true, false, true}
	fmt.Println(values.JoinToString(
		strconv.FormatBool,
		", ",
		"{",
		"}",
	)) // {true, false, true}
}

func ExampleMap() {
	numbers := []int{69, 88, 65, 77, 80, 76, 69}
	toString := func(i int) string { return string(rune(i)) }
	fmt.Println(genfuncs.Map(numbers, toString)) // [E X A M P L E]
}

func ExampleSortBy() {
	var numbers genfuncs.Slice[int] = []int{1, 0, 9, 6, 0}
	sorted := numbers.SortBy(genfuncs.OrderedComparator[int]())
	fmt.Println(numbers) // [1 0 9 6 0]
	fmt.Println(sorted)  // [0 0 1 6 9]
}

func ExampleSwap() {
	var words genfuncs.Slice[string] = []string{"world", "hello"}
	words.Swap(0, 1)
	fmt.Println(words) // [hello world]
}
