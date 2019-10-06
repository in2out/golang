package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []int{1, 2, 3, 4, 5}

	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))

	sl := []int{} // defualt value is nil
	fmt.Println(sl)

	sli := make([]int, 5)
	fmt.Println(sli)

	slic := make([]int, 5, 10)
	fmt.Println(slic)

	// slice 의 slice는 ptr로 연결되어 있다.
	// slice 의 구조가 ptr/len/cap 으로 되어 있어서, ptr 이 연결되는 구조.

	// func으로 전달할때, value가 넘어가지만, 내부 값이 포인터기때문에 함수안에서 값 변조가 된다.

	// slice 는 slice끼리 비교가 안된다.
	// 비교를 위해서는 방문을 통해 직접 비교 해야 한다.

	b := []int{1, 20, 300, 4000}

	fmt.Println(len(b), cap(b))
}
