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
	"os"
	"strings"
	"testing"
)

var words container.GSlice[string] = []string{"hello", "world"}

func TestFunctionExamples(t *testing.T) {
	if _, ok := os.LookupEnv("RUN_EXAMPLES"); !ok {
		t.Skip("skipping: RUN_EXAMPLES not set")
	}
	ExampleAssociate()
	ExampleAssociateWith()
	ExampleDistinct()
	ExampleFlatMap()
	ExampleFold()
	ExampleGroupBy()
	ExampleMap()
}

func ExampleAssociate() {
	byLastName := func(n string) (string, string) {
		parts := strings.Split(n, " ")
		return parts[1], n
	}
	names := []string{"fred flintstone", "barney rubble"}
	nameMap := gslices.Associate(names, byLastName)
	fmt.Println(nameMap["rubble"])
	// Output: barney rubble
}

func ExampleAssociateWith() {
	oddEven := func(i int) string {
		if i%2 == 0 {
			return "EVEN"
		}
		return "ODD"
	}
	numbers := []int{1, 2, 3, 4}
	odsEvensMap := gslices.AssociateWith(numbers, oddEven)
	fmt.Println(odsEvensMap[2])
	fmt.Println(odsEvensMap[3])
	// Output:
	// EVEN
	// ODD
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

func ExampleFold() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := func(a int, b int) int { return a + b }
	fmt.Println(gslices.Fold(numbers, 0, sum))
	// Output: 15
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
