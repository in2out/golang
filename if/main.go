package main

import "fmt"

func add(ad, b int) int {
	return a + b
}

func main() {
	var b bool

	if b {
		fmt.Println("b is true")
	} else {
		fmt.Println("b is false")
	}

	// num scope 	는 if clause 범위
	if num := add(1, 2); num == 3 {
		fmt.Println("num is 3")
	}
}
