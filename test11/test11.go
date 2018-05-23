package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func doWork(i int, wg worker) {
	for {
		fmt.Printf("Worker %d received %c \n", i, <-wg.in)
		wg.done()
	}
}

func createWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{make(chan int), func() {
		wg.Done()
	}}
	go doWork(i, w)
	return w
}

func chanDemo() {
	workers := [10]worker{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	chanDemo()
}
