package main

import (
	"fmt"
	"reflect"
)

type Runner interface {
	Run()
}

type Gopher struct {
	name string
	age  int
}

func (g Gopher) Run() {
	fmt.Println("RUN", g.name)
}

type Turtle struct {
	name string
}

func (t Turtle) Run() {
	fmt.Println("SLOW", t.name)
}

//func PrintErr(err error) {
//	fmt.Println("ERR:", err.Error())
//}

func main() {
	var runner Runner = Gopher{
		name: "jin",
		age:  100,
	}
	fmt.Println(runner)
	runner.Run()
	fmt.Println(reflect.TypeOf(runner))

	runner = Turtle{
		name: "TT",
	}
	runner.Run()
	fmt.Println(reflect.TypeOf(runner))

	// err := errors.New("err")
	// fmt.Println(err)

	var i interface{}

	i = true
	i = 1000
	i = "string"

	fmt.Println(i)
}
