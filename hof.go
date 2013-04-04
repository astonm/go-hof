// Package hof is an experimental implementation of common higher-order functions in Go.
// This package requires a version of Go that includes reflect.MakeFunc and reflect.MakeSlice.
package hof

import "reflect"

func _map(in []reflect.Value) []reflect.Value {
	f := in[0]
	args := in[1]

	outType := reflect.SliceOf(f.Type().Out(0))

	out := reflect.MakeSlice(outType, args.Len(), args.Len())
	for i := 0; i < args.Len(); i++ {
		ret := f.Call([]reflect.Value{args.Index(i)})
		out.Index(i).Set(ret[0])
	}

	return []reflect.Value{out}
}

func filter(in []reflect.Value) []reflect.Value {
	f := in[0]
	args := in[1]

	out := reflect.MakeSlice(args.Type(), 0, args.Len())
	for i := 0; i < args.Len(); i++ {
		val := args.Index(i)
		shouldInclude := f.Call([]reflect.Value{val})[0]

		if shouldInclude.Bool() {
			out = reflect.Append(out, val)
		}
	}

	return []reflect.Value{out}
}

func reduce(in []reflect.Value) []reflect.Value {
	f := in[0]
	args := in[1]

	outType := f.Type().Out(0)
	out := reflect.Zero(outType)

	haveInit := false
	if len(in) > 2 {
		out = in[2]
		haveInit = true
	}

	for i := 0; i < args.Len(); i++ {
		if !haveInit && i == 0 {
			out = args.Index(i)
		} else {
			out = f.Call([]reflect.Value{out, args.Index(i)})[0]
		}
	}

	return []reflect.Value{out}
}

/*
MakeMapFunc takes pointer to a map function zero value and substitutes in the appropriate implementation.
Generally, such a function will take the following form:
    var mapper func (func (A) B, []A) []B
*/
func MakeMapFunc(mapFnPtr interface{}) {
	f := reflect.ValueOf(mapFnPtr).Elem()
	f.Set(reflect.MakeFunc(f.Type(), _map))
}

/*
MakeFilterFunc takes pointer to a filter function zero value and substitutes in the appropriate implementation.
Generally, such a function will take the following form:
    var filter func (func (A) bool, []A) []A
*/
func MakeFilterFunc(filterFnPtr interface{}) {
	f := reflect.ValueOf(filterFnPtr).Elem()
	f.Set(reflect.MakeFunc(f.Type(), filter))
}

/*
MakeFilterFunc takes pointer to a filter function zero value and substitutes in the appropriate implementation.
Generally, such a function will take one of the following forms:
    var reduce func (func (A, A) A, []A) A // takes two args: the reducer and the slice

    var reduce func (func (A, B) A, []B, A) A // takes three args: the reducer, the slice, and an initial value
*/
func MakeReduceFunc(reduceFnPtr interface{}) {
	f := reflect.ValueOf(reduceFnPtr).Elem()
	f.Set(reflect.MakeFunc(f.Type(), reduce))
}
