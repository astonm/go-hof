package hof_test

import "fmt"
import "github.com/astonm/go-hof"

func ExampleMakeMapFunc() {
	var mapper func(func(int) int, []int) []int
	hof.MakeMapFunc(&mapper)

	in := []int{1, 2, 3, 4, 5}
	f := func(x int) int { return x * 2 }

	for _, n := range(mapper(f, in)) {
		fmt.Println(n)
	}
	// Output:
	// 2
	// 4
	// 6
	// 8
	// 10
}

func ExampleMakeFilterFunc() {
	var filter func(func(int) bool, []int) []int
	hof.MakeFilterFunc(&filter)

	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := func(x int) bool { return x%2 == 1 }

	for _, n := range(filter(f, in)) {
		fmt.Println(n)
	}
	// Output:
	// 1
	// 3
	// 5
	// 7
	// 9
}


func ExampleMakeReduceFunc() {
	var reduce func(func(int, int) int, []int) int
	hof.MakeReduceFunc(&reduce)

	in := []int{1, 4, 9, 16, 25}
	f := func(x, y int) int { return x + y }

	fmt.Println(reduce(f, in))
	// Output:
	// 55
}

func ExampleMakeReduceFunc_withInitial() {
	var reduce func(func(string, int) string, []int, string) string
	hof.MakeReduceFunc(&reduce)

	in := []int{5,4,3,2,1}
	f := func(x string, y int) string { return x + string('0'+y) }

	fmt.Println(reduce(f, in, ""))
	// Output:
	// 54321
}