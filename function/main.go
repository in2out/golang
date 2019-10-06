package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(sub(1, 2))
	fmt.Println(multipleResults(1, 2))
	fmt.Println(namedReturn(1, 2))
}

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func multipleResults(a, b int) (int, int) {
	return a + b, a - b
}

func namedReturn(a, b int) (c int, d int) {
	c = a + b
	d = a - b
	return
}
