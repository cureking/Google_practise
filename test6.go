package main

import (
	"fmt"
	"time"
)

func dowork(i int, c chan int) {
	for {
		fmt.Printf("worker %d received:%c \n", i, <-c)
	}
}

func createworker(i int) chan<- int {
	c := make(chan int)
	go dowork(i, c)
	return c
}

func chanDemo() {
	//c := make(chan int)
	c := [10]chan<- int{}
	for i := 0; i < 10; i++ {
		c[i] = createworker(i)
	}
	for i := 0; i < 10; i++ {
		c[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		c[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
