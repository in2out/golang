package main

import (
	"fmt"
)

type Sender struct {
	val string
}

// c <-chan read only
func (s Sender) Send(c chan<- string) { // <- write only
	//for {
	c <- s.val
	//}
}

func main() {
	s1 := Sender{
		val: "AAA",
	}
	s2 := Sender{
		val: "BBB",
	}

	c1 := make(chan string)
	c2 := make(chan string)

	go s1.Send(c1)
	go s2.Send(c2)

	for {
		select {
		case val := <-c1:
			fmt.Println("C1:", val)
		case val := <-c2:
			fmt.Println("C2:", val)
		default:
			fmt.Println("default")
		}
	}
}
