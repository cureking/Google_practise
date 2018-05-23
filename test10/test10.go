package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func doWork(i int, in chan int, wg *sync.WaitGroup) {
	for {
		fmt.Printf("Worker %d received %c \n", i, <-in)
		wg.Done()
	}
}

func createWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWork(i, w.in, wg)
	return w
}

func chanDemo() {
	workers := [10]worker{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	//在这里作wg.Add(20)，比在下面循环内做wg.Add(1)靠谱。后者可能出现deadlock情况，原因自行思考。
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
