<!-- Code generated by gomarkdoc. DO NOT EDIT -->

[![License](https://img.shields.io/github/license/nwillc/genfuncs.svg)](https://tldrlegal.com/license/-isc-license)
[![CI](https://github.com/nwillc/genfuncs/workflows/CI/badge.svg)](https://github.com/nwillc/genfuncs/actions/workflows/CI.yml)
[![codecov.io](https://codecov.io/github/nwillc/genfuncs/coverage.svg?branch=master)](https://codecov.io/github/nwillc/genfuncs?branch=master)
[![goreportcard.com](https://goreportcard.com/badge/github.com/nwillc/genfuncs)](https://goreportcard.com/report/github.com/nwillc/genfuncs)
[![Reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/nwillc/genfuncs)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Releases](https://img.shields.io/github/tag/nwillc/genfuncs.svg)](https://github.com/nwillc/genfuncs/tags)

# Genfuncs Package

Genfuncs implements various functions utilizing Go's Generics to help avoid writing boilerplate code,
in particular when working with containers like heap, list, map, queue, set, slice, etc. Many of the functions are
based on Kotlin's Sequence and Map. Some functional patterns like Result and Promises are presents. Attempts were also 
made to introduce more polymorphism into Go's containers. This package, while very usable, is primarily a 
proof-of-concept since it is likely Go will provide similar before long. In fact, golang.org/x/exp/slices 
and golang.org/x/exp/maps offer some similar functions and I incorporate them here.

## Code Style

The coding style is not always idiomatic, in particular:

 - All functions have named return values and those variable are used in the return statements.
 - Some places where the `range` build-in could have been used instead use explicit indexing.

Both of these, while less idiomatic, were done because they measurably improve performance.

## General notes:
 - A Map interface is provided to allow both Go's normal map and it's sync.Map to be used polymorphically.
 - The bias of these functions where appropriate is to be pure, without side effects, at the cost of copying data.
 - Examples are found in `*examples_test.go` files or projects like [gordle](https://github.com/nwillc/gordle), 
[gorelease](https://github.com/nwillc/gorelease) or [gotimer](https://github.com/nwillc/gotimer).

## License

The code is under the [ISC License](https://github.com/nwillc/genfuncs/blob/master/LICENSE.md).

## Requirements

Build with Go 1.18+

## Getting

```bash
go get github.com/nwillc/genfuncs
```

# Packages
- [genfuncs](<#genfuncs>)
- [genfuncs/container](<#container>)
- [genfuncs/container/gmaps](<#gmaps>)
- [genfuncs/container/gslices](<#gslices>)
- [genfuncs/container/sequences](<#sequences>)
- [genfuncs/promises](<#promises>)
- [genfuncs/results](<#results>)


# genfuncs

```go
import "github.com/nwillc/genfuncs"
```

## Index

- [Variables](<#variables>)
- [func Empty[T any]() (empty T)](<#func-empty>)
- [func Max[T constraints.Ordered](v ...T) (max T)](<#func-max>)
- [func Min[T constraints.Ordered](v ...T) (min T)](<#func-min>)
- [func Ordered[T constraints.Ordered](a, b T) (order int)](<#func-ordered>)
- [func OrderedEqual[O constraints.Ordered](a, b O) (orderedEqualTo bool)](<#func-orderedequal>)
- [func OrderedGreater[O constraints.Ordered](a, b O) (orderedGreaterThan bool)](<#func-orderedgreater>)
- [func OrderedLess[O constraints.Ordered](a, b O) (orderedLess bool)](<#func-orderedless>)
- [type BiFunction](<#type-bifunction>)
  - [func TransformArgs[T1, T2, R any](transform Function[T1, T2], operation BiFunction[T2, T2, R]) (fn BiFunction[T1, T1, R])](<#func-transformargs>)
- [type Function](<#type-function>)
  - [func Curried[A, R any](operation BiFunction[A, A, R], a A) (fn Function[A, R])](<#func-curried>)
  - [func Not[T any](predicate Function[T, bool]) (fn Function[T, bool])](<#func-not>)
  - [func OrderedEqualTo[O constraints.Ordered](a O) (fn Function[O, bool])](<#func-orderedequalto>)
  - [func OrderedGreaterThan[O constraints.Ordered](a O) (fn Function[O, bool])](<#func-orderedgreaterthan>)
  - [func OrderedLessThan[O constraints.Ordered](a O) (fn Function[O, bool])](<#func-orderedlessthan>)
- [type Promise](<#type-promise>)
  - [func NewPromise[T any](ctx context.Context, action func(context.Context) *Result[T]) *Promise[T]](<#func-newpromise>)
  - [func NewPromiseFromResult[T any](result *Result[T]) *Promise[T]](<#func-newpromisefromresult>)
  - [func (p *Promise[T]) Cancel()](<#func-promiset-cancel>)
  - [func (p *Promise[T]) OnError(action func(error)) *Promise[T]](<#func-promiset-onerror>)
  - [func (p *Promise[T]) OnSuccess(action func(t T)) *Promise[T]](<#func-promiset-onsuccess>)
  - [func (p *Promise[T]) Wait() *Result[T]](<#func-promiset-wait>)
- [type Result](<#type-result>)
  - [func NewError[T any](err error) *Result[T]](<#func-newerror>)
  - [func NewResult[T any](t T) *Result[T]](<#func-newresult>)
  - [func NewResultError[T any](t T, err error) *Result[T]](<#func-newresulterror>)
  - [func (r *Result[T]) Error() error](<#func-resultt-error>)
  - [func (r *Result[T]) MustGet() T](<#func-resultt-mustget>)
  - [func (r *Result[T]) Ok() bool](<#func-resultt-ok>)
  - [func (r *Result[T]) OnError(action func(e error)) *Result[T]](<#func-resultt-onerror>)
  - [func (r *Result[T]) OnSuccess(action func(t T)) *Result[T]](<#func-resultt-onsuccess>)
  - [func (r *Result[T]) OrElse(v T) T](<#func-resultt-orelse>)
  - [func (r *Result[T]) OrEmpty() T](<#func-resultt-orempty>)
  - [func (r *Result[T]) String() string](<#func-resultt-string>)
  - [func (r *Result[T]) Then(action func(t T) *Result[T]) *Result[T]](<#func-resultt-then>)
- [type ToString](<#type-tostring>)
  - [func StringerToString[T fmt.Stringer]() (fn ToString[T])](<#func-stringertostring>)


## Variables

```go
var (
    // Orderings
    LessThan    = -1
    EqualTo     = 0
    GreaterThan = 1

    // Predicates
    IsBlank    = OrderedEqualTo("")
    IsNotBlank = Not(IsBlank)
    F32IsZero  = OrderedEqualTo(float32(0.0))
    F64IsZero  = OrderedEqualTo(0.0)
    IIsZero    = OrderedEqualTo(0)
)
```

```go
var (
    // PromiseNoActionErrorMsg indicates a Promise was created for no action.
    PromiseNoActionErrorMsg = "promise requested with no action"
    // PromisePanicErrorMsg indicates the action of a Promise caused a panic.
    PromisePanicErrorMsg = "promise action panic"
)
```

```go
var IllegalArguments = fmt.Errorf("illegal arguments")
```

NoSuchElement error is used by panics when attempts are made to access out of bounds.

```go
var NoSuchElement = fmt.Errorf("no such element")
```

## func [Empty](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L39>)

```go
func Empty[T any]() (empty T)
```

Empty return an empty value of type T.

## func [Max](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L80>)

```go
func Max[T constraints.Ordered](v ...T) (max T)
```

Max returns max value one or more constraints.Ordered values,

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

## func [Min](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L94>)

```go
func Min[T constraints.Ordered](v ...T) (min T)
```

Min returns min value of one or more constraints.Ordered values,

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

## func [Ordered](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L134>)

```go
func Ordered[T constraints.Ordered](a, b T) (order int)
```

Ordered performs old school \-1/0/1 comparison of constraints.Ordered arguments.

## func [OrderedEqual](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L44>)

```go
func OrderedEqual[O constraints.Ordered](a, b O) (orderedEqualTo bool)
```

OrderedEqual returns true jf a is ordered equal to b.

## func [OrderedGreater](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L56>)

```go
func OrderedGreater[O constraints.Ordered](a, b O) (orderedGreaterThan bool)
```

OrderedGreater returns true if a is ordered greater than b.

## func [OrderedLess](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L68>)

```go
func OrderedLess[O constraints.Ordered](a, b O) (orderedLess bool)
```

OrderedLess returns true if a is ordered less than b.

## type [BiFunction](<https://github.com/nwillc/genfuncs/blob/master/functions.go#L20>)

BiFunction accepts two arguments and produces a result.

```go
type BiFunction[T, U, R any] func(T, U) R
```

### func [TransformArgs](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L114>)

```go
func TransformArgs[T1, T2, R any](transform Function[T1, T2], operation BiFunction[T2, T2, R]) (fn BiFunction[T1, T1, R])
```

TransformArgs uses the function to transform the arguments to be passed to the operation.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"time"

	"github.com/nwillc/genfuncs"
)

func main() {
	var unixTime = func(t time.Time) int64 { return t.Unix() }
	var chronoOrder = genfuncs.TransformArgs(unixTime, genfuncs.OrderedLess[int64])
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

Function is a single argument function.

```go
type Function[T, R any] func(T) R
```

### func [Curried](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L122>)

```go
func Curried[A, R any](operation BiFunction[A, A, R], a A) (fn Function[A, R])
```

Curried takes a BiFunction and one argument, and Curries the function to return a single argument Function.

### func [Not](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L128>)

```go
func Not[T any](predicate Function[T, bool]) (fn Function[T, bool])
```

Not takes a predicate returning and inverts the result.

### func [OrderedEqualTo](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L50>)

```go
func OrderedEqualTo[O constraints.Ordered](a O) (fn Function[O, bool])
```

OrderedEqualTo return a function that returns true if its argument is ordered equal to a.

### func [OrderedGreaterThan](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L62>)

```go
func OrderedGreaterThan[O constraints.Ordered](a O) (fn Function[O, bool])
```

OrderedGreaterThan returns a function that returns true if its argument is ordered greater than a.

### func [OrderedLessThan](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L74>)

```go
func OrderedLessThan[O constraints.Ordered](a O) (fn Function[O, bool])
```

OrderedLessThan returns a function that returns true if its argument is ordered less than a.

## type [Promise](<https://github.com/nwillc/genfuncs/blob/master/promise.go#L26-L32>)

Promise provides asynchronous Result of an action.

```go
type Promise[T any] struct {
    // contains filtered or unexported fields
}
```

### func [NewPromise](<https://github.com/nwillc/genfuncs/blob/master/promise.go#L42>)

```go
func NewPromise[T any](ctx context.Context, action func(context.Context) *Result[T]) *Promise[T]
```

NewPromise creates a Promise for an action.

### func [NewPromiseFromResult](<https://github.com/nwillc/genfuncs/blob/master/promise.go#L77>)

```go
func NewPromiseFromResult[T any](result *Result[T]) *Promise[T]
```

NewPromiseFromResult returns a completed Promise with the specified result.

### func \(\*Promise\[T\]\) [Cancel](<https://github.com/nwillc/genfuncs/blob/master/promise.go#L84>)

```go
func (p *Promise[T]) Cancel()
```

Cancel the Promise which will allow any action that is listening on \`\<\-ctx.Done\(\)\` to complete.

### func \(\*Promise\[T\]\) [OnError](<https://github.com/nwillc/genfuncs/blob/master/promise.go#L89>)

```go
func (p *Promise[T]) OnError(action func(error)) *Promise[T]
```

OnError returns a new Promise with an error handler waiting on the original Promise.

### func \(\*Promise\[T\]\) [OnSuccess](<https://github.com/nwillc/genfuncs/blob/master/promise.go#L96>)

```go
func (p *Promise[T]) OnSuccess(action func(t T)) *Promise[T]
```

OnSuccess returns a new Promise with a success handler waiting on the original Promise.

### func \(\*Promise\[T\]\) [Wait](<https://github.com/nwillc/genfuncs/blob/master/promise.go#L103>)

```go
func (p *Promise[T]) Wait() *Result[T]
```

Wait on the completion of a Promise.

## type [Result](<https://github.com/nwillc/genfuncs/blob/master/result.go#L26-L29>)

Result is an implementation of the Maybe pattern. This is mostly for experimentation as it is a poor fit with Go's traditional idiomatic error handling.

```go
type Result[T any] struct {
    // contains filtered or unexported fields
}
```

### func [NewError](<https://github.com/nwillc/genfuncs/blob/master/result.go#L46>)

```go
func NewError[T any](err error) *Result[T]
```

NewError for an error.

### func [NewResult](<https://github.com/nwillc/genfuncs/blob/master/result.go#L36>)

```go
func NewResult[T any](t T) *Result[T]
```

NewResult for a value.

### func [NewResultError](<https://github.com/nwillc/genfuncs/blob/master/result.go#L41>)

```go
func NewResultError[T any](t T, err error) *Result[T]
```

NewResultError creates a Result from a value, error tuple.

### func \(\*Result\[T\]\) [Error](<https://github.com/nwillc/genfuncs/blob/master/result.go#L55>)

```go
func (r *Result[T]) Error() error
```

Error of the Result, nil if Ok\(\).

### func \(\*Result\[T\]\) [MustGet](<https://github.com/nwillc/genfuncs/blob/master/result.go#L112>)

```go
func (r *Result[T]) MustGet() T
```

MustGet returns the value of the Result if Ok\(\) or if not, panics with the error.

### func \(\*Result\[T\]\) [Ok](<https://github.com/nwillc/genfuncs/blob/master/result.go#L69>)

```go
func (r *Result[T]) Ok() bool
```

Ok returns the status of Result, is it ok, or an error.

### func \(\*Result\[T\]\) [OnError](<https://github.com/nwillc/genfuncs/blob/master/result.go#L74>)

```go
func (r *Result[T]) OnError(action func(e error)) *Result[T]
```

OnError performs the action if Result is not Ok\(\).

### func \(\*Result\[T\]\) [OnSuccess](<https://github.com/nwillc/genfuncs/blob/master/result.go#L82>)

```go
func (r *Result[T]) OnSuccess(action func(t T)) *Result[T]
```

OnSuccess performs action if Result is Ok\(\).

### func \(\*Result\[T\]\) [OrElse](<https://github.com/nwillc/genfuncs/blob/master/result.go#L99>)

```go
func (r *Result[T]) OrElse(v T) T
```

OrElse returns the value of the Result if Ok\(\), or the value v if not.

### func \(\*Result\[T\]\) [OrEmpty](<https://github.com/nwillc/genfuncs/blob/master/result.go#L107>)

```go
func (r *Result[T]) OrEmpty() T
```

OrEmpty will return the value of the Result or the empty value if Error.

### func \(\*Result\[T\]\) [String](<https://github.com/nwillc/genfuncs/blob/master/result.go#L90>)

```go
func (r *Result[T]) String() string
```

String returns a string representation of Result, either the value or error.

### func \(\*Result\[T\]\) [Then](<https://github.com/nwillc/genfuncs/blob/master/result.go#L60>)

```go
func (r *Result[T]) Then(action func(t T) *Result[T]) *Result[T]
```

Then performs the action on the Result.

## type [ToString](<https://github.com/nwillc/genfuncs/blob/master/functions.go#L26>)

ToString is used to create string representations, it accepts any type and returns a string.

```go
type ToString[T any] func(T) string
```

### func [StringerToString](<https://github.com/nwillc/genfuncs/blob/master/dry.go#L108>)

```go
func StringerToString[T fmt.Stringer]() (fn ToString[T])
```

StringerToString creates a ToString for any type that implements fmt.Stringer.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	"time"

	"github.com/nwillc/genfuncs"
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



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
