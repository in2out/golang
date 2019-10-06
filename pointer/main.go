package main

import "fmt"

func main() {
	a := 0
	p := &a

	// pointer연산은 안된다.
	fmt.Println(&a, *p)

	c := 10
	d := 20

	swap(&c, &d)
	fmt.Println(c, d)
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
