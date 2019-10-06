package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}

	fmt.Println(a)
	fmt.Println(len(a), cap(a))

	a = append(a, 5)

	fmt.Println(a)
	fmt.Println(len(a), cap(a)) // cap 이 보통의 경우 2배로 증가한다.

	// 일반적으로 make을 통해 할당해서, 사용하는게 좋다
}
