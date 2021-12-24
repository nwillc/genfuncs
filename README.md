<!-- Code generated by gomarkdoc. DO NOT EDIT -->

[![License](https://img.shields.io/github/license/nwillc/genfuncs.svg)](https://tldrlegal.com/license/-isc-license)
[![Releases](https://img.shields.io/github/tag/nwillc/genfuncs.svg)](https://github.com/nwillc/genfuncs/tags)

# genfuncs

```go
import "github.com/nwillc/genfuncs"
```

Package genfuncs implements various functions utilizing Go's Generics to help avoid writing boilerplate code\, in particular when working with slices\. Many of the functions are based on Kotlin's Sequence\. This package\, though usable\, is primarily a proof\-of\-concept since it is likely Go will provide similar at some point soon\.

The code is under the ISC License: https://github.com/nwillc/genfuncs/blob/master/LICENSE.md

## Index

- [func All[T any](slice []T, predicate Predicate[T]) bool](<#func-all>)
- [func Any[T any](slice []T, predicate Predicate[T]) bool](<#func-any>)
- [func Associate[T, V any, K comparable](slice []T, keyValueFor KeyValueFor[T, K, V]) map[K]V](<#func-associate>)
- [func AssociateWith[K comparable, V any](slice []K, valueFor ValueFor[K, V]) map[K]V](<#func-associatewith>)
- [func Contains[T comparable](slice []T, element T) bool](<#func-contains>)
- [func Distinct[T comparable](slice []T) []T](<#func-distinct>)
- [func Filter[T any](slice []T, predicate Predicate[T]) []T](<#func-filter>)
- [func Find[T any](slice []T, predicate Predicate[T]) (T, bool)](<#func-find>)
- [func FindLast[T any](slice []T, predicate Predicate[T]) (T, bool)](<#func-findlast>)
- [func FlatMap[T, R any](slice []T, function Function[T, []R]) []R](<#func-flatmap>)
- [func Fold[T, R any](slice []T, initial R, biFunction BiFunction[R, T, R]) R](<#func-fold>)
- [func GroupBy[T any, K comparable](slice []T, keyFor KeyFor[T, K]) map[K][]T](<#func-groupby>)
- [func JoinToString[T any](slice []T, stringer Stringer[T], separator string, prefix string, postfix string) string](<#func-jointostring>)
- [func Map[T, R any](slice []T, function Function[T, R]) []R](<#func-map>)
- [func Sort[T any](slice []T, comparator Comparator[T])](<#func-sort>)
- [func SortBy[T any](slice []T, comparator Comparator[T]) []T](<#func-sortby>)
- [func Swap[T any](slice []T, i, j int)](<#func-swap>)
- [type BiFunction](<#type-bifunction>)
- [type Comparator](<#type-comparator>)
  - [func FunctionComparator[T, R any](transform Function[T, R], comparator Comparator[R]) Comparator[T]](<#func-functioncomparator>)
  - [func OrderedComparator[T constraints.Ordered]\(\) Comparator[T]](<#func-orderedcomparator>)
  - [func ReverseComparator[T any](comparator Comparator[T]) Comparator[T]](<#func-reversecomparator>)
- [type ComparedOrder](<#type-comparedorder>)
- [type Function](<#type-function>)
- [type Heap](<#type-heap>)
  - [func NewHeap[T any](comparator Comparator[T]) *Heap[T]](<#func-newheap>)
  - [func (h *Heap[T]) Len() int](<#func-heap-len>)
  - [func (h *Heap[T]) Pop() T](<#func-heap-pop>)
  - [func (h *Heap[T]) Push(v T)](<#func-heap-push>)
  - [func (h *Heap[T]) PushAll(values ...T)](<#func-heap-pushall>)
- [type KeyFor](<#type-keyfor>)
- [type KeyValueFor](<#type-keyvaluefor>)
- [type Predicate](<#type-predicate>)
- [type Stringer](<#type-stringer>)
  - [func StringerStringer[T fmt.Stringer]\(\) Stringer[T]](<#func-stringerstringer>)
- [type ValueFor](<#type-valuefor>)


## func All

```go
func All[T any](slice []T, predicate Predicate[T]) bool
```

All returns true if all elements of slice match the predicate\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

func main() {
	numbers := []float32{1, 2.2, 3.0, 4}
	positive := func(i float32) bool { return i > 0 }
	fmt.Println(genfuncs.All(numbers, positive)) // true
}
```

</p>
</details>

## func Any

```go
func Any[T any](slice []T, predicate Predicate[T]) bool
```

Any returns true if any element of the slice matches the predicate\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

func main() {
	fruits := []string{"apple", "banana", "grape"}
	isApple := func(fruit string) bool { return fruit == "apple" }
	isPear := func(fruit string) bool { return fruit == "pear" }
	fmt.Println(genfuncs.Any(fruits, isApple)) // true
	fmt.Println(genfuncs.Any(fruits, isPear))  // false
}
```

</p>
</details>

## func Associate

```go
func Associate[T, V any, K comparable](slice []T, keyValueFor KeyValueFor[T, K, V]) map[K]V
```

Associate returns a map containing key/values created by applying a function to elements of the slice\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"strings"
)

func main() {
	byLastName := func(n string) (string, string) {
		parts := strings.Split(n, " ")
		return parts[1], n
	}
	names := []string{"fred flintstone", "barney rubble"}
	nameMap := genfuncs.Associate(names, byLastName)
	fmt.Println(nameMap["rubble"]) // barney rubble
}
```

</p>
</details>

## func AssociateWith

```go
func AssociateWith[K comparable, V any](slice []K, valueFor ValueFor[K, V]) map[K]V
```

AssociateWith returns a Map where keys are elements from the given sequence and values are produced by the valueSelector function applied to each element\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

func main() {
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
```

</p>
</details>

## func Contains

```go
func Contains[T comparable](slice []T, element T) bool
```

Contains returns true if element is found in slice\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

func main() {
	values := []float32{1.0, .5, 42}
	fmt.Println(genfuncs.Contains(values, .5))    // true
	fmt.Println(genfuncs.Contains(values, 3.142)) // false
}
```

</p>
</details>

## func Distinct

```go
func Distinct[T comparable](slice []T) []T
```

Distinct returns a slice containing only distinct elements from the given slice\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

func main() {
	values := []int{1, 2, 2, 3, 1, 3}
	fmt.Println(genfuncs.Distinct(values)) // [1 2 3]
}
```

</p>
</details>

## func Filter

```go
func Filter[T any](slice []T, predicate Predicate[T]) []T
```

Filter returns a slice containing only elements matching the given predicate\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

func main() {
	values := []int{1, -2, 2, -3}
	isPositive := func(i int) bool { return i > 0 }
	fmt.Println(genfuncs.Filter(values, isPositive)) // [1 2]
}
```

</p>
</details>

## func Find

```go
func Find[T any](slice []T, predicate Predicate[T]) (T, bool)
```

Find returns the first element matching the given predicate and true\, or false when no such element was found\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

func main() {
	values := []int{-1, -2, 2, -3}
	isPositive := func(i int) bool { return i > 0 }
	fmt.Println(genfuncs.Find(values, isPositive)) // 2 true
}
```

</p>
</details>

## func FindLast

```go
func FindLast[T any](slice []T, predicate Predicate[T]) (T, bool)
```

FindLast returns the last element matching the given predicate and true\, or false when no such element was found\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

func main() {
	values := []int{-1, -2, 2, 3}
	isPositive := func(i int) bool { return i > 0 }
	fmt.Println(genfuncs.FindLast(values, isPositive)) // 3 true
}
```

</p>
</details>

## func FlatMap

```go
func FlatMap[T, R any](slice []T, function Function[T, []R]) []R
```

FlatMap returns a slice of all elements from results of transform function being invoked on each element of original slice\, and those resultant slices concatenated\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"strings"
)

func main() {
	words := []string{"hello", " ", "world"}
	slicer := func(s string) []string { return strings.Split(s, "") }
	fmt.Println(genfuncs.FlatMap(words, slicer)) // [h e l l o   w o r l d]
}
```

</p>
</details>

## func Fold

```go
func Fold[T, R any](slice []T, initial R, biFunction BiFunction[R, T, R]) R
```

Fold accumulates a value starting with initial value and applying operation from left to right to current accumulated value and each element\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := func(a int, b int) int { return a + b }
	fmt.Println(genfuncs.Fold(numbers, 0, sum)) // 15
}
```

</p>
</details>

## func GroupBy

```go
func GroupBy[T any, K comparable](slice []T, keyFor KeyFor[T, K]) map[K][]T
```

GroupBy groups elements of the slice by the key returned by the given keySelector function applied to each element and returns a map where each group key is associated with a slice of corresponding elements\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

func main() {
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
```

</p>
</details>

## func JoinToString

```go
func JoinToString[T any](slice []T, stringer Stringer[T], separator string, prefix string, postfix string) string
```

JoinToString creates a string from all the elements using the stringer on each\, separating them using separator\, and using the given prefix and postfix\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"strconv"
)

func main() {
	values := []bool{true, false, true}
	fmt.Println(genfuncs.JoinToString(
		values,
		strconv.FormatBool,
		", ",
		"{",
		"}",
	)) // {true, false, true}
}
```

</p>
</details>

## func Map

```go
func Map[T, R any](slice []T, function Function[T, R]) []R
```

Map returns a slice containing the results of applying the given transform function to each element in the original slice\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

func main() {
	numbers := []int{69, 88, 65, 77, 80, 76, 69}
	toString := func(i int) string { return string(rune(i)) }
	fmt.Println(genfuncs.Map(numbers, toString)) // [E X A M P L E]
}
```

</p>
</details>

## func Sort

```go
func Sort[T any](slice []T, comparator Comparator[T])
```

Sort sorts a slice by Comparator order\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"strings"
)

var alphaOrder = genfuncs.OrderedComparator[string]()

func main() {
	letters := strings.Split("example", "")

	genfuncs.Sort(letters, alphaOrder)
	fmt.Println(letters) // [a e e l m p x]
	genfuncs.Sort(letters, genfuncs.ReverseComparator(alphaOrder))
	fmt.Println(letters) // [x p m l e e a]
}
```

</p>
</details>

## func SortBy

```go
func SortBy[T any](slice []T, comparator Comparator[T]) []T
```

SortBy copies a slice\, sorts the copy applying the Comparator and returns it\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

func main() {
	numbers := []int{1, 0, 9, 6, 0}
	sorted := genfuncs.SortBy(numbers, genfuncs.OrderedComparator[int]())
	fmt.Println(numbers) // [1 0 9 6 0]
	fmt.Println(sorted)  // [0 0 1 6 9]
}
```

</p>
</details>

## func Swap

```go
func Swap[T any](slice []T, i, j int)
```

Swap two values in the slice\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

func main() {
	words := []string{"world", "hello"}
	genfuncs.Swap(words, 0, 1)
	fmt.Println(words) // [hello world]
}
```

</p>
</details>

## type BiFunction

BiFunction accepts two arguments and produces a result\.

```go
type BiFunction[T, U, R any] func(T, U) R
```

## type Comparator

Comparator compares two arguments of the same type and returns LessThan\, EqualTo or GreaterThan based relative order\.

```go
type Comparator[T any] BiFunction[T, T, ComparedOrder]
```

### func FunctionComparator

```go
func FunctionComparator[T, R any](transform Function[T, R], comparator Comparator[R]) Comparator[T]
```

FunctionComparator composites an existing Comparator\[R\] and Function\[T\,R\] into a new Comparator\[T\]\.

### func OrderedComparator

```go
func OrderedComparator[T constraints.Ordered]() Comparator[T]
```

OrderedComparator will create a Comparator from any type included in the constraints\.Ordered constraint\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

var (
	lexicalOrder   = genfuncs.OrderedComparator[string]()
	reverseLexical = genfuncs.ReverseComparator(lexicalOrder)
)

func main() {
	fmt.Println(lexicalOrder("a", "b")) // -1
	fmt.Println(lexicalOrder("a", "a")) // 0
	fmt.Println(lexicalOrder("b", "a")) // 1
}
```

</p>
</details>

### func ReverseComparator

```go
func ReverseComparator[T any](comparator Comparator[T]) Comparator[T]
```

ReverseComparator reverses a Comparator to facilitate switching sort orderings\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

var (
	lexicalOrder   = genfuncs.OrderedComparator[string]()
	reverseLexical = genfuncs.ReverseComparator(lexicalOrder)
)

func main() {
	fmt.Println(lexicalOrder("a", "b"))   // -1
	fmt.Println(reverseLexical("a", "b")) // 1
}
```

</p>
</details>

## type ComparedOrder

ComparedOrder is the type returned by a Comparator\.

```go
type ComparedOrder int
```

```go
var (
    LessThan    ComparedOrder = -1
    EqualTo     ComparedOrder = 0
    GreaterThan ComparedOrder = 1
)
```

## type Function

Function accepts one argument and produces a result\.

```go
type Function[T, R any] func(T) R
```

## type Heap

Heap implements either a min or max ordered heap of any type\.

```go
type Heap[T any] struct {
    // contains filtered or unexported fields
}
```

### func NewHeap

```go
func NewHeap[T any](comparator Comparator[T]) *Heap[T]
```

NewHeap return a heap ordered based on the Comparator\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

var (
	ascendingOrder = genfuncs.OrderedComparator[int]()
)

func main() {
	heap := genfuncs.NewHeap(ascendingOrder)
	heap.PushAll(3, 1, 4, 2)
	for heap.Len() > 0 {
		fmt.Print(heap.Pop()) // 1234
	}
	fmt.Println()
}
```

</p>
</details>

### func \(\*Heap\) Len

```go
func (h *Heap[T]) Len() int
```

Len returns current length of the heap\.

### func \(\*Heap\) Pop

```go
func (h *Heap[T]) Pop() T
```

Pop an item off the heap\.

### func \(\*Heap\) Push

```go
func (h *Heap[T]) Push(v T)
```

Push a value onto the heap\.

### func \(\*Heap\) PushAll

```go
func (h *Heap[T]) PushAll(values ...T)
```

PushAll the values onto the Heap\.

## type KeyFor

KeyFor is used for generating keys from types\, it accepts any type and returns a comparable key for it\.

```go
type KeyFor[T any, K comparable] Function[T, K]
```

## type KeyValueFor

KeyValueFor is used to generate a key and value from a type\, it accepts any type\, and returns a comparable key and any value\.

```go
type KeyValueFor[T any, K comparable, V any] func(T) (K, V)
```

## type Predicate

Predicate is used evaluate a value\, it accepts any type and returns a bool\.

```go
type Predicate[T any] func(T) bool
```

## type Stringer

Stringer is used to create string representations\, it accepts any type and returns a string\.

```go
type Stringer[T any] func(T) string
```

### func StringerStringer

```go
func StringerStringer[T fmt.Stringer]() Stringer[T]
```

StringerStringer creates a Stringer for any type that implements fmt\.Stringer\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"time"
)

func main() {
	var epoch time.Time
	fmt.Println(epoch.String()) // 0001-01-01 00:00:00 +0000 UTC
	stringer := genfuncs.StringerStringer[time.Time]()
	fmt.Println(stringer(epoch)) // 0001-01-01 00:00:00 +0000 UTC
}
```

</p>
</details>

## type ValueFor

ValueFor given a comparable key will return a value for it\.

```go
type ValueFor[K comparable, T any] Function[K, T]
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
