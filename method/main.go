package main

import (
	"fmt"
)

type Gopher struct {
	name string
	age  int
}

// mehtod receiver
func (g Gopher) Info() {
	fmt.Printf("[%s] %d\n", g.name, g.age)
}

func (g *Gopher) Rename(newName string) {
	g.name = newName
}

func Info(g Gopher) {
	fmt.Printf("[%s] %d\n", g.name, g.age)
}

func main() {
	g := Gopher{
		name: "tom",
		age:  10,
	}
	Info(g)
	g.Info()
	Gopher.Info(g)

	g.Rename("jung")
	g.Info()

	p := &g
	p.Info()

	p.Rename("jin")
	p.Info()
	g.Info()

	k := p
	k.Info()
	k.Rename("jin")
	k.Info()
	g.Info()

}
