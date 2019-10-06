package main

import "fmt"

func main() {
	m := map[string]int{
		"A": 1,
		"B": 3000,
		"C": 50000,
	}
	m["D"] = 6000

	fmt.Println(m)

	val := m["D"]
	fmt.Println(val)
	val1, ok := m["E"]
	fmt.Println(val1, ok)

	if _, ok := m["E"]; ok == false {
		fmt.Println("E in not in")
	}

	delete(m, "D")
	val2 := m["D"]
	fmt.Println(val2)
}
