package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done *sync.WaitGroup
}

func doWork(i int, in chan int, wg *sync.WaitGroup) {
	for {
		fmt.Printf("Worker %d received %c \n", i, <-in)
		wg.Done()
	}
}

func createWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{
		in:   make(chan int),
		done: wg,
	}
	go doWork(i, w.in, wg)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	workers := [10]worker{}
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
	//wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}
}

func main() {
	chanDemo()
}
