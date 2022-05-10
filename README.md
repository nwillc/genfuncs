<!-- Code generated by gomarkdoc. DO NOT EDIT -->

[![License](https://img.shields.io/github/license/nwillc/genfuncs.svg)](https://tldrlegal.com/license/-isc-license)
[![CI](https://github.com/nwillc/genfuncs/workflows/CI/badge.svg)](https://github.com/nwillc/genfuncs/actions/workflows/CI.yml)
[![codecov.io](https://codecov.io/github/nwillc/genfuncs/coverage.svg?branch=master)](https://codecov.io/github/nwillc/genfuncs?branch=master)
[![goreportcard.com](https://goreportcard.com/badge/github.com/nwillc/genfuncs)](https://goreportcard.com/report/github.com/nwillc/genfuncs)
[![Reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/nwillc/genfuncs)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Releases](https://img.shields.io/github/tag/nwillc/genfuncs.svg)](https://github.com/nwillc/genfuncs/tags)

# Genfuncs

Genfuncs implements various functions utilizing Go's Generics to help avoid writing boilerplate code,
in particular when working with containers like heaps, maps, queues, sets, slices, etc. Many of the functions are 
based on Kotlin's Sequence and Map. This package, while very usable, is primarily a proof-of-concept since it is likely 
Go will provide similar before long. In fact, golang.org/x/exp/slices and golang.org/x/exp/maps offer some similar 
functions and I incorporate them here.

Examples are found in `*examples_test.go` files or projects like [gordle](https://github.com/nwillc/gordle).

The code is under the [ISC License](https://github.com/nwillc/genfuncs/blob/master/LICENSE.md).

# Requirements

Build with Go 1.18+

# Getting

```bash
go get github.com/nwillc/genfuncs
```

# Packages
- [genfuncs](<#genfuncs>)
- [genfuncs/container](<#container>)

# genfuncs

```go
import "github.com/nwillc/genfuncs"
```

## Index

- [Variables](<#variables>)
- [func EqualOrder[O constraints.Ordered](a, b O) bool](<#func-equalorder>)
- [func GreaterOrdered[O constraints.Ordered](a, b O) bool](<#func-greaterordered>)
- [func LessOrdered[O constraints.Ordered](a, b O) bool](<#func-lessordered>)
- [func Max[T constraints.Ordered](v ...T) T](<#func-max>)
- [func Min[T constraints.Ordered](v ...T) T](<#func-min>)
- [func Order[T constraints.Ordered](a, b T) int](<#func-order>)
- [type BiFunction](<#type-bifunction>)
  - [func TransformArgs[T1, T2, R any](transform Function[T1, T2], operation BiFunction[T2, T2, R]) BiFunction[T1, T1, R]](<#func-transformargs>)
- [type Function](<#type-function>)
  - [func Curried[A, R any](operation BiFunction[A, A, R], a A) Function[A, R]](<#func-curried>)
  - [func IsEqualOrdered[O constraints.Ordered](a O) Function[O, bool]](<#func-isequalordered>)
  - [func IsGreaterOrdered[O constraints.Ordered](a O) Function[O, bool]](<#func-isgreaterordered>)
  - [func IsLessOrdered[O constraints.Ordered](a O) Function[O, bool]](<#func-islessordered>)
  - [func Not[T any](predicate Function[T, bool]) Function[T, bool]](<#func-not>)
- [type MapKeyFor](<#type-mapkeyfor>)
- [type MapKeyValueFor](<#type-mapkeyvaluefor>)
- [type MapValueFor](<#type-mapvaluefor>)
- [type ToString](<#type-tostring>)
  - [func StringerToString[T fmt.Stringer]() ToString[T]](<#func-stringertostring>)


## Variables

```go
var (
    // Orderings
    OrderedLess    = -1
    OrderedEqual   = 0
    OrderedGreater = 1

    // Predicates
    IsBlank    = IsEqualOrdered("")
    IsNotBlank = Not(IsBlank)
    F32IsZero  = IsEqualOrdered(float32(0.0))
    F64IsZero  = IsEqualOrdered(0.0)
    IIsZero    = IsEqualOrdered(0)
)
```

```go
var IllegalArguments = fmt.Errorf("illegal arguments")
```

NoSuchElement error is used by panics when attempts are made to access out of bounds\.

```go
var NoSuchElement = fmt.Errorf("no such element")
```

## func [EqualOrder](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L39>)

```go
func EqualOrder[O constraints.Ordered](a, b O) bool
```

EqualOrder tests if constraints\.Ordered a equal to b\.

## func [GreaterOrdered](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L49>)

```go
func GreaterOrdered[O constraints.Ordered](a, b O) bool
```

GreaterOrdered tests if constraints\.Ordered a is greater than b\.

## func [LessOrdered](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L59>)

```go
func LessOrdered[O constraints.Ordered](a, b O) bool
```

LessOrdered tests if constraints\.Ordered a is less than b\.

## func [Max](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L69>)

```go
func Max[T constraints.Ordered](v ...T) T
```

Max returns max value one or more constraints\.Ordered values\,

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

func main() {
	fmt.Println(genfuncs.Max(1, 2))
	words := []string{"dog", "cat", "gorilla"}
	fmt.Println(genfuncs.Max(words...))
}
```

#### Output

```
2
gorilla
```

</p>
</details>

## func [Min](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L83>)

```go
func Min[T constraints.Ordered](v ...T) T
```

Min returns min value of one or more constraints\.Ordered values\,

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
)

func main() {
	fmt.Println(genfuncs.Min(1, 2))
	words := []string{"dog", "cat", "gorilla"}
	fmt.Println(genfuncs.Min(words...))
}
```

#### Output

```
1
cat
```

</p>
</details>

## func [Order](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L119>)

```go
func Order[T constraints.Ordered](a, b T) int
```

Order old school \-1/0/1 order of constraints\.Ordered\.

## type [BiFunction](<https://github.com/nwillc/genfuncs/blob/master/functions.go#L20>)

BiFunction accepts two arguments and produces a result\.

```go
type BiFunction[T, U, R any] func(T, U) R
```

### func [TransformArgs](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L102>)

```go
func TransformArgs[T1, T2, R any](transform Function[T1, T2], operation BiFunction[T2, T2, R]) BiFunction[T1, T1, R]
```

TransformArgs uses the function to transform the arguments to be passed to the operation\.

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
	var unixTime = func(t time.Time) int64 { return t.Unix() }
	var chronoOrder = genfuncs.TransformArgs(unixTime, genfuncs.LessOrdered[int64])
	now := time.Now()
	fmt.Println(chronoOrder(now, now.Add(time.Second)))
}
```

#### Output

```
true
```

</p>
</details>

## type [Function](<https://github.com/nwillc/genfuncs/blob/master/functions.go#L23>)

Function is a single argument function\.

```go
type Function[T, R any] func(T) R
```

### func [Curried](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L109>)

```go
func Curried[A, R any](operation BiFunction[A, A, R], a A) Function[A, R]
```

Curried takes a BiFunction and one argument\, and Curries the function to return a single argument Function\.

### func [IsEqualOrdered](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L44>)

```go
func IsEqualOrdered[O constraints.Ordered](a O) Function[O, bool]
```

IsEqualOrdered return a EqualOrder for a\.

### func [IsGreaterOrdered](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L54>)

```go
func IsGreaterOrdered[O constraints.Ordered](a O) Function[O, bool]
```

IsGreaterOrdered returns a function that returns true if its argument is greater than a\.

### func [IsLessOrdered](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L64>)

```go
func IsLessOrdered[O constraints.Ordered](a O) Function[O, bool]
```

IsLessOrdered returns a function that returns true if its argument is less than a\.

### func [Not](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L114>)

```go
func Not[T any](predicate Function[T, bool]) Function[T, bool]
```

Not takes a predicate returning and inverts the result\.

## type [MapKeyFor](<https://github.com/nwillc/genfuncs/blob/master/functions.go#L26>)

MapKeyFor is used for generating keys from types\, it accepts any type and returns a comparable key for it\.

```go
type MapKeyFor[T any, K comparable] func(T) K
```

## type [MapKeyValueFor](<https://github.com/nwillc/genfuncs/blob/master/functions.go#L30>)

MapKeyValueFor is used to generate a key and value from a type\, it accepts any type\, and returns a comparable key and any value\.

```go
type MapKeyValueFor[T any, K comparable, V any] func(T) (K, V)
```

## type [MapValueFor](<https://github.com/nwillc/genfuncs/blob/master/functions.go#L33>)

MapValueFor given a comparable key will return a value for it\.

```go
type MapValueFor[K comparable, T any] func(K) T
```

## type [ToString](<https://github.com/nwillc/genfuncs/blob/master/functions.go#L36>)

ToString is used to create string representations\, it accepts any type and returns a string\.

```go
type ToString[T any] func(T) string
```

### func [StringerToString](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L97>)

```go
func StringerToString[T fmt.Stringer]() ToString[T]
```

StringerToString creates a ToString for any type that implements fmt\.Stringer\.

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
	fmt.Println(epoch.String())
	stringer := genfuncs.StringerToString[time.Time]()
	fmt.Println(stringer(epoch))
}
```

#### Output

```
0001-01-01 00:00:00 +0000 UTC
0001-01-01 00:00:00 +0000 UTC
```

</p>
</details>

# container

```go
import "github.com/nwillc/genfuncs/container"
```

## Index

- [type Container](<#type-container>)
- [type Deque](<#type-deque>)
  - [func NewDeque[T any](t ...T) *Deque[T]](<#func-newdeque>)
  - [func (d *Deque[T]) Add(t T)](<#func-dequet-add>)
  - [func (d *Deque[T]) AddAll(t ...T)](<#func-dequet-addall>)
  - [func (d *Deque[T]) AddLeft(t T)](<#func-dequet-addleft>)
  - [func (d *Deque[T]) AddRight(t T)](<#func-dequet-addright>)
  - [func (d *Deque[T]) Cap() int](<#func-dequet-cap>)
  - [func (d *Deque[T]) Len() int](<#func-dequet-len>)
  - [func (d *Deque[T]) Peek() T](<#func-dequet-peek>)
  - [func (d *Deque[T]) PeekLeft() T](<#func-dequet-peekleft>)
  - [func (d *Deque[T]) PeekRight() T](<#func-dequet-peekright>)
  - [func (d *Deque[T]) Remove() T](<#func-dequet-remove>)
  - [func (d *Deque[T]) RemoveLeft() T](<#func-dequet-removeleft>)
  - [func (d *Deque[T]) RemoveRight() T](<#func-dequet-removeright>)
  - [func (d *Deque[T]) Values() GSlice[T]](<#func-dequet-values>)
- [type GMap](<#type-gmap>)
  - [func (m GMap[K, V]) All(predicate genfuncs.Function[V, bool]) bool](<#func-gmapk-v-all>)
  - [func (m GMap[K, V]) Any(predicate genfuncs.Function[V, bool]) bool](<#func-gmapk-v-any>)
  - [func (m GMap[K, V]) Contains(key K) bool](<#func-gmapk-v-contains>)
  - [func (m GMap[K, V]) Filter(predicate genfuncs.Function[V, bool]) GMap[K, V]](<#func-gmapk-v-filter>)
  - [func (m GMap[K, V]) FilterKeys(predicate genfuncs.Function[K, bool]) GMap[K, V]](<#func-gmapk-v-filterkeys>)
  - [func (m GMap[K, V]) ForEach(action func(k K, v V))](<#func-gmapk-v-foreach>)
  - [func (m GMap[K, V]) GetOrElse(k K, defaultValue func() V) V](<#func-gmapk-v-getorelse>)
  - [func (m GMap[K, V]) Keys() GSlice[K]](<#func-gmapk-v-keys>)
  - [func (m GMap[K, V]) Len() int](<#func-gmapk-v-len>)
  - [func (m GMap[K, V]) Values() GSlice[V]](<#func-gmapk-v-values>)
- [type GSlice](<#type-gslice>)
  - [func (s GSlice[T]) All(predicate genfuncs.Function[T, bool]) bool](<#func-gslicet-all>)
  - [func (s GSlice[T]) Any(predicate genfuncs.Function[T, bool]) bool](<#func-gslicet-any>)
  - [func (s GSlice[T]) Compare(s2 GSlice[T], comparison genfuncs.BiFunction[T, T, int]) int](<#func-gslicet-compare>)
  - [func (s GSlice[T]) Equal(s2 GSlice[T], comparison genfuncs.BiFunction[T, T, int]) bool](<#func-gslicet-equal>)
  - [func (s GSlice[T]) Filter(predicate genfuncs.Function[T, bool]) GSlice[T]](<#func-gslicet-filter>)
  - [func (s GSlice[T]) Find(predicate genfuncs.Function[T, bool]) (T, bool)](<#func-gslicet-find>)
  - [func (s GSlice[T]) FindLast(predicate genfuncs.Function[T, bool]) (T, bool)](<#func-gslicet-findlast>)
  - [func (s GSlice[T]) ForEach(action func(i int, t T))](<#func-gslicet-foreach>)
  - [func (s GSlice[T]) JoinToString(stringer genfuncs.ToString[T], separator string, prefix string, postfix string) string](<#func-gslicet-jointostring>)
  - [func (s GSlice[T]) Len() int](<#func-gslicet-len>)
  - [func (s GSlice[T]) Random() T](<#func-gslicet-random>)
  - [func (s GSlice[T]) SortBy(lessThan genfuncs.BiFunction[T, T, bool]) GSlice[T]](<#func-gslicet-sortby>)
  - [func (s GSlice[T]) Swap(i, j int)](<#func-gslicet-swap>)
  - [func (s GSlice[T]) Values() GSlice[T]](<#func-gslicet-values>)
- [type HasValues](<#type-hasvalues>)
- [type Heap](<#type-heap>)
  - [func NewHeap[T any](compare genfuncs.BiFunction[T, T, bool], values ...T) *Heap[T]](<#func-newheap>)
  - [func (h *Heap[T]) Add(v T)](<#func-heapt-add>)
  - [func (h *Heap[T]) AddAll(values ...T)](<#func-heapt-addall>)
  - [func (h *Heap[T]) Len() int](<#func-heapt-len>)
  - [func (h *Heap[T]) Peek() T](<#func-heapt-peek>)
  - [func (h *Heap[T]) Remove() T](<#func-heapt-remove>)
  - [func (h *Heap[T]) Values() GSlice[T]](<#func-heapt-values>)
- [type MapSet](<#type-mapset>)
  - [func (h *MapSet[T]) Add(t T)](<#func-mapsett-add>)
  - [func (h *MapSet[T]) AddAll(t ...T)](<#func-mapsett-addall>)
  - [func (h *MapSet[T]) Contains(t T) bool](<#func-mapsett-contains>)
  - [func (h *MapSet[T]) Len() int](<#func-mapsett-len>)
  - [func (h *MapSet[T]) Remove(t T)](<#func-mapsett-remove>)
  - [func (h *MapSet[T]) Values() GSlice[T]](<#func-mapsett-values>)
- [type Queue](<#type-queue>)
- [type Set](<#type-set>)
  - [func NewMapSet[T comparable](t ...T) Set[T]](<#func-newmapset>)


## type [Container](<https://github.com/nwillc/genfuncs/blob/master/container/container.go#L20-L26>)

Container is a minimal container that HasValues and accepts additional elements\.

```go
type Container[T any] interface {

    // Add an element to the Container.
    Add(t T)
    // AddAll elements to the Container.
    AddAll(t ...T)
    // contains filtered or unexported methods
}
```

## type [Deque](<https://github.com/nwillc/genfuncs/blob/master/container/deque.go#L24-L30>)

Deque is a doubly ended implementation of Queue with default behavior of a Fifo but provides left and right access\.

```go
type Deque[T any] struct {
    // contains filtered or unexported fields
}
```

### func [NewDeque](<https://github.com/nwillc/genfuncs/blob/master/container/deque.go#L33>)

```go
func NewDeque[T any](t ...T) *Deque[T]
```

NewDeque creates a Deque containing any provided elements\.

### func \(\*Deque\[T\]\) [Add](<https://github.com/nwillc/genfuncs/blob/master/container/deque.go#L40>)

```go
func (d *Deque[T]) Add(t T)
```

Add an element to the right of the Deque\.

### func \(\*Deque\[T\]\) [AddAll](<https://github.com/nwillc/genfuncs/blob/master/container/deque.go#L45>)

```go
func (d *Deque[T]) AddAll(t ...T)
```

AddAll elements to the right of the Deque\.

### func \(\*Deque\[T\]\) [AddLeft](<https://github.com/nwillc/genfuncs/blob/master/container/deque.go#L52>)

```go
func (d *Deque[T]) AddLeft(t T)
```

AddLeft an element to the left of the Deque\.

### func \(\*Deque\[T\]\) [AddRight](<https://github.com/nwillc/genfuncs/blob/master/container/deque.go#L60>)

```go
func (d *Deque[T]) AddRight(t T)
```

AddRight an element to the right of the Deque\.

### func \(\*Deque\[T\]\) [Cap](<https://github.com/nwillc/genfuncs/blob/master/container/deque.go#L123>)

```go
func (d *Deque[T]) Cap() int
```

Cap returns the capacity of the Deque\.

### func \(\*Deque\[T\]\) [Len](<https://github.com/nwillc/genfuncs/blob/master/container/deque.go#L68>)

```go
func (d *Deque[T]) Len() int
```

Len reports the length of the Deque\.

### func \(\*Deque\[T\]\) [Peek](<https://github.com/nwillc/genfuncs/blob/master/container/deque.go#L73>)

```go
func (d *Deque[T]) Peek() T
```

Peek returns the left most element in the Deque without removing it\.

### func \(\*Deque\[T\]\) [PeekLeft](<https://github.com/nwillc/genfuncs/blob/master/container/deque.go#L78>)

```go
func (d *Deque[T]) PeekLeft() T
```

PeekLeft returns the left most element in the Deque without removing it\.

### func \(\*Deque\[T\]\) [PeekRight](<https://github.com/nwillc/genfuncs/blob/master/container/deque.go#L84>)

```go
func (d *Deque[T]) PeekRight() T
```

PeekRight returns the right most element in the Deque without removing it\.

### func \(\*Deque\[T\]\) [Remove](<https://github.com/nwillc/genfuncs/blob/master/container/deque.go#L91>)

```go
func (d *Deque[T]) Remove() T
```

Remove and return the left most element in the Deque\.

### func \(\*Deque\[T\]\) [RemoveLeft](<https://github.com/nwillc/genfuncs/blob/master/container/deque.go#L96>)

```go
func (d *Deque[T]) RemoveLeft() T
```

RemoveLeft and return the left most element in the Deque\.

### func \(\*Deque\[T\]\) [RemoveRight](<https://github.com/nwillc/genfuncs/blob/master/container/deque.go#L106>)

```go
func (d *Deque[T]) RemoveRight() T
```

RemoveRight and return the right most element in the Deque\.

### func \(\*Deque\[T\]\) [Values](<https://github.com/nwillc/genfuncs/blob/master/container/deque.go#L116>)

```go
func (d *Deque[T]) Values() GSlice[T]
```

Values in the Deque returned in a new GSlice\.

## type [GMap](<https://github.com/nwillc/genfuncs/blob/master/container/gmap.go#L27>)

GMap is a generic type corresponding to a standard Go map and implements HasValues\.

```go
type GMap[K comparable, V any] map[K]V
```

### func \(GMap\[K\, V\]\) [All](<https://github.com/nwillc/genfuncs/blob/master/container/gmap.go#L30>)

```go
func (m GMap[K, V]) All(predicate genfuncs.Function[V, bool]) bool
```

All returns true if all values in GMap satisfy the predicate\.

### func \(GMap\[K\, V\]\) [Any](<https://github.com/nwillc/genfuncs/blob/master/container/gmap.go#L40>)

```go
func (m GMap[K, V]) Any(predicate genfuncs.Function[V, bool]) bool
```

Any returns true if any values in GMap satisfy the predicate\.

### func \(GMap\[K\, V\]\) [Contains](<https://github.com/nwillc/genfuncs/blob/master/container/gmap.go#L50>)

```go
func (m GMap[K, V]) Contains(key K) bool
```

Contains returns true if the GMap contains the given key\.

### func \(GMap\[K\, V\]\) [Filter](<https://github.com/nwillc/genfuncs/blob/master/container/gmap.go#L56>)

```go
func (m GMap[K, V]) Filter(predicate genfuncs.Function[V, bool]) GMap[K, V]
```

Filter a GMap by a predicate\, returning a new GMap that contains only values that satisfy the predicate\.

### func \(GMap\[K\, V\]\) [FilterKeys](<https://github.com/nwillc/genfuncs/blob/master/container/gmap.go#L68>)

```go
func (m GMap[K, V]) FilterKeys(predicate genfuncs.Function[K, bool]) GMap[K, V]
```

FilterKeys returns a new GMap that contains only values whose key satisfy the predicate\.

### func \(GMap\[K\, V\]\) [ForEach](<https://github.com/nwillc/genfuncs/blob/master/container/gmap.go#L80>)

```go
func (m GMap[K, V]) ForEach(action func(k K, v V))
```

ForEach performs the given action on each entry in the GMap\.

### func \(GMap\[K\, V\]\) [GetOrElse](<https://github.com/nwillc/genfuncs/blob/master/container/gmap.go#L87>)

```go
func (m GMap[K, V]) GetOrElse(k K, defaultValue func() V) V
```

GetOrElse returns the value at the given key if it exists or returns the result of defaultValue\.

### func \(GMap\[K\, V\]\) [Keys](<https://github.com/nwillc/genfuncs/blob/master/container/gmap.go#L96>)

```go
func (m GMap[K, V]) Keys() GSlice[K]
```

Keys return a GSlice containing the keys of the GMap\.

### func \(GMap\[K\, V\]\) [Len](<https://github.com/nwillc/genfuncs/blob/master/container/gmap.go#L101>)

```go
func (m GMap[K, V]) Len() int
```

Len is the number of elements in the GMap\.

### func \(GMap\[K\, V\]\) [Values](<https://github.com/nwillc/genfuncs/blob/master/container/gmap.go#L106>)

```go
func (m GMap[K, V]) Values() GSlice[V]
```

Values returns a GSlice of all the values in the GMap\.

## type [GSlice](<https://github.com/nwillc/genfuncs/blob/master/container/gslice.go#L33>)

GSlice is a generic type corresponding to a standard Go slice that implements HasValues\.

```go
type GSlice[T any] []T
```

### func \(GSlice\[T\]\) [All](<https://github.com/nwillc/genfuncs/blob/master/container/gslice.go#L36>)

```go
func (s GSlice[T]) All(predicate genfuncs.Function[T, bool]) bool
```

All returns true if all elements of slice match the predicate\.

### func \(GSlice\[T\]\) [Any](<https://github.com/nwillc/genfuncs/blob/master/container/gslice.go#L46>)

```go
func (s GSlice[T]) Any(predicate genfuncs.Function[T, bool]) bool
```

Any returns true if any element of the slice matches the predicate\.

### func \(GSlice\[T\]\) [Compare](<https://github.com/nwillc/genfuncs/blob/master/container/gslice.go#L57>)

```go
func (s GSlice[T]) Compare(s2 GSlice[T], comparison genfuncs.BiFunction[T, T, int]) int
```

Compare one GSlice to another\, applying a comparison to each pair of corresponding entries\. Compare returns 0 if all the pair's match\, \-1 if this GSlice is less\, or 1 if it is greater\.

### func \(GSlice\[T\]\) [Equal](<https://github.com/nwillc/genfuncs/blob/master/container/gslice.go#L63>)

```go
func (s GSlice[T]) Equal(s2 GSlice[T], comparison genfuncs.BiFunction[T, T, int]) bool
```

Equal compares this GSlice to another\, applying a comparison to each pair\, if the lengths are equal and all the values are then true is returned\.

### func \(GSlice\[T\]\) [Filter](<https://github.com/nwillc/genfuncs/blob/master/container/gslice.go#L68>)

```go
func (s GSlice[T]) Filter(predicate genfuncs.Function[T, bool]) GSlice[T]
```

Filter returns a slice containing only elements matching the given predicate\.

### func \(GSlice\[T\]\) [Find](<https://github.com/nwillc/genfuncs/blob/master/container/gslice.go#L79>)

```go
func (s GSlice[T]) Find(predicate genfuncs.Function[T, bool]) (T, bool)
```

Find returns the first element matching the given predicate and true\, or false when no such element was found\.

### func \(GSlice\[T\]\) [FindLast](<https://github.com/nwillc/genfuncs/blob/master/container/gslice.go#L90>)

```go
func (s GSlice[T]) FindLast(predicate genfuncs.Function[T, bool]) (T, bool)
```

FindLast returns the last element matching the given predicate and true\, or false when no such element was found\.

### func \(GSlice\[T\]\) [ForEach](<https://github.com/nwillc/genfuncs/blob/master/container/gslice.go#L104>)

```go
func (s GSlice[T]) ForEach(action func(i int, t T))
```

ForEach element of the GSlice invoke given function with the element\. Syntactic sugar for a range that intends to traverse all the elements\, i\.e\. no exiting midway through\.

### func \(GSlice\[T\]\) [JoinToString](<https://github.com/nwillc/genfuncs/blob/master/container/gslice.go#L112>)

```go
func (s GSlice[T]) JoinToString(stringer genfuncs.ToString[T], separator string, prefix string, postfix string) string
```

JoinToString creates a string from all the elements using the stringer on each\, separating them using separator\, and using the given prefix and postfix\.

### func \(GSlice\[T\]\) [Len](<https://github.com/nwillc/genfuncs/blob/master/container/gslice.go#L127>)

```go
func (s GSlice[T]) Len() int
```

Len is the number of elements in the GSlice\.

### func \(GSlice\[T\]\) [Random](<https://github.com/nwillc/genfuncs/blob/master/container/gslice.go#L132>)

```go
func (s GSlice[T]) Random() T
```

Random returns a random element of the GSlice\.

### func \(GSlice\[T\]\) [SortBy](<https://github.com/nwillc/genfuncs/blob/master/container/gslice.go#L137>)

```go
func (s GSlice[T]) SortBy(lessThan genfuncs.BiFunction[T, T, bool]) GSlice[T]
```

SortBy copies a slice\, sorts the copy applying the Order and returns it\.

### func \(GSlice\[T\]\) [Swap](<https://github.com/nwillc/genfuncs/blob/master/container/gslice.go#L145>)

```go
func (s GSlice[T]) Swap(i, j int)
```

Swap two values in the slice\.

### func \(GSlice\[T\]\) [Values](<https://github.com/nwillc/genfuncs/blob/master/container/gslice.go#L150>)

```go
func (s GSlice[T]) Values() GSlice[T]
```

Values is the GSlice itself\.

## type [HasValues](<https://github.com/nwillc/genfuncs/blob/master/container/has_values.go#L20-L25>)

HasValues is an interface that indicates a struct contains values that can counted and be retrieved\.

```go
type HasValues[T any] interface {
    // Len returns length of the Container.
    Len() int
    // Values returns a copy of the current values in the Container without modifying the contents.
    Values() GSlice[T]
}
```

## type [Heap](<https://github.com/nwillc/genfuncs/blob/master/container/heap.go#L28-L32>)

Heap implements an ordered heap of any type which can be min heap or max heap depending on the compare provided\. Heap implements Queue\.

```go
type Heap[T any] struct {
    // contains filtered or unexported fields
}
```

### func [NewHeap](<https://github.com/nwillc/genfuncs/blob/master/container/heap.go#L35>)

```go
func NewHeap[T any](compare genfuncs.BiFunction[T, T, bool], values ...T) *Heap[T]
```

NewHeap return a heap ordered based on the compare and adds any values provided\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
)

func main() {
	heap := container.NewHeap[int](genfuncs.LessOrdered[int], 3, 1, 4, 2)
	for heap.Len() > 0 {
		fmt.Print(heap.Remove())
	}
	fmt.Println()
}
```

#### Output

```
1234
```

</p>
</details>

### func \(\*Heap\[T\]\) [Add](<https://github.com/nwillc/genfuncs/blob/master/container/heap.go#L45>)

```go
func (h *Heap[T]) Add(v T)
```

Add a value onto the heap\.

### func \(\*Heap\[T\]\) [AddAll](<https://github.com/nwillc/genfuncs/blob/master/container/heap.go#L52>)

```go
func (h *Heap[T]) AddAll(values ...T)
```

AddAll the values onto the Heap\.

### func \(\*Heap\[T\]\) [Len](<https://github.com/nwillc/genfuncs/blob/master/container/heap.go#L42>)

```go
func (h *Heap[T]) Len() int
```

Len returns current length of the heap\.

### func \(\*Heap\[T\]\) [Peek](<https://github.com/nwillc/genfuncs/blob/master/container/heap.go#L61>)

```go
func (h *Heap[T]) Peek() T
```

Peek returns the next element without removing it\.

### func \(\*Heap\[T\]\) [Remove](<https://github.com/nwillc/genfuncs/blob/master/container/heap.go#L76>)

```go
func (h *Heap[T]) Remove() T
```

Remove an item off the heap\.

### func \(\*Heap\[T\]\) [Values](<https://github.com/nwillc/genfuncs/blob/master/container/heap.go#L84>)

```go
func (h *Heap[T]) Values() GSlice[T]
```

Values returns a slice of the values in the Heap in no particular order\.

## type [MapSet](<https://github.com/nwillc/genfuncs/blob/master/container/map_set.go#L25-L27>)

MapSet is a Set implementation based on a map\. MapSet implements Set\.

```go
type MapSet[T comparable] struct {
    // contains filtered or unexported fields
}
```

### func \(\*MapSet\[T\]\) [Add](<https://github.com/nwillc/genfuncs/blob/master/container/map_set.go#L37>)

```go
func (h *MapSet[T]) Add(t T)
```

Add element to MapSet\.

### func \(\*MapSet\[T\]\) [AddAll](<https://github.com/nwillc/genfuncs/blob/master/container/map_set.go#L42>)

```go
func (h *MapSet[T]) AddAll(t ...T)
```

AddAll elements to MapSet\.

### func \(\*MapSet\[T\]\) [Contains](<https://github.com/nwillc/genfuncs/blob/master/container/map_set.go#L49>)

```go
func (h *MapSet[T]) Contains(t T) bool
```

Contains returns true if MapSet contains element\.

### func \(\*MapSet\[T\]\) [Len](<https://github.com/nwillc/genfuncs/blob/master/container/map_set.go#L55>)

```go
func (h *MapSet[T]) Len() int
```

Len returns the length of the MapSet\.

### func \(\*MapSet\[T\]\) [Remove](<https://github.com/nwillc/genfuncs/blob/master/container/map_set.go#L60>)

```go
func (h *MapSet[T]) Remove(t T)
```

Remove an element from the MapSet\.

### func \(\*MapSet\[T\]\) [Values](<https://github.com/nwillc/genfuncs/blob/master/container/map_set.go#L65>)

```go
func (h *MapSet[T]) Values() GSlice[T]
```

Values returns the elements in the MapSet as a GSlice\.

## type [Queue](<https://github.com/nwillc/genfuncs/blob/master/container/queue.go#L20-L25>)

Queue is a container providing some define order when accessing elements\. Queue implements Container\.

```go
type Queue[T any] interface {

    // Peek returns the next element without removing it.
    Peek() T
    Remove() T
    // contains filtered or unexported methods
}
```

## type [Set](<https://github.com/nwillc/genfuncs/blob/master/container/set.go#L20-L25>)

Set is a Container that contains no duplicate elements\.

```go
type Set[T comparable] interface {

    // Contains returns true if the Set contains a given element.
    Contains(t T) bool
    Remove(T)
    // contains filtered or unexported methods
}
```

### func [NewMapSet](<https://github.com/nwillc/genfuncs/blob/master/container/map_set.go#L30>)

```go
func NewMapSet[T comparable](t ...T) Set[T]
```

NewMapSet returns a new Set containing given values\.

# gmaps

```go
import "github.com/nwillc/genfuncs/container/gmaps"
```

## Index

- [func Map[K comparable, V any, R any](m container.GMap[K, V], transform genfuncs.BiFunction[K, V, R]) container.GSlice[R]](<#func-map>)
- [func MapMerge[K comparable, V any](mv ...container.GMap[K, container.GSlice[V]]) container.GMap[K, container.GSlice[V]]](<#func-mapmerge>)


## func [Map](<https://github.com/nwillc/genfuncs/blob/master/container/gmaps/gmap_functiions.go#L25>)

```go
func Map[K comparable, V any, R any](m container.GMap[K, V], transform genfuncs.BiFunction[K, V, R]) container.GSlice[R]
```

Map returns a GSlice containing the results of applying the given transform function to each element in the GMap\.

## func [MapMerge](<https://github.com/nwillc/genfuncs/blob/master/container/gmaps/gmap_functiions.go#L36>)

```go
func MapMerge[K comparable, V any](mv ...container.GMap[K, container.GSlice[V]]) container.GMap[K, container.GSlice[V]]
```

MapMerge merges maps of container\.GSlice's together into a new map appending the container\.GSlice's when collisions occur\.

# gslices

```go
import "github.com/nwillc/genfuncs/container/gslices"
```

## Index

- [func Associate[T, V any, K comparable](slice container.GSlice[T], keyValueFor genfuncs.MapKeyValueFor[T, K, V]) container.GMap[K, V]](<#func-associate>)
- [func AssociateWith[T comparable, V any](slice container.GSlice[T], valueFor genfuncs.MapValueFor[T, V]) container.GMap[T, V]](<#func-associatewith>)
- [func Distinct[T comparable](slice container.GSlice[T]) container.GSlice[T]](<#func-distinct>)
- [func FlatMap[T, R any](slice container.GSlice[T], transform genfuncs.Function[T, container.GSlice[R]]) container.GSlice[R]](<#func-flatmap>)
- [func Fold[T, R any](slice container.GSlice[T], initial R, operation genfuncs.BiFunction[R, T, R]) R](<#func-fold>)
- [func GroupBy[T any, K comparable](slice container.GSlice[T], keyFor genfuncs.MapKeyFor[T, K]) container.GMap[K, container.GSlice[T]]](<#func-groupby>)
- [func Map[T, R any](slice container.GSlice[T], transform genfuncs.Function[T, R]) container.GSlice[R]](<#func-map>)
- [func ToSet[T comparable](slice container.GSlice[T]) container.Set[T]](<#func-toset>)


## func [Associate](<https://github.com/nwillc/genfuncs/blob/master/container/gslices/gslice_functions.go#L25>)

```go
func Associate[T, V any, K comparable](slice container.GSlice[T], keyValueFor genfuncs.MapKeyValueFor[T, K, V]) container.GMap[K, V]
```

Associate returns a map containing key/values created by applying a function to elements of the slice\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs/container/gslices"
	"strings"
)

func main() {
	byLastName := func(n string) (string, string) {
		parts := strings.Split(n, " ")
		return parts[1], n
	}
	names := []string{"fred flintstone", "barney rubble"}
	nameMap := gslices.Associate(names, byLastName)
	fmt.Println(nameMap["rubble"])
}
```

#### Output

```
barney rubble
```

</p>
</details>

## func [AssociateWith](<https://github.com/nwillc/genfuncs/blob/master/container/gslices/gslice_functions.go#L36>)

```go
func AssociateWith[T comparable, V any](slice container.GSlice[T], valueFor genfuncs.MapValueFor[T, V]) container.GMap[T, V]
```

AssociateWith returns a Map where keys are elements from the given sequence and values are produced by the valueSelector function applied to each element\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs/container/gslices"
)

func main() {
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
}
```

#### Output

```
EVEN
ODD
```

</p>
</details>

## func [Distinct](<https://github.com/nwillc/genfuncs/blob/master/container/gslices/gslice_functions.go#L46>)

```go
func Distinct[T comparable](slice container.GSlice[T]) container.GSlice[T]
```

Distinct returns a slice containing only distinct elements from the given slice\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs/container/gslices"
)

func main() {
	values := []int{1, 2, 2, 3, 1, 3}
	gslices.Distinct(values).ForEach(func(_, i int) {
		fmt.Println(i)
	})
}
```

#### Output

```
1
2
3
```

</p>
</details>

## func [FlatMap](<https://github.com/nwillc/genfuncs/blob/master/container/gslices/gslice_functions.go#L52>)

```go
func FlatMap[T, R any](slice container.GSlice[T], transform genfuncs.Function[T, container.GSlice[R]]) container.GSlice[R]
```

FlatMap returns a slice of all elements from results of transform being invoked on each element of original slice\, and those resultant slices concatenated\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/container/gslices"
	"strings"
)

var words container.GSlice[string] = []string{"hello", "world"}

func main() {
	slicer := func(s string) container.GSlice[string] { return strings.Split(s, "") }
	fmt.Println(gslices.FlatMap(words.SortBy(genfuncs.LessOrdered[string]), slicer))
}
```

#### Output

```
[h e l l o w o r l d]
```

</p>
</details>

## func [Fold](<https://github.com/nwillc/genfuncs/blob/master/container/gslices/gslice_functions.go#L62>)

```go
func Fold[T, R any](slice container.GSlice[T], initial R, operation genfuncs.BiFunction[R, T, R]) R
```

Fold accumulates a value starting with initial value and applying operation from left to right to current accumulated value and each element\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs/container/gslices"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := func(a int, b int) int { return a + b }
	fmt.Println(gslices.Fold(numbers, 0, sum))
}
```

#### Output

```
15
```

</p>
</details>

## func [GroupBy](<https://github.com/nwillc/genfuncs/blob/master/container/gslices/gslice_functions.go#L72>)

```go
func GroupBy[T any, K comparable](slice container.GSlice[T], keyFor genfuncs.MapKeyFor[T, K]) container.GMap[K, container.GSlice[T]]
```

GroupBy groups elements of the slice by the key returned by the given keySelector function applied to each element and returns a map where each group key is associated with a slice of corresponding elements\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs/container/gslices"
)

func main() {
	oddEven := func(i int) string {
		if i%2 == 0 {
			return "EVEN"
		}
		return "ODD"
	}
	numbers := []int{1, 2, 3, 4}
	grouped := gslices.GroupBy(numbers, oddEven)
	fmt.Println(grouped["ODD"])
}
```

#### Output

```
[1 3]
```

</p>
</details>

## func [Map](<https://github.com/nwillc/genfuncs/blob/master/container/gslices/gslice_functions.go#L83>)

```go
func Map[T, R any](slice container.GSlice[T], transform genfuncs.Function[T, R]) container.GSlice[R]
```

Map returns a new container\.GSlice containing the results of applying the given transform function to each element in the original slice\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/nwillc/genfuncs/container/gslices"
)

func main() {
	numbers := []int{69, 88, 65, 77, 80, 76, 69}
	toString := func(i int) string { return string(rune(i)) }
	fmt.Println(gslices.Map(numbers, toString))
}
```

#### Output

```
[E X A M P L E]
```

</p>
</details>

## func [ToSet](<https://github.com/nwillc/genfuncs/blob/master/container/gslices/gslice_functions.go#L92>)

```go
func ToSet[T comparable](slice container.GSlice[T]) container.Set[T]
```

ToSet creates a Set from the elements of the GSlice\.

# version

```go
import "github.com/nwillc/genfuncs/gen/version"
```

## Index

- [Constants](<#constants>)


## Constants

Version number for official releases\.

```go
const Version = "v0.12.0"
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
