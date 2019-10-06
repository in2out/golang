package main

import (
	"fmt"
	"reflect"
)

func printResult(f func(int, int) int) {
	fmt.Println(f(1, 2))
}

func printResult2(f func(int, int) int) func() {
	return func() {
		fmt.Println(f(1, 2))
	}
}

func main() {
	var add = func(a, b int) int {
		return a + b
	}

	fmt.Println(add(1, 2))
	fmt.Println(reflect.TypeOf(add))

	printResult(add)

	f := printResult2(add)
	f()
}
