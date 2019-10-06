package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr [10]int

	brr := [5]int{1, 2, 3, 4, 5}
	crr := [5]int{1, 2, 3, 4, 5}

	fmt.Println(brr == crr)
	fmt.Println(arr)
	fmt.Println(brr)
	fmt.Println(brr[3])

	slice := brr[0:2]

	fmt.Println(slice) // 가변적 길이 가진 array

	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(slice))
}
