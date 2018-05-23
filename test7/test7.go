package main

import (
	"fmt"
	"time"
)

func dowork(i int, c chan int) {
	for x := range c {
		fmt.Printf("worker %d received %c \n", i, x)
	}
}

func createworker(i int) chan<- int {
	c := make(chan int)
	go dowork(i, c)
	return c
}

func chanDemo() {
	channels := [10]chan<- int{}
	for i := 0; i < 10; i++ {
		channels[i] = createworker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go dowork(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go dowork(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}
