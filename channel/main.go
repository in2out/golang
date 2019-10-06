package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	Name string
}

func (w Worker) Work(c chan int, wg *sync.WaitGroup) {
	for task := range c {
		time.Sleep(time.Second)
		fmt.Printf("%s: DONE[%d]\n", w.Name, task)
		wg.Done()
	}
	fmt.Printf("%s: STOP\n", w.Name)
}

func main() {
	//c := make(chan int, 10)
	//go Send(c, 100)
	//fmt.Println(<-c)

	workers := make([]Worker, 0, 10)
	//w := Worker{
	//Name: "Worker_1",
	//}
	c := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		workers = append(workers, Worker{
			Name: fmt.Sprintf("Worker_%d", i),
		})
	}

	for _, worker := range workers {
		go worker.Work(c, wg)
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		c <- i
	}

	close(c)
	wg.Wait()

	val, ok := <-c
	fmt.Println(val, ok)
}
