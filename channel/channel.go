package main

import (
	"fmt"
	"time"
)

func createworker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			//n:=<-c
			fmt.Printf("worker %d received: %c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {
	//var c chan int	//c==nil
	//c := make(chan int)
	var channels [10]chan<- int
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
	//设置缓存为3
	go func() {
		for {
			fmt.Println(<-c)
		}

	}()
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go func() {
		for n := range c {
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//Channel as first-class citizen
	chanDemo()
	//Buffered Channel
	//bufferedChannel()
	//Channel Close
	//channelClose()

	//GO的并发是基于CSP模型。

	//PS：不要通过共享内存来通信；通过通信来共享内存
	//Don't communicate by sharing memory;share memory by communicating

}
