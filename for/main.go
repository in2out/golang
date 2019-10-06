package main

import "fmt"

func main() {
	var cnt int

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		cnt++
	}
	prt := fmt.Println("cnt is ", cnt)

	for {
		cnt++
		if cnt >= 20 {
			break
		}
	}
	prt()
}
