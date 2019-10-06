package main

import (
	"fmt"
)

// 메쏘드의 집합
type Sounder interface {
	Sound() string
}

type Pig struct {
	Name string
}

func (p Pig) Sound() string {
	return "꿀"
}

type Duck struct {
	Name string
}

func (d Duck) Sound() string {
	return "꽥"
}

//func PlaySound(s Sounder) {
//	var, ok := s.(Pig)
//
//	if !ok {
//		fmt.Println("It's not a pig.")
//	}
//
//	fmt.Println(s.Sound())
//}

func PlaySound(s Sounder) {
	switch val := s.(type) {
	case Pig:
		fmt.Println(val.Sound())
	case Duck:
		fmt.Println(val.Sound())
	default:
		fmt.Println("unkown", s.Sound())
	}
}

func main() {
	p := Pig{}
	d := Duck{}

	PlaySound(p)
	PlaySound(d)
}
