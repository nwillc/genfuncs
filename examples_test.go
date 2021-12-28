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
	"os"
	"strings"
	"testing"
	"time"
)

var floatCmp = genfuncs.OrderedComparator[float32]()
var greaterThanZero = genfuncs.IsGreaterThan(0)
var intCmp = genfuncs.OrderedComparator[int]()
var lexicalOrder = genfuncs.OrderedComparator[string]()
var reverseLexical = genfuncs.ReverseComparator(lexicalOrder)
var wordPositions = map[string]int{"hello": 1, "world": 2}
var words genfuncs.Slice[string] = []string{"hello", "world"}

func TestFunctionExamples(t *testing.T) {
	if _, ok := os.LookupEnv("RUN_EXAMPLES"); !ok {
		t.Skip("skipping: RUN_EXAMPLES not set")
	}
	// Functions
	ExampleOrderedComparator()
	ExampleReverseComparator()
	ExampleStringerStringer()
	ExampleFunctionComparator()
	// Heap
	ExampleNewHeap()
	// Map
	ExampleKeys()
	ExampleValues()
	// Slice fluent
	ExampleSlice_All()
	ExampleSlice_Any()
	ExampleSlice_Contains()
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
	ExampleFold()
	ExampleFlatMap()
	ExampleGroupBy()
	ExampleMap()
	// Sort
	ExampleSort()
}

func ExampleOrderedComparator() {
	fmt.Println(lexicalOrder("a", "b")) // -1
	fmt.Println(lexicalOrder("a", "a")) // 0
	fmt.Println(lexicalOrder("b", "a")) // 1
}

func ExampleReverseComparator() {
	fmt.Println(lexicalOrder("a", "b"))   // -1
	fmt.Println(reverseLexical("a", "b")) // 1
}

func ExampleStringerStringer() {
	var epoch time.Time
	fmt.Println(epoch.String()) // 0001-01-01 00:00:00 +0000 UTC
	stringer := genfuncs.StringerStringer[time.Time]()
	fmt.Println(stringer(epoch)) // 0001-01-01 00:00:00 +0000 UTC
}

func ExampleFunctionComparator() {
	var unixTime = func(t time.Time) int64 { return t.Unix() }
	var timeComparator = genfuncs.FunctionComparator(unixTime, genfuncs.OrderedComparator[int64]())

	now := time.Now()
	fmt.Println(timeComparator(now, now.Add(time.Second))) // -1
}

func ExampleNewHeap() {
	heap := genfuncs.NewHeap(intCmp)
	heap.PushAll(3, 1, 4, 2)
	for heap.Len() > 0 {
		fmt.Print(heap.Pop()) // 1234
	}
	fmt.Println()
}

func ExampleKeys() {
	keys := genfuncs.Keys(wordPositions)
	fmt.Println(keys) // [hello, world]
}

func ExampleValues() {
	values := genfuncs.Values(wordPositions)
	fmt.Println(values) // [1, 2]
}

func ExampleSlice_All() {
	var numbers genfuncs.Slice[int] = []int{1, 2, 3, 4}
	fmt.Println(numbers.All(greaterThanZero)) // true
}

func ExampleSlice_Any() {
	var fruits genfuncs.Slice[string] = []string{"apple", "banana", "grape"}
	isPear := genfuncs.IsEqualTo("pear")
	isNotPear := isPear.Not()
	fmt.Println(fruits.Any(isNotPear)) // true
	fmt.Println(fruits.Any(isPear))    // false
}

func ExampleSlice_Contains() {
	var floats genfuncs.Slice[float32] = []float32{1.0, .5, 42}
	fmt.Println(floats.Contains(.5, floatCmp))    // true
	fmt.Println(floats.Contains(3.142, floatCmp)) // false
}

func ExampleSlice_Filter() {
	var values genfuncs.Slice[int] = []int{1, -2, 2, -3}
	fmt.Println(values.Filter(greaterThanZero)) // [1 2]
}

func ExampleSlice_Find() {
	var values genfuncs.Slice[int] = []int{-1, -2, 2, -3}
	fmt.Println(values.Find(greaterThanZero)) // 2 true
}

func ExampleSlice_FindLast() {
	var values genfuncs.Slice[int] = []int{-1, -2, 2, 3}
	fmt.Println(values.FindLast(greaterThanZero)) // 3 true
}

func ExampleSlice_JoinToString() {
	var toString genfuncs.Stringer[string] = func(s string) string { return s }
	fmt.Println(words.JoinToString(
		toString,
		", ",
		"{",
		"}",
	)) // {hello, world}
}

func ExampleSlice_SortBy() {
	var numbers genfuncs.Slice[int] = []int{1, 0, 9, 6, 0}
	sorted := numbers.SortBy(intCmp)
	fmt.Println(numbers) // [1 0 9 6 0]
	fmt.Println(sorted)  // [0 0 1 6 9]
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

func ExampleDistinct() {
	values := []int{1, 2, 2, 3, 1, 3}
	fmt.Println(genfuncs.Distinct(values)) // [1 2 3]
}

func ExampleFlatMap() {
	slicer := func(s string) genfuncs.Slice[string] { return strings.Split(s, "") }
	fmt.Println(genfuncs.FlatMap(words.SortBy(lexicalOrder), slicer)) // [h e l l o w o r l d]
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

func ExampleMap() {
	numbers := []int{69, 88, 65, 77, 80, 76, 69}
	toString := func(i int) string { return string(rune(i)) }
	fmt.Println(genfuncs.Map(numbers, toString)) // [E X A M P L E]
}

func ExampleSort() {
	var letters genfuncs.Slice[string] = strings.Split("example", "")

	letters.Sort(lexicalOrder)
	fmt.Println(letters) // [a e e l m p x]
	letters.Sort(reverseLexical)
	fmt.Println(letters) // [x p m l e e a]
}
