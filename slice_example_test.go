package genfuncs_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/nwillc/genfuncs"
)

func TestExamples(t *testing.T) {
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
	ExampleReduce()
}

func ExampleAll() {
	numbers := []float32{1, 2.2, 3.0, 4}
	positive := func(i float32) bool { return i > 0 }
	fmt.Println(genfuncs.All(numbers, positive)) // true
}

func ExampleAny() {
	fruits := []string{"apple", "banana", "grape"}
	isApple := func(fruit string) bool { return fruit == "apple" }
	isPear := func(fruit string) bool { return fruit == "pear" }
	fmt.Println(genfuncs.Any(fruits, isApple)) // true
	fmt.Println(genfuncs.Any(fruits, isPear))  // false
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
	values := []float32{1.0, .5, 42}
	fmt.Println(genfuncs.Contains(values, .5))    // true
	fmt.Println(genfuncs.Contains(values, 3.142)) // false
}

func ExampleDistinct() {
	values := []int{1, 2, 2, 3, 1, 3}
	fmt.Println(genfuncs.Distinct(values)) // [1 2 3]
}

func ExampleFilter() {
	values := []int{1, -2, 2, -3}
	isPositive := func(i int) bool { return i > 0 }
	fmt.Println(genfuncs.Filter(values, isPositive)) // [1 2]
}

func ExampleFind() {
	values := []int{-1, -2, 2, -3}
	isPositive := func(i int) bool { return i > 0 }
	fmt.Println(genfuncs.Find(values, isPositive)) // 2 true
}

func ExampleFindLast() {
	values := []int{-1, -2, 2, 3}
	isPositive := func(i int) bool { return i > 0 }
	fmt.Println(genfuncs.FindLast(values, isPositive)) // 3 true
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
	values := []bool{true, false, true}
	btos := func(b bool) string { return strconv.FormatBool(b) }
	fmt.Println(genfuncs.JoinToString(values, btos, ", ", "{", "}")) // {true, false, true}
}

func ExampleMap() {
	numbers := []int{69, 88, 65, 77, 80, 76, 69}
	toString := func(i int) string { return string(rune(i)) }
	fmt.Println(genfuncs.Map(numbers, toString)) // [E X A M P L E]
}

func ExampleReduce() {
	type item struct {
		name  string
		price float64
	}
	items := []item{
		{name: "Salmon", price: 5},
		{name: "Tuna", price: 3.5},
	}
	fmt.Println(genfuncs.Reduce(items, func(total float64, current item) float64 { return total + current.price }, 0)) // 8.5
}
