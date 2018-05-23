package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func createWorker(id int) chan int {
	c := make(chan int)
	go doWork(id, c)
	return c
}

func doWork(id int, c chan int) {
	for {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d \n", id, <-c)
	}
}

func main() {
	c1, c2 := generator(), generator()
	worker := createWorker(0)
	values := []int{}

	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)

	for {
		var activeValue int
		var activeWorker chan int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-tm:
			fmt.Println("bye")
			return
		case <-tick:
			fmt.Println("values.length:", len(values))
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		}
	}

}
