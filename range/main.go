package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := []int{100, 200, 300, 400, 500}
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[1])
	}

	for i := range a {
		fmt.Println(i, a[i])
	}

	for i, val := range a {
		fmt.Println(i, val)
	}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[1])
	}
	for key, val := range m {
		fmt.Println(key, val)
	}

	// map 에는 순서가 없다.
	for key := range m {
		delete(m, key)
	}
	fmt.Println(m)
}
