package main

import (
	"fmt"
)

func main() {
	fmt.Println("main 1")

	// stack 에 입력되고, 실행시킨다.
	defer fmt.Println("defer 1")

	fmt.Println("main 2")

	defer fmt.Println("defer 2")

	fmt.Println("main 3")

}
