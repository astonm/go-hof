## Overview

go-hof is an experimental implementation of common higher-order functions in Go (`map`, `reduce` and `filter`). Go has no generics, so typically emulating this functionality would involve either operating on slices of `interface{}` (requiring type assertions to use the results) or writing similar code for each type of data being operated on, with the code differing only by the types involved.

This library uses Go's `reflect.MakeFunc` functionality to remove much of the boilerplate code involved in implementing a higher order function across multiple types. `reflect.MakeFunc` is unfortunately not available in Go 1.0, but is available as a part of [go1.1beta1](https://code.google.com/p/go/downloads/list?q=go1.1beta1).

## Installation
This code is go-gettable. Run `go get github.com/astonm/go-hof` and then, in your code, `import "github.com/astonm/go-hof"` to use the `hof` package.

## Usage
This package currently supports `map`, `filter` and `reduce`. Accordingly, it exports three functions to support creating higher order functions.

    func MakeMapFunc(mapPtr interface{})

    func MakeFilterFunc(filterPtr interface{})

    func MakeReduceFunc(reducePtr interface{})

Each method takes a pointer to a zero value of a properly-typed function for that operation and replaces that zero value with working code for that particular higher order function. Kind of a mouthful. Here's an example:

    var mapper func (func (int) int, []int) []int
    hof.MakeMapFunc(&mapper)

    in := []int{1,2,3,4,5}
    f := func(x int) int { return x*2 }
    fmt.Println(mapper(f, in)) // prints [2 4 6 8 10], a slice holding ints

The equivalent mapper without `hof` might be used this way (assuming a `mapper` written using `interface{}` based "generification":

    in := []int{1,2,3,4,5}
    f := func (x interface{}) interface{} {
        return x.(int)*2
    }
    interfaceIn := make([]interface{}, 0)
    for _, x := range(in) {
        interfaceIn = append(interfaceIn, x)
    }
    fmt.Println(mapper(f, interfaceIn)) // prints [2 4 6 8 10], a slice holding interfaces holding ints

Check out hof_test.go for a few more examples.

## Usage

The zero types for the higher order functions can get a bit messy, but the basic formulas for them are as follows (capital letters are types):

    // map
    var mapper func (func (A) B, []A) []B // map is already taken...

    // filter
    var filter func (func (A) bool, []A) []A

    // reduce
    var reduce func (func (A, A) A, []A) A

    // reduce with initial value passed as third parameter
    var reduce func (func (A, B) A, []B, A) A

At the moment only slices are supported.

## TODO

* Add sanity-checking and friendly error messages for the arguments to functions. Currently the code blows up with mysterious-seeming messaging if the types aren't correct.
* Add support for operations on arrays.
