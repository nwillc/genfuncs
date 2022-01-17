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

var greaterThanZero = genfuncs.IsGreaterThanOrdered(0)
var wordPositions = map[string]int{"hello": 1, "world": 2}
var words container.Slice[string] = []string{"hello", "world"}

func TestFunctionExamples(t *testing.T) {
	if _, ok := os.LookupEnv("RUN_EXAMPLES"); !ok {
		t.Skip("skipping: RUN_EXAMPLES not set")
	}
	// Heap
	ExampleNewHeap()
	// Map
	ExampleContains()
	ExampleKeys()
	ExampleValues()
	// Slice fluent
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
	// Sort
	ExampleSort()
}

func ExampleNewHeap() {
	heap := container.NewHeap(genfuncs.INumericOrder, 3, 1, 4, 2)
	for heap.Len() > 0 {
		fmt.Print(heap.Remove())
	}
	fmt.Println()
	// 1234
}

func ExampleContains() {
	fmt.Println(container.Contains(wordPositions, "hello")) // true
	fmt.Println(container.Contains(wordPositions, "no"))    // false
}

func ExampleKeys() {
	keys := container.Keys(wordPositions)
	fmt.Println(keys) // [hello, world]
}

func ExampleValues() {
	values := container.Values(wordPositions)
	fmt.Println(values) // [1, 2]
}

func ExampleSlice_All() {
	var numbers container.Slice[int] = []int{1, 2, 3, 4}
	fmt.Println(numbers.All(greaterThanZero)) // true
}

func ExampleSlice_Any() {
	var fruits container.Slice[string] = []string{"apple", "banana", "grape"}
	isPear := genfuncs.IsEqualComparable("pear")
	fmt.Println(fruits.Any(isPear))               // false
	fmt.Println(fruits.Any(genfuncs.Not(isPear))) // true
}

func ExampleSlice_Filter() {
	var values container.Slice[int] = []int{1, -2, 2, -3}
	fmt.Println(values.Filter(greaterThanZero)) // [1 2]
}

func ExampleSlice_Find() {
	var values container.Slice[int] = []int{-1, -2, 2, -3}
	fmt.Println(values.Find(greaterThanZero)) // 2 true
}

func ExampleSlice_FindLast() {
	var values container.Slice[int] = []int{-1, -2, 2, 3}
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
	var numbers container.Slice[int] = []int{1, 0, 9, 6, 0}
	fmt.Println(numbers)                                // [1 0 9 6 0]
	fmt.Println(numbers.SortBy(genfuncs.INumericOrder)) // [0 0 1 6 9]
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
	slicer := func(s string) container.Slice[string] { return strings.Split(s, "") }
	fmt.Println(container.FlatMap(words.SortBy(genfuncs.SLexicalOrder), slicer)) // [h e l l o w o r l d]
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
	fmt.Println(container.Map(numbers, toString)) // [E X A M P L E]
}

func ExampleSort() {
	var letters container.Slice[string] = strings.Split("example", "")

	letters.Sort(genfuncs.SLexicalOrder)
	fmt.Println(letters) // [a e e l m p x]
	letters.Sort(genfuncs.SReverseLexicalOrder)
	fmt.Println(letters) // [x p m l e e a]
}
