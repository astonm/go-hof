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

func MakeMap(mapPtr interface{}) {
	f := reflect.ValueOf(mapPtr).Elem()
	f.Set(reflect.MakeFunc(f.Type(), _map))
}

func MakeFilter(filterPtr interface{}) {
	f := reflect.ValueOf(filterPtr).Elem()
	f.Set(reflect.MakeFunc(f.Type(), filter))
}

func MakeReduce(reducePtr interface{}){
	f := reflect.ValueOf(reducePtr).Elem()
	f.Set(reflect.MakeFunc(f.Type(), reduce))
}