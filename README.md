# go-fp
![example workflow](https://github.com/j-dumbell/go-fp/actions/workflows/test-build.yml/badge.svg?branch=main)

A Golang library to enable functional programming (FP).  Includes utilities for function composition, currying,
collection functions, and fp-friendly equivalents of standard library functions.
```go
import (
    "strconv"
	fp "github.com/j-dumbell/go-fp"
)

var ints = []int{1, 2, 3}

func square(i int) int {
    return i * i
}

func sum(ints []int) int {
    total := 0
    for _, i := range ints {
        total += i
    }
    return total
}

// "14"
var x = fp.Pipe3(ints, fp.Map(square), sum, strconv.Itoa)
```


## Function composition

### `Flow<X>`
Composes X unary functions from left to right.  X = 2, ..., 5.
```go
func double(i int) int {
    return 2 * i
}

// func(int) string
var doubleThenString = fp.Flow2(double, strconv.Itoa)

// "4"
var flowed = doubleThenString(2) 
```

### `Pipe<X>`
Pipes a value through X unary functions from left to right.  X = 2, ..., 5.
```go
func add3(i int) int {
    return i + 3
}

// "8"
var piped = fp.Pipe2(5, add3, strconv.Itoa)
```

## Currying

### `Curry<X>`
Curries a function of arity X from left to right.  X = 2, 3, 4.
```go
func add(i, j int) int {
    return i + j
}

// func(int) func(int) int
var curriedAdd = fp.Curry2(add)

// 3
var oneAddTwo = curriedAdd(1)(2)
```

## Collection functions
All collection functions are curried with data last.

### `Map`
Creates a new slice from another slice by calling the provided unary callback function on each element.
```go
var numbers = []int{1, 2, 3}

// []string{"1", "2", "3"}
var toStrings = fp.Map(strconv.Itoa)(numbers) 
```

### `MapI`
As-per Map, but the provided callback function must an include a 2nd argument for the index of the slice element.
```go
var uints = []int{1, 2, 3}

// []string{1, 3, 5}
var result = fp.MapI(func(i int, index int) int {
    return i + index
})(uints)
```

## FP-friendly standard library functions 
Standard library equivalents are auto-curried and data last.  The package structure matches the standard library
structure, with the package name suffixed with `fp`.  The FP-friendly functions have exactly the same names.
E.g. for fp equivalents of functions from the `strings` package, use `"github.com/j-dumbell/go-fp/stringsfp"`.
```go
import (
    "github.com/j-dumbell/go-fp/stringsfp"
)

var data = []string{"1", "2", "3"}

// Equivalent to strings.Join(data, ",")
var joinedData = stringsfp.Join(",")(data)
```
