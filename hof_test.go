package hof

import "testing"
import "reflect"

func TestIntMap(t *testing.T) {
	var mapper func(func(int) int, []int) []int
	MakeMapFunc(&mapper)

	in := []int{1, 2, 3, 4, 5}
	f := func(x int) int { return x * 2 }
	exp := []int{2, 4, 6, 8, 10}

	out := mapper(f, in)
	if !reflect.DeepEqual(exp, out) {
		t.Fatal("expected", exp, ", got", out)
	}
}

func TestStringIntMap(t *testing.T) {
	var mapper func(func(string) int, []string) []int
	MakeMapFunc(&mapper)

	in := []string{"try", "this", "thing"}
	f := func(x string) int { return len(x) }
	exp := []int{3, 4, 5}

	out := mapper(f, in)

	if !reflect.DeepEqual(exp, out) {
		t.Fatal("expected", exp, ", got", out)
	}
}

func TestIntStringMap(t *testing.T) {
	var mapper func(func(int) string, []int) []string
	MakeMapFunc(&mapper)

	in := []int{1, 2, 3}
	exp := []string{"x", "xx", "xxx"}
	f := func(x int) string {
		out := ""
		for i := 0; i < x; i++ {
			out += "x"
		}
		return out
	}

	out := mapper(f, in)

	if !reflect.DeepEqual(exp, out) {
		t.Fatal("expected", exp, ", got", out)
	}
}

func TestEmptyMap(t *testing.T) {
	var mapper func(func(int) int, []int) []int
	MakeMapFunc(&mapper)

	in := []int{}
	f := func(x int) int { return x * 2 }
	exp := []int{}

	out := mapper(f, in)
	if !reflect.DeepEqual(exp, out) {
		t.Fatal("expected", exp, ", got", out)
	}
}

func TestFilter(t *testing.T) {
	var filter func(func(int) bool, []int) []int
	MakeFilterFunc(&filter)

	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := func(x int) bool { return x%2 == 0 }
	exp := []int{2, 4, 6, 8, 10}

	out := filter(f, in)
	if !reflect.DeepEqual(exp, out) {
		t.Fatal("expected", exp, ", got", out)
	}
}

func TestReduce(t *testing.T) {
	var reduce func(func(int, int) int, []int) int
	MakeReduceFunc(&reduce)

	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := func(x, y int) int { return x + y }
	exp := 55

	out := reduce(f, in)
	if exp != out {
		t.Fatal("expected", exp, ", got", out)
	}
}

func TestReduceInit(t *testing.T) {
	var reduce func(func(int, int) int, []int, int) int
	MakeReduceFunc(&reduce)

	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := func(x, y int) int { return x + y }
	init := 45
	exp := 100

	out := reduce(f, in, init)
	if exp != out {
		t.Fatal("expected", exp, ", got", out)
	}
}

func TestReduceTwoTypes(t *testing.T) {
	var reduce func(func(string, int) string, []int, string) string
	MakeReduceFunc(&reduce)

	in := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	f := func(x string, y int) string { return x + string('0'+y) }
	init := ""
	exp := "0123456789"

	out := reduce(f, in, init)
	if exp != out {
		t.Fatal("expected", exp, ", got", out)
	}
}

func sliceRange(n int) []int {
	out := make([]int, n, n)
	for i := 0; i < n; i++ {
		out[i] = i
	}
	return out
}

var benchmarkIn = sliceRange(100)

func BenchmarkForMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := len(benchmarkIn)
		out := make([]int, l, l)
		for i := 0; i < l; i++ {
			out[i] = benchmarkIn[i] * 2
		}
		if out[1] != 2 {
			panic("wrong result")
		}
	}
}

func BenchmarkMakeMapFunc(b *testing.B) {
	var mapper func(func(int) int, []int) []int
	MakeMapFunc(&mapper)

	for n := 0; n < b.N; n++ {
		f := func(x int) int { return x * 2 }
		out := mapper(f, benchmarkIn)
		if out[1] != 2 {
			panic("wrong result")
		}
	}
}

func interfaceMapper(f func(interface{}) interface{}, in []interface{}) []interface{} {
	out := make([]interface{}, len(in), len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}

func BenchmarkInterfaceMapFunc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := len(benchmarkIn)

		interfaceIn := make([]interface{}, l, l)
		for i, x := range benchmarkIn {
			interfaceIn[i] = x
		}

		f := func(x interface{}) interface{} {
			return x.(int) * 2
		}

		interfaceOut := interfaceMapper(f, interfaceIn)

		out := make([]int, l, l)
		for i, n := range interfaceOut {
			out[i] = n.(int)
		}
	}
}
