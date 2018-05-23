package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

func dowork(i int,
	c chan int, done chan bool) {
	for x := range c {
		fmt.Printf("worker %d received %c \n", i, x)
		done <- true
	}
}

func createworker(i int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go dowork(i, w.in, w.done)
	return w
}

func chanDemo() {
	workers := [10]worker{}
	for i := 0; i < 10; i++ {
		workers[i] = createworker(i)
	}
	for i := 0; i < 10; i++ {
		//这里采用的是无缓冲通道，即阻塞式通信
		//当workers[i].in传递数据给dowork，必须得等dowork完成工作后，才能返回workers[i].done
		workers[i].in <- 'a' + i
		//只有当该处接受到workers[i].done，才可以进行下一步。（无缓冲通道特点：同步性）
		<-workers[i].done
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		<-workers[i].done
	}
}

func main() {
	chanDemo()
}
