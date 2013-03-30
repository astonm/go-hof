package hof

import "testing"
import "reflect"

func TestIntMap(t *testing.T) {
	var mapper func (func (int) int, []int) []int
	MakeMap(&mapper)

	in := []int{1,2,3,4,5}
	f := func(x int) int { return x*2 }
	exp := []int{2,4,6,8,10}

	out := mapper(f, in)
	if !reflect.DeepEqual(exp, out) {
		t.Fatal("expected", exp, ", got", out)
	}
}

func TestStringIntMap(t *testing.T) {
	var mapper func (func (string) int, []string) []int
	MakeMap(&mapper)

	in := []string{"try", "this", "thing"}
	f := func(x string) int { return len(x) }
	exp := []int{3, 4, 5}

	out := mapper(f, in)

	if !reflect.DeepEqual(exp, out) {
		t.Fatal("expected", exp, ", got", out)
	}
}

func TestIntStringMap(t *testing.T) {
	var mapper func (func (int) string, []int) []string
	MakeMap(&mapper)

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
	var mapper func (func (int) int, []int) []int
	MakeMap(&mapper)

	in := []int{}
	f := func(x int) int { return x*2 }
	exp := []int{}

	out := mapper(f, in)
	if !reflect.DeepEqual(exp, out) {
		t.Fatal("expected", exp, ", got", out)
	}
}

func TestFilter(t *testing.T) {
	var filter func (func (int) bool, []int) []int
	MakeFilter(&filter)

	in := []int{1,2,3,4,5,6,7,8,9,10}
	f := func(x int) bool { return x % 2 == 0 }
	exp := []int{2,4,6,8,10}

	out := filter(f, in)
	if !reflect.DeepEqual(exp, out) {
		t.Fatal("expected", exp, ", got", out)
	}
}

func TestReduce(t *testing.T) {
	var reduce func (func (int, int) int, []int) int
	MakeReduce(&reduce)

	in := []int{1,2,3,4,5,6,7,8,9,10}
	f := func(x, y int) int { return x + y }
	exp := 55

	out := reduce(f, in)
	if exp != out {
		t.Fatal("expected", exp, ", got", out)
	}
}

func TestReduceInit(t *testing.T) {
	var reduce func (func (int, int) int, []int, int) int
	MakeReduce(&reduce)

	in := []int{1,2,3,4,5,6,7,8,9,10}
	f := func(x, y int) int { return x + y }
	init := 45
	exp := 100

	out := reduce(f, in, init)
	if exp != out {
		t.Fatal("expected", exp, ", got", out)
	}
}