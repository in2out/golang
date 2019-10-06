package main

import (
	"fmt"
	"sync"
	"time"
)

func Work(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Println("start", id)
	time.Sleep(time.Second)
	fmt.Println("done", id)
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go Work(wg, i)
	}
	wg.Wait()
	fmt.Println("Main")
}
