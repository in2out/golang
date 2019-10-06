package main

import "fmt"


// embeding 을 지원한다.
// 선언시는 embeding형태로 작성, 엑세스 할때는 literal로 써도 된다.
type Gopher struct {
	name string
	age  int
}

func main() {
	a := Gopher{"name", 10}
	b := Gopher{
		name: "name",
		age:  20,
	}

	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("defer")
	}

	fmt.Println(a)
	fmt.Println(b)

	c Gopher
	fmt.Println(c)
}
