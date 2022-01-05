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

- [func Associate[T, V any, K comparable](slice Slice[T], keyValueFor KeyValueFor[T, K, V]) map[K]V](<#func-associate>)
- [func AssociateWith[K comparable, V any](slice Slice[K], valueFor ValueFor[K, V]) map[K]V](<#func-associatewith>)
- [func Fold[T, R any](slice Slice[T], initial R, biFunction BiFunction[R, T, R]) R](<#func-fold>)
- [func GroupBy[T any, K comparable](slice Slice[T], keyFor KeyFor[T, K]) map[K]Slice[T]](<#func-groupby>)
- [type BiFunction](<#type-bifunction>)
- [type Function](<#type-function>)
- [type Heap](<#type-heap>)
  - [func NewHeap[T any](lessThan LessThan[T]) *Heap[T]](<#func-newheap>)
  - [func (h *Heap[T]) Len() int](<#func-heap-len>)
  - [func (h *Heap[T]) Pop() T](<#func-heap-pop>)
  - [func (h *Heap[T]) Push(v T)](<#func-heap-push>)
  - [func (h *Heap[T]) PushAll(values ...T)](<#func-heap-pushall>)
- [type KeyFor](<#type-keyfor>)
- [type KeyValueFor](<#type-keyvaluefor>)
- [type LessThan](<#type-lessthan>)
  - [func OrderedLessThan[T constraints.Ordered]\(\) LessThan[T]](<#func-orderedlessthan>)
  - [func Reverse[T any](lessThan LessThan[T]) LessThan[T]](<#func-reverse>)
  - [func TransformLessThan[T, R any](transform Function[T, R], lessThan LessThan[R]) LessThan[T]](<#func-transformlessthan>)
- [type Predicate](<#type-predicate>)
  - [func IsGreaterThan[T constraints.Ordered](a T) Predicate[T]](<#func-isgreaterthan>)
  - [func IsLessThan[T constraints.Ordered](a T) Predicate[T]](<#func-islessthan>)
  - [func (p Predicate[T]) Not() Predicate[T]](<#func-predicate-not>)
- [type Slice](<#type-slice>)
  - [func Distinct[T comparable](slice Slice[T]) Slice[T]](<#func-distinct>)
  - [func FlatMap[T, R any](slice Slice[T], function Function[T, Slice[R]]) Slice[R]](<#func-flatmap>)
  - [func Keys[K comparable, V any](m map[K]V) Slice[K]](<#func-keys>)
  - [func Map[T, R any](slice Slice[T], function Function[T, R]) Slice[R]](<#func-map>)
  - [func Values[K comparable, V any](m map[K]V) Slice[V]](<#func-values>)
  - [func (s Slice[T]) All(predicate Predicate[T]) bool](<#func-slice-all>)
  - [func (s Slice[T]) Any(predicate Predicate[T]) bool](<#func-slice-any>)
  - [func (s Slice[T]) Contains(element T, lessThan LessThan[T]) bool](<#func-slice-contains>)
  - [func (s Slice[T]) Filter(predicate Predicate[T]) Slice[T]](<#func-slice-filter>)
  - [func (s Slice[T]) Find(predicate Predicate[T]) (T, bool)](<#func-slice-find>)
  - [func (s Slice[T]) FindLast(predicate Predicate[T]) (T, bool)](<#func-slice-findlast>)
  - [func (s Slice[T]) JoinToString(stringer Stringer[T], separator string, prefix string, postfix string) string](<#func-slice-jointostring>)
  - [func (s Slice[T]) Sort(lessThan LessThan[T])](<#func-slice-sort>)
  - [func (s Slice[T]) SortBy(lessThan LessThan[T]) Slice[T]](<#func-slice-sortby>)
  - [func (s Slice[T]) Swap(i, j int)](<#func-slice-swap>)
- [type Stringer](<#type-stringer>)
  - [func StringerStringer[T fmt.Stringer]\(\) Stringer[T]](<#func-stringerstringer>)
- [type ValueFor](<#type-valuefor>)


## func Associate

```go
func Associate[T, V any, K comparable](slice Slice[T], keyValueFor KeyValueFor[T, K, V]) map[K]V
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
func AssociateWith[K comparable, V any](slice Slice[K], valueFor ValueFor[K, V]) map[K]V
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

## func Fold

```go
func Fold[T, R any](slice Slice[T], initial R, biFunction BiFunction[R, T, R]) R
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
func GroupBy[T any, K comparable](slice Slice[T], keyFor KeyFor[T, K]) map[K]Slice[T]
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

## type BiFunction

BiFunction accepts two arguments and produces a result\.

```go
type BiFunction[T, U, R any] func(T, U) R
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
func NewHeap[T any](lessThan LessThan[T]) *Heap[T]
```

NewHeap return a heap ordered based on the LessThan\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

var intCmp = genfuncs.OrderedLessThan[int]()

func main() {
	heap := genfuncs.NewHeap(intCmp)
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

## type LessThan

LessThan compares two arguments of the same type and returns true if the first is less than the second\.

```go
type LessThan[T any] BiFunction[T, T, bool]
```

### func OrderedLessThan

```go
func OrderedLessThan[T constraints.Ordered]() LessThan[T]
```

OrderedLessThan will create a LessThan from any type included in the constraints\.Ordered constraint\.

### func Reverse

```go
func Reverse[T any](lessThan LessThan[T]) LessThan[T]
```

Reverse reverses a LessThan to facilitate reverse sort ordering\.

### func TransformLessThan

```go
func TransformLessThan[T, R any](transform Function[T, R], lessThan LessThan[R]) LessThan[T]
```

TransformLessThan composites an existing LessThan\[R\] and transform Function\[T\,R\] into a new LessThan\[T\]\. The transform is used to convert the arguments before they are passed to the lessThan\.

## type Predicate

Predicate is used evaluate a value\, it accepts any type and returns a bool\.

```go
type Predicate[T any] func(T) bool
```

### func IsGreaterThan

```go
func IsGreaterThan[T constraints.Ordered](a T) Predicate[T]
```

IsGreaterThan creates a Predicate that tests if its argument is greater than a given value\.

### func IsLessThan

```go
func IsLessThan[T constraints.Ordered](a T) Predicate[T]
```

IsLessThan creates a Predicate that tests if its argument is less than a given value\.

### func \(Predicate\) Not

```go
func (p Predicate[T]) Not() Predicate[T]
```

## type Slice

```go
type Slice[T any] []T
```

### func Distinct

```go
func Distinct[T comparable](slice Slice[T]) Slice[T]
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

### func FlatMap

```go
func FlatMap[T, R any](slice Slice[T], function Function[T, Slice[R]]) Slice[R]
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

var lexicalOrder = genfuncs.OrderedLessThan[string]()

var words genfuncs.Slice[string] = []string{"hello", "world"}

func main() {
	slicer := func(s string) genfuncs.Slice[string] { return strings.Split(s, "") }
	fmt.Println(genfuncs.FlatMap(words.SortBy(lexicalOrder), slicer)) // [h e l l o w o r l d]
}
```

</p>
</details>

### func Keys

```go
func Keys[K comparable, V any](m map[K]V) Slice[K]
```

Keys returns a slice of all the keys in the map\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

var wordPositions = map[string]int{"hello": 1, "world": 2}

func main() {
	keys := genfuncs.Keys(wordPositions)
	fmt.Println(keys) // [hello, world]
}
```

</p>
</details>

### func Map

```go
func Map[T, R any](slice Slice[T], function Function[T, R]) Slice[R]
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

### func Values

```go
func Values[K comparable, V any](m map[K]V) Slice[V]
```

Values returns a slice of all the values in the map\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

var wordPositions = map[string]int{"hello": 1, "world": 2}

func main() {
	values := genfuncs.Values(wordPositions)
	fmt.Println(values) // [1, 2]
}
```

</p>
</details>

### func \(Slice\) All

```go
func (s Slice[T]) All(predicate Predicate[T]) bool
```

All returns true if all elements of slice match the predicate\.

### func \(Slice\) Any

```go
func (s Slice[T]) Any(predicate Predicate[T]) bool
```

Any returns true if any element of the slice matches the predicate\.

### func \(Slice\) Contains

```go
func (s Slice[T]) Contains(element T, lessThan LessThan[T]) bool
```

Contains returns true if element is found in slice\.

### func \(Slice\) Filter

```go
func (s Slice[T]) Filter(predicate Predicate[T]) Slice[T]
```

Filter returns a slice containing only elements matching the given predicate\.

### func \(Slice\) Find

```go
func (s Slice[T]) Find(predicate Predicate[T]) (T, bool)
```

Find returns the first element matching the given predicate and true\, or false when no such element was found\.

### func \(Slice\) FindLast

```go
func (s Slice[T]) FindLast(predicate Predicate[T]) (T, bool)
```

FindLast returns the last element matching the given predicate and true\, or false when no such element was found\.

### func \(Slice\) JoinToString

```go
func (s Slice[T]) JoinToString(stringer Stringer[T], separator string, prefix string, postfix string) string
```

JoinToString creates a string from all the elements using the stringer on each\, separating them using separator\, and using the given prefix and postfix\.

### func \(Slice\) Sort

```go
func (s Slice[T]) Sort(lessThan LessThan[T])
```

Sort sorts a slice by the LessThan order\.

### func \(Slice\) SortBy

```go
func (s Slice[T]) SortBy(lessThan LessThan[T]) Slice[T]
```

SortBy copies a slice\, sorts the copy applying the Comparator and returns it\.

### func \(Slice\) Swap

```go
func (s Slice[T]) Swap(i, j int)
```

Swap two values in the slice\.

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
