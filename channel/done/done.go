package main

import (
	"fmt"
	"sync"
)

func doWork(id int,
	c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("Worker %d received %c \n", id, n)
		go func() {
			//这里使用goroutine是因为这里的done是来自两个不同的channel
			//channel是阻塞式的.
			wg.Done()
		}()
	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func createworker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWork(id, w.in, wg)
	return w
}

func chanDemo() {
	//var c chan int	//c==nil
	//c := make(chan int)
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createworker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()

	// wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//	//这里有两个done，导致阻塞
	//}
}

func main() {
	chanDemo()
}
