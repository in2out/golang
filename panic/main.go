package main

import "fmt"

func div(a, b int) (res int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			res = -1
		}
	}()

	return a / b
}

func main() {
	fmt.Println(div(1, 2))
	fmt.Println(div(10, 0))
}
