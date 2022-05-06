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

package container_test

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"os"
	"strings"
	"testing"
)

var greaterThanZero = genfuncs.IsGreaterOrdered(0)
var wordPositions = container.GMap[string, int]{"hello": 1, "world": 2}
var words container.GSlice[string] = []string{"hello", "world"}

func TestFunctionExamples(t *testing.T) {
	if _, ok := os.LookupEnv("RUN_EXAMPLES"); !ok {
		t.Skip("skipping: RUN_EXAMPLES not set")
	}
	// Map
	ExampleContains()
	ExampleKeys()
	ExampleValues()
	// GSlice fluent
	ExampleSlice_All()
	ExampleSlice_Any()
	ExampleSlice_Filter()
	ExampleSlice_Find()
	ExampleSlice_FindLast()
	ExampleSlice_JoinToString()
	ExampleSlice_SortBy()
	ExampleSlice_Swap()
	// slice functions
	ExampleAssociate()
	ExampleAssociateWith()
	ExampleDistinct()
	ExampleFlatMap()
	ExampleFold()
	ExampleGroupBy()
	ExampleMap()
}

func ExampleContains() {
	fmt.Println(wordPositions.Contains("hello")) // true
	fmt.Println(wordPositions.Contains("no"))    // false
}

func ExampleKeys() {
	fmt.Println(wordPositions.Keys()) // [hello, world]
}

func ExampleValues() {
	fmt.Println(wordPositions.Values()) // [1, 2]
}

func ExampleSlice_All() {
	var numbers container.GSlice[int] = []int{1, 2, 3, 4}
	fmt.Println(numbers.All(greaterThanZero)) // true
}

func ExampleSlice_Any() {
	var fruits container.GSlice[string] = []string{"apple", "banana", "grape"}
	isPear := genfuncs.IsEqualOrdered("pear")
	fmt.Println(fruits.Any(isPear))               // false
	fmt.Println(fruits.Any(genfuncs.Not(isPear))) // true
}

func ExampleSlice_Filter() {
	var values container.GSlice[int] = []int{1, -2, 2, -3}
	fmt.Println(values.Filter(greaterThanZero)) // [1 2]
}

func ExampleSlice_Find() {
	var values container.GSlice[int] = []int{-1, -2, 2, -3}
	fmt.Println(values.Find(greaterThanZero)) // 2 true
}

func ExampleSlice_FindLast() {
	var values container.GSlice[int] = []int{-1, -2, 2, 3}
	fmt.Println(values.FindLast(greaterThanZero)) // 3 true
}

func ExampleSlice_JoinToString() {
	var toString genfuncs.ToString[string] = func(s string) string { return s }
	fmt.Println(words.JoinToString(
		toString,
		" ",
		"> ",
		" <",
	)) // > hello world <
}

func ExampleSlice_SortBy() {
	var numbers container.GSlice[int] = []int{1, 0, 9, 6, 0}
	fmt.Println(numbers)                                   // [1 0 9 6 0]
	fmt.Println(numbers.SortBy(genfuncs.LessOrdered[int])) // [0 0 1 6 9]
}

func ExampleSlice_Swap() {
	words.Swap(0, 1)
	fmt.Println(words) // [world hello]
}

// slice functions

func ExampleAssociate() {
	byLastName := func(n string) (string, string) {
		parts := strings.Split(n, " ")
		return parts[1], n
	}
	names := []string{"fred flintstone", "barney rubble"}
	nameMap := container.Associate(names, byLastName)
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
	odsEvensMap := container.AssociateWith(numbers, oddEven)
	fmt.Println(odsEvensMap[2]) // EVEN
	fmt.Println(odsEvensMap[3]) // ODD
}

func ExampleDistinct() {
	values := []int{1, 2, 2, 3, 1, 3}
	fmt.Println(container.Distinct(values)) // [1 2 3]
}

func ExampleFlatMap() {
	slicer := func(s string) container.GSlice[string] { return strings.Split(s, "") }
	fmt.Println(container.GSliceFlatMap(words.SortBy(genfuncs.LessOrdered[string]), slicer)) // [h e l l o w o r l d]
}

func ExampleFold() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := func(a int, b int) int { return a + b }
	fmt.Println(container.Fold(numbers, 0, sum)) // 15
}

func ExampleGroupBy() {
	oddEven := func(i int) string {
		if i%2 == 0 {
			return "EVEN"
		}
		return "ODD"
	}
	numbers := []int{1, 2, 3, 4}
	grouped := container.GroupBy(numbers, oddEven)
	fmt.Println(grouped["ODD"]) // [1 3]
}

func ExampleMap() {
	numbers := []int{69, 88, 65, 77, 80, 76, 69}
	toString := func(i int) string { return string(rune(i)) }
	fmt.Println(container.GSliceMap(numbers, toString)) // [E X A M P L E]
}
