package genfuncs_test

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"strings"
	"testing"
)

func TestExamples(t *testing.T) {
	ExampleAll()
	ExampleAny()
	ExampleAssociate()
	ExampleAssociateWith()
	ExampleContains()
	ExampleDistinct()
}

func ExampleAll() {
	numbers := []int{1, 2, 3, 4}
	positive := func(i int) bool { return i >= 0 }
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
	odsEvens := func(i int) string {
		if i%2 == 0 {
			return "EVEN"
		}
		return "ODD"
	}
	numbers := []int{1, 2, 3, 4}
	odsEvensMap := genfuncs.AssociateWith(numbers, odsEvens)
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
